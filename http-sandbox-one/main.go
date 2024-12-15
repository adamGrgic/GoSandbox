package main

// @@@@@@@@@@ WIP @@@@@@@@@@@@
// started on this, then realized I wanted to take a step back so I started `http-sandbox-0`

// The purpose of this sandbox module is to test various http concepts in Go
//
// This project emulates an IoT cluster where multiple nodes communicate back and forth with a master node
// ...
//
// included:
// goroutines for simulating mini nodes (access points) that constantly output temperature readings
// one primary node that can handle connections between all other nodes
// a master node to handle connections with primary nodes

// TODOs
// create a simple algorithm for outputting a range of numbers to simulate an IoT temperature device
// create a server to use this algorithm

// create a primary server node that connects to the mini nodes to listen to and retrieve temperature readings
//
// ... maybe implement a master node, undecided.

import (
	"fmt"
)

func main() {
	fmt.Println("Starting up HTTP One Sandbox")

	// sensor.CreateSensor()
}
