package main

import (
	"flag"
	"log"
	"net/http"
	"path/filepath"
)

func main() {
	addr := flag.String("addr", "127.0.0.1:4000", "Сетевой адрес HTTP")
	flag.Parse()

	mux := http.NewServeMux()
	mux.HandleFunc("/", home)
	mux.HandleFunc("/snippet", showSnippet)
	mux.HandleFunc("/snippet/create", createSnippet)

	fileServer := http.FileServer(neuteredFileSystem{http.Dir("./ui/static/")})
	//mux.Handle("/static", http.NotFoundHandler())
	mux.Handle("/static/", http.StripPrefix("/static", fileServer))

	log.Printf("Запуск сервера на %s", *addr)
	err := http.ListenAndServe(*addr, mux)
	log.Fatal(err)
}

// Проверка страницы на вывод каталогов которые должны быть закрыты
type neuteredFileSystem struct {
	fs http.FileSystem
}

func (nfs neuteredFileSystem) Open(path string) (http.File, error) {
	f, err := nfs.fs.Open(path)
	if err != nil {
		return nil, err
	}

	s, err := f.Stat()
	if s.IsDir() {
		index := filepath.Join(path, "index.html")
		if _, err := nfs.fs.Open(index); err != nil {
			closeErr := f.Close()
			if closeErr != nil {
				return nil, closeErr
			}

			return nil, err
		}
	}

	return f, nil
}
