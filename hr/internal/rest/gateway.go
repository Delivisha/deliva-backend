package rest

import (
	"context"
	"deliva/hr/hrpb"
	"github.com/go-chi/chi/v5"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func RegisterGateway(ctx context.Context, mux *chi.Mux, grpcAddr string) error {
	const apiRoot = "/api/hr"

	gateway := runtime.NewServeMux()
	err := hrpb.RegisterHrServiceHandlerFromEndpoint(ctx, gateway, grpcAddr, []grpc.DialOption{
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	})
	if err != nil {
		return err
	}
	// mount the GRPC gateway
	mux.Mount(apiRoot, gateway)
	return nil
}
