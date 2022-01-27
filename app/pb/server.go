package pb

import (
	"context"
	"fmt"
	"github.com/hinha/PAM-Trello/app/pb/trello"
	log "github.com/sirupsen/logrus"
	"google.golang.org/grpc"
)

type GrpcClient struct {
	Conn GrpcConn
}

const MaxSizeFile = 1024 * 1024 * 250 // 25mb in bytes

func NewGrpc(address, port string) *GrpcClient {
	return &GrpcClient{Conn: GrpcConn{address, port}}
}

type GrpcConn struct {
	address string
	port    string
}

func (g *GrpcConn) connect() (*grpc.ClientConn, error) {
	var client *grpc.ClientConn
	var err error

	portLine := fmt.Sprintf("%s:%s", g.address, g.port)
	client, err = grpc.Dial(portLine,
		grpc.WithDefaultCallOptions(grpc.MaxCallRecvMsgSize(MaxSizeFile)),
		grpc.WithInsecure())

	if err != nil {
		return nil, err
	}
	log.Printf("connected grpc on %s", portLine)

	return client, err
}

type ClientServices struct {
	GrpcConn
	clientTrello trello.TrelloClient
}

func NewService(c GrpcConn) Services {
	connect, err := c.connect()
	if err != nil {
		panic(err)
	}
	return &ClientServices{
		clientTrello: trello.NewTrelloClient(connect),
	}
}

func (c *ClientServices) Analyze(ctx context.Context, req *trello.PamInput) (*trello.AnalyzeResponse, error) {
	response, err := c.clientTrello.Analyze(ctx, req)
	if err != nil {
		return nil, err
	}

	if len(response.Error) > 0 {
		return nil, fmt.Errorf(response.Error)
	}

	return response, nil
}

func (c *ClientServices) GetCard(ctx context.Context, req *trello.Request) (*trello.Response, error) {
	return c.clientTrello.GetCard(ctx, req)
}
