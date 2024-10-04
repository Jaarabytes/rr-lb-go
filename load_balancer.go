package main

import (
	"crypto/sha256"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"sync"
)

// 4 databases to handle 16 possible characters, 0123456789ABCDEF, based on hash
var (
	databases = []string{
		"1db.txt",
		"2db.txt",
		"3db.txt",
		"4db.txt",
	}
	currentDatabase int
	mutex          sync.Mutex
)


// Start the HTTP Server
func main() {
	http.HandleFunc("/", handleRequest)
	log.Println("Server starting on :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func handleRequest(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Only POST requests are allowed", http.StatusMethodNotAllowed)
		return
	}

	body, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Error reading request body", http.StatusInternalServerError)
		return
	}

	hash := sha256.Sum256(body)
	hashedString := fmt.Sprintf("%x", hash)

  databaseContent := hashedString + " " + string(body)
	database := getNextDatabase()
	err = writeToFile(database, databaseContent)
	if err != nil {
		http.Error(w, "Error writing to file", http.StatusInternalServerError)
		return
	}

	fmt.Fprintf(w, "Request processed and written to %s", database)
}

func getNextDatabase() string {
	mutex.Lock()
	defer mutex.Unlock()
	database := databases[currentDatabase]
	currentDatabase = (currentDatabase + 1) % len(databases)
	return database
}

func writeToFile(filename, content string) error {
	f, err := os.OpenFile(filename, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return err
	}
	defer f.Close()

	_, err = f.WriteString(content + "\n")
	return err
}
