# Archivio Brasil One / ArchivioOneApiServer

A simple program using google's gopacket package to communicate with the pcap.h library, a tcpdump on Go.
You can capture packets on your network.

Pretty simple, just to illustrate the power of the package, and countless things you can do.

## Technologies 
	- Back-end Golang: version 1.11.5

## Pacakges

github.com/google/gopacket

```sh

$ go get -u github.com/google/gopacket

```

## Install libcpa in S.O

You need to install the libcpa library on your operating system for the program to compile properly. 
I'm using ubuntu server, so I'll leave it to you like install on ubuntu.

```bash

$ sudo apt-get install libpcap-dev

```

## Block Instructions

Smelling a new block of instructions before compiling.
You need to generate a block of instructions and put in your code to work correctly, which I put is from my machine.

Run the tcpdump command copy and paste the code into Go, then compile...

```bash
$ sudo tcpdump -i lo -dd port 5010

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

```

## Compile Go

Now just compile

```bash

go build

```

Example:
```bash

$ sudo ./sniff -i lo

```