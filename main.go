package main

import (
	"flag"
	"fmt"
	"github.com/mpetavy/common"
)

var (
	discoverClient  *string
	discoverServer  *string
	discoverTimeout *int
	discoverUID     *string
	discoverInfo    *string

	server *common.DiscoverServer
)

func init() {
	common.Init(true, "1.0.0", "", "", "2019", "service discovery", "mpetavy", fmt.Sprintf("https://github.com/mpetavy/%s", common.Title()), common.APACHE, nil, start, stop, run, 0)

	discoverClient = flag.String("c", "", "discover client")
	discoverServer = flag.String("s", "", "discover server")
	discoverTimeout = flag.Int("t", 1000, "discover timeout")
	discoverUID = flag.String("uid", "", "discover uid")
	discoverInfo = flag.String("info", "", "discover info")

	common.Events.NewFuncReceiver(common.EventFlagsParsed{}, func(event common.Event) {
		if *discoverClient != "" {
			common.App().StartFunc = nil
			common.App().StopFunc = nil
		} else {
			common.App().RunFunc = nil
		}
	})
}

func start() error {
	var err error

	server, err = common.NewDiscoverServer(*discoverServer, common.MillisecondToDuration(*discoverTimeout), *discoverUID, *discoverInfo)
	if err != nil {
		return err
	}

	err = server.Start()
	if err != nil {
		return err
	}

	return nil
}

func run() error {
	discoveredIps, err := common.Discover(*discoverClient, common.MillisecondToDuration(*discoverTimeout), *discoverUID)
	if err != nil {
		return err
	}

	for k, v := range discoveredIps {
		common.Info("discovered #%s info: %s", k, v)
	}

	return nil
}

func stop() error {
	return server.Stop()
}

func main() {
	defer common.Done()

	common.Run([]string{"c|s"})
}
