package fchost

import (
	"github.com/libp2p/go-libp2p"
	"github.com/libp2p/go-libp2p-core/peer"
	"github.com/libp2p/go-libp2p/config"
	"github.com/multiformats/go-multiaddr"
)

var (
	addrs = []string{
		"/dns4/bootstrap-0-sin.fil-test.net/tcp/1347/p2p/12D3KooWPdUquftaQvoQEtEdsRBAhwD6jopbF2oweVTzR59VbHEd",
		"/ip4/86.109.15.57/tcp/1347/p2p/12D3KooWPdUquftaQvoQEtEdsRBAhwD6jopbF2oweVTzR59VbHEd",
		"/dns4/bootstrap-0-dfw.fil-test.net/tcp/1347/p2p/12D3KooWQSCkHCzosEyrh8FgYfLejKgEPM5VB6qWzZE3yDAuXn8d",
		"/ip4/139.178.84.45/tcp/1347/p2p/12D3KooWQSCkHCzosEyrh8FgYfLejKgEPM5VB6qWzZE3yDAuXn8d",
		"/dns4/bootstrap-0-fra.fil-test.net/tcp/1347/p2p/12D3KooWEXN2eQmoyqnNjde9PBAQfQLHN67jcEdWU6JougWrgXJK",
		"/ip4/136.144.49.17/tcp/1347/p2p/12D3KooWEXN2eQmoyqnNjde9PBAQfQLHN67jcEdWU6JougWrgXJK",
		"/dns4/bootstrap-1-sin.fil-test.net/tcp/1347/p2p/12D3KooWLmJkZd33mJhjg5RrpJ6NFep9SNLXWc4uVngV4TXKwzYw",
		"/ip4/86.109.15.123/tcp/1347/p2p/12D3KooWLmJkZd33mJhjg5RrpJ6NFep9SNLXWc4uVngV4TXKwzYw",
		"/dns4/bootstrap-1-dfw.fil-test.net/tcp/1347/p2p/12D3KooWGXLHjiz6pTRu7x2pkgTVCoxcCiVxcNLpMnWcJ3JiNEy5",
		"/ip4/139.178.86.3/tcp/1347/p2p/12D3KooWGXLHjiz6pTRu7x2pkgTVCoxcCiVxcNLpMnWcJ3JiNEy5",
		"/dns4/bootstrap-1-fra.fil-test.net/tcp/1347/p2p/12D3KooW9szZmKttS9A1FafH3Zc2pxKwwmvCWCGKkRP4KmbhhC4R",
		"/ip4/136.144.49.131/tcp/1347/p2p/12D3KooW9szZmKttS9A1FafH3Zc2pxKwwmvCWCGKkRP4KmbhhC4R",
	}
)

func getBootstrapPeers() []peer.AddrInfo {
	maddrs := make([]multiaddr.Multiaddr, len(addrs))
	for i, addr := range addrs {
		var err error
		maddrs[i], err = multiaddr.NewMultiaddr(addr)
		if err != nil {
			panic(err)
		}
	}
	peers, err := peer.AddrInfosFromP2pAddrs(maddrs...)
	if err != nil {
		panic(err)
	}
	return peers
}

func getDefaultOpts() []config.Option {
	return []config.Option{libp2p.Defaults}
}
