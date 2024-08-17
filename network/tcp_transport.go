package network

import (
	"fmt"
	"net"
)

type TCPTransport struct {
	listenAddr				string
	listener					net.Listener
}

type TCPPeer struct {
	conn 							net.Conn
}

func NewTcpTransport(addr string) *TCPTransport {
	return &TCPTransport{
		listenAddr: addr,
	}
}

func (t *TCPTransport) readLoop(peer *TCPPeer) {
	buf := make([]byte, 2048)
	for {
		n, err := peer.conn.Read(buf)
		if err != nil {
			fmt.Printf("read error: %+v\n", err)
			continue
		}

		msg := buf[:n]
		fmt.Println(string(msg))
		// HandleMsg => server
	}
}

func (t *TCPTransport) acceptLoop() {
	for {
		conn, err := t.listener.Accept()
		if err != nil {
			fmt.Printf("Error accepting connection from : %+v\n ", conn)
			continue
		}

		peer := &TCPPeer{
			conn: conn,
		}

		fmt.Printf("Accepted connection => %+v\n", conn)

		go t.readLoop(peer)
	}
}

func (t *TCPTransport) Start() error {
	ln, err := net.Listen("tcp", t.listenAddr)
	if err != nil {
		return err
	}
	t.listener = ln

	go t.acceptLoop()

	fmt.Println("TCP transport listening to port: ", t.listenAddr)

	return nil
}
