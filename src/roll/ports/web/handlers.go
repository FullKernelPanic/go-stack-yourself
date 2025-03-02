package web

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"go-stack-yourself/src/roll/services"
	"go.opentelemetry.io/contrib/bridges/otelslog"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/metric"
)

var name = "demo.roll.ports.http"

var (
	tracer   = otel.Tracer(name)
	meter    = otel.Meter(name)
	logger   = otelslog.NewLogger(name)
	rollCnt  metric.Int64Counter
	execTime metric.Float64Counter
)

func init() {
	var err error
	rollCnt, err = meter.Int64Counter(
		"dice_rolls",
		metric.WithDescription("The number of rolls by roll value"),
		metric.WithUnit("{roll_unit}"),
	)
	execTime, err = meter.Float64Counter(
		"dice_rolls_exec_time",
		metric.WithDescription("The execution time of the handler"),
		metric.WithUnit("{roll_exec_time}"),
	)
	if err != nil {
		panic(err)
	}
}

func RolldiceHandler(w http.ResponseWriter, r *http.Request) {
	t := time.Now()

	ctx, span := tracer.Start(r.Context(), "RolldiceHandler")

	if span.SpanContext().HasTraceID() {
		w.Header().Set("X-Trace-Id", span.SpanContext().TraceID().String())
	}

	defer func() {
		if span != nil {
			span.End()
		}
	}()

	player := r.URL.Query().Get("player")

	roll := services.RollDice(player)

	var msg string
	if player != "" {
		msg = fmt.Sprintf("%s is rolling the dice", player)
	} else {
		msg = "Anonymous player is rolling the dice"
	}

	logger.InfoContext(ctx, msg, "result", roll)

	rollValueAttr := attribute.Int("diceroll.value", roll)
	span.SetAttributes(rollValueAttr)

	rollCnt.Add(ctx, 1, metric.WithAttributes(rollValueAttr))

	var err error
	if player != "" {
		err = showResultUser(w, r, player, roll)
	} else {
		err = showResultAnonym(w, r, roll)
	}

	if err != nil {
		log.Printf("Write failed: %v\n", err)
	}

	execTime.Add(ctx, time.Now().Sub(t).Seconds())
}

func RollHomeHandler(w http.ResponseWriter, r *http.Request) {
	_ = showDiceRollHome(w, r)
}
