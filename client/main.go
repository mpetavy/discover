package main

import (
	"flag"
	"time"

	"github.com/mpetavy/common"
	"github.com/mpetavy/discover"
)

var (
	discoverAddress *string
	discoverTimeout *time.Duration
	discoverUID     *string
)

func init() {
	common.Init(common.Title(), "1.0.0", "2019", "discover demo client", "mpetavy", common.APACHE, "https://github.com/mpetavy/"+common.Title(), true, start, nil, nil, 0)

	discoverAddress = flag.String("c", ":9999", "discover address")
	discoverTimeout = flag.Duration("t", time.Millisecond*1000, "discover read timeout")
	discoverUID = flag.String("uid", "my-uid", "discover uid")
}

func start() error {
	discoveredIps, err := discover.Discover(*discoverAddress, *discoverTimeout, *discoverUID)
	if err != nil {
		return err
	}

	for k, v := range discoveredIps {
		common.Info("discovered #%s: %s", k, v)
	}

	return nil
}

func main() {
	defer common.Done()

	common.Run(nil)
}
