## This docker uses as draft

FROM golang:1.12

WORKDIR /go/src/github.com/gospeak/auth-service

COPY . .

CMD ["go", "run", "main.go"]
