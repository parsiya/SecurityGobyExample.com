// [net package](https://godoc.org/net) contains most of the networking code
// in Go.TCP connections are made with [net.Dial](https://godoc.org/net#Dial).

package main

import (
	"fmt"
	"net"
	"time"
)

func main() {

	// Address can be created with [net.JoinHostPort](https://godoc.org/net#JoinHostPort).
	add1 := net.JoinHostPort("127.0.0.1", "1234")

	// "127.0.0.1:1234"
	fmt.Println(add1)

	// It also takes care of IPv6 addresses
	add2 := net.JoinHostPort("2001:db8::1", "1111")

	// [2001:db8::1]:1111
	fmt.Println(add2)

	// net.Dial accepts a network and an address.
	// TCP networks are "tcp", "tcp4" and "tcp6." In most cases we will be using
	// the general "tcp" network and let the package choose one.
	// Address must be a combination of "host:port."
	conn, err := net.Dial("tcp", "127.0.0.1:22")

	// Host can also be a resolvable host name.
	conn, err := net.Dial("tcp", "example.com:80")

	// DialTimeout sets a time out automatically
	ts := time.Second * 5
	conn, err := net.DialTimeout("tcp", "example.com:80", ts)

	// Return local and remote addresses
	laddr := conn.LocalAddr()
	raddr := conn.RemoteAddr()

	// Set time out after Dial - 0 means connection will never time out
	timeout := time.Second * 5
	err := conn.SetDeadline(timeout)

	// It's possible to have different time outs for read/write
	err := conn.SetReadDeadline(timeout)
	err := conn.SetWriteDeadline(timeout)

	// Read, reads from the connection and writes into a []byte. Returns the
	// number of bytes read.
	var buf []byte
	n, err := conn.Read(buf)

	// Write is similar. Returns the number of bytes written.
	n, err := conn.Write(buf)

	// Create a bufio.Reader from connection
	connReader := bufio.NewReader(conn)

	// ReadString until a null byte
	str, err := connReader.ReadString(byte(0x00))

	// ReadBytes until a null byte
	buf, err := connReader.ReadBytes(byte(0x00))

	// Or create a scanner from bufio.Reader
	scanner := bufio.NewScanner(connReader)

	// Read from scanner
	for scanner.Scan() {
		fmt.Printf("%s\n", scanner.Text())
	}

	// Check if scanner has quit with an error
	if err := scanner.Err(); err != nil {
		fmt.Println("Scanner error", err)
	}

	// net contains TCP specific methods.

	// [ResolveTCPAddr](https://golang.org/pkg/net/#ResolveTCPAddr) converts
	// a host:port address to a *TCPAddr.
	tcpAddr, err := net.ResolveTCPAddr("tcp", "127.0.0.1:22")
	tcpAddr, err := net.ResolveTCPAddr("tcp",
		net.JoinHostPort("127.0.0.1", "22"))

	// TCPAddr can be used in [net.DialTCP](https://godoc.org/net#DialTCP).
	// DialTCP is similar to Dial with additional methods.
	// tcpConn can be fed into any function that accepts net.Conn
	tcpConn, err := DialTCP("tcp", tcpAddr)
}
