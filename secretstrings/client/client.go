package main

import (
	"bufio"
	"flag"
	"fmt"
	"net/rpc"
	"os"
	"secretstrings/stubs"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	server := flag.String("server", "127.0.0.1:8030", "IP:port string to connect to as server")
	flag.Parse()
	fmt.Println("Server: ", *server)

	client, _ := rpc.Dial("tcp", *server)
	defer client.Close()

	// Read the words from wordlist
	f, err := os.Open("../wordlist")
	check(err)

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		msg := scanner.Text()
		request := stubs.Request{Message: msg}
		response := new(stubs.Response)

		client.Call(stubs.PremiumReverseHandler, request, response)
		fmt.Println("Request: " + request.Message)
		fmt.Println("Responded: " + response.Message)
	}
}
