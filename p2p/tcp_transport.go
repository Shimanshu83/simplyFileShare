package p2p

import (
	"fmt"
	"net"
	"sync"
)

type TCPTransport struct {
	listenAddress string
	listener      net.Listener
	shakeHands    HandshakeFunc
	decoder       Decoder

	mu    sync.RWMutex // protects the peer but I don't know why it is needed will be figuring it out later on please makre sure that you will be figure it out later
	peers map[net.Addr]Peer
}

// Tcp peer reprsents the remote node over a tcp established connection
type TCPPeer struct {
	// conn is the underlying connection of the peer
	conn net.Conn

	// if we dial and retrieve a conn -> outbound == true
	// if we accept and retrieve a conn -> outbound == false
	outbound bool
}

func NewPeer(conn net.Conn, outbound bool) *TCPPeer {
	return &TCPPeer{
		conn:     conn,
		outbound: outbound,
	}
}

func NewTCPTransport(listenAddress string) *TCPTransport {
	return &TCPTransport{
		listenAddress: listenAddress,
		shakeHands: NOPHandshakeFunc() ,
	}
}

func (t *TCPTransport) ListenAndAccept() error {
	var err error
	t.listener, err = net.Listen("tcp", t.listenAddress)
	if err != nil {
		return err
	}

	go t.startAcceptLoop()

	return nil

}

func (t *TCPTransport) startAcceptLoop() {

	for {
		conn, err := t.listener.Accept()
		if err != nil {
			fmt.Printf("Tcp error accep error: %s\n", err)
		}
		go t.handleConnection(conn)

	}

}

func (t *TCPTransport) handleConnection(conn net.Conn) {
	peer := NewPeer(conn, true)
	fmt.Printf(" nwe incoming connection: %v\n", peer)

	if err := t.shakeHands(conn); err != nil {

	}

}
