package main

import (
	"html/template"
	"log"
	"net/http"
	"strconv"
)

var tmpl = template.Must(template.ParseGlob("templates/*.html"))
var votes = make(map[int]int)

const password = "200817"

// Kandidat adalah struct untuk setiap kandidat voting
type Kandidat struct {
	ID    int
	Nama  string
	Votes int
}

// KandidatList adalah daftar kandidat
var KandidatList = []Kandidat{
	{ID: 1, Nama: "Kandidat 1"},
	{ID: 2, Nama: "Kandidat 2"},
	{ID: 3, Nama: "Kandidat 3"},
	{ID: 4, Nama: "Kandidat 4"},
	{ID: 5, Nama: "Kandidat 5"},
	{ID: 6, Nama: "Kandidat 6"},
	{ID: 7, Nama: "Kandidat 7"},
	{ID: 8, Nama: "Kandidat 8"},
	{ID: 9, Nama: "Kandidat 9"},
	{ID: 10, Nama: "Kandidat 10"},
	{ID: 11, Nama: "Kandidat 11"},
	{ID: 12, Nama: "Kandidat 12"},
	{ID: 13, Nama: "Kandidat 13"},
	{ID: 14, Nama: "Kandidat 14"},
	{ID: 15, Nama: "Kandidat 15"},
	{ID: 16, Nama: "Kandidat 16"},
	{ID: 17, Nama: "Kandidat 17"},
	{ID: 18, Nama: "Kandidat 18"},
	{ID: 19, Nama: "Kandidat 19"},
	{ID: 20, Nama: "Kandidat 20"},
	{ID: 21, Nama: "Kandidat 21"},
	{ID: 22, Nama: "Kandidat 22"},
	{ID: 23, Nama: "Kandidat 23"},
	{ID: 24, Nama: "Kandidat 24"},
	{ID: 25, Nama: "Kandidat 25"},
	{ID: 26, Nama: "Kandidat 26"},
	{ID: 27, Nama: "Kandidat 27"},
	{ID: 28, Nama: "Kandidat 28"},
	{ID: 29, Nama: "Kandidat 29"},
	{ID: 30, Nama: "Kandidat 30"},
	{ID: 31, Nama: "Kandidat 31"},
	{ID: 32, Nama: "Kandidat 32"},
	{ID: 33, Nama: "Kandidat 33"},
	{ID: 34, Nama: "Kandidat 34"},
	{ID: 35, Nama: "Kandidat 35"},
	{ID: 36, Nama: "Kandidat 36"},
	{ID: 37, Nama: "Kandidat 37"},
	{ID: 38, Nama: "Kandidat 38"},
	{ID: 39, Nama: "Kandidat 39"},
	{ID: 40, Nama: "Kandidat 40"},
}

func main() {
	http.HandleFunc("/", votingHandler)
	http.HandleFunc("/results", resultsHandler)
	http.HandleFunc("/thank-you", thankYouHandler) // New handler
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))

	log.Println("Server berjalan di :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

// votingHandler menangani halaman voting
func votingHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		id, err := strconv.Atoi(r.FormValue("kandidat"))
		if err == nil && id >= 1 && id <= len(KandidatList) {
			votes[id]++
			http.Redirect(w, r, "/thank-you", http.StatusSeeOther) // Redirect to thank-you page
			return
		}
	}
	tmpl.ExecuteTemplate(w, "voting.html", KandidatList)
}

// resultsHandler menangani halaman hasil voting yang dilindungi dengan password
func resultsHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		inputPassword := r.FormValue("password")
		if inputPassword == password {
			for i := range KandidatList {
				KandidatList[i].Votes = votes[KandidatList[i].ID]
			}
			tmpl.ExecuteTemplate(w, "results.html", KandidatList)
			return
		}
		http.Error(w, "Password salah!", http.StatusUnauthorized)
		return
	}
	tmpl.ExecuteTemplate(w, "results.html", nil)
}

// thankYouHandler menangani halaman terima kasih
func thankYouHandler(w http.ResponseWriter, r *http.Request) {
	tmpl.ExecuteTemplate(w, "thank-you.html", nil) // Serve the thank-you template
}
