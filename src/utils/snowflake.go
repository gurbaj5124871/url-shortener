package utils

import (
	"github.com/gurbaj5124871/url-shortener/config"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"github.com/sony/sonyflake"
)

var logger zerolog.Logger = log.With().Str("logger", "utils").Logger()
var Snowflake *sonyflake.Sonyflake

func InitIDGenerator() {
	settings := sonyflake.Settings{}

	conf := config.GetConfig()
	if conf.Server.MachineID != nil {
		settings.MachineID = func() (uint16, error) {
			return *conf.Server.MachineID, nil
		}
	}

	Snowflake = sonyflake.NewSonyflake(settings)
	if Snowflake == nil {
		logger.Fatal().Msg("error while initiating snowflake id generator")
	}
}
