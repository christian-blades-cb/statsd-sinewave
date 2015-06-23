package main // github.com/christian-blades-cb/statsd-sinewave

import (
	"math"
	"net"

	"fmt"
	"github.com/jessevdk/go-flags"
	"log"
	"time"
)

func main() {
	var opts struct {
		StatsdServer string `short:"s" long:"host" env:"STATSD_HOST" default:"localhost:8125" required:"true" description:"statsd hostname/ip and port"`
		Delay        int64  `short:"d" long:"delay" required:"true" default:"50" description:"delay (in milliseconds) between stats emissions"`
	}
	_, err := flags.Parse(&opts)
	if err != nil {
		log.Fatal("unable to parse command line arguments")
	}

	conn, err := net.Dial("udp", opts.StatsdServer)
	if err != nil {
		log.Fatal(err)
	}

	log.Println("Emitting")
	ticker := time.NewTicker(time.Millisecond * time.Duration(opts.Delay))
	x := 0.0
	for _ = range ticker.C {
		fmt.Fprintf(conn, "sinewave:%f|c\n", x)
		x = math.Mod(x, 360)
	}
}
