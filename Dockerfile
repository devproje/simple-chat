FROM golang:alpine

WORKDIR /opt/chat

COPY . .

RUN go mod tidy
RUN go build -o simple_chat server.go

EXPOSE 3000

ENTRYPOINT ["/opt/chat/simple_chat"]
