package telnet

import (
	"bufio"
	"fmt"
	"log"
	"net"
)

func Check() {
	var (
		err        error
		connection net.Conn
		command    string
	)
	// connect to this socket
	connection, err = net.Dial("tcp", "192.168.56.100:22")
	if err != nil {
		log.Println(err.Error())
	}
	command = "ping\n\n"
	fmt.Fprintf(connection, command)
	_, err = bufio.NewReader(connection).ReadString('\n')
	if err != nil {
		log.Println(err.Error())
	}
}
