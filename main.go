package main

import (
	"flag"
	"github.com/mpetavy/common"
)

var (
	discoverAddress *string
	discoverTimeout *int
	discoverUID     *string
	discoverInfo    *string

	discoverServer *common.Server
)

func init() {
	common.Init(common.Title(), "1.0.0", "2019", "discover demo server", "mpetavy", common.APACHE, "https://github.com/mpetavy/"+common.Title(), true, start, stop, nil, 0)

	discoverAddress = flag.String("c", ":9999", "discover address")
	discoverTimeout = flag.Int("t", 1000, "discover timeout")
	discoverUID = flag.String("uid", "discover-uid", "discover uid")
	discoverInfo = flag.String("info", "discover-info", "discover info")
}

func start() error {
	if common.IsRunningAsService() {
		var err error

		discoverServer, err = common.New(*discoverAddress, common.MsecToDuration(*discoverTimeout), *discoverUID, *discoverInfo)
		if err != nil {
			return err
		}

		return discoverServer.Start()
	}

	discoveredIps, err := common.Discover(*discoverAddress, common.MsecToDuration(*discoverTimeout), *discoverUID)
	if err != nil {
		return err
	}

	for k, v := range discoveredIps {
		common.Info("discovered #%s: %s", k, v)
	}

	return nil
}

func stop() error {
	if common.IsRunningAsService() {
		return discoverServer.Stop()
	}

	return nil
}

func main() {
	defer common.Done()

	common.Run(nil)
}
