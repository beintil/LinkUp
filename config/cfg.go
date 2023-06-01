package config

import (
	"fmt"
	"github.com/joho/godotenv"
	"log"
	"os"
	"strconv"
)

type config interface {
	ToInt() int
	ToString() string
	ToBool() bool
	envRequire(key string) *Cfg
}

type Cfg struct {
	key string
}

func Get(key string) *Cfg {
	var cfg = &Cfg{}
	return cfg.envRequire(key)
}

func (cfg *Cfg) ToString() string {
	return cfg.key
}

func (cfg *Cfg) ToInt() int {
	f, err := strconv.ParseFloat(cfg.key, 64)
	if err != nil {
		log.Fatal(err)
	}
	return int(f)
}

func (cfg *Cfg) ToBool() bool {
	return cfg.key == "t" || cfg.key == "true"
}

func (cfg *Cfg) envRequire(key string) *Cfg {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal(fmt.Errorf("error loading .env file: %s", err))
		return cfg
	}

	cfg.key = os.Getenv(key)
	if cfg.key == "" {
		log.Fatal(fmt.Errorf("in the file .env %s is incorrect or missing", cfg.key))
	}
	return cfg
}
