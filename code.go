package main

import (
	"fmt"
	"log"
	"net/http"
)

// Обработчик для эндпоинта /api/about-me
func aboutMeHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/plain; charset=utf-8")
	fmt.Fprintln(w, "Привет! Меня зовут Богдан, я занимаюсь программированием последние 6 лет. Работал с C++ и немного с JavaScript. Люблю backend-разработку и автоматизацию.")
}

// Обработчик для эндпоинта /api/why-go
func whyGoHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/plain; charset=utf-8")
	fmt.Fprintln(w, "Мне интересно изучать Go, потому что это быстрый, минималистичный язык с отличной поддержкой многопоточности. Хочу научиться писать производительные веб-сервисы и понимать, как строить масштабируемые системы.")
}

// Обработчик для корня "/"
// Чтобы при заходе на https://...tuna.am не было 404
func rootHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/plain; charset=utf-8")
	fmt.Fprintln(w, "Добро пожаловать! Доступные эндпоинты:\n/api/about-me\n/api/why-go")
}

func main() {
	// Регистрируем обработчики для роутов
	http.HandleFunc("/", rootHandler)
	http.HandleFunc("/api/about-me", aboutMeHandler)
	http.HandleFunc("/api/why-go", whyGoHandler)

	// Запускаем сервер на порту 8080
	fmt.Println("Сервер запущен на http://localhost:8080")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal("Ошибка запуска сервера: ", err)
	}
}
