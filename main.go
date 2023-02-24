package main

import (
	"log"
	"net/http"
)

// Обработчик главной страницы.
func home(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}
	_, err := w.Write([]byte("Привет из Snippetbox"))
	if err != nil {
		return
	}
}

// Обработчик для отображения содержимого заметки.
func showSnippet(w http.ResponseWriter, _ *http.Request) {
	_, err := w.Write([]byte("Отображение заметки..."))
	if err != nil {
		return
	}
}

// Обработчик для создания новой заметки.
func createSnippet(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		w.Header().Set("Allow", http.MethodPost)
		http.Error(w, "Метод запрещен!", 405)
		return
	}
	_, err := w.Write([]byte("Форма для создания новой заметки..."))
	if err != nil {
		return
	}
}

func main() {
	// Регистрируем два новых обработчика и соответствующие URL-шаблоны в маршрутизаторе servemux
	mux := http.NewServeMux()
	mux.HandleFunc("/", home)
	mux.HandleFunc("/snippet", showSnippet)
	mux.HandleFunc("/snippet/create", createSnippet)

	log.Println("Запуск веб-сервера на http://127.0.0.1:4000")
	err := http.ListenAndServe(":4000", mux)
	log.Fatal(err)
}
