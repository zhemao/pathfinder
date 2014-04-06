package main

import (
	"net/http"
	"flag"
	"log"
	"fmt"
	"bufio"
	"os"
	"io/ioutil"
)

var rootdir = flag.String("rootdir", "./",
			  "which directory to serve files from")
var address = flag.String("address", ":8080",
			  "address to bind server to")
var noconsole = flag.Bool("noconsole", false,
			"if present, don't take user input")

const CommandBufferSize = 100
var commandchan = make(chan string, CommandBufferSize)

func getLine(reader *bufio.Reader) (string, error) {
	line := make([]byte, 0)
	for {
		linepart, hasMore, err := reader.ReadLine()
		if err != nil {
			return "", err
		}
		line = append(line, linepart...)
		if !hasMore {
			break
		}
	}
	return string(line), nil
}

func handleInput(reader *bufio.Reader) {
	for {
		fmt.Print("> ")
		line, err := getLine(reader)
		if err != nil || line == "quit" {
			os.Exit(0)
		}
		commandchan <- line
	}
}

func handleDebug(writer http.ResponseWriter, request *http.Request) {
	header := writer.Header()
	header["Content-Type"] = []string{"text/plain"};
	if request.Method == "GET" {
		message := <-commandchan
		writer.Write([]byte(message))
	} else if request.Method == "POST" {
		body, err := ioutil.ReadAll(request.Body)
		if err != nil {
			writer.WriteHeader(400)
			writer.Write([]byte("Bad request"))
			return
		}
		fmt.Println(string(body))
		if !*noconsole {
			fmt.Print("> ")
		}
		writer.Write([]byte("success"))
	} else {
		writer.WriteHeader(405)
		writer.Write([]byte("Method not allowed\n"))
	}
}


func main() {
	flag.Parse()

	http.HandleFunc("/debug", handleDebug)
	http.Handle("/", http.FileServer(http.Dir(*rootdir)))

	reader := bufio.NewReader(os.Stdin)

	if !*noconsole {
		go handleInput(reader)
	}

	fmt.Printf("serving %s from %s\n", *rootdir, *address)
	err := http.ListenAndServe(*address, nil)
	if err != nil {
		log.Fatal(err)
	}
}
