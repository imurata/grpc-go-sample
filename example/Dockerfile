FROM golang:latest

WORKDIR /go/src/app

COPY client/client ./client
COPY server/server ./server

EXPOSE 8080
EXPOSE 9000

# サーバーを実行
CMD ["./server"]

