package main

import (
	"fmt"
	"log"
	"net"
	"os"
	"strings"
	"strconv"
	"time"
)

func main() {
	// Prompt the user for a list of ports to listen on
	ports := getPorts()

	// Listen for incoming connections on the specified ports
	for _, port := range ports {
		ln, err := net.Listen("tcp", fmt.Sprintf(":%d", port))
		if err != nil {
			log.Fatalf("Error listening on port %d: %s", port, err)
		}
		defer ln.Close()

		// Open a file for logging connection attempts on this port
		f, err := os.OpenFile(fmt.Sprintf("honeypot-%d.log", port), os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0644)
		if err != nil {
			log.Fatalf("Error opening honeypot log file: %s", err)
		}
		defer f.Close()

		// Accept incoming connections and log connection attempts
		for {
			conn, err := ln.Accept()
			if err != nil {
				log.Fatalf("Error accepting connection: %s", err)
			}

			// Log the connection attempt to the honeypot log file
			logMessage := time.Now().Format("2006-01-02 15:04:05") + " Connection from " + conn.RemoteAddr().String() + "\n"
			if _, err := f.WriteString(logMessage); err != nil {
				log.Fatalf("Error writing to honeypot log file: %s", err)
			}

			// Close the connection
			conn.Close()
		}
	}
}

func getPorts() []int {
	var ports []int

	// Prompt the user for a list of ports to listen on
	fmt.Print("Enter a comma-separated list of ports to listen on (default is 22): ")
	var input string
	fmt.Scanln(&input)
	if input == "" {
		ports = []int{22}
	} else {
		for _, portString := range strings.Split(input, ",") {
			portString = strings.TrimSpace(portString)
			port, err := strconv.Atoi(portString)
			if err != nil {
				log.Fatalf("Error parsing port number %s: %s", portString, err)
			}
			ports = append(ports, port)
		}
	}

	return ports
}
