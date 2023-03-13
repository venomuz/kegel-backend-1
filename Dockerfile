FROM golang:1.18.10-alpine3.17

RUN mkdir -p /root/public

RUN mkdir -p /root/go/src/github.com/venomuz/kegel-backend

WORKDIR /root/go/src/github.com/venomuz/kegel-backend

COPY . .


RUN go mod download

RUN go build -o main cmd/app/main.go

EXPOSE 9090

CMD ./main