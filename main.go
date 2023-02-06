package main

import (
	"flag"
	"github.com/mpetavy/common"
)

var (
	LDFLAG_DEVELOPER = "mpetavy"                             // will be replaced with ldflag
	LDFLAG_HOMEPAGE  = "https://github.com/mpetavy/discover" // will be replaced with ldflag
	LDFLAG_LICENSE   = common.APACHE                         // will be replaced with ldflag
	LDFLAG_VERSION   = "1.0.10"                              // will be replaced with ldflag
	LDFLAG_EXPIRE    = ""                                    // will be replaced with ldflag
	LDFLAG_GIT       = ""                                    // will be replaced with ldflag
	LDFLAG_BUILD     = ""                                    // will be replaced with ldflag

	discoverClient  *string
	discoverServer  *string
	discoverTimeout *int
	discoverUID     *string
	discoverInfo    *string

	server *common.DiscoverServer
)

// Server
// Without host: go run . -s :5000 -uid test -log.verbose
// With host (use broadcast ip): go run . -s 192.168.1.255:5000 -uid test -log.verbose
//
// Client
// Without host: go run . -c :5000 -uid test -log.verbose
// With host (use ip): go run . -s 192.168.1.3:5000 -uid test -log.verbose

func init() {
	common.Init(LDFLAG_VERSION, LDFLAG_GIT, LDFLAG_BUILD, "2019", "service discovery", LDFLAG_DEVELOPER, LDFLAG_HOMEPAGE, LDFLAG_LICENSE, nil, start, stop, run, 0)

	discoverClient = flag.String("c", "", "discover client")
	discoverServer = flag.String("s", "", "discover server")
	discoverTimeout = flag.Int("t", 1000, "discover timeout")
	discoverUID = flag.String("uid", "discover", "discover uid")
	discoverInfo = flag.String("info", "<host>", "discover info")

	common.Events.NewFuncReceiver(common.EventFlagsParsed{}, func(event common.Event) {
		if *discoverClient != "" {
			common.App().StartFunc = nil
			common.App().StopFunc = nil
		} else {
			common.App().RunFunc = nil
			if !common.IsRunningAsService() {
				common.Panic(flag.Set(common.FlagNameService, common.SERVICE_SIMULATE))
			}
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

	common.Info("Listen on %v...", *discoverServer)

	return nil
}

func run() error {
	list, err := common.Discover(*discoverClient, common.MillisecondToDuration(*discoverTimeout), *discoverUID)
	if err != nil {
		return err
	}

	for i, v := range list {
		common.Info("discovered #%d: %s", i, v)
	}

	return nil
}

func stop() error {
	return server.Stop()
}

func main() {
	defer common.Done()

	common.Run([]string{"c|s", "uid"})
}
