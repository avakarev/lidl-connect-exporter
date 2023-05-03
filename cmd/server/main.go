// Command server starts server, monitors balance / consumption status, and exposes these data as metrics
package main

import (
	"github.com/avakarev/go-buildmeta"
	"github.com/avakarev/go-logutil"
	"github.com/avakarev/go-timeutil"
	"github.com/go-co-op/gocron"
	"github.com/rs/zerolog/log"

	"github.com/avakarev/lidl-connect-exporter/internal/lidlconnect"
	"github.com/avakarev/lidl-connect-exporter/internal/metrics"
)

var client *lidlconnect.Client

func meterConsumptions() {
	consumptions, err := client.GetConsumptions()
	if err != nil {
		log.Error().Err(err).Send()
		return
	}
	metrics.MeterConsumption(consumptions)
}

func meterBalance() {
	balance, err := client.GetBalanceInfo()
	if err != nil {
		log.Error().Err(err).Send()
		return
	}
	metrics.MeterBalance(balance)
}

func meterTariff() {
	tariff, err := client.GetBookedTariff()
	if err != nil {
		log.Error().Err(err).Send()
		return
	}
	metrics.MeterBookedTariff(tariff)
}

func main() {
	logutil.MustInit()
	log.Info().Str("ref", buildmeta.Ref).Str("commit", buildmeta.Commit).Msg("build meta")

	client = lidlconnect.DefaultClient()

	scheduler := gocron.NewScheduler(timeutil.Location)
	if _, err := scheduler.Every(5).Minutes().Do(meterConsumptions); err != nil {
		log.Error().Err(err).Send()
	}
	if _, err := scheduler.Every(30).Minutes().Do(meterBalance); err != nil {
		log.Error().Err(err).Send()
	}
	if _, err := scheduler.Every(1).Hour().Do(meterTariff); err != nil {
		log.Error().Err(err).Send()
	}
	scheduler.StartAsync()

	if err := metrics.Serve(); err != nil {
		log.Fatal().Err(err).Send()
	}
}
