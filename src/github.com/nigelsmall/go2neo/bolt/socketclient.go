package bolt

import (
  "fmt"
  "net"
)

var preamble []byte = []byte{0x60, 0x60, 0xB0, 0x17}
var handshakeRequest []byte = []byte{0x00, 0x00, 0x00, 0x01, 0x00, 0x00, 0x00, 0x00,
                                     0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00}

type Client struct {
  readBuffer []byte
}

func Driver(address string) *Client {
  client := Client{readBuffer: make([]byte, 65536)}

  conn, error := net.Dial("tcp", address)
  if error != nil {
    fmt.Println("Cannot listen to server: ", error)
    return nil
  }

  // perform handshake
  conn.Write(preamble)
  conn.Write(handshakeRequest)
  size, error := conn.Read(client.readBuffer)
  if error != nil {
    fmt.Println("Cannot receive data: ", error)
    return nil
  }
  if size != 4 {
    fmt.Println("Not enough data: ", size)
    return nil
  }

  fmt.Println("Handshake response: ", client.readBuffer[:4])

  return &client
}
