package addrutil

import (
	"testing"

	ma "github.com/multiformats/go-multiaddr"
)

func TestSubtractAndNegFilter(t *testing.T) {
	localhost := newMultiaddr(t, "/ip4/127.0.0.1/tcp/1234")
	private := newMultiaddr(t, "/ip4/192.168.1.1/tcp/1234")
	toRemoveFilter := SubtractFilter(
		localhost,
		newMultiaddr(t, "/ip6/::1/tcp/1234"),
		newMultiaddr(t, "/ip4/1.2.3.4/udp/1234/utp"),
	)
	result := FilterAddrs([]ma.Multiaddr{localhost, private}, toRemoveFilter)
	if len(result) != 1 || !result[0].Equal(private) {
		t.Errorf("Expected only one remaining address: %s", private.String())
	}

	// Negate original filter
	result = FilterAddrs([]ma.Multiaddr{localhost, private}, FilterNeg(toRemoveFilter))
	if len(result) != 1 || !result[0].Equal(localhost) {
		t.Errorf("Expected only one remaining address: %s", localhost.String())
	}
}
