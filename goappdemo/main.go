package main

import (
	"database/sql"
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"

	_ "github.com/lib/pq"
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
	fmt.Println("Build nr:", Version)
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

	data := getFromDB()
	w.Write(byteME(data))
}

func byteME(message string) []byte {
	return []byte(message)
}

func getFromDB() string {
	username := os.Getenv("POSTGRES_USER")
	password := os.Getenv("POSTGRES_PASSWORD")
	server := os.Getenv("POSTGRES_SERVER")
	connStr := fmt.Sprintf("postgresql://%s:%s@%s/todos?sslmode=disable",
		username,
		password,
		server)

	// Connect to database
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}

	rows, err := db.Query("SELECT * FROM users")
	defer rows.Close()
	if err != nil {
		log.Fatal(err)
	}

	var line string
	var buffer []string

	for rows.Next() {
		rows.Scan(&line)
		buffer = append(buffer, line)
	}
	return strings.Join(buffer, "\n")
}
