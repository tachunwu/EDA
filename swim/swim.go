package main

import (
	"flag" // or "github.com/spf13/pflag" for more features
	"log"
	"time"

	"github.com/hashicorp/memberlist"
)

// Define command-line flags
var (
	clusterAddr string
	bindPort    int
	bindAddr    string
)
func main() {

	flag.StringVar(&clusterAddr, "cluster", "127.0.0.1:5000", "Address of the cluster to join")
	flag.StringVar(&bindAddr, "addr", "127.0.0.1", "Address to bind")
	flag.IntVar(&bindPort, "port", 7946, "Port to bind on")
	flag.Parse()

	log.Println("Starting node.")

	// Configure memberlist with the bind port
	config := memberlist.DefaultLocalConfig()
	config.BindPort = bindPort
	config.BindAddr = bindAddr
	list, err := memberlist.Create(config)
	if err != nil {
		log.Println("Failed to create memberlist: " + err.Error())
	}

	// Join the cluster
	clusterCount, err := list.Join([]string{clusterAddr})
	log.Println("Joining cluster of size", clusterCount)
	if err != nil {
		log.Println("Failed to join cluster: " + err.Error())
	}

	// Log the members list every two seconds
	for {
		time.Sleep(10 * time.Second)
		log.Println("Members:")
		for _, member := range list.Members() {
			log.Printf("  %s %s\n", member.Name, member.Addr)
		}
	}
}
