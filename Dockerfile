FROM golang:latest

WORKDIR /app
COPY . .

RUN go build -o main .

EXPOSE 9797

CMD ["./main"]

