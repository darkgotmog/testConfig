package config

import (
	"os"
	"strconv"
	"strings"
	"time"
)

type Config struct {
	Host     string
	IP       string
	Port     string
	ID       int64
	FlagSend bool
}

func NewConfig() *Config {
	con := &Config{}
	con.Load()
	return con
}

func (con *Config) Load() {

	port := os.Getenv("T_PORT")
	if port == "" {
		port = "6701"
	}
	con.Port = port
	ip := os.Getenv("T_IP")
	if ip == "" {
		ip = "127.0.0.1"
	}
	k := strings.LastIndex(ip, ".")

	ipnew := ip[:k+1] + "255"

	con.IP = ipnew

	host := os.Getenv("T_HOST")
	if host == "" {
		host = "0.0.0.0"
	}
	con.Host = host

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
}
