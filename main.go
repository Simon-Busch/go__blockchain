package main

import (
	"bytes"
	// "encoding/gob"
	// "fmt"
	"log"
	// "time"
	// "time"

	"github.com/Simon-Busch/go__blockchain/core"
	"github.com/Simon-Busch/go__blockchain/crypto"

	"github.com/Simon-Busch/go__blockchain/network"
	// "github.com/sirupsen/logrus"

	"net"
)

func main() {
	privKey := crypto.GeneratePrivateKey()
	localNode := makeServer("LOCAL_NODE", &privKey, ":3000", []string{":3001"})
	go localNode.Start()

	remoteNode := makeServer("REMOTE_NODE", nil, ":3001", []string{":3002"})
	go remoteNode.Start()

	remoteNodeB := makeServer("REMOTE_NODE_B", nil, ":3002", nil)
	go remoteNodeB.Start()

	// time.Sleep(1 * time.Second)

	// tcpTester()

	select {}
}


func makeServer(id string, pk *crypto.PrivateKey, addr string, seedNodes []string) *network.Server {
	opts := network.ServerOpts{
		SeedNodes:  seedNodes,
		ListenAddr: addr,
		PrivateKey: pk,
		ID:         id,
	}

	s, err := network.NewServer(opts)
	if err != nil {
		log.Fatal(err)
	}

	return s
}


func tcpConnector() {
	conn , err := net.Dial("tcp", ":3000")
	if err != nil {
		panic(err)
	}

	//Build a transaction
	privKey := crypto.GeneratePrivateKey()

	tx := core.NewTransaction(contract())
	tx.Sign(privKey)
	buf := &bytes.Buffer{}
	if err := tx.Encode(core.NewGobTxEncoder(buf)); err != nil {
		panic(err)
	}

	msg := network.NewMessage(network.MessageTypeTx, buf.Bytes())


	_, err = conn.Write(msg.Bytes())
	if err != nil {
		panic(err)
	}
}

// var transports = []network.Transport{
// 	network.NewLocalTransport("LOCAL"),
// 	// network.NewLocalTransport("REMOTE_B"),
// 	// network.NewLocalTransport("REMOTE_C"),

// }

// func main() {
// 	initRemoteServers(transports)

// 	localNode := transports[0]
// 	lateTr :=	network.NewLocalTransport("LATE_NODE")
// 	// remoteNodeA := transports[1]
// 	// remoteNodeC := transports[3]

// 	// go func() {
// 	// 	for {
// 	// 		if err := sendTransaction(remoteNodeA, localNode.Addr()); err != nil {
// 	// 			logrus.Error(err)
// 	// 		}
// 	// 		time.Sleep(2 * time.Second)
// 	// 	}
// 	// }()

// 	go func() {
// 		time.Sleep(7 * time.Second)
// 		lateServer := makeServer(string(lateTr.Addr()), lateTr, nil)
// 		// if err := localNode.Connect(lateTr); err != nil {
// 		// 	fmt.Println(err)
// 		// }
// 		go lateServer.Start()
// 	}()

// 	privKey := crypto.GeneratePrivateKey()
// 	localServer := makeServer("LOCAL", localNode, &privKey)
// 	localServer.Start()
// }

// func initRemoteServers(trs []network.Transport) {
// 	for i := 0; i < len(trs); i++ {
// 		id := fmt.Sprintf("REMOTE_%d", i)
// 		s := makeServer(id, trs[i], nil)
// 		go s.Start()
// 	}
// }

// func makeServer(id string, tr network.Transport, pk *crypto.PrivateKey) *network.Server {
// 	opts := network.ServerOpts{
// 		PrivateKey: pk,
// 		ID:         id,
// 		Transports: transports,
// 		Transport:  tr,
// 	}

// 	s, err := network.NewServer(opts)
// 	if err != nil {
// 		log.Fatal(err)
// 	}

// 	return s
// }

// func sendTransaction(tr network.Transport, to network.NetAddr) error {
// 	privKey := crypto.GeneratePrivateKey()

// 	tx := core.NewTransaction(contract())
// 	tx.Sign(privKey)
// 	buf := &bytes.Buffer{}
// 	if err := tx.Encode(core.NewGobTxEncoder(buf)); err != nil {
// 		return err
// 	}

// 	msg := network.NewMessage(network.MessageTypeTx, buf.Bytes())

// 	return tr.SendMessage(to, msg.Bytes())
// }

func contract() []byte {
	data := []byte{0x02, 0x0a, 0x03, 0x0a, 0x0b, 0x4f, 0x0c, 0x4f, 0x0c, 0x46, 0x0c, 0x03, 0x0a, 0x0d, 0x0f}
	pushFoo := []byte{0x4f, 0x0c, 0x4f, 0x0c, 0x46, 0x0c, 0x03, 0x0a, 0x0d, 0xae}
	data = append(data, pushFoo...)
	return data
}

// func sendGetStatusMessage(tr network.Transport, to network.NetAddr) error {
// 	var (
// 		getStatusMsg = new(network.GetStatusMessage)
// 		buf = new(bytes.Buffer)
// 	)

// 	if err := gob.NewEncoder(buf).Encode(getStatusMsg); err != nil {
// 		return err
// 	}

// 	msg := network.NewMessage(network.MessageTypeGetStatus, buf.Bytes())
// 	return tr.SendMessage(to, msg.Bytes())
// }
