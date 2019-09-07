package discover

import (
	"fmt"
	"log"
	"net"
	"sync"
	"time"

	"github.com/mpetavy/common"
)

type Server struct {
	address  string
	timeout  time.Duration
	uid      string
	info     string
	quitCh   chan struct{}
	listener net.PacketConn
}

const (
	maxInfoLength = 1024
)

func New(address string, timeout time.Duration, uid string, info string) (*Server, error) {
	if len(info) > maxInfoLength {
		return nil, fmt.Errorf("max UDP info length exceeded. max length expected: %d received: %d", maxInfoLength, len(info))
	}

	return &Server{address: address, timeout: timeout, uid: uid, info: info}, nil
}

func (server *Server) Start() error {
	common.DebugFunc("discover server: %+v", *server)

	if server.quitCh != nil {
		return fmt.Errorf("Server already started")
	}

	b := make([]byte, maxInfoLength)

	var err error

	server.listener, err = net.ListenPacket("udp4", server.address)
	if err != nil {
		return err
	}

	server.quitCh = make(chan struct{})

	go func() {
		for !common.AppStopped() {
			select {
			case <-server.quitCh:
				break
			default:
				err := server.listener.SetDeadline(common.DeadlineByDuration(server.timeout))
				if err != nil {
					common.Error(err)

					break
				}

				n, peer, err := server.listener.ReadFrom(b)
				if err != nil {
					if err, ok := err.(net.Error); ok && err.Timeout() {
						break
					} else {
						common.DebugError(err)

						break
					}
				}

				receivedUID := string(b[:n])

				common.Debug("received UDP broadcast from %+v: %s\n", peer, receivedUID)

				if receivedUID != server.uid {
					common.Debug("not matching uid, expected: %s received: %s -> ignore", server.uid, receivedUID)

					break
				}

				common.Debug("answer positive discover with info %s to %+v", server.info, peer)

				if _, err := server.listener.WriteTo([]byte(server.info), peer); err != nil {
					common.Error(err)
				}
			}
		}
	}()

	return nil
}

func (server *Server) Stop() error {
	common.DebugError(server.listener.Close())

	if server.quitCh == nil {
		return fmt.Errorf("Server already stopped")
	}

	close(server.quitCh)

	server.quitCh = nil

	common.DebugFunc("discover server: %+v", *server)

	return nil
}

func Discover(address string, timeout time.Duration, uid string) (map[string]string, error) {
	common.DebugFunc("discover uid: %s", uid)

	localIps, err := common.FindActiveIPs()
	if err != nil {
		panic(err)
	}

	var wg sync.WaitGroup

	discoveredIps := make(map[string]string)

	c, err := net.ListenPacket("udp4", ":0")
	if err != nil {
		return discoveredIps, err
	}
	defer func() {
		common.IgnoreError(c.Close())
	}()

	_, discoverPort, err := net.SplitHostPort(address)
	if err != nil {
		return discoveredIps, err
	}

	for _, localIp := range localIps {
		ip, ipNet, err := net.ParseCIDR(localIp)
		if err != nil {
			panic(err)
		}

		ip = ip.To4()

		if ip == nil {
			continue
		}

		wg.Add(1)

		go func(ip net.IP, ipNet *net.IPNet) {
			defer wg.Done()

			ones, bits := ipNet.Mask.Size()
			mask := net.CIDRMask(ones, bits)

			broadcast := net.IP(make([]byte, 4))
			for i := range ip {
				broadcast[i] = ip[i] | ^mask[i]
			}

			common.Debug("UDP broadcast: %v for ip: %v on port: %s", broadcast.String(), ipNet, discoverPort)

			dst, err := net.ResolveUDPAddr("udp4", broadcast.String()+":"+discoverPort)
			if err != nil {
				log.Fatal(err)
			}

			if _, err := c.WriteTo([]byte(uid), dst); err != nil {
				log.Fatal(err)
			}
		}(ip, ipNet)
	}

	wg.Wait()

	common.Debug("reading answers ...")

	b := make([]byte, maxInfoLength)
	for {
		err := c.SetDeadline(common.DeadlineByDuration(timeout))
		if err != nil {
			common.Error(err)

			break
		}

		n, peer, err := c.ReadFrom(b)
		if err != nil {
			if err, ok := err.(net.Error); ok && err.Timeout() {
				break
			} else {
				log.Fatal(err)
			}
		}

		client := peer.String()
		info := string(b[:n])

		discoveredIps[client] = info

		common.Debug("%d bytes read from %+v: %s\n", n, client, info)
	}

	return discoveredIps, nil
}
