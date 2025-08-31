content = """# Go Webserver Test

Простой веб-сервер на Go, выполненный в рамках тестового задания.

##  Функциональность

Сервер реализует три эндпоинта:

- `/` — корневая страница (HTML) со списком доступных API.  
- `/api/about-me` — возвращает информацию обо мне в формате JSON.  
- `/api/why-go` — возвращает информацию о том, почему я выбрал Go, в формате JSON.  

##  Структура проекта

.
├── main.go # основной код сервера
└── README.md # документация проекта



##  Запуск

1. Убедитесь, что Go установлен:
go version


2. Клонируйте репозиторий:
git clone https://github.com/username/go-webserver-test.git
cd go-webserver-test


3. Запустите сервер:
go run main.go


4. Сервер будет доступен по адресу:
http://localhost:8080

##  Доступ из интернета

Для глобального доступа можно использовать сервис **tuna.am**.  
После запуска туннеля (`tuna http 8080`) сервер доступен по выданному адресу, например:

https://y7wrkb-2a10-c943-100--a22.ru.tuna.am



##  Примеры запросов


### GET /api/about-me  

Ответ:
{
"name": "Богдан",
"experience": "6 лет в программировании",
"skills": ["C++", "JavaScript", "Backend", "Автоматизация"],
"interest": "Люблю backend-разработку и автоматизацию"
}


### GET /api/why-go  

Ответ:
{
"reason": "Go — быстрый, минималистичный язык с отличной поддержкой многопоточности",
"expectations": "Хочу научиться писать производительные веб-сервисы и понимать, как строить масштабируемые системы"
}

##  Документация API (OpenAPI 3.0)

openapi: 3.0.0
info:
title: Go Webserver Test
version: 1.0.0
description: |
Простой веб-сервер на Go с тремя эндпоинтами:
- /
- /api/about-me
- /api/why-go
servers:

url: http://localhost:8080

url: https://y7wrkb-2a10-c943-100--a22.ru.tuna.am
paths:
/api/about-me:
get:
summary: Информация обо мне
responses:
'200':
description: JSON с информацией обо мне
content:
application/json:
schema:
type: object
properties:
name:
type: string
experience:
type: string
skills:
type: array
items:
type: string
interest:
type: string
/api/why-go:
get:
summary: Почему выбрал Go
responses:
'200':
description: JSON с причинами изучения Go и ожиданиями
content:
application/json:
schema:
type: object
properties:
reason:
type: string
expectations:
type: string


##  Автор

**Богдан**
"""

with open("/mnt/data/README.md", "w", encoding="utf-8") as f:
    f.write(content)

"/mnt/data/README.md"
Результат
'/mnt/data/README.md'


