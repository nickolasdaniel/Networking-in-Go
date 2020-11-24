package main

import (
	"net"
	"os"
	"fmt"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Println("Usage:", os.Args[0], "hostname")
		os.Exit(1)
	}

	name := os.Args[1]
	cname, err := net.LookupCNAME(name) 
	if err != nil {
		fmt.Println("Error:", err.Error())
		os.Exit(2)
	}
	fmt.Println("Canonical Name: ", cname)
	addrs, err := net.LookupHost(name)
	if err != nil {
		fmt.Println("Error:", err.Error())
		os.Exit(2)
	}

	for _, s := range addrs {
		fmt.Println(s)
	}
	os.Exit(0)
}