package appinsights

import (
	"net/http"
	"os"
	"runtime"
	"strings"
	"time"
)

// Configuration data used to initialize a new TelemetryClient.
type TelemetryConfiguration struct {
	// Instrumentation key for the client.
	InstrumentationKey string

	// Endpoint URL where data will be submitted.
	EndpointUrl string

	// Maximum number of telemetry items that can be submitted in each
	// request.  If this many items are buffered, the buffer will be
	// flushed before MaxBatchInterval expires.
	MaxBatchSize int

	// Maximum time to wait before sending a batch of telemetry.
	MaxBatchInterval time.Duration

	// Customized http client if desired (will use http.DefaultClient otherwise)
	Client *http.Client
}

// Creates a new TelemetryConfiguration object with the specified
// connection string and default values.
func NewTelemetryConfiguration(instrumentationKey string) *TelemetryConfiguration {
	config := parseConnectionString(instrumentationKey)
	config.MaxBatchSize = 1024
	config.MaxBatchInterval = time.Duration(10) * time.Second
	return config
}

func (config *TelemetryConfiguration) setupContext() *TelemetryContext {
	context := NewTelemetryContext(config.InstrumentationKey)
	context.Tags.Internal().SetSdkVersion(sdkName + ":" + Version)
	context.Tags.Device().SetOsVersion(runtime.GOOS)

	if hostname, err := os.Hostname(); err == nil {
		context.Tags.Device().SetId(hostname)
		context.Tags.Cloud().SetRoleInstance(hostname)
	}

	return context
}

// Parse connection string into TelemetryConfiguration. Handles passing just an instrumentation key as well.
func parseConnectionString(connectionStringOrInstrumentationKey string) *TelemetryConfiguration {
	parts := strings.Split(connectionStringOrInstrumentationKey, ";")

	if len(parts) == 1 {
		// If there are no semi-colons, parse it as an instrumentation key
		return &TelemetryConfiguration{
			InstrumentationKey: connectionStringOrInstrumentationKey,
			EndpointUrl:        "https://dc.services.visualstudio.com/v2/track",
		}
	}

	var instrumentationKey string
	var endpointUrl string
	for _, part := range parts {
		kv := strings.Split(part, "=")
		if len(kv) == 2 {
			switch kv[0] {
			case "InstrumentationKey":
				instrumentationKey = kv[1]
			case "IngestionEndpoint":
				endpointUrl = kv[1]
			}
		}
	}

	return &TelemetryConfiguration{
		InstrumentationKey: instrumentationKey,
		EndpointUrl:        strings.Trim(endpointUrl, "/") + "/v2/track",
	}
}
