package main

import (
	"context"
	"flag"
	"fmt"
	"net"
	"os"
	"os/signal"

	"github.com/gurbaj5124871/url-shortener/config"
	grpcLogger "github.com/philip-bui/grpc-zerolog"
	"github.com/rs/zerolog"
	logger "github.com/rs/zerolog/log"
	"google.golang.org/grpc"
)

var log zerolog.Logger = logger.With().Str("logger", "main").Logger()

func main() {
	// Init config
	environment := flag.String("e", "development", "")
	flag.Usage = func() {
		fmt.Println("Usage: server -e {mode}")
		os.Exit(1)
	}
	log.Info().Msgf("Starting server in %s mode", *environment)
	flag.Parse()
	config.Init()
	conf := config.GetConfig()

	// Init tcp/http server
	lis, err := net.Listen("tcp", conf.Server.Port)
	if err != nil {
		log.Fatal().Err(err).Msg("Error while initiating server")
	}

	// Serve gRPC
	s := grpc.NewServer(
		grpcLogger.UnaryInterceptorWithLogger(&log),
	)
	ctx := context.Background()
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	go func() {
		for range c {
			logger.Warn().Msg("Shutting down gRPC server")
			s.GracefulStop()
			<-ctx.Done()
		}
	}()

	log.Info().Msgf("Server listening on port %s", conf.Server.Port)
	s.Serve(lis)
}
