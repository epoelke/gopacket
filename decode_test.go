// Copyright (c) 2012 Google, Inc. All rights reserved.
// Copyright (c) 2009-2012 Andreas Krennmair. All rights reserved.

package gopacket

import (
	"fmt"
	"math/rand"
	"os"
	"reflect"
	"testing"
	"time"
)

var testSimpleTCPPacket []byte = []byte{
	0x00, 0x00, 0x0c, 0x9f, 0xf0, 0x20, 0xbc, 0x30, 0x5b, 0xe8, 0xd3, 0x49,
	0x08, 0x00, 0x45, 0x00, 0x01, 0xa4, 0x39, 0xdf, 0x40, 0x00, 0x40, 0x06,
	0x55, 0x5a, 0xac, 0x11, 0x51, 0x49, 0xad, 0xde, 0xfe, 0xe1, 0xc5, 0xf7,
	0x00, 0x50, 0xc5, 0x7e, 0x0e, 0x48, 0x49, 0x07, 0x42, 0x32, 0x80, 0x18,
	0x00, 0x73, 0xab, 0xb1, 0x00, 0x00, 0x01, 0x01, 0x08, 0x0a, 0x03, 0x77,
	0x37, 0x9c, 0x42, 0x77, 0x5e, 0x3a, 0x47, 0x45, 0x54, 0x20, 0x2f, 0x20,
	0x48, 0x54, 0x54, 0x50, 0x2f, 0x31, 0x2e, 0x31, 0x0d, 0x0a, 0x48, 0x6f,
	0x73, 0x74, 0x3a, 0x20, 0x77, 0x77, 0x77, 0x2e, 0x66, 0x69, 0x73, 0x68,
	0x2e, 0x63, 0x6f, 0x6d, 0x0d, 0x0a, 0x43, 0x6f, 0x6e, 0x6e, 0x65, 0x63,
	0x74, 0x69, 0x6f, 0x6e, 0x3a, 0x20, 0x6b, 0x65, 0x65, 0x70, 0x2d, 0x61,
	0x6c, 0x69, 0x76, 0x65, 0x0d, 0x0a, 0x55, 0x73, 0x65, 0x72, 0x2d, 0x41,
	0x67, 0x65, 0x6e, 0x74, 0x3a, 0x20, 0x4d, 0x6f, 0x7a, 0x69, 0x6c, 0x6c,
	0x61, 0x2f, 0x35, 0x2e, 0x30, 0x20, 0x28, 0x58, 0x31, 0x31, 0x3b, 0x20,
	0x4c, 0x69, 0x6e, 0x75, 0x78, 0x20, 0x78, 0x38, 0x36, 0x5f, 0x36, 0x34,
	0x29, 0x20, 0x41, 0x70, 0x70, 0x6c, 0x65, 0x57, 0x65, 0x62, 0x4b, 0x69,
	0x74, 0x2f, 0x35, 0x33, 0x35, 0x2e, 0x32, 0x20, 0x28, 0x4b, 0x48, 0x54,
	0x4d, 0x4c, 0x2c, 0x20, 0x6c, 0x69, 0x6b, 0x65, 0x20, 0x47, 0x65, 0x63,
	0x6b, 0x6f, 0x29, 0x20, 0x43, 0x68, 0x72, 0x6f, 0x6d, 0x65, 0x2f, 0x31,
	0x35, 0x2e, 0x30, 0x2e, 0x38, 0x37, 0x34, 0x2e, 0x31, 0x32, 0x31, 0x20,
	0x53, 0x61, 0x66, 0x61, 0x72, 0x69, 0x2f, 0x35, 0x33, 0x35, 0x2e, 0x32,
	0x0d, 0x0a, 0x41, 0x63, 0x63, 0x65, 0x70, 0x74, 0x3a, 0x20, 0x74, 0x65,
	0x78, 0x74, 0x2f, 0x68, 0x74, 0x6d, 0x6c, 0x2c, 0x61, 0x70, 0x70, 0x6c,
	0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2f, 0x78, 0x68, 0x74, 0x6d,
	0x6c, 0x2b, 0x78, 0x6d, 0x6c, 0x2c, 0x61, 0x70, 0x70, 0x6c, 0x69, 0x63,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2f, 0x78, 0x6d, 0x6c, 0x3b, 0x71, 0x3d,
	0x30, 0x2e, 0x39, 0x2c, 0x2a, 0x2f, 0x2a, 0x3b, 0x71, 0x3d, 0x30, 0x2e,
	0x38, 0x0d, 0x0a, 0x41, 0x63, 0x63, 0x65, 0x70, 0x74, 0x2d, 0x45, 0x6e,
	0x63, 0x6f, 0x64, 0x69, 0x6e, 0x67, 0x3a, 0x20, 0x67, 0x7a, 0x69, 0x70,
	0x2c, 0x64, 0x65, 0x66, 0x6c, 0x61, 0x74, 0x65, 0x2c, 0x73, 0x64, 0x63,
	0x68, 0x0d, 0x0a, 0x41, 0x63, 0x63, 0x65, 0x70, 0x74, 0x2d, 0x4c, 0x61,
	0x6e, 0x67, 0x75, 0x61, 0x67, 0x65, 0x3a, 0x20, 0x65, 0x6e, 0x2d, 0x55,
	0x53, 0x2c, 0x65, 0x6e, 0x3b, 0x71, 0x3d, 0x30, 0x2e, 0x38, 0x0d, 0x0a,
	0x41, 0x63, 0x63, 0x65, 0x70, 0x74, 0x2d, 0x43, 0x68, 0x61, 0x72, 0x73,
	0x65, 0x74, 0x3a, 0x20, 0x49, 0x53, 0x4f, 0x2d, 0x38, 0x38, 0x35, 0x39,
	0x2d, 0x31, 0x2c, 0x75, 0x74, 0x66, 0x2d, 0x38, 0x3b, 0x71, 0x3d, 0x30,
	0x2e, 0x37, 0x2c, 0x2a, 0x3b, 0x71, 0x3d, 0x30, 0x2e, 0x33, 0x0d, 0x0a,
	0x0d, 0x0a,
}

// A few benchmarks for figuring out exactly how fast some underlying Go
// things are.

func BenchmarkTypeAssertion(b *testing.B) {
	var eth LinkLayer = &Ethernet{}
	c := 0
	for i := 0; i < b.N; i++ {
		if _, ok := eth.(*Ethernet); ok {
			c++
		}
	}
}

func BenchmarkMapLookup(b *testing.B) {
	m := map[LayerType]bool{
		LayerTypeTCP:      true,
		LayerTypeEthernet: true,
	}
	for i := 0; i < b.N; i++ {
		_ = m[LayerTypeIPv4]
	}
}

func BenchmarkNilMapLookup(b *testing.B) {
	var m map[LayerType]bool
	for i := 0; i < b.N; i++ {
		_ = m[LayerTypeIPv4]
	}
}

func BenchmarkNilMapLookupWithNilCheck(b *testing.B) {
	var m map[LayerType]bool
	for i := 0; i < b.N; i++ {
		if m != nil {
			_ = m[LayerTypeIPv4]
		}
	}
}

func BenchmarkArrayLookup(b *testing.B) {
	m := make([]bool, 100)
	for i := 0; i < b.N; i++ {
		_ = m[LayerTypeIPv4]
	}
}

// Benchmarks for actual gopacket code

func BenchmarkLayerClassSliceContains(b *testing.B) {
	lc := NewLayerClassSlice([]LayerType{LayerTypeTCP, LayerTypeEthernet})
	for i := 0; i < b.N; i++ {
		_ = lc.Contains(LayerTypeTCP)
	}
}

func BenchmarkLayerClassMapContains(b *testing.B) {
	lc := NewLayerClassMap([]LayerType{LayerTypeTCP, LayerTypeEthernet})
	for i := 0; i < b.N; i++ {
		_ = lc.Contains(LayerTypeTCP)
	}
}

func BenchmarkLazyNoCopyEthLayer(b *testing.B) {
	for i := 0; i < b.N; i++ {
		NewPacket(testSimpleTCPPacket, LinkTypeEthernet, DecodeOptions{Lazy: true, NoCopy: true}).Layer(LayerTypeEthernet)
	}
}

func BenchmarkLazyNoCopyIPLayer(b *testing.B) {
	for i := 0; i < b.N; i++ {
		NewPacket(testSimpleTCPPacket, LinkTypeEthernet, DecodeOptions{Lazy: true, NoCopy: true}).Layer(LayerTypeIPv4)
	}
}

func BenchmarkLazyNoCopyTCPLayer(b *testing.B) {
	for i := 0; i < b.N; i++ {
		NewPacket(testSimpleTCPPacket, LinkTypeEthernet, DecodeOptions{Lazy: true, NoCopy: true}).Layer(LayerTypeTCP)
	}
}

func BenchmarkLazyNoCopyAllLayers(b *testing.B) {
	for i := 0; i < b.N; i++ {
		NewPacket(testSimpleTCPPacket, LinkTypeEthernet, DecodeOptions{Lazy: true, NoCopy: true}).Layers()
	}
}

func BenchmarkDefault(b *testing.B) {
	for i := 0; i < b.N; i++ {
		NewPacket(testSimpleTCPPacket, LinkTypeEthernet, Default)
	}
}

func BenchmarkLazy(b *testing.B) {
	for i := 0; i < b.N; i++ {
		NewPacket(testSimpleTCPPacket, LinkTypeEthernet, Lazy)
	}
}

func BenchmarkNoCopy(b *testing.B) {
	for i := 0; i < b.N; i++ {
		NewPacket(testSimpleTCPPacket, LinkTypeEthernet, NoCopy)
	}
}

func BenchmarkLazyNoCopy(b *testing.B) {
	for i := 0; i < b.N; i++ {
		NewPacket(testSimpleTCPPacket, LinkTypeEthernet, DecodeOptions{Lazy: true, NoCopy: true})
	}
}

func BenchmarkAlloc(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = &TCP{}
	}
}

func BenchmarkFlow(b *testing.B) {
	p := NewPacket(testSimpleTCPPacket, LinkTypeEthernet, DecodeOptions{Lazy: true, NoCopy: true})
	net := p.NetworkLayer()
	for i := 0; i < b.N; i++ {
		net.NetFlow()
	}
}

func BenchmarkEndpoints(b *testing.B) {
	p := NewPacket(testSimpleTCPPacket, LinkTypeEthernet, DecodeOptions{Lazy: true, NoCopy: true})
	flow := p.NetworkLayer().NetFlow()
	for i := 0; i < b.N; i++ {
		flow.Endpoints()
	}
}

func BenchmarkTCPLayerFromDecodedPacket(b *testing.B) {
	b.StopTimer()
	p := NewPacket(testSimpleTcpPacket, LinkTypeEthernet, Default)
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		_ = p.Layer(LayerTypeTCP)
	}
}

func BenchmarkTCPLayerClassFromDecodedPacket(b *testing.B) {
	b.StopTimer()
	p := NewPacket(testSimpleTcpPacket, LinkTypeEthernet, Default)
	lc := NewLayerClass([]LayerType{LayerTypeTCP})
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		_ = p.LayerClass(lc)
	}
}

// TestFlowMapKey makes sure a flow and an endpoint can be used as map keys.
func TestFlowMapKey(t *testing.T) {
	_ = map[Flow]bool{}
	_ = map[Endpoint]bool{}
	_ = map[[2]Flow]bool{}
	if NewUDPPortEndpoint(53) != NewUDPPortEndpoint(53) {
		t.Error("Endpoint equality seems to be broken")
	}
}

func BenchmarkCheckEthernetPrefix(b *testing.B) {
	key := [3]byte{5, 5, 5}
	for i := 0; i < b.N; i++ {
		_ = ValidMACPrefixMap[key]
	}
}

func TestDecodeSimpleTCPPacket(t *testing.T) {
	equal := func(desc, want string, got fmt.Stringer) {
		if want != got.String() {
			t.Errorf("%s: got %q want %q", desc, got.String(), want)
		}
	}
	p := NewPacket(testSimpleTCPPacket, LinkTypeEthernet, DecodeOptions{Lazy: true, NoCopy: true})
	if eth := p.LinkLayer(); eth == nil {
		t.Error("No ethernet layer found")
	} else {
		equal("Eth Src", "bc:30:5b:e8:d3:49", eth.LinkFlow().Src())
		equal("Eth Dst", "00:00:0c:9f:f0:20", eth.LinkFlow().Dst())
	}
	if net := p.NetworkLayer(); net == nil {
		t.Error("No net layer found")
	} else if ip, ok := net.(*IPv4); !ok {
		t.Error("Net layer is not IP layer")
	} else {
		equal("IP Src", "172.17.81.73", net.NetFlow().Src())
		equal("IP Dst", "173.222.254.225", net.NetFlow().Dst())
		want := &IPv4{
			Version:    4,
			IHL:        5,
			TOS:        0,
			Length:     420,
			Id:         14815,
			Flags:      0x02,
			FragOffset: 0,
			TTL:        64,
			Protocol:   6,
			Checksum:   0x555A,
			SrcIP:      []byte{172, 17, 81, 73},
			DstIP:      []byte{173, 222, 254, 225},
		}
		if !reflect.DeepEqual(ip, want) {
			t.Errorf("IP layer mismatch, \ngot  %#v\nwant %#v\n", ip, want)
		}
	}
	if trans := p.TransportLayer(); trans == nil {
		t.Error("No transport layer found")
	} else if tcp, ok := trans.(*TCP); !ok {
		t.Error("Transport layer is not TCP layer")
	} else {
		equal("TCP Src", "50679", trans.AppFlow().Src())
		equal("TCP Dst", "80", trans.AppFlow().Dst())
		want := &TCP{
			SrcPort:    50679,
			DstPort:    80,
			Seq:        0xc57e0e48,
			Ack:        0x49074232,
			DataOffset: 8,
			Flags:      0x18,
			Window:     0x73,
			Checksum:   0xabb1,
			Urgent:     0,
			sPort:      []byte{0xc5, 0xf7},
			dPort:      []byte{0x0, 0x50},
		}
		if !reflect.DeepEqual(tcp, want) {
			t.Errorf("TCP layer mismatch\ngot  %#v\nwant %#v", tcp, want)
		}
	}
	if payload, ok := p.Layer(LayerTypePayload).(*Payload); payload == nil || !ok {
		t.Error("No payload layer found")
	} else {
		if string(payload.Data) != "GET / HTTP/1.1\r\nHost: www.fish.com\r\nConnection: keep-alive\r\nUser-Agent: Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/535.2 (KHTML, like Gecko) Chrome/15.0.874.121 Safari/535.2\r\nAccept: text/html,application/xhtml+xml,application/xml;q=0.9,*/*;q=0.8\r\nAccept-Encoding: gzip,deflate,sdch\r\nAccept-Language: en-US,en;q=0.8\r\nAccept-Charset: ISO-8859-1,utf-8;q=0.7,*;q=0.3\r\n\r\n" {
			t.Error("--- Payload STRING ---\n", string(payload.Data), "\n--- Payload BYTES ---\n", payload.Data)
		}
	}
}

// Makes sure packet payload doesn't display the 6 trailing null of this packet
// as part of the payload.  They're actually the ethernet trailer.
func TestDecodeSmallTCPPacketHasEmptyPayload(t *testing.T) {
	p := NewPacket(
		[]byte{
			0xbc, 0x30, 0x5b, 0xe8, 0xd3, 0x49, 0xb8, 0xac, 0x6f, 0x92, 0xd5, 0xbf,
			0x08, 0x00, 0x45, 0x00, 0x00, 0x28, 0x00, 0x00, 0x40, 0x00, 0x40, 0x06,
			0x3f, 0x9f, 0xac, 0x11, 0x51, 0xc5, 0xac, 0x11, 0x51, 0x49, 0x00, 0x63,
			0x9a, 0xef, 0x00, 0x00, 0x00, 0x00, 0x2e, 0xc1, 0x27, 0x83, 0x50, 0x14,
			0x00, 0x00, 0xc3, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
		}, LinkTypeEthernet, Default)

	if payload := p.Layer(LayerTypePayload); payload != nil {
		t.Error("Payload found for empty TCP packet")
	}
}

func TestDecodeVLANPacket(t *testing.T) {
	p := NewPacket(
		[]byte{
			0x00, 0x10, 0xdb, 0xff, 0x10, 0x00, 0x00, 0x15, 0x2c, 0x9d, 0xcc, 0x00,
			0x81, 0x00, 0x01, 0xf7, 0x08, 0x00, 0x45, 0x00, 0x00, 0x28, 0x29, 0x8d,
			0x40, 0x00, 0x7d, 0x06, 0x83, 0xa0, 0xac, 0x1b, 0xca, 0x8e, 0x45, 0x16,
			0x94, 0xe2, 0xd4, 0x0a, 0x00, 0x50, 0xdf, 0xab, 0x9c, 0xc6, 0xcd, 0x1e,
			0xe5, 0xd1, 0x50, 0x10, 0x01, 0x00, 0x5a, 0x74, 0x00, 0x00, 0x00, 0x00,
			0x00, 0x00, 0x00, 0x00,
		}, LinkTypeEthernet, Default)
	if err := p.ErrorLayer(); err != nil {
		t.Error("Error while parsing vlan packet:", err)
	}
	if vlan := p.Layer(LayerTypeDot1Q); vlan == nil {
		t.Error("Didn't detect vlan")
	} else if _, ok := vlan.(*Dot1Q); !ok {
		t.Error("LayerTypeDot1Q layer is not a Dot1Q object")
	}
	for i, l := range p.Layers() {
		t.Logf("Layer %d: %#v", i, l)
	}
	want := []LayerType{LayerTypeEthernet, LayerTypeDot1Q, LayerTypeIPv4, LayerTypeTCP}
	if len(p.Layers()) != len(want) {
		t.Fatal("Incorrect number of headers:", len(p.Layers()))
	}
	for i, l := range p.Layers() {
		if l.LayerType() != want[i] {
			t.Errorf("At index %d, got layer type %s, want %s", i, l.LayerType(), want[i])
		}
	}
}

func TestDecoderSecurity(t *testing.T) {
	seed := time.Now().UnixNano()
	fmt.Fprintf(os.Stderr, "If you see a crash here, it's serious business.  Report it!\n"+
		"Send this number with any crash reports: %v\n", seed)
	r := rand.New(rand.NewSource(seed))

	testCases := []struct {
		s string
		d decoderFunc
	}{
		{"ARP", decodeARP},
		{"Dot1Q", decodeDot1Q},
		{"Ethernet", decodeEthernet},
		{"ICMP", decodeICMP},
		{"IPv4", decodeIPv4},
		{"IPv6", decodeIPv6},
		{"PPP", decodePPP},
		{"TCP", decodeTCP},
		{"UDP", decodeUDP},
		{"MPLS", decodeMPLS},
		{"PPPoE", decodePPPoE},
	}
	for _, tc := range testCases {
		// Fuzz-test the decoder tc.d by feeding it random inputs.
		// We're fine with errors occurring here... what we're looking
		// for is actual program crashes, which Shouldn't Happen (tm) due
		// to golang slice range checking... but we're paranoid.
		t.Logf("Testing %s", tc.s)
		for i := 0; i < 1000; i++ {
			b := make([]byte, 0, 256)
			for r.Int()%100 != 0 {
				b = append(b, byte(r.Int()))
			}
			NewPacket(b, tc.d, Default)
		}
	}
	fmt.Fprintln(os.Stderr, "No crash to see here... continuing with testing")
}
