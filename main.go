package main

import (
	"log"
	"net/http"
	"github.com/Jeffail/gabs"
	"fmt"
	"os/exec"
	"os"
)

func api(rw http.ResponseWriter, req *http.Request) {
	parsed, err := gabs.ParseJSONBuffer(req.Body)
	if err != nil {
		panic(err)
	}
	defer req.Body.Close()
	cmd := parsed.S("repository", "name").Data().(string)
	out, err := exec.Command("/bin/sh", cmd+".sh").Output()
	if err != nil {
		//log.Fatal(err)
		log.Printf("result:%s\n", err)
		fmt.Fprintf(rw, "result:%s\n", err)
	} else {
		log.Printf("result:%s\n", out)
		fmt.Fprintf(rw, "result:%s\n", out)
	}
}

func main() {
	http.HandleFunc("/api", api)
	argsWithoutProg := os.Args[1:]
	port := "8081"
	if len(argsWithoutProg) > 0 {
		port = argsWithoutProg[0]
	}

	f, err := os.OpenFile("logs", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalf("error opening file: %v", err)
	}
	defer f.Close()

	log.SetOutput(f)

	log.Printf("server is running on %s", port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", port), nil))
}
