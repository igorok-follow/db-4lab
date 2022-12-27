package repository

import (
	"context"

	"database-service/config"
	"github.com/go-redis/redis"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

type Connection struct {
	databaseURI string
	DB          *sqlx.DB
}

func NewConnection(uri string) *Connection {
	return &Connection{
		databaseURI: uri,
	}
}

func (conn *Connection) Open() error {
	ctx := context.Background()

	db, err := sqlx.ConnectContext(ctx, "mysql", conn.databaseURI)
	if err != nil {
		return err
	}

	conn.DB = db

	return nil
}

func NewRedisClient(redisClient *config.Redis) (*redis.Client, error) {
	client := redis.NewClient(&redis.Options{
		Addr:     redisClient.Address,
		Password: redisClient.Password,
	})

	_, err := client.Ping().Result()
	if err != nil {
		return nil, err
	}

	return client, nil
}
