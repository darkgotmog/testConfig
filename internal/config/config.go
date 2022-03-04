package config

import (
	"os"
	"strconv"
	"time"
)

type Config struct {
	Host       string
	IP         string
	Port       string
	ID         int64
	FlagSend   bool
	SharedName string
	PodName    string
}

func NewConfig() *Config {
	con := &Config{}
	con.Load()
	return con
}

func (con *Config) Load() {

	port := os.Getenv("T_PORT")
	if port == "" {
		port = "9999"
	}
	con.Port = port
	ip := os.Getenv("T_IP")
	if ip == "" {
		ip = "239.0.0.0"
	}

	con.IP = ip

	host := os.Getenv("T_HOST")
	if host == "" {
		host = "239.0.0.0"
	}
	con.Host = host

	sharedName := os.Getenv("SHARED_NAME")
	if sharedName == "" {
		sharedName = "test"
	}
	con.SharedName = sharedName

	id, err := strconv.ParseInt(os.Getenv("T_ID"), 10, 64)
	if err != nil {
		id = time.Now().Unix()
	}
	con.ID = id

	flag, err := strconv.ParseBool(os.Getenv("T_FLAG_SEND"))
	if err != nil {
		con.FlagSend = false
	} else {
		con.FlagSend = flag
	}

	podName := os.Getenv("MY_POD_NAME")
	if podName == "" {
		podName = "podTestUdp"
	}
	con.PodName = podName

}
