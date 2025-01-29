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
	addr := fmt.Sprintf("%s:%s", cfg.Server.Host, cfg.Server.Port)
	conn, err := grpc.DialContext(context.Background(), addr, grpc.WithBlock(), grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("could not connect to order service: %v", err)
	}
	defer conn.Close()

	mux := runtime.NewServeMux(
		runtime.WithIncomingHeaderMatcher(runtime.DefaultHeaderMatcher),
		runtime.WithOutgoingHeaderMatcher(runtime.DefaultHeaderMatcher),
		runtime.WithErrorHandler(runtime.DefaultHTTPErrorHandler),
	)

	if err = pb.RegisterMovieServiceHandler(context.Background(), mux, conn); err != nil {
		log.Fatalf("failed to register the order server: %v", err)
	}

	clientAddr := fmt.Sprintf("%s:%s", cfg.Client.Host, cfg.Client.Port)
	log.Println("API gateway server is running on " + clientAddr)
	if err = http.ListenAndServe(clientAddr, mux); err != nil {
		log.Fatal("gateway server closed abruptly: ", err)
	}

}
