package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type AboutMe struct {
	Name       string   `json:"name"`
	Experience string   `json:"experience"`
	Skills     []string `json:"skills"`
	Interest   string   `json:"interest"`
}

type WhyGo struct {
	Reason       string `json:"reason"`
	Expectations string `json:"expectations"`
}

// Обработчик для эндпоинта /api/about-me
func aboutMeHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	//fmt.Fprintln(w, "Привет! Меня зовут Богдан, я занимаюсь программированием последние 6 лет. Работал с C++ и немного с JavaScript. Люблю backend-разработку и автоматизацию.")
	response := AboutMe{
		Name:       "Богдан",
		Experience: "6 лет в программировании",
		Skills:     []string{"C++", "JavaScript", "Backend", "Автоматизация"},
		Interest:   "Люблю backend-разработку и автоматизацию",
	}

	json.NewEncoder(w).Encode(response)
}

// Обработчик для эндпоинта /api/why-go
func whyGoHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	//fmt.Fprintln(w, "Мне интересно изучать Go, потому что это быстрый, минималистичный язык с отличной поддержкой многопоточности. Хочу научиться писать производительные веб-сервисы и понимать, как строить масштабируемые системы.")
	response := WhyGo{
		Reason:       "Go — быстрый, минималистичный язык с отличной поддержкой многопоточности",
		Expectations: "Хочу научиться писать производительные веб-сервисы и понимать, как строить масштабируемые системы",
	}

	json.NewEncoder(w).Encode(response)
}

// Обработчик для корня "/"
// Чтобы при заходе на https://...tuna.am не было 404
func rootHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	fmt.Fprintln(w, `
		<!DOCTYPE html>
		<html lang = "ru">
		<head>
			<meta charset="UTF-8">
			<title>Go Webserver</title>
		</head>
		<body style="font-family: sans-serif; padding: 20px;">
			<h2>Добро пожаловать!</h2>
			<p>Доступные эндпоинты:</p>
			<ul>
				<li><a href="/api/about-me">/api/about-me</a></li>
				<li><a href="/api/why-go">/api/why-go</a></li>
			</ul>
		</body>
		</html>
	`)
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
