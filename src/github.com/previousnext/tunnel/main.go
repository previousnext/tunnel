package main

import (
	"fmt"
	"io"
	"log"
	"net"
	"strings"
	"sync"

	"gopkg.in/alecthomas/kingpin.v2"
)

var (
    ports   = kingpin.Arg("ports", "Ports to forward").Required().String()
    backend = kingpin.Arg("backend", "Backend to receive requests").Required().String()
)

func main() {
	kingpin.Parse()

	var wg sync.WaitGroup

	ps := strings.Split(*ports, ",")

	for _, p := range ps {
		wg.Add(1)

		go func(p string) {
			defer wg.Done()

			l, err := net.Listen("tcp", ":"+p)
			if err != nil {
				log.Fatalf("Failed to setup listener: %v", err)
			}
			defer l.Close()

			fmt.Printf("0.0.0.0:%s -> %s:%s\n", p, *backend, p)

			for {
				conn, err := l.Accept()
				if err != nil {
					log.Fatalf("ERROR: failed to accept listener: %v", err)
				}
				go forward(conn, *backend+":"+p)
			}
		}(p)
	}

	wg.Wait()
}

func forward(conn net.Conn, b string) {
	client, err := net.Dial("tcp", b)
	if err != nil {
		log.Fatalf("Dial failed: %v", err)
	}
	go func() {
		defer client.Close()
		defer conn.Close()
		io.Copy(client, conn)
	}()
	go func() {
		defer client.Close()
		defer conn.Close()
		io.Copy(conn, client)
	}()
}
