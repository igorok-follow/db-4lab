package app

import (
	"context"
	"net"
	"net/http"
	"time"

	"database-service/helpers/logger"
	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	"github.com/jmoiron/sqlx"
	"github.com/rs/cors"
	"google.golang.org/grpc"

	"database-service/config"
	"database-service/helpers"
	"database-service/pkg/api"
	"database-service/pkg/repository"
	"database-service/pkg/service"
)

func Run(config *config.Config) {
	serverMeta := &helpers.ServerMeta{
		LogLevel: config.Logger.LogLevel,
	}

	conn := repository.NewConnection(config.Database.Uri)
	err := conn.Open()
	if err != nil {
		logger.FatalError("Configure connection error", err)
	}

	defer func(DB *sqlx.DB) {
		err := DB.Close()
		if err != nil {
			logger.FatalError("Defer db close error", err)
		}
	}(conn.DB)

	logger.Info("Connected to users database")

	repositories := repository.NewRepositories(conn.DB)

	hasher := helpers.NewHasher(config.Token.Salt)

	jwtManager, err := helpers.NewJWT(config.Token.SecretKey)
	if err != nil {
		logger.FatalError("JWT manager error", err)
	}

	redisClient, err := repository.NewRedisClient(config.Redis)
	if err != nil {
		logger.FatalError("Configure redis connection error", err)
	}

	defer redisClient.Close()

	logger.Info("Connected to redis")

	deps := &service.Dependencies{
		Repository: repositories,
		Hasher:     hasher,
		JWTManager: jwtManager,
	}
	services := service.NewServices(deps)

	s := grpc.NewServer()
	api.RegisterDatabaseServer(s, services.DatabaseService)

	logger.Info("Starting database-service...",
		logger.String("host", config.Server.Host),
		logger.String("port", config.Server.Port))

	l, err := net.Listen("tcp", config.Server.Port)
	if err != nil {
		logger.Error("listen tcp error", err)
	}

	go func() {
		logger.Error("serve error", s.Serve(l))
	}()

	serverMeta.GrpcServerStarted = time.Now().String()

	gwconn, err := grpc.DialContext(
		context.Background(),
		"localhost"+config.Server.Port,
		grpc.WithBlock(),
		grpc.WithInsecure(),
	)
	if err != nil {
		logger.FatalError("Failed to dial server", err)
	}

	gwmux := runtime.NewServeMux(runtime.WithMarshalerOption(runtime.MIMEWildcard, &runtime.JSONPb{OrigName: false}))

	handler := cors.New(cors.Options{
		AllowedOrigins:   []string{"*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "HEAD", "OPTIONS"},
		AllowedHeaders:   []string{"Content-Type", "X-Requested-With", "Authorization", "x-forwarded-for"},
		AllowCredentials: true,
	}).Handler(gwmux)
	err = api.RegisterDatabaseHandler(context.Background(), gwmux, gwconn)
	if err != nil {
		logger.FatalError("Failed to register gateway", err)
	}

	gwServer := &http.Server{
		Addr:    config.Gateway.Port,
		Handler: gwmux,
	}

	serverMeta.GrpcGatewayStarted = time.Now().String()
	serverMeta.Running = time.Now().String()

	logger.Info("Serving gRPC-Gateway...",
		logger.String("host", config.Gateway.Host),
		logger.String("port", config.Gateway.Port))
	logger.FatalError("gRPC-Gateway serving error", http.ListenAndServe(gwServer.Addr, handler))

}
