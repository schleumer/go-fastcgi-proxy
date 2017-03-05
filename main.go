package main

import (
	"fmt"
	"net"
	"os"
)

type Request struct {
	Connection       net.Conn
	BeginRequestBody BeginRequestBody
}

type BeginRequestBody struct {
	Role0    int
	Role1    int
	Flags    int
	Reserved [5]int
}

type Header struct {
	Version         int
	Type            int
	RequestIDB1     int
	RequestIDB0     int
	ContentLengthB1 int
	ContentLengthB0 int
	PaddingLength   int
	Reserved        int
	ContentLength   int
}

type EndRequestBody struct {
	AppStatus0     int
	AppStatus1     int
	AppStatus2     int
	AppStatus3     int
	ProtocolStatus int
	Reserved       [3]int
}

func makeRequest(conn net.Conn) Request {
	return Request{Connection: conn}
}

func main() {
	l, err := net.Listen("tcp", "127.0.0.1:2000")

	if err != nil {
		fmt.Println("Error opening server:", err.Error())
		os.Exit(1)
	}

	defer l.Close()

	for {
		// Listen for an incoming connection.
		conn, err := l.Accept()

		if err != nil {
			fmt.Println("Error accepting:", err.Error())
			os.Exit(1)
		}
		// Handle connections in a new goroutine.
		go handleRequest(conn)
	}
}

func handleRequest(conn net.Conn) {
	request := makeRequest(conn)

	fmt.Println(request)
}

//
// func fetchBeginRequestBody(conn net.Conn, content *[]byte) BeginRequestBody {
// 	buf := make([]byte, 8)
//
// 	_, err := conn.Read(buf)
//
// 	if err != nil {
// 		fmt.Println("Error reading:", err.Error())
// 	}
//
// }
//
// func fetchHeader(conn net.Conn, content *[]byte) RequestHeader {
// 	buf := make([]byte, 8)
//
// 	_, err := conn.Read(buf)
//
// 	if err != nil {
// 		fmt.Println("Error reading:", err.Error())
// 	}
//
// 	*content = append(*content, buf...)
//
// 	header := RequestHeader{
// 		int(buf[0]),
// 		int(buf[1]),
// 		int(buf[2]),
// 		int(buf[3]),
// 		int(buf[4]),
// 		int(buf[5]),
// 		int(buf[6]),
// 		int(buf[7]),
// 		0,
// 	}
//
// 	header.ContentLength = (header.ContentLength1 << 8) | header.ContentLength0
//
// 	return header
// }
//
// func fetchBody(conn net.Conn, header RequestHeader, content *[]byte) []byte {
// 	buf := make([]byte, header.ContentLength)
//
// 	_, err := conn.Read(buf)
//
// 	if err != nil {
// 		fmt.Println("Error reading:", err.Error())
// 	}
//
// 	return buf
// }
//
// func fetchEndRequestBody(conn net.Conn) {
// 	buf := make([]byte, 16)
//
// 	_, err := conn.Read(buf)
//
// 	if err != nil {
// 		fmt.Println("Error reading:", err.Error())
// 	}
//
// 	fmt.Println(buf)
// }
