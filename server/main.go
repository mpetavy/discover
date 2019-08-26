package main

import (
	"flag"
	"github.com/mpetavy/common"
	"github.com/mpetavy/discover"
	"time"
)

var (
	discoverAddress     *string
	discoverReadTimeout *time.Duration
	discoverUID         *string
	discoverInfo        *string

	discoverServer *discover.Server
)

func init() {
	common.Init(common.Title(), "1.0.0", "2019", "discover demo server", "mpetavy", common.APACHE, "https://github.com/mpetavy/"+common.Title(), true, start, stop, nil, 0)

	discoverAddress = flag.String("discover.address", ":9999", "discover address")
	discoverReadTimeout = flag.Duration("discover.readtimeout", time.Millisecond*1000, "discover read timeout")
	discoverUID = flag.String("discover.uid", "my-uid", "discover uid")
	discoverInfo = flag.String("discover.info", "my-info", "discover info")
}

func start() error {
	var err error

	discoverServer, err = discover.New(*discoverAddress, *discoverReadTimeout, *discoverUID, *discoverInfo)
	if err != nil {
		return err
	}

	return discoverServer.Start()
}

func stop() error {
	return discoverServer.Stop()
}

func main() {
	defer common.Done()

	common.Run(nil)
}
