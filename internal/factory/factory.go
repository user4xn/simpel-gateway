package factory

import (
	"simpel-gateway/database"
	"simpel-gateway/internal/repository"
	"simpel-gateway/pkg/util"

	"github.com/redis/go-redis/v9"
)

type Factory struct {
	ServerRepository repository.Server
}

func NewFactory() *Factory {
	db := database.GetConnection()
	redisClient := redis.NewClient(&redis.Options{
		Addr:     util.GetEnv("REDIS_URL", "localhost") + ":" + util.GetEnv("REDIS_PORT", "6379"),
		Password: util.GetEnv("REDIS_PASS", ""),
		DB:       0,
	})

	return &Factory{
		// Pass the db connection to the repository package for database query calling
		ServerRepository: repository.NewServerRepository(db, redisClient),
		// Assign the appropriate implementation of the ReturInsightRepository
	}
}
