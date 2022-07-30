//Веб приложение через которое можно будет создавать и делится заметками.

package main

import (
	"log"
	"net/http"
)

//Создается функция-обработчик домашней страницы "home"
func home(w http.ResponseWriter, r *http.Request) {
	//Проверка на catch-all
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}
	w.Write([]byte("Привет из Snippetbox"))
}

//Создается функция-обработчик для страницы "showSnippet"
func showSnippet(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Выводи заметки из Snippetbox"))
}

//Создается функция-обработчик для страницы "creatSnippet" которая отвечает только на POST запросы
func creatSnippet(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		w.Header().Set("Allow", http.MethodPost)
		http.Error(w, "Метод запрещён", 405)
		return
	}
	w.Write([]byte("Создаёт заметки в  Snippetbox"))
}

func main() {
	//http.NewServeMux() для инициализации нового рутера
	//Регистрируется как обработчик для URL-шаблонов.
	mux := http.NewServeMux()
	mux.HandleFunc("/", home)
	mux.HandleFunc("/snippet", showSnippet)
	mux.HandleFunc("/snippet/create", creatSnippet)

	//Функция http.ListenAndServe() используется для запуска нового веб-сервера.
	log.Println("Запуск веб-сервера http://127.0.0.1:4000")
	err := http.ListenAndServe("127.0.0.1:4000", mux)
	log.Fatal(err)
}
