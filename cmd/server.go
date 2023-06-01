package cmd

import (
	"LinkUp_Update/config"
	"LinkUp_Update/var/logs"
	"errors"
	"github.com/gin-gonic/gin"
	"sync"

	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

type server interface {
	StartServer(r *gin.Engine)
	stopServer(serverErrCh chan error, done chan<- bool) error
	runServer(r *gin.Engine, addr string) error
}

type MyServer struct {
	log *logs.Logging
	mx  *sync.RWMutex
}

func (s *MyServer) StartServer(r *gin.Engine) {
	// Starting the server
	serverErrCh := make(chan error, 1)
	doneCh := make(chan bool, 1)

	addr := fmt.Sprint(config.Get("HTTP_HOST").ToString(), ":", config.Get("HTTP_PORT").ToString())
	go func() {
		defer func() {
			if err := recover(); err != nil {
				s.log.LogApi(fmt.Errorf("panic occurred, stopping server: %v", err))
				if err = s.stopServer(serverErrCh, doneCh); err != nil {
					s.log.LogApi(errors.New(fmt.Sprint(err)))
				}
			}
		}()
		serverErrCh <- s.runServer(r, addr)
	}()

	// Waiting for a signal to stop the server.
	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt, syscall.SIGTERM)

	select {
	case <-stop:
		log.Print("interrupt signal received, stopping server")
		if err := s.stopServer(serverErrCh, doneCh); err != nil {
			s.log.LogApi(fmt.Errorf("failed to stop server: %v", err))
			os.Exit(1)
		}
	case <-doneCh:
		return
	}
}

func (s *MyServer) stopServer(serverErrCh chan error, done chan<- bool) error {
	log.Print("stopping server...")

	// Send a signal to stop the server by closing the server error channel.
	close(serverErrCh)

	select {
	case <-time.After(5 * time.Second):
		// If the server does not stop within 5 seconds, return an error.
		return errors.New("server did not stop gracefully")
	case err := <-serverErrCh:
		if err != nil {
			return fmt.Errorf("server error: %w", err)
		}
	}

	log.Print("server closed")

	// Signal that the server has stopped.
	done <- true

	return nil
}

func (s *MyServer) runServer(r *gin.Engine, addr string) error {
	goServer := http.Server{
		Addr:         addr,
		Handler:      r,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
	}

	log.Printf("starting server at %s", addr)

	// Creating a context with the possibility of termination when receiving a signal to stop the server.
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	errCh := make(chan error, 1)
	go func() {
		if err := goServer.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			errCh <- fmt.Errorf("error starting server: %w", err)
		}
	}()

	// Waiting for a signal to stop the server.
	<-ctx.Done()

	s.log.LogApi(errors.New("server closed"))

	// Shutdown the server with a timeout of 5 seconds.
	ctxShutdown, cancelShutdown := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancelShutdown()

	if err := goServer.Shutdown(ctxShutdown); err != nil {
		errCh <- fmt.Errorf("error stopping server: %w", err)
	}

	select {
	case err := <-errCh:
		return err
	default:
		return nil
	}
}
