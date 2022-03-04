package main

import (
	"fmt"
	"regexp"
	"strconv"
)

func main() {

	// server := udp.NewServerUdp(context.Background(), "0.0.0.0", "6701")
	// err := server.Start()
	// if err != nil {
	// 	fmt.Println(err)
	// 	os.Exit(0)
	// }

	// defer server.Close()

	// c := make(chan os.Signal)

	// signal.Notify(c, os.Interrupt, syscall.SIGTERM)
	// go func() {
	// 	<-c
	// 	fmt.Println("\r- Ctrl+C pressed in Terminal")
	// 	os.Exit(0)
	// }()

	// for {
	// 	select {
	// 	case msg := <-server.ChanMessage:
	// 		{
	// 			fmt.Println("Msg: ", msg.Id, msg.Data)

	// 		}
	// 	}
	// }

	fmt.Println(checkToChangeEnvoyFileterGprcTimeout("4.4"))
	fmt.Println(checkToChangeEnvoyFileterGprcTimeout("4.4s"))
	fmt.Println(checkToChangeEnvoyFileterGprcTimeout("0.4s"))
	fmt.Println(checkToChangeEnvoyFileterGprcTimeout("1006"))
	fmt.Println(checkToChangeEnvoyFileterGprcTimeout("1.1"))
	fmt.Println(checkToChangeEnvoyFileterGprcTimeout("jgjg"))

}

func checkToChangeEnvoyFileterGprcTimeout(s string) string {
	str := "1s"

	match, _ := regexp.MatchString(`(^\d+[s]$)|(^\d+[.]\d+[s]$)`, s)
	if match {
		return s
	}

	match, _ = regexp.MatchString(`(^\d+$)|(^\d+[.]\d+$)`, s)
	if match {

		value, err := strconv.ParseFloat(s, 64)
		if err != nil {
			return str
		}

		value = value / 1000
		strValue := strconv.FormatFloat(value, 'g', 6, 64) + "s"
		return strValue
	}

	return str
}
