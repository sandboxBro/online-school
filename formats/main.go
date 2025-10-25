package main

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
	"net/http"
)

// Простая структура для JSON и XML
type Message struct {
	Text string `json:"text" xml:"text"`
}

func textHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/plain; charset=utf-8")
	fmt.Fprintln(w, "Привет! Это ответ в формате TEXT.")
}

func htmlHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	fmt.Fprintln(w, "<h1>Привет !</h1><p>Это ответ в формате HTML.</p>")
}

func jsonHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	msg := Message{Text: "Привет! Это ответ в формате JSON."}
	json.NewEncoder(w).Encode(msg)
}

func xmlHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/xml")
	msg := Message{Text: "Привет! Это ответ в формате XML."}
	xml.NewEncoder(w).Encode(msg)
}

func main() {
	http.HandleFunc("/text", textHandler)
	http.HandleFunc("/html", htmlHandler)
	http.HandleFunc("/json", jsonHandler)
	http.HandleFunc("/xml", xmlHandler)
	fmt.Println("Сервер запущен на http://127.0.0.1:8080 ...")
	http.ListenAndServe(":8080", nil)
}
