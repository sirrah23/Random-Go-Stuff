/*getheadinfo
 */

package main

import (
	"fmt"
	"net"
	"os"
)

func main() {
	var buf [512]byte
	if len(os.Args) != 2 {
		fmt.Fprintf(os.Stderr, "Usage: %s host:port ", os.Args[0])
		os.Exit(1)
	}
	service := os.Args[1]
	//Take the DNS/ip and port number to create TCPaddr struct
	tcpAddr, err := net.ResolveTCPAddr("tcp4", service)
	checkError(err)
	//Connect to the address at port number via TCPaddr struct
	conn, err := net.DialTCP("tcp", nil, tcpAddr)
	checkError(err)
	//Write to the server (make a request, in this case a HEAD request from HTTP)
	_, err = conn.Write([]byte("AYY LMAO"))
	checkError(err)
	//Read what the server gave back to us
	_, err = conn.Read(buf[0:])
	checkError(err)
	//Result
	fmt.Println(string(buf[0:]))

	os.Exit(0)

}

func checkError(err error) {
	if err != nil {
		fmt.Fprintf(os.Stderr, "Fatal error: %s", err.Error())
		os.Exit(1)
	}
}
