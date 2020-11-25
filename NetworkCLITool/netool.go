package main

import (
	"log"
	"os"
	"github.com/urfave/cli"
	"net"
	"fmt"
)

func main() {
	if len(os.Args) != 3 {
		fmt.Println("[*] Usage:", os.Args[0][len(os.Args[0])-len("netool"):], "command", "hostname")
		os.Exit(1)
	}

	/* Create instance of CLI application and add a short description of its purpose */
	app := cli.NewApp()
	app.Name = "Netool CLI"
	app.Usage = "Let's you query IPs, CNAMEs, MX records and Name Servers" 

	/* Used the same flag for all of the commands */
	myFlags := []cli.Flag {
		&cli.StringFlag {
			Name: "host",
			Value: string(os.Args[2]),
		},
	}

	app.Commands = []*cli.Command {
		{
			Name: "ns",
			Usage: "Looks up the Name Servers for a particular host",
			Flags: myFlags,
			/* The block of code inside Action will be triggered 
			   whenever 'ns' will be given as an argument */
			Action: func(c *cli.Context) error {
				ns, err := net.LookupNS(c.String("host"))
				if err != nil {
					fmt.Fprintf(os.Stderr, "Error:", err)
				}
				for i := 0; i < len(ns); i++ {
					fmt.Println(ns[i].Host)
				}
				return nil
			},
		},
		{
			Name: "ip",
			Usage: "Looks up the IP addresses for a particular host",
			Flags: myFlags,
			Action: func (c *cli.Context) error {
				ip, err := net.LookupIP(c.String("host"))
				if err != nil {
					fmt.Fprintf(os.Stderr, "Error:", err)
				}
				for i := 0; i < len(ip); i++ {
					fmt.Println(ip[i])
				}
				return nil
			},
		},
		{
			Name: "cname",
			Usage: "Looks up to the CNAME for a particular host",
			Flags: myFlags,
			Action: func (c *cli.Context) error {
				cname, err := net.LookupCNAME(c.String("host"))
				if err != nil {
					fmt.Fprintf(os.Stderr, "Error:", err)
				}
				fmt.Println(cname)
				return nil 
			},
		},
		{
			Name: "mx",
			Usage: "Looks up the MX records for a particular host",
			Flags: myFlags,
			Action: func(c *cli.Context) error {
				mx, err := net.LookupMX(c.String("host"))
				if err != nil {
					fmt.Fprintf(os.Stderr, "Error:", err)
				}
				for i := 0; i < len(mx); i++ {
					fmt.Println(mx[i].Host, mx[i].Pref)
				}
				return nil
			},
		},
	}
	/* Runs the instance and checks for any possible error */
	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}