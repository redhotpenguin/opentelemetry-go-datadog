package datadog

import (
	"context"
	"log"
	"time"

	"go.opentelemetry.io/otel/api/global"
	export "go.opentelemetry.io/otel/sdk/export/metric"
	"go.opentelemetry.io/otel/sdk/metric/batcher/ungrouped"
	"go.opentelemetry.io/otel/sdk/metric/controller/push"
	"go.opentelemetry.io/otel/sdk/metric/selector/simple"
)

func InstallNewPipeline() (*push.Controller, error) {
	controller, err := NewExportPipeline()
	if err != nil {
		return controller, err
	}

	global.SetMeterProvider(controller)

	return controller, err
}

func NewExportPipeline() (*push.Controller, error) {
	exp, err := NewMeterExporter()
	if err != nil {
		return nil, err
	}

	selector := simple.NewWithExactMeasure()

	batcher := ungrouped.New(selector, export.NewDefaultLabelEncoder(), true)

	pusher := push.New(batcher, exp, time.Minute)
	pusher.Start()

	return pusher, nil
}

// MeterExporter exports metrics to DataDog
type MeterExporter struct {
}

// NewMeterExporter constructs a NewMeterExporter
func NewMeterExporter() (*MeterExporter, error) {
	return &MeterExporter{}, nil
}

// Export exports the provide metric record to DataDog.
func (e *MeterExporter) Export(ctx context.Context, checkpoint export.CheckpointSet) error {
	log.Printf("Export Checkpoint: %+v\n", checkpoint)
	return nil
}
