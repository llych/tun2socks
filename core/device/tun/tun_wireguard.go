// +build !linux !amd64,!arm64

package tun

import (
	"fmt"

	"github.com/xjasonlyu/tun2socks/core/device"
	"github.com/xjasonlyu/tun2socks/core/device/rwbased"

	"golang.zx2c4.com/wireguard/tun"
)

type TUN struct {
	*rwbased.Endpoint

	nt   *tun.NativeTun
	mtu  uint32
	name string
}

func Open(opts ...Option) (device.Device, error) {
	t := &TUN{}

	for _, opt := range opts {
		opt(t)
	}

	nt, err := tun.CreateTUN(t.name, int(t.mtu))
	if err != nil {
		return nil, fmt.Errorf("create tun: %w", err)
	}
	t.nt = nt.(*tun.NativeTun)

	mtu, err := nt.MTU()
	if err != nil {
		return nil, fmt.Errorf("get mtu: %w", err)
	}
	t.mtu = uint32(mtu)

	ep, err := rwbased.New(t, uint32(mtu))
	if err != nil {
		return nil, fmt.Errorf("create endpoint: %w", err)
	}
	t.Endpoint = ep

	return t, nil
}

func (t *TUN) Name() string {
	name, _ := t.nt.Name()
	return name
}

func (t *TUN) Close() error {
	return t.nt.Close()
}