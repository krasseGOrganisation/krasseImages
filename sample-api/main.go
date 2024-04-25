package main

import (
	"encoding/json"
	"log"
	"net/http"
	"os"
	"sample-api/db"
	"sample-api/db/entities"
	"time"

	"github.com/goombaio/namegenerator"
	"github.com/joho/godotenv"
)

var nameGenerator namegenerator.Generator

func main() {

	godotenv.Load()

	db.Initialize()

	addr := os.Getenv("ADDR")
	if addr == "" {
		addr = ":80"
	}

	seed := time.Now().UTC().UnixNano()
	nameGenerator = namegenerator.NewNameGenerator(seed)

	http.HandleFunc("/", index)
	http.HandleFunc("/list", listAllEntries)
	http.HandleFunc("/create", createEntry)

	log.Println("HTTP Server will start on " + addr)
	err := http.ListenAndServe(addr, nil)
	if err != nil {
		log.Println("ERROR ON STARTUP")
	}
}

func listAllEntries(w http.ResponseWriter, r *http.Request) {
	entries := []entities.Entry{}
	if err := db.DB.Find(&entries).Error; err != nil {
		log.Println(err)
	}

	data, err := json.Marshal(entries)

	if err != nil {
		log.Println("ERROR", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Write(data)
}

func createEntry(w http.ResponseWriter, r *http.Request) {
	entry := entities.Entry{
		Name: nameGenerator.Generate(),
	}

	if err := db.DB.Save(&entry).Error; err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Write([]byte("Created Successfully"))
}

func index(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte(`<h1>Sample Application</h1><hr/><a href="/list">List all Persons</a><br /><a href="/create">Create a new Person</a>`))
}
