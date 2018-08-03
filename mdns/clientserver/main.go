package main

import (
	"fmt"
	"os"

	"github.com/micro/mdns"
)

func main() {
	if len(os.Args) != 2 {
		panic("Usage: clientserver name-of-this-server")
	}
	// Setup our service export
	host, _ := os.Hostname()
	serviceName := os.Args[1]
	fmt.Println("Running", serviceName)
	info := []string{"My awesome service"}
	service, _ := mdns.NewMDNSService(host, serviceName, "", "", 8000, nil, info)

	// Create the mDNS server, defer shutdown
	server, _ := mdns.NewServer(&mdns.Config{Zone: service})
	defer server.Shutdown()
	// Make a channel for results and start listening
	entriesCh := make(chan *mdns.ServiceEntry, 4)
	go func() {
		for entry := range entriesCh {
			fmt.Printf("%s Got new entry: %v\n", serviceName, entry)
		}
	}()

	// Start the lookup
	mdns.Lookup("_foobar._tcp", entriesCh)
	close(entriesCh)
}
