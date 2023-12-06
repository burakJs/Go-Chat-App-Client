package model

import (
	"GoSocketChatClientApp/utils"
	"bufio"
	"encoding/json"
	"fmt"
	"net"
	"strings"
)

type Server struct {
	NetworkType    string
	NetworkAddress string
}

func (s Server) StartListen() (net.Conn, error) {
	conn, err := net.Dial(s.NetworkType, s.NetworkAddress)
	return conn, err
}

func (s Server) CloseListen(conn net.Conn) {
	err := conn.Close()
	if err != nil {
		fmt.Printf("got error while closing connection: %v\n", err)
	}
}

func (s Server) Handle(client *Client, conn net.Conn, reader *bufio.Reader) {
	client.Message = "/enter"
	err := s.sendDataToServer(conn, client)

	if err != nil {
		fmt.Printf("got error while send data with /enter command: %v\n", err)
		return
	}

	for {
		text, err := utils.ReadMessageFromConsole(reader)
		if err != nil {
			fmt.Printf("got error while reading message of client: %v\n", err)
			return
		}
		client.Message = text
		err = s.sendDataToServer(conn, client)
		if err != nil {
			fmt.Println(err)
			return
		}

		if strings.Compare(text, "/exit") == 0 {
			break
		}
	}
}

func (s Server) sendDataToServer(conn net.Conn, client *Client) error {
	jsonData, err := json.Marshal(client)
	_, err = conn.Write(jsonData)
	return err
}
