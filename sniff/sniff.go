// Go Api server
// @jeffotoni
// 2019-01-04

////////////////////////////////////
// Important Note:
// //////////////////////
// You need the libcpa package installed on your machine
// INSTALL libcpa to found pcap.h
// Ubuntu:
// sudo apt-get install libpcap-dev

package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/google/gopacket"
	"github.com/google/gopacket/pcap"
)

func main() {

	inteth := flag.String("i", "", "example: -i eth0")
	if len(os.Args) < 3 {
		flag.PrintDefaults()
		return
	}

	flag.Parse()
	fmt.Println("Run Sniff in: " + *inteth)
	handle, err := pcap.OpenLive(*inteth, 1600, true, pcap.BlockForever)
	if err != nil {
		fmt.Println(err)
		return
	}

	// to function needs to capture
	// tcpdump -i lo -dd port 5010
	bpfInstructions := []pcap.BPFInstruction{
		{0x28, 0, 0, 0x0000000c},
		{0x15, 0, 8, 0x000086dd},
		{0x30, 0, 0, 0x00000014},
		{0x15, 2, 0, 0x00000084},
		{0x15, 1, 0, 0x00000006},
		{0x15, 0, 17, 0x00000011},
		{0x28, 0, 0, 0x00000036},
		{0x15, 14, 0, 0x00001392},
		{0x28, 0, 0, 0x00000038},
		{0x15, 12, 13, 0x00001392},
		{0x15, 0, 12, 0x00000800},
		{0x30, 0, 0, 0x00000017},
		{0x15, 2, 0, 0x00000084},
		{0x15, 1, 0, 0x00000006},
		{0x15, 0, 8, 0x00000011},
		{0x28, 0, 0, 0x00000014},
		{0x45, 6, 0, 0x00001fff},
		{0xb1, 0, 0, 0x0000000e},
		{0x48, 0, 0, 0x0000000e},
		{0x15, 2, 0, 0x00001392},
		{0x48, 0, 0, 0x00000010},
		{0x15, 0, 1, 0x00001392},
		{0x6, 0, 0, 0x00040000},
		{0x6, 0, 0, 0x00000000},
	}

	if err := handle.SetBPFInstructionFilter(bpfInstructions); err != nil {
		panic(err)
	}

	packetSource := gopacket.NewPacketSource(handle, handle.LinkType())
	for packet := range packetSource.Packets() {
		fmt.Printf("%s", packet.Dump())
	}
}
