package main

import (
	"bufio"
	"html/template"
	"log"
	"net/http"
	"os"
)

func home(w http.ResponseWriter, r *http.Request) {
	// Write txt file
	f, err := os.OpenFile("./txt/list.txt", os.O_RDWR|os.O_APPEND, 0660)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	// Read txt file
	file, err := os.Open("./txt/list.txt")
	if err != nil {
		log.Fatal(err)
	}

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	var txtlines []string
	for scanner.Scan() {
		txtlines = append(txtlines, scanner.Text())
	}

	file.Close()

	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}

	if r.Method == "GET" {
		t, _ := template.ParseFiles("./ui/html/form.html")
		t.Execute(w, nil)
	} else {
		w.Write([]byte("Thanks for subscribing!"))
		r.ParseForm()
		f.WriteString("\n")
		f.WriteString(r.Form.Get("phone-number"))
	}

	for _, eachline := range txtlines {
		mail(eachline)
	}
}
