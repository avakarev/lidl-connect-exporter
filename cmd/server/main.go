// Command server starts server, monitors balance / consumption status, and exposes these data as metrics
package main

import (
	"time"

	"github.com/avakarev/go-util/buildmeta"
	"github.com/avakarev/go-util/envutil"
	"github.com/avakarev/go-util/timeutil"
	"github.com/avakarev/go-util/zerologutil"
	"github.com/go-co-op/gocron/v2"
	"github.com/rs/zerolog/log"

	"github.com/avakarev/lidl-connect-exporter/internal/lidlconnect"
	"github.com/avakarev/lidl-connect-exporter/internal/metrics"
)

func meterConsumptions(client *lidlconnect.Client) {
	consumptions, err := client.GetConsumptions()
	if err != nil {
		log.Error().Err(err).Send()
		return
	}
	metrics.MeterConsumption(consumptions)
}

func meterBalance(client *lidlconnect.Client) {
	balance, err := client.GetBalanceInfo()
	if err != nil {
		log.Error().Err(err).Send()
		return
	}
	metrics.MeterBalance(balance)
}

func meterTariff(client *lidlconnect.Client) {
	tariff, err := client.GetBookedTariff()
	if err != nil {
		log.Error().Err(err).Send()
		return
	}
	metrics.MeterBookedTariff(tariff)
}

func main() {
	zerologutil.MustInit()
	log.Info().Str("ref", buildmeta.Ref).Str("commit", buildmeta.Commit).Msg("build meta")

	client := lidlconnect.NewClient(&lidlconnect.Account{
		Username: envutil.MustStr("LIDL_CONNECT_USERNAME"),
		Password: envutil.MustStr("LIDL_CONNECT_PASSWORD"),
		Name:     "default",
	})

	scheduler, err := gocron.NewScheduler(gocron.WithLocation(timeutil.Location))
	if err != nil {
		log.Fatal().Err(err).Send()
	}
	if _, err := scheduler.NewJob(gocron.DurationJob(5*time.Minute), gocron.NewTask(meterConsumptions, client)); err != nil {
		log.Fatal().Err(err).Send()
	}
	if _, err := scheduler.NewJob(gocron.DurationJob(30*time.Minute), gocron.NewTask(meterBalance, client)); err != nil {
		log.Fatal().Err(err).Send()
	}
	if _, err := scheduler.NewJob(gocron.DurationJob(1*time.Hour), gocron.NewTask(meterTariff, client)); err != nil {
		log.Fatal().Err(err).Send()
	}
	scheduler.Start()

	if err := metrics.Serve(); err != nil {
		log.Fatal().Err(err).Send()
	}
}
