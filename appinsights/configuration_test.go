package appinsights

import (
	"testing"
	"time"
)

func TestTelemetryConfiguration(t *testing.T) {
	testKey := "test"
	defaultEndpoint := "https://dc.services.visualstudio.com/v2/track"

	config := NewTelemetryConfiguration(testKey)

	if config.InstrumentationKey != testKey {
		t.Errorf("InstrumentationKey is %s, want %s", config.InstrumentationKey, testKey)
	}

	if config.EndpointUrl != defaultEndpoint {
		t.Errorf("EndpointUrl is %s, want %s", config.EndpointUrl, defaultEndpoint)
	}

	if config.Client != nil {
		t.Errorf("Client is not nil, want nil")
	}
}

func TestConnectionString(t *testing.T) {
	connectionString := "InstrumentationKey=00000000-0000-0000-0000-000000000001;IngestionEndpoint=https://westeurope-1.in.applicationinsights.azure.com/;LiveEndpoint=https://westeurope.livediagnostics.monitor.azure.com/"
	expectedInstrumentationKey := "00000000-0000-0000-0000-000000000001"
	expectedEndpoint := "https://westeurope-1.in.applicationinsights.azure.com/v2/track"
	defaultBatchSize := 1024
	defaultBatchInterval := time.Duration(10) * time.Second

	config := NewTelemetryConfiguration(connectionString)

	if config.InstrumentationKey != expectedInstrumentationKey {
		t.Errorf("InstrumentationKey is %s, want %s", config.InstrumentationKey, expectedInstrumentationKey)
	}

	if config.EndpointUrl != expectedEndpoint {
		t.Errorf("EndpointUrl is %s, want %s", config.EndpointUrl, expectedEndpoint)
	}

	if config.Client != nil {
		t.Errorf("Client is not nil, want nil")
	}

	if config.MaxBatchSize != defaultBatchSize {
		t.Errorf("MaxBatchSize is %d, want %d", config.MaxBatchSize, defaultBatchSize)
	}

	if config.MaxBatchInterval != defaultBatchInterval {
		t.Errorf("MaxBatchInterval is %s, want %s", config.MaxBatchInterval, defaultBatchInterval)
	}
}
