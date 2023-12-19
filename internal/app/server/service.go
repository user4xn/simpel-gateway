package server

import (
	"context"
	"fmt"
	"simpel-gateway/internal/dto"
	"simpel-gateway/internal/factory"
	"simpel-gateway/internal/model"
	"simpel-gateway/internal/repository"
	"simpel-gateway/pkg/constants"
)

type service struct {
	serverRepository repository.Server
}

type Service interface {
	Store(ctx context.Context, payload dto.PayloadServer) error
	GetListServer(ctx context.Context) ([]dto.ListServer, error)
	CheckServerIp(ctx context.Context, payload dto.PayloadCheckServerIp) dto.CheckServerIp
}

func NewService(f *factory.Factory) Service {
	return &service{
		serverRepository: f.ServerRepository,
	}
}

func (s *service) GetListServer(ctx context.Context) ([]dto.ListServer, error) {
	var ListServer []dto.ListServer

	servers, err := s.serverRepository.GetAll(ctx)
	if err != nil {
		return nil, err
	}
	for _, dserver := range servers {
		dserver := dto.ListServer{
			ServerIp:   dserver.ServerIp,
			ServerName: dserver.ServerName,
			CreatedAt:  dserver.CreatedAt,
			UpdatedAt:  dserver.UpdatedAt,
		}
		ListServer = append(ListServer, dserver)
	}

	return ListServer, nil
}

func (s *service) Store(ctx context.Context, payload dto.PayloadServer) error {
	_, err := s.serverRepository.FindOne(ctx, "server_ip", "server_ip = ?", payload.ServerIp)

	if err != nil {
		dataStore := model.Server{
			ServerIp:   payload.ServerIp,
			ServerName: payload.ServerName,
		}
		s.serverRepository.Store(ctx, dataStore)
		return nil
	}

	return constants.DuplicateStoreServer
}

func (s *service) CheckServerIp(ctx context.Context, payload dto.PayloadCheckServerIp) dto.CheckServerIp {
	fmt.Println(payload)
	_, err := s.serverRepository.FindOne(ctx, "server_ip", "server_ip = ?", payload.ServerIp)

	if err != nil {
		data := dto.CheckServerIp{
			IsRegistered: false,
		}
		return data
	}
	data := dto.CheckServerIp{
		IsRegistered: true,
	}
	return data
}
