//go:build windows
// +build windows

package netproviders

import (
	"github.com/moby/buildkit/util/bklog"
	"github.com/moby/buildkit/util/network"
)

func getHostProvider() (network.Provider, bool) {
	return nil, false
}

func getFallback() (network.Provider, string) {
	bklog.L.Warn("using null network as the default")
	return network.NewNoneProvider(), ""
}
