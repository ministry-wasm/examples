package main

import (
	"fmt"
	"math/rand"
	"os"
	"strings"
	"time"

	"github.com/micro/mdns"
)

func main() {
	rand.Seed(time.Now().UnixNano())
	serviceName := "myService"
	if len(os.Args) == 2 {
		serviceName = os.Args[1]
	} else {
		serviceName = fmt.Sprintf("Random%06d", rand.Intn(999999))
	}
	// Setup our service export
	host, _ := os.Hostname()
	info := []string{"My awesome service"}
	port := 8000 + rand.Intn(10)
	port = 8000
	fmt.Println("Running", serviceName, port)
	service, err := mdns.NewMDNSService(host, serviceName, "", "", port, nil, info)
	if err != nil {
		fmt.Println("error creating service", err.Error())
	}

	// Create the mDNS server, defer shutdown
	server, err := mdns.NewServer(&mdns.Config{Zone: service})
	if err != nil {
		fmt.Println("error starting server", err.Error())
	}
	defer server.Shutdown()
	// Make a channel for results and start listening
	entriesCh := make(chan *mdns.ServiceEntry, 4)
	others := map[string]string{}
	go func() {
		for entry := range entriesCh {
			if !strings.Contains(entry.Name, serviceName) {
				fmt.Printf("%s Got new entry: %v\n", serviceName, entry.Name)
				others[entry.Name] = entry.Host
			}
		}
	}()

	// Start the lookup
	mdns.Lookup("_foobar._tcp", entriesCh)
	for {
		fmt.Println("................")
		time.Sleep(time.Second * 2)
		for k, _ := range others {
			fmt.Println(">>", serviceName, "knows", k)
		}
	}
	close(entriesCh)
}
