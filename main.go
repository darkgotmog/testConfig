package main

import (
	"fmt"
	"os"
	"os/signal"

	"path/filepath"
	"syscall"

	"github.com/lyft/goruntime/loader"
	stats "github.com/lyft/gostats"
)

type DirectRefresher struct {
	currDir string
}

func (d *DirectRefresher) WatchDirectory(runtimePath string, appDirPath string) string {
	d.currDir = filepath.Join(runtimePath, appDirPath)
	return d.currDir
}

func (d *DirectRefresher) ShouldRefresh(path string, op loader.FileSystemOp) bool {
	if filepath.Dir(path) == d.currDir &&
		(op == loader.Write || op == loader.Create) {
		return true
	}
	return false
}

func main() {
	os.Setenv("USE_STATSD", "false")

	// sink := mock.NewSink()
	// store := stats.NewStore(sink, false)

	store := stats.NewDefaultStore()
	refresher := DirectRefresher{}
	runtime, err := loader.New2("/test", "config", store.ScopeWithTags("test", map[string]string{}), &refresher)
	if err != nil {
		// Handle error
		fmt.Println(err)
		os.Exit(0)
	}
	chanChaneFile := make(chan int)

	runtime.AddUpdateCallback(chanChaneFile)
	s := runtime.Snapshot()
	fmt.Println(s.Entries())
	// s.Get("conf1")

	c := make(chan os.Signal)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)
	go func() {
		<-c
		fmt.Println("\r- Ctrl+C pressed in Terminal")
		os.Exit(0)
	}()

	for {
		select {
		case <-chanChaneFile:
			{
				fmt.Println("config Change!")
				snap := runtime.Snapshot().Entries()

				for key, value := range snap {
					// if value.Uint64Valid {
					fmt.Println("Name:", key, value.Modified, value.Uint64Valid)
					// }
				}

			}
		}
	}
	// fmt.Println("configTest!")
}

// volumeMounts:
// - name: testconfig
//   mountPath: /test/config/

//   volumes:
//   - name: testconfig
// 	configMap:
// 	  name: testconfig
// 	  defaultMode: 256
