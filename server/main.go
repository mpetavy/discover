package main

import (
	"flag"
	"time"

	"github.com/mpetavy/common"
	"github.com/mpetavy/discover"
)

var (
	discoverPort    *string
	discoverTimeout *time.Duration
	discoverUID     *string
	discoverInfo    *string

	discoverServer *discover.Server
)

func init() {
	common.Init(common.Title(), "1.0.0", "2019", "discover demo server", "mpetavy", common.APACHE, "https://github.com/mpetavy/"+common.Title(), true, start, stop, nil, 0)

	discoverPort = flag.String("port", ":9999", "discover address")
	discoverTimeout = flag.Duration("timeout", time.Millisecond*1000, "discover timeout")
	discoverUID = flag.String("uid", "my-uid", "discover uid")
	discoverInfo = flag.String("info", "my-info", "discover info")
}

func start() error {
	var err error

	discoverServer, err = discover.New(*discoverPort, *discoverTimeout, *discoverUID, *discoverInfo)
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
