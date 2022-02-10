package udp

import (
	"configTest/internal/message"
	"context"
	"encoding/json"
	"fmt"
	"net"
)

const maxBufferSize = 1024

type ServerUdp struct {
	ctx         context.Context
	ip          string
	port        string
	conn        *net.UDPConn
	ChanMessage chan message.Message
}

func NewServerUdp(ctx context.Context, ip string, port string) *ServerUdp {
	server := &ServerUdp{
		ctx:         ctx,
		ip:          ip,
		port:        port,
		ChanMessage: make(chan message.Message),
	}
	return server
}

func (c *ServerUdp) Start() error {

	address := c.ip + ":" + c.port

	addr, err := net.ResolveUDPAddr("udp", address)
	if err != nil {
		return err
	}

	conn, err := net.ListenMulticastUDP("udp", nil, addr)
	if err != nil {
		return err
	}

	c.conn = conn

	go c.runLoop()

	return nil
}

func (c *ServerUdp) Close() error {
	close(c.ChanMessage)
	err := c.conn.Close()
	return err
}

func (c *ServerUdp) runLoop() {
	buffer := make([]byte, maxBufferSize)

	for {
		n, addr, err := c.conn.ReadFromUDP(buffer)
		if err != nil {
			continue
		}

		fmt.Printf("packet-received: bytes=%d from=%s\n",
			n, addr.String())

		fmt.Println(string(buffer[:n]))

		messa := message.Message{}
		json.Unmarshal(buffer[:n], &messa)

		c.ChanMessage <- messa

		select {
		case <-c.ctx.Done():
			fmt.Println("serverUdp contex cancelled")
			// c.Close()
			return
		default:
			continue
		}

	}
}
