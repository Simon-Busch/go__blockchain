package main

import (
	"time"

	"github.com/Simon-Busch/go__blockchain/network"
)

// Server
// Transport ==> tcp, udp
// Block
// Tx
// Keypair


func main() {

	trLocal := network.NewLocalTransport("Local")
	trRemote := network.NewLocalTransport("Remote")

	trLocal.Connect(trRemote)
	trRemote.Connect(trLocal)

	go func() {
		for {
			trRemote.SendMessage(trLocal.Addr(), []byte("Hello world"))
			time.Sleep(1 * time.Second)
		}
	}()


	opts := network.ServerOpts{
		Transports: []network.Transport{trLocal},
	}

	s := network.NewServer(opts)

	s.Start()
}
