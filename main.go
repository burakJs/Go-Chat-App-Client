package main

import (
	"GoSocketChatClientApp/model"
	"GoSocketChatClientApp/utils"
	"bufio"
	"fmt"
	"net"
	"os"
)

const (
	networkType        = "tcp"
	networkHost        = "localhost"
	defaultNetworkPort = "3000"
)

func main() {
	server := model.Server{NetworkType: networkType, NetworkAddress: networkHost + ":" + defaultNetworkPort}
	reader := bufio.NewReader(os.Stdin)
	client := model.Client{}
	var conn net.Conn
	text, err := utils.ReadNameFromConsole(reader)
	if err != nil {
		fmt.Printf("got error while reading name of client: %v\n", err)
		return
	}

	client.Name = text

	conn, err = server.StartListen()
	defer server.CloseListen(conn)
	if err != nil {
		fmt.Printf("got error while connecting to server %v\n", err)
		return
	}
	server.Handle(&client, conn, reader)

}
