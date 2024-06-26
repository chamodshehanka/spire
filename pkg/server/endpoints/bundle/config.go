package bundle

import (
	"net"
	"time"

	"github.com/spiffe/spire/pkg/common/diskcertmanager"
)

type EndpointConfig struct {
	// Address is the address on which to serve the federation bundle endpoint.
	Address *net.TCPAddr

	// ACME is the ACME configuration for the bundle endpoint.
	// If unset, the bundle endpoint will use SPIFFE auth.
	ACME *ACMEConfig

	DiskCertManager *diskcertmanager.DiskCertManager

	RefreshHint time.Duration
}
