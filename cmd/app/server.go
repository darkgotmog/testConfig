package main

import (
	"configTest/internal/udp"
	"context"
	"fmt"
	"os"
	"os/signal"
	"syscall"
)

func main() {

	server := udp.NewServerUdp(context.Background(), "0.0.0.0", "6701")
	err := server.Start()
	if err != nil {
		fmt.Println(err)
		os.Exit(0)
	}

	defer server.Close()

	c := make(chan os.Signal)

	signal.Notify(c, os.Interrupt, syscall.SIGTERM)
	go func() {
		<-c
		fmt.Println("\r- Ctrl+C pressed in Terminal")
		os.Exit(0)
	}()

	for {
		select {
		case msg := <-server.ChanMessage:
			{
				fmt.Println("Msg: ", msg.Id, msg.Data)

			}
		}
	}

}
