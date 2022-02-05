package proxy

import (
	"context"
	"fmt"
	"net"
	"sync"
	"time"

	"github.com/xjasonlyu/tun2socks/v2/log"
	M "github.com/xjasonlyu/tun2socks/v2/metadata"
	"github.com/xjasonlyu/tun2socks/v2/proxy/proto"
	"golang.org/x/crypto/ssh"
)

type SSH struct {
	*Base

	user        string
	pass        string
	client      *ssh.Client
	mux         sync.Mutex
	alive       bool
	retryWaters []chan error
}

func NewSSH(address, username, password string) (*SSH, error) {
	// ssh 没有连接池，所以每次都需要重新建立连接
	client, err := newSSHConn(username, password, address)
	if err != nil {
		return nil, err
	}

	proxy := &SSH{
		Base: &Base{
			addr:  address,
			proto: proto.SSH,
		},
		user:        username,
		pass:        password,
		client:      client,
		alive:       true,
		retryWaters: make([]chan error, 0),
	}
	go proxy.Ping()
	return proxy, nil
}

func (S *SSH) getAlive() bool {
	S.mux.Lock()
	defer S.mux.Unlock()
	return S.alive
}

func (S *SSH) setAlive(flag bool) {
	S.alive = flag
}

func (S *SSH) DialContext(ctx context.Context, metadata *M.Metadata) (net.Conn, error) {
	// TODO: ssh 没有连接池
	var err error
	if !S.getAlive() {
		// 重新连接ssh
		if err := <-S.Reconnect(); err != nil {
			return nil, err
		}
	}
	c, err := S.client.Dial("tcp", metadata.DestinationAddress())
	if err != nil {
		return nil, err
	}

	return c, nil
}

func (S *SSH) Reconnect() chan error {
	S.mux.Lock()
	isReconnecting := len(S.retryWaters) > 0
	errChan := make(chan error, 0)
	S.retryWaters = append(S.retryWaters, errChan)
	S.mux.Unlock()

	if isReconnecting {
		return errChan
	}

	go func() {
		client, err := newSSHConn(S.user, S.pass, S.addr)
		if err == nil {
			S.client = client
			S.setAlive(true)
		}

		for _, c := range S.retryWaters {
			c <- err
		}
		S.retryWaters = make([]chan error, 0)
	}()
	return errChan
}

func newSSHConn(user, pass, addr string) (*ssh.Client, error) {
	var err error
	sshConf := &ssh.ClientConfig{
		User:            user,
		Auth:            []ssh.AuthMethod{ssh.Password(pass)},
		HostKeyCallback: ssh.InsecureIgnoreHostKey(),
	}
	client, err := ssh.Dial("tcp", addr, sshConf)
	if err != nil {
		return nil, fmt.Errorf("ssh -> %s %s %s %v", addr, user, pass, err)
	}
	log.Infof("[SSH] new -> %s", addr)
	return client, nil
}

func (S *SSH) CloseSSH() {
	_ = S.client.Close()
	S.setAlive(false)
}

func (S *SSH) Ping() {
	errCount := 0
	for {
		interval := 10
		//log.Infof("[SSH] ping -> %s %v", S.addr, S.getAlive())
		if S.getAlive() {
			_, _, err := S.client.SendRequest("ping", true, nil)
			if err != nil {
				errCount++
				log.Errorf("[SSH] ping -> %s err %v", S.addr, err)
				if errCount >= 2 {
					log.Warnf("[SSH] close -> %s", S.addr)
					S.CloseSSH()
				}
				interval = 5
			} else {
				errCount = 0
			}
		}

		time.Sleep(time.Duration(interval) * time.Second)
	}

}
