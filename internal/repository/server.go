package repository

import (
	"context"
	"simpel-gateway/internal/dto"
	"simpel-gateway/internal/model"
	"simpel-gateway/pkg/helper"
	"simpel-gateway/pkg/util"

	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
)

type Server interface {
	Store(ctx context.Context, data model.Server) error
	FindOne(ctx context.Context, selectedFields string, query string, args ...any) (model.Server, error)
	GetAll(ctx context.Context) ([]dto.ListServer, error)
}

type server struct {
	Db           *gorm.DB
	RedisClient  *redis.Client
	CacheEnabled bool
}

func NewServerRepository(db *gorm.DB, redisClient *redis.Client) Server {
	return &server{
		Db:           db,
		RedisClient:  redisClient,
		CacheEnabled: true,
	}
}

func (r *server) Store(ctx context.Context, data model.Server) error {
	tx := r.Db.WithContext(ctx)
	if err := tx.Model(model.Server{}).Create(&data).Error; err != nil {
		return err
	}

	cacheKey := "server_list-*"

	if err := helper.DeleteRedisKeysByPattern(r.RedisClient, cacheKey); err != nil {
		return nil
	}

	return nil
}

func (r *server) FindOne(ctx context.Context, selectedFields string, query string, args ...interface{}) (model.Server, error) {
	var res model.Server

	db := r.Db.WithContext(ctx).Model(&model.Server{})

	db = util.SetSelectFields(db, selectedFields)

	// Gunakan First() untuk menemukan satu data
	if err := db.Where(query, args...).First(&res).Error; err != nil {
		return model.Server{}, err
	}

	return res, nil
}

func (r *server) GetAll(ctx context.Context) ([]dto.ListServer, error) {
	tx := r.Db.WithContext(ctx).Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	const timeFormat = "2006-01-02 15:04:05"

	query := tx.Model(&model.Server{})
	var res []model.Server
	if err := query.Find(&res).Error; err != nil {
		tx.Rollback()
		return nil, err
	}

	var ListServer []dto.ListServer
	for _, dserver := range res {
		formatCreatedAt := dserver.CreatedAt.Format(timeFormat)
		formatUpdatedAt := dserver.UpdatedAt.Format(timeFormat)

		serverDTO := dto.ListServer{
			ServerIp:   dserver.ServerIp,
			ServerName: dserver.ServerName,
			CreatedAt:  formatCreatedAt,
			UpdatedAt:  formatUpdatedAt,
		}

		ListServer = append(ListServer, serverDTO)
	}

	if err := tx.Commit().Error; err != nil {
		tx.Rollback()
		return nil, err
	}
	return ListServer, nil
}
