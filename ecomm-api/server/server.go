package server

import (
	"Micro-services/ecomm-api/storer"
	"context"
)

type Server struct {
	storer *storer.MySQLStorer
}

func NewServer(storer *storer.MySQLStorer) *Server {
	return &Server{
		storer: storer,
	}
}

func (s *Server) CreateProduct(ctx context.Context, p *storer.Product) (*storer.Product, error) {
	return s.storer.CreateProduct(ctx, p)
}

func (s *Server) GetProduct(ctx context.Context, p *storer.Product) (*storer.Product, error) {
	return s.storer.GetProduct(ctx, p)
}

func (s *Server) UpdateProduct(ctx context.Context, p *storer.Product) (*storer.Product, error) {
	return s.storer.UpdateProduct(ctx, p)
}

func (s *Server) ListProducts(ctx context.Context) ([]storer.Product, error) {
	return s.storer.ListProducts(ctx)
}

func (s *Server) DeletePRoduct(ctx context.Context, id int64) error {
	return s.storer.DeleteProduct(ctx, id)
}
