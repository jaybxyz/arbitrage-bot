package main

import (
	"context"
	"fmt"
	"os"

	"github.com/kogisin/arbitrage-bot/client"
	"github.com/kogisin/arbitrage-bot/config"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

var (
	configPath = "config.toml"
)

func init() {
	zerolog.SetGlobalLevel(zerolog.DebugLevel)
	log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stderr}) // pretty logging
}

func main() {
	cfg, err := config.Read(configPath)
	if err != nil {
		log.Err(err).Msg("failed to read config file")
		return
	}

	client, err := client.NewClient(cfg.RPC.Address, cfg.GRPC.Address)
	if err != nil {
		log.Err(err).Msg("failed to create new client")
		return
	}

	s, err := client.RPC.GetStatus(context.Background())
	if err != nil {
		log.Err(err).Msg("failed to get status")
		return
	}

	fmt.Println(s.SyncInfo.LatestBlockHeight)
}
