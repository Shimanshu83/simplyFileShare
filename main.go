package main

import (
	"fmt"
	"log"

	"github.com/Shimanshu83/simplyFile/p2p"
)

func main() {

	fmt.Println("Side project started...")
	tr := p2p.NewTCPTransport(":3000")

	if err := tr.ListenAndAccept(); err != nil {
		log.Fatal(err)
	}

	select {}
}
