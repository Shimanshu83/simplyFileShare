package p2p

// Peer is an interface that represents the remote node
type Peer interface {
}

// Transport is anything that handles the communication between
// the nodes int the network. This can be of the form
// (Tcp , udp , websockets ...)
type Transport interface {
	ListenAndAccept() error
}