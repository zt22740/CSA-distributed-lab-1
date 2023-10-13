package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
)

func Read(conn net.Conn) {
	reader := bufio.NewReader(conn)
	msg, _ := reader.ReadString('\n')
	fmt.Printf(msg)
}
func main() {
	//msgP := flag.String("msg", "Default message", "The message you want to send")
	//flag.Parse()
	stdin := bufio.NewReader(os.Stdin)
	for {
		conn, _ := net.Dial("tcp", "127.0.0.1:8030")
		fmt.Println("Enter text ->")
		msg, _ := stdin.ReadString('\n')
		fmt.Fprintln(conn, msg)
		Read(conn)
	}
}
