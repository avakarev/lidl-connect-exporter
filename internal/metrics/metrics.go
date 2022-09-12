package metrics

import (
	"github.com/prometheus/client_golang/prometheus"
	"github.com/rs/zerolog/log"

	"github.com/avakarev/lidl-connect-exporter/internal/lidlconnect"
)

// Namespace is prometheus metrics namespace
const Namespace = "lidl_connect"

var balance = prometheus.NewGauge(prometheus.GaugeOpts{
	Namespace: Namespace,
	Name:      "balance",
	Help:      "The state of the balance",
})

var bookedTariff = prometheus.NewGaugeVec(prometheus.GaugeOpts{
	Namespace: Namespace,
	Name:      "booked_tariff_fee",
	Help:      "Booked tariff fee",
}, []string{"name"})

var consumptionConsumed = prometheus.NewGaugeVec(prometheus.GaugeOpts{
	Namespace: Namespace,
	Name:      "consumption_consumed",
	Help:      "Consumption consumed",
}, []string{"unit", "type"})

var consumptionLeft = prometheus.NewGaugeVec(prometheus.GaugeOpts{
	Namespace: Namespace,
	Name:      "consumption_left",
	Help:      "Consumption left",
}, []string{"unit", "type"})

var consumptionMax = prometheus.NewGaugeVec(prometheus.GaugeOpts{
	Namespace: Namespace,
	Name:      "consumption_max",
	Help:      "Consumption max",
}, []string{"unit", "type"})

// MeterBalance records balance metrics
func MeterBalance(b *lidlconnect.BalanceInfo) {
	value := float64(b.CurrentCustomer.Balance) / 100
	balance.Set(value)
	log.Debug().Float64("balance", value).Msg("balance metering")
}

// MeterBookedTariff records booked tariff metrics
func MeterBookedTariff(t *lidlconnect.BookedTariff) {
	value := float64(t.BasicFee) / 100
	bookedTariff.With(prometheus.Labels{"name": t.Name}).Set(value)
	log.Debug().Float64("fee", value).Str("name", t.Name).Msg("tariff metering")
}

// MeterConsumption records consumptions metrics
func MeterConsumption(consumptions []lidlconnect.ConsumptionsForUnit) {
	for _, c := range consumptions {
		labels := prometheus.Labels{"unit": c.FormattedUnit, "type": c.Type}
		consumptionConsumed.With(labels).Set(c.Consumed)
		consumptionLeft.With(labels).Set(c.Left)
		consumptionMax.With(labels).Set(c.Max)
		log.Debug().
			Float64("consumed", c.Consumed).
			Float64("left", c.Left).
			Float64("max", c.Max).
			Str("unit", c.FormattedUnit).
			Str("type", c.Type).
			Msg("consumptions metering")
	}
}

func init() {
	prometheus.MustRegister(balance)
	prometheus.MustRegister(bookedTariff)
	prometheus.MustRegister(consumptionConsumed)
	prometheus.MustRegister(consumptionLeft)
	prometheus.MustRegister(consumptionMax)
}
