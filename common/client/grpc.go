package client

import (
	"context"
	"errors"

	"github.com/bektosh03/test-crud/genprotos/datapb"
	"github.com/bektosh03/test-crud/genprotos/postpb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func NewDataServiceClient(ctx context.Context, addr string) (client datapb.DataServiceClient, close func() error, err error) {
	if addr == "" {
		return nil, func() error { return nil }, errors.New("empty address")
	}

	conn, err := grpc.DialContext(ctx, addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, func() error { return nil }, err
	}

	return datapb.NewDataServiceClient(conn), conn.Close, nil
}

func NewPostServiceClient(ctx context.Context, addr string) (client postpb.PostServiceClient, close func() error, err error) {
	if addr == "" {
		return nil, func() error { return nil }, errors.New("empty address")
	}

	conn, err := grpc.DialContext(ctx, addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, func() error { return nil }, err
	}

	return postpb.NewPostServiceClient(conn), conn.Close, nil
}
