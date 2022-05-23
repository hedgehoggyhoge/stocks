package server

import (
	"context"
	"math/rand"
	pb "stocks/grpc_utils/proto/PortfolioStorage"
)

func CreateServer() *PortfolioServer {
	return &PortfolioServer{db: db.CreateDB()}
}

type PortfolioServer struct {
	pb.UnimplementedPortfolioStorageServer
	db MongoDBWrapper
}

func (p *PortfolioServer) UpdatePortfolio(ctx context.Context, updatePortfolioRequest *pb.UpdatePortfolioRequest) (*pb.UpdatePortfolioResponse, error) {
	err := p.db.UpdateData(updatePortfolioRequest.UserId, updatePortfolioRequest.PortfolioItems)
	if err != nil {
		return nil, err
	}
	return &pb.UpdatePortfolioResponse{}, nil
}

func (p *PortfolioServer) GetPortfolio(ctx context.Context, getPortfolioRequest *pb.GetPortfolioRequest) (*pb.GetPortfolioResponse, error) {
	resp, err := p.db.GetPortfolioData(getPortfolioRequest.UserId)
	if err != nil {
		return nil, err
	}
	var price float32 = 0
	for _, item := range resp.PortfolioItems {
		price += getPrice(item.ticker) * item.quantity
	}
	return &pb.GetPortfolioResponse{
		PortfolioItems: resp.PortfolioItems,
		Price:          price,
	}, nil
}

func getPrice(ticker string) float32 {
	return 1 + rand.Float32()
}
