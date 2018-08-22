package main

import (
	"fmt"
	"log"
	clients "./lib"
)

/*
Description: 
start up bitcoincash chain rpc client
 * Author: architect.bian
 * Date: 2018/08/19 19:33
 */
func main() {
	fmt.Println("start up bitcoincash rpc client")
	cfg := &clients.Config{
		Host:         "172.16.2.41:8332",
		User:         "user",
		Pass:         "password",
		HTTPPostMode: true, // Bitcoin core only supports HTTP POST mode
		DisableTLS:   true, // Bitcoin core does not provide TLS by default
	}
	client, err := clients.NewClient(cfg)

	if err != nil {
		log.Fatal(err)
	}
	defer client.Shutdown()

	// Get the current block count.
	blockCount, err := client.GetBlockCount()
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Block count: %d", blockCount)
}
