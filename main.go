package main

import (
	"bytes"
	"encoding/json"
	"log"

	"net/http"
	"time"

	"github.com/Simon-Busch/go__blockchain/core"
	"github.com/Simon-Busch/go__blockchain/crypto"
	"github.com/Simon-Busch/go__blockchain/network"
	"github.com/Simon-Busch/go__blockchain/types"
	"github.com/Simon-Busch/go__blockchain/util"
)


func main() {
	validatorPrivKey := crypto.GeneratePrivateKey()
	localNode := makeServer("LOCAL_NODE", &validatorPrivKey, ":3000", []string{":3001"}, ":9000")
	go localNode.Start()

	remoteNode := makeServer("REMOTE_NODE", nil, ":3001", []string{":3002"}, "")
	go remoteNode.Start()

	remoteNodeB := makeServer("REMOTE_NODE_B", nil, ":3002", nil, "")
	go remoteNodeB.Start()

	go func() {
		time.Sleep(11 * time.Second)
		lateNode := makeServer("LATE_NODE", nil, ":3003", []string{":3001"}, "")
		go lateNode.Start()
	}()


	time.Sleep(1 * time.Second)

	// if err := sendTransaction(validatorPrivKey); err != nil {
	// 	panic(err)
	// }

	// collectionOwnerPrivKey := crypto.GeneratePrivateKey()
	// txSendTicker := time.NewTicker(1 * time.Second)

	// collectionHash := createCollectionTx(collectionOwnerPrivKey)
	// go func() {
	// 	for i := 0; i < 20 ; i++ {
	// 		nftMinter(collectionOwnerPrivKey, collectionHash)

	// 		<- txSendTicker.C // Wait for the channel to be ready
	// 	}
	// }()

	select {}
}


func makeServer(id string, pk *crypto.PrivateKey, addr string, seedNodes []string, apiListenAddr string) *network.Server {
	opts := network.ServerOpts{
		APIListenAddr: apiListenAddr,
		SeedNodes: 		 seedNodes,
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

func sendTransaction(privKey crypto.PrivateKey) error {
	toPrivKey := crypto.GeneratePrivateKey()

	tx := core.NewTransaction(nil)
	tx.To = toPrivKey.PublicKey()
	tx.Value = 100

	if err := tx.Sign(privKey); err != nil {
		return err
	}

	buf := &bytes.Buffer{}
	if err := tx.Encode(core.NewGobTxEncoder(buf)); err != nil {
		panic(err)
	}

	req, err := http.NewRequest("POST", "http://localhost:9000/tx", buf)
	if err != nil {
		panic(err)
	}

	client := http.DefaultClient
	_, err = client.Do(req)
	return err
}

func createCollectionTx(privKey crypto.PrivateKey) types.Hash {
	tx := core.NewTransaction(nil)
	tx.TxInner = core.CollectionTx{
		Fee: 				100,
		MetaData: 	[]byte("foo collection"),
	}

	tx.Sign(privKey)

	buf := &bytes.Buffer{}
	if err := tx.Encode(core.NewGobTxEncoder(buf)); err != nil {
		panic(err)
	}

	req, err := http.NewRequest("POST", "http://localhost:9000/tx", buf)
	if err != nil {
		panic(err)
	}

	client := http.DefaultClient
	_, err = client.Do(req)
	if err != nil {
		panic(err)
	}

	return tx.Hash(core.TxHasher{})
}


func nftMinter(privKey crypto.PrivateKey, collection types.Hash) {
	metaData := map[string]any{
		"power": 		8,
		"health": 	100,
		"stamina": 	20,
		"color": 		"orange",
		"rare": 		"yes",
	}

	metaBuf := new(bytes.Buffer)
	if err := json.NewEncoder(metaBuf).Encode(metaData); err != nil {
		panic(err)
	}

	tx := core.NewTransaction(nil)
	tx.TxInner = core.MintTx{
		Fee: 								100,
		NFT: 								util.RandomHash(),
		MetaData: 					metaBuf.Bytes(),
		Collection: 				collection,
		CollectionOwner:		privKey.PublicKey(),
	}

	tx.Sign(privKey)

	buf := &bytes.Buffer{}
	if err := tx.Encode(core.NewGobTxEncoder(buf)); err != nil {
		panic(err)
	}

	req, err := http.NewRequest("POST", "http://localhost:9000/tx", buf)
	if err != nil {
		panic(err)
	}

	client := http.DefaultClient
	_, err = client.Do(req)
	if err != nil {
		panic(err)
	}
}
