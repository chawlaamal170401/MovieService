package main

import (
	"context"
	"fmt"
	"github.com/razorpay/movie-service/internals/config"
	pb "github.com/razorpay/movie-service/internals/proto"
	"google.golang.org/grpc/credentials/insecure"
	"log"
	"net/http"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
)

func main() {
	cfg, err := config.LoadConfig()
	addr := fmt.Sprintf(cfg.Server.Host + ":" + cfg.Server.Port)
	conn, err := grpc.Dial(addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("could not connect to order service: %v", err)
	}
	defer conn.Close()

	mux := runtime.NewServeMux(
		runtime.WithIncomingHeaderMatcher(func(headerName string) (string, bool) {
			return headerName, true
		}),
	)

	if err = pb.RegisterMovieServiceHandler(context.Background(), mux, conn); err != nil {
		log.Fatalf("failed to register the order server: %v", err)
	}

	clientAddr := fmt.Sprintf(cfg.Client.Host + ":" + cfg.Client.Port)
	fmt.Println("API gateway server is running on " + clientAddr)
	if err = http.ListenAndServe(clientAddr, mux); err != nil {
		log.Fatal("gateway server closed abruptly: ", err)
	}
}
