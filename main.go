package main

import (
	"configTest/internal/config"
	"configTest/internal/message"
	"configTest/internal/udp"
	"context"
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {

	conf := config.NewConfig()

	fmt.Println(conf)

	server := udp.NewServerUdp(context.Background(), conf.Host, conf.Port)
	err := server.Start()
	if err != nil {
		fmt.Println(err)
		os.Exit(0)
	}

	defer server.Close()
	// conf.IP = "127.0.0.1"
	conf.FlagSend = true

	client := udp.NewClientUdp(context.Background(), conf.IP, conf.Port, 5*time.Second)

	err = client.Connect()
	if err != nil {
		fmt.Println("connect", err)
	}

	defer client.Close()

	if conf.FlagSend {
		go LoopSendUdpMessage(client, conf.ID)
	}

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
				if msg.Id != conf.ID {
					fmt.Println("Msg: ", msg.Id, msg.Data)
				}

			}
		}
	}

}

func LoopSendUdpMessage(client *udp.ClientUdp, id int64) {

	ticker := time.NewTicker(3 * time.Second)
	for _ = range ticker.C {
		err := client.Send(&message.Message{Id: id, Data: []byte(time.Now().GoString())})

		if err != nil {
			fmt.Println("send", err)
		}
	}
}

// type DirectRefresher struct {
// 	currDir string
// }

// func (d *DirectRefresher) WatchDirectory(runtimePath string, appDirPath string) string {
// 	d.currDir = filepath.Join(runtimePath, appDirPath)
// 	return d.currDir
// }

// func (d *DirectRefresher) ShouldRefresh(path string, op loader.FileSystemOp) bool {
// 	if filepath.Dir(path) == d.currDir &&
// 		(op == loader.Write || op == loader.Create) {
// 		return true
// 	}
// 	return false
// }

// func main() {
// 	os.Setenv("USE_STATSD", "false")

// 	// sink := mock.NewSink()
// 	// store := stats.NewStore(sink, false)

// 	store := stats.NewDefaultStore()
// 	refresher := DirectRefresher{}
// 	runtime, err := loader.New2("test", "config", store.ScopeWithTags("test", map[string]string{}), &refresher)
// 	if err != nil {
// 		// Handle error
// 		fmt.Println(err)
// 		os.Exit(0)
// 	}
// 	chanChaneFile := make(chan int)

// 	runtime.AddUpdateCallback(chanChaneFile)
// 	s := runtime.Snapshot()
// 	fmt.Println(s.Entries())

// 	client := udp.NewClientUdp(context.Background(), "192.168.31.255", "6701", 5*time.Second)

// 	err = client.Connect()
// 	if err != nil {
// 		fmt.Println("connect", err)
// 	}

// 	defer client.Close()

// 	err = client.Send(&message.Message{Id: 3, Data: []byte("sdfasdf")})

// 	if err != nil {
// 		fmt.Println("send", err)
// 	}

// 	c := make(chan os.Signal)
// 	signal.Notify(c, os.Interrupt, syscall.SIGTERM)
// 	go func() {
// 		<-c
// 		fmt.Println("\r- Ctrl+C pressed in Terminal")
// 		os.Exit(0)
// 	}()

// 	for {
// 		select {
// 		case <-chanChaneFile:
// 			{
// 				fmt.Println("config Change!")
// 				snap := runtime.Snapshot().Entries()

// 				for key, value := range snap {
// 					// if value.Uint64Valid {
// 					fmt.Println("Name:", key, value.Modified, value.Uint64Valid)
// 					// }
// 				}

// 			}
// 		}
// 	}
// 	// fmt.Println("configTest!")
// }

// volumeMounts:
// - name: testconfig
//   mountPath: /test/config/

//   volumes:
//   - name: testconfig
// 	configMap:
// 	  name: testconfig
// 	  defaultMode: 256
