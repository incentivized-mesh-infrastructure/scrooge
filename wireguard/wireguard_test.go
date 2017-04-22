package wireguard

import (
	"fmt"
	"net"
	"testing"

	"github.com/agl/ed25519"
	"github.com/incentivized-mesh-infrastructure/scrooge/types"
)

var (
	iface = &net.Interface{
		Name: "foo0",
	}
	account1 = &types.Account{
		PublicKey:  [ed25519.PublicKeySize]byte{0x3b, 0xee, 0xb8, 0xd0, 0x2, 0x7c, 0x31, 0x38, 0x1a, 0xc2, 0x28, 0xdc, 0xe1, 0x23, 0x2d, 0x62, 0x9c, 0xcd, 0x68, 0x1e, 0xde, 0x7d, 0x45, 0xbb, 0xc0, 0xec, 0x10, 0x87, 0x94, 0x8d, 0xfe, 0xa},
		PrivateKey: [ed25519.PrivateKeySize]byte{0x45, 0xc2, 0x72, 0x9, 0x8d, 0xc7, 0x63, 0x2f, 0xff, 0xe1, 0x43, 0x1, 0x72, 0x90, 0x8a, 0x6c, 0x34, 0xa2, 0x11, 0x50, 0xf3, 0x2, 0x55, 0xa3, 0xae, 0x4d, 0x1d, 0x8f, 0x9e, 0x1f, 0xa6, 0x58, 0x3b, 0xee, 0xb8, 0xd0, 0x2, 0x7c, 0x31, 0x38, 0x1a, 0xc2, 0x28, 0xdc, 0xe1, 0x23, 0x2d, 0x62, 0x9c, 0xcd, 0x68, 0x1e, 0xde, 0x7d, 0x45, 0xbb, 0xc0, 0xec, 0x10, 0x87, 0x94, 0x8d, 0xfe, 0xa},
		Seqnum:     16,
		// TunnelPublicKey:  "lrWazDvT07U5oOzCA7CRbpG5ULAEXNkMGvNyAhN34E8=",
		// TunnelPrivateKey: "ABuSdM2Z3V5Pc+4G3EtdIC5RN2ksOYFin2IvPMVbu0s=",
	}
	account2 = &types.Account{
		PublicKey:  [ed25519.PublicKeySize]byte{0x9b, 0xbe, 0x22, 0x49, 0xca, 0x84, 0x70, 0xb4, 0xda, 0x9a, 0xed, 0x36, 0xd2, 0xec, 0x62, 0x75, 0x28, 0x7d, 0xac, 0x3d, 0x1, 0x5e, 0x3d, 0xf7, 0xa1, 0x2f, 0xd1, 0xc6, 0xcb, 0x96, 0xa5, 0x86},
		PrivateKey: [ed25519.PrivateKeySize]byte{0xf6, 0x4, 0x2e, 0x29, 0xbe, 0x99, 0xde, 0x68, 0xfc, 0x1b, 0x41, 0x58, 0xe0, 0xc9, 0xab, 0xc6, 0x81, 0xa5, 0x2a, 0x79, 0x76, 0x5a, 0xae, 0x59, 0x79, 0x58, 0x64, 0x5f, 0x14, 0xa3, 0x4a, 0xcb, 0x9b, 0xbe, 0x22, 0x49, 0xca, 0x84, 0x70, 0xb4, 0xda, 0x9a, 0xed, 0x36, 0xd2, 0xec, 0x62, 0x75, 0x28, 0x7d, 0xac, 0x3d, 0x1, 0x5e, 0x3d, 0xf7, 0xa1, 0x2f, 0xd1, 0xc6, 0xcb, 0x96, 0xa5, 0x86},
		Seqnum:     16,
		// TunnelPublicKey:  "94FWZtMpGojuNReJoBKz8KrcCODd+1uNGhzX8aeqKtw=",
		// TunnelPrivateKey: "KVtaPVp8ZqtNrVZk+lxgL3OKj2POSYrT13s4S6EyzCH3gVZm0ykaiO41F4mgErPwqtwI4N37W40aHNfxp6oq3A==",
	}
)

func TestCreateTunnel(t *testing.T) {
	tunnel := types.Tunnel{
		PublicKey:        account2.PublicKey,
		ListenPort:       4500,
		Endpoint:         "localhost:4500",
		VirtualInterface: net.Interface{Name: "foo2"},
	}
	err := CreateTunnel(&tunnel, account1.PrivateKey)
	if err != nil {
		t.Fatal(err)
	}
}

func TestParseConfig(t *testing.T) {
	testConfig := `[Interface]
PrivateKey = yAnz5TF+lXXJte14tji3zlMNq+hd2rYUIgJBgB3fBmk=
ListenPort = 51820

[Peer]
PublicKey = xTIBA5rboUvnH4htodjb6e697QjLERt1NAB4mZqp8Dg=
Endpoint = 192.95.5.67:1234
AllowedIPs = 10.192.122.3/32, 10.192.124.1/24`

	config, err := ParseConfig(testConfig)
	if err != nil {
		t.Fatal(err)
	}

	if config.PrivateKey != "yAnz5TF+lXXJte14tji3zlMNq+hd2rYUIgJBgB3fBmk=" {
		t.Fatal("config.PrivateKey incorrect")
	}

	if config.ListenPort != 51820 {
		t.Fatal("config.ListenPort incorrect")
	}

	if config.Peer.AllowedIPs != "10.192.122.3/32, 10.192.124.1/24" {
		t.Fatal("config.Peer.AllowedIPs[0] incorrect: ", config.Peer.AllowedIPs)
	}

	fmt.Println(config.Peer.AllowedIPs[0])
}
