package main

import (
	"flag"
	"github.com/xjasonlyu/tun2socks/v2/engine"
	"github.com/xjasonlyu/tun2socks/v2/log"
	"github.com/xjasonlyu/tun2socks/v2/tun"
	"os"
	"os/signal"
	"syscall"
	"time"

	_ "github.com/xjasonlyu/tun2socks/v2/common/automaxprocs"
)

var key = new(engine.Key)

func init() {
	flag.IntVar(&key.Mark, "fwmark", 0, "Set firewall MARK (Linux only)")
	flag.IntVar(&key.MTU, "mtu", 0, "Set device maximum transmission unit (MTU)")
	flag.IntVar(&key.UDPTimeout, "udp-timeout", 0, "Set timeout for each UDP session")
	flag.BoolVar(&key.Version, "version", false, "Show version information and quit")
	flag.StringVar(&key.Config, "config", "", "YAML format configuration file")
	flag.StringVar(&key.Device, "device", "", "Use this device [driver://]name")
	flag.StringVar(&key.Interface, "interface", "", "Use network INTERFACE (Linux/MacOS only)")
	flag.StringVar(&key.LogLevel, "loglevel", "info", "Log level [debug|info|warning|error|silent]")
	flag.StringVar(&key.Proxy, "proxy", "", "Use this proxy [protocol://]host[:port]")
	flag.StringVar(&key.Stats, "stats", "", "HTTP statistic server listen address")
	flag.StringVar(&key.Token, "token", "", "HTTP statistic server auth token")
	flag.StringVar(&key.Net, "net", "", "路由表 例: 172.17.0.0/24 多个逗号分隔")
	flag.StringVar(&key.Gateway, "gateway", "10.254.17.1", "tun ip, 也作为网关")
	flag.Parse()
}

func main() {
	engine.Insert(key)

	checkErr := func(msg string, f func() error) {
		if err := f(); err != nil {
			log.Fatalf("Failed to %s: %v", msg, err)
		}
	}

	checkErr("start engine", engine.Start)

	time.Sleep(1 * time.Second)

	checkErr("tun", func() error {
		return tun.CreateTun(key.Device, key.Gateway)
	})

	if key.Net != "" {
		checkErr("route", func() error {
			return tun.CreateRoute(key.Net, key.Gateway)
		})
	}

	defer checkErr("stop engine", engine.Stop)

	sigCh := make(chan os.Signal, 1)
	signal.Notify(sigCh, syscall.SIGINT, syscall.SIGTERM)
	<-sigCh
}
