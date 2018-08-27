package main

import (
	"fmt"
	"log"
	"./client"
)

/*
Description: 
start up bitcoincash chain rpc client
 * Author: architect.bian
 * Date: 2018/08/19 19:33
 */
func main() {
	fmt.Println("start up bitcoincash rpc client")
	cfg := &client.Config{
		Host:         "172.16.2.41:8332",
		User:         "user",
		Pass:         "password",
	}
	//cfg := &client.Config{
	//	Host:         "192.168.137.169:18332",
	//	User:         "bch",
	//	Pass:         "bch",
	//}
	client := client.NewClient(cfg).Startup()
	defer client.Shutdown()

	// Get the current block count.
	blockCount, err := client.GetBlockCount()
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Block count: %d", blockCount)

}
