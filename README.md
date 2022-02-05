![tun2socks](docs/logo.png)

[![GitHub Workflow][1]](https://github.com/xjasonlyu/tun2socks/actions)
[![Go Version][2]](https://github.com/xjasonlyu/tun2socks/blob/main/go.mod)
[![Go Report][3]](https://goreportcard.com/badge/github.com/xjasonlyu/tun2socks)
[![GitHub License][4]](https://github.com/xjasonlyu/tun2socks/blob/main/LICENSE)
[![Releases][5]](https://github.com/xjasonlyu/tun2socks/releases)

[1]: https://img.shields.io/github/workflow/status/xjasonlyu/tun2socks/Go?style=flat-square
[2]: https://img.shields.io/github/go-mod/go-version/xjasonlyu/tun2socks/main?style=flat-square
[3]: https://goreportcard.com/badge/github.com/xjasonlyu/tun2socks?style=flat-square
[4]: https://img.shields.io/github/license/xjasonlyu/tun2socks?style=flat-square
[5]: https://img.shields.io/github/v/release/xjasonlyu/tun2socks?include_prereleases&style=flat-square

English | [简体中文](README_ZH.md)

## 2022-02-06
增加ssh协议(单连接，性能有些差)
```bash
# -net 为自动添加路由表(网段路由到tun), 密码如有特殊字符，可以先 urlencode 再使用
sudo ./tun2socks -device tun://utun3 -proxy ssh://用户:密码@ip:端口 -net 100.0.0.1/8,10.0.0.1/8,192.168.2.0/24
```
## Features

- **Network Support**
  - Dualstack: `IPv4/IPv6`
  - Forwarder: `TCP/UDP`
  - Ping Echo: `ICMP`
- **Platform Support**
  - Linux
  - MacOS
  - Windows
  - FreeBSD
  - OpenBSD
- **Proxy Protocol**
  - HTTP
  - Socks4
  - Socks5
  - Shadowsocks
- **Extra Feature**
  - Improved stability without CGO
  - Optimized UDP transmission for game
  - Performed with >2.5Gbps throughput
  - TCP/IP stack powered by **[gVisor](https://github.com/google/gvisor)**

## Documentation

Docs and quick start guides can be found at [Github Wiki](https://github.com/xjasonlyu/tun2socks/wiki).

## Community

Welcome and feel free to ask any questions at [Github Discussions](https://github.com/xjasonlyu/tun2socks/discussions).

## Credits

- [Dreamacro/clash](https://github.com/Dreamacro/clash) - A rule-based tunnel in Go
- [google/gvisor](https://github.com/google/gvisor) - Application Kernel for Containers
- [wireguard-go](https://git.zx2c4.com/wireguard-go) - Go Implementation of WireGuard

## License

[GPL-3.0](https://github.com/xjasonlyu/tun2socks/blob/main/LICENSE)

[![FOSSA Status](https://app.fossa.com/api/projects/git%2Bgithub.com%2Fxjasonlyu%2Ftun2socks.svg?type=large)](https://app.fossa.com/projects/git%2Bgithub.com%2Fxjasonlyu%2Ftun2socks?ref=badge_large)


## Stargazers over time

[![Stargazers over time](https://starchart.cc/xjasonlyu/tun2socks.svg)](https://starchart.cc/xjasonlyu/tun2socks)
