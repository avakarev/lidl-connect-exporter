// Command cli prints current balance/consumptions state
package main

import (
	"github.com/avakarev/go-util/buildmeta"
	"github.com/avakarev/go-util/zerologutil"
	"github.com/rs/zerolog/log"

	"github.com/avakarev/lidl-connect-exporter/internal/lidlconnect"
)

func main() {
	zerologutil.MustInit()
	log.Info().Fields(buildmeta.Fields()).Msg("build meta")

	client := lidlconnect.DefaultClient()

	balance, err := client.GetBalanceInfo()
	if err != nil {
		log.Error().Err(err).Send()
	}
	log.Info().Interface("balance", balance).Send()

	tariff, err := client.GetBookedTariff()
	if err != nil {
		log.Error().Err(err).Send()
	}
	log.Info().Interface("tariff", tariff).Send()

	consumptions, err := client.GetConsumptions()
	if err != nil {
		log.Error().Err(err).Send()
	}
	log.Info().Interface("consumptions", consumptions).Send()
}
