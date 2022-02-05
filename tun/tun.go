package tun

import (
	"fmt"
	"github.com/xjasonlyu/tun2socks/v2/common/command"
	"net/url"
	"strings"
)
import "github.com/xjasonlyu/tun2socks/v2/log"

func CreateTun(device, gateway string) error {
	tunName := getTunName(device)
	log.Infof("[TUN] create tun -> %s(%s)", tunName, gateway)
	cmd := fmt.Sprintf(`ifconfig %s %s %s netmask 255.255.255.0`, tunName, gateway, gateway)
	log.Debugf("[TUN] create tun -> command: %s", cmd)
	output, err := command.Run(cmd)
	if err != nil {
		log.Errorf("[tun] create fail: %s", output)
		return err
	}
	return nil
}

func CreateRoute(nets, gateway string) error {
	netList := strings.Split(nets, ",")
	for _, net := range netList {
		net = strings.TrimSpace(net)
		log.Infof("[ROUTE] create route -> %s <-> %s", net, gateway)
		cmd := fmt.Sprintf(`route -n add -net %s -gateway %s`, net, gateway)
		log.Debugf("[ROUTE] create route -> command: %s", cmd)
		output, err := command.Run(cmd)
		if err != nil {
			log.Errorf("[tun] create route: %s", output)
			return err
		}
	}

	return nil
}

func getTunName(device string) string {
	u, _ := url.Parse(device)
	return u.Host
}
