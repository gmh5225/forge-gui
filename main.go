package main

import (
	"embed"
	"html/template"
	"log"
	"net/http"

	"github.com/jchv/go-webview2"
)

//go:embed internal/ui/template.html
var content embed.FS

func main() {
	w := webview2.NewWithOptions(webview2.WebViewOptions{
		Debug:     true,
		AutoFocus: true,
		WindowOptions: webview2.WindowOptions{
			Title:  "Forge GUI",
			Width:  800,
			Height: 600,
		},
	})
	defer w.Destroy()

	tmpl, err := template.ParseFS(content, "internal/ui/template.html")
	if err != nil {
		log.Fatal(err)
	}

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		tmpl.Execute(w, nil)
	})

	go http.ListenAndServe("127.0.0.1:8080", nil)

	w.Navigate("http://127.0.0.1:8080")
	w.Run()
}
