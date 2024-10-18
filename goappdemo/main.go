package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"
)

var (
	port        int
	versionFlag bool
	Version     string = "development"
)

func init() {
	flag.IntVar(&port, "port", 8080, "Application port to listen.")
	flag.BoolVar(&versionFlag, "version", false, "Application version from git hash.")
}

func main() {
	flag.Parse()
	if versionFlag {
		fmt.Println("version: ", Version)
		return
	}
	fmt.Println("Using port:", port)
	fmt.Println("Build nr:")
	fmt.Println("Environment:")
	for _, v := range os.Environ() {
		data := strings.Split(v, "=")
		fmt.Println(fmt.Sprintf(" * %s=%s", data[0], data[1]))
	}

	http.HandleFunc("/auth", authenticationAPI)
	http.HandleFunc("/api", generalAPI)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", port), nil))
}

func authenticationAPI(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Authentication connection from:", r.RemoteAddr)
	w.Write(byteME("authentication API"))
	//not authenticated right now
}

func generalAPI(w http.ResponseWriter, r *http.Request) {
	fmt.Println("API connection from:", r.RemoteAddr)
	w.Write(byteME("general API"))
}

func byteME(message string) []byte {
	return []byte(message)
}
