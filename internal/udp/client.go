package udp

import (
	"configTest/internal/message"
	"context"
	"encoding/json"
	"fmt"
	"net"
	"time"
)

type ClientUdp struct {
	ctx            context.Context
	ip             string
	port           string
	conn           net.Conn
	requestMessage chan message.RequestMessage
	timeout        time.Duration
}

func NewClientUdp(ctx context.Context, ip string, port string, timeout time.Duration) *ClientUdp {
	client := &ClientUdp{
		ctx:            ctx,
		ip:             ip,
		port:           port,
		requestMessage: make(chan message.RequestMessage),
		timeout:        timeout,
	}
	return client
}

func (c *ClientUdp) Connect() error {

	address := c.ip + ":" + c.port

	udp, err := net.DialTimeout("udp", address, c.timeout)
	if err != nil {
		return err
	}
	c.conn = udp

	go c.runLoop()

	return nil
}

func (c *ClientUdp) Send(mess *message.Message) error {
	responseCh := make(chan message.Response)
	c.requestMessage <- message.RequestMessage{Message: mess, ResponseCh: responseCh}
	r := <-responseCh
	return r.Err
}

func (c *ClientUdp) Close() error {
	close(c.requestMessage)
	err := c.conn.Close()
	return err
}

func (c *ClientUdp) runLoop() {
	for {
		select {
		case <-c.ctx.Done():
			fmt.Println("clientUdp contex cancelled")
			// c.Close()
			return
		case req := <-c.requestMessage:
			mess := req.Message
			responseCh := req.ResponseCh

			err := c.sendData(mess)
			responseCh <- message.Response{Err: err}
		}

	}
}

func (c *ClientUdp) sendData(message *message.Message) error {
	err := c.conn.SetDeadline(time.Now().Add(c.timeout))
	if err != nil {
		return err
	}
	data, err := json.Marshal(message)
	if err != nil {
		return err
	}
	if _, err = c.conn.Write(data); err != nil {
		return err
	}
	return nil
}
