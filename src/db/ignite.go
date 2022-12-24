package db

import (
	"crypto/tls"
	"net"
	"time"

	ignite "github.com/amsokol/ignite-go-client/binary/v1"
	"github.com/gurbaj5124871/url-shortener/config"
	"github.com/rs/zerolog"
	logger "github.com/rs/zerolog/log"
)

var log zerolog.Logger = logger.With().Str("logger", "db").Logger()

func IgniteConnect() ignite.Client {
	conf := config.GetConfig()

	opts := ignite.ConnInfo{
		Network:  "tcp",
		Host:     conf.Ignite.Host,
		Port:     conf.Ignite.Post,
		Major:    1,
		Minor:    1,
		Patch:    0,
		Username: conf.Ignite.Username,
		Password: conf.Ignite.Password,
		Dialer: net.Dialer{
			Timeout: 10 * time.Second,
		},
	}
	if conf.Ignite.TLS {
		opts.TLSConfig = &tls.Config{
			// You should only set this to true for testing purposes.
			InsecureSkipVerify: true,
		}
	}

	c, err := ignite.Connect(opts)
	if err != nil {
		log.Fatal().Err(err).Msg("Failed to connect to ignite")
	}
	defer c.Close()

	log.Info().Msg("Apache Ignite connected")

	return c
}
