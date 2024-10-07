// NOTE: This is the HashBased version of the load balancer.
// It attempts to write to all 4 databases using hashes
//
// Only 4 databases are used since hashes are hexadecimal based thus 16 is the number of all possible characters.
//
// The request string is hashed and inserted into either of the databases if it lies on the below ranges. If it starts with 0, 1, 2 or 3, the request
// would therefore be located in 1db.txt, and so on.
// 0123 ~ 1db.txt
// 4567 ~ 2db.txt
// 89AB ~ 3db.txt 
// CDEF ~ 4db.txt

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
	database := getAppropriateDatabase(hashedString)
  if database == "" {
		http.Error(w, "Error, inappropriate database provided", http.StatusInternalServerError)
		return
	}
	err = writeToFile(database, databaseContent)
  if err != nil {
      http.Error(w, "Error writing to file", http.StatusInternalServerError)
      return
    }

	fmt.Fprintf(w, "Request processed and written to %s\n", database)
}

// Checks range of first character of the hash, and returns the correct/ appropriate database
// TODO: fix ths, make it hash based
// Introduce a switch case
// introduce the for range\
func getAppropriateDatabase(hashedString string) string {
	mutex.Lock()
	defer mutex.Unlock()
	switch hashedString[0] {
	case '0', '1', '2', '3':
		return "1db.txt"
	case '4', '5', '6', '7':
		return "2db.txt"
	case '8', '9', 'a', 'b':
		return "3db.txt"
	case 'c', 'd', 'e', 'f':
		return "4db.txt"
	default:
		return "" // Return empty string to signify an error
	}}

// If file contains content, it appends to the file
// If file contains no content, it writes to the file
// If file does not exist, it creates the file

func writeToFile(filename, content string) error {
	f, err := os.OpenFile(filename, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return err
	}
	defer f.Close()

	_, err = f.WriteString(content + "\n")
	return err
}
