package lib

import "fmt"

type TransportType string

const (
	HTTP  TransportType = "http"
	GRPC  TransportType = "grpc"
	RMQ   TransportType = "rmq"
	Kafka TransportType = "kafka"
)

// TransportConfig holds the configuration for the transport
type TransportConfig struct {
	Type        TransportType      // Transport type (HTTP, GRPC, RMQ, Kafka)
	Address     string             // Address to bind the transport
	ExtraConfig map[string]any     // Extra configuration parameters
}

func (tc TransportConfig) String() string {
	return fmt.Sprintf("Type: %s, Address: %s, ExtraConfig: %+v", tc.Type, tc.Address, tc.ExtraConfig)
}