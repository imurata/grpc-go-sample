package main

import (
	"github.com/example/grpc_sample"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"log"
	"flag"
	"net/http"
)


func main() {
	serverAddress := flag.String("server", "123.123.123.123:9000", "Server address in the format 'host:port")
	flag.Parse()

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		conn, err := grpc.Dial(*serverAddress, grpc.WithInsecure())
		if err != nil {
			log.Fatalf("Did not connect: %s", err)
		}
		c := grpc_sample.NewSampleServiceClient(conn)

		response, err := c.GetData(context.Background(), &grpc_sample.Message{Body: "send data"})
		if err != nil {
			log.Fatalf("Error: %s", err)
		}
		log.Print(response.Body)

		defer conn.Close()
	})

	log.Print("HTTP server is listening on :8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatalf("HTTP server error: %s", err)
	}
}








