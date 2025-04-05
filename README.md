tags:
golang
rest-api
gorm
gin
REST API: Go мен Gin пайдаланып көлік басқару жүйесі
Бұл құжат Go тілінде жазылған көлік басқару жүйесіне арналған REST API жобасының құрылымын түсіндіреді. Жоба Gin фреймворкі мен GORM ORM кітапханасын қолдана отырып, PostgreSQL дерекқорымен жұмыс істейді. Біз көлік туралы ақпаратты CRUD операциялары арқылы өңдейміз.
1. 
   Жобаның мақсаты
   Бұл жобаның мақсаты — көлік деректерін сақтау, оқу, жаңарту және жою үшін API жасау. Пайдаланушы көлік туралы мәліметтерді енгізіп, оларды өңдей алады.

Қолданылатын құралдар:
Go — негізгі бағдарламалау тілі.
Gin — HTTP сервері мен маршрутизация үшін қолданылатын фреймворк.
GORM — дерекқормен жұмыс істеу үшін ORM.
PostgreSQL — дерекқор ретінде пайдаланылады.
Docker — жобаны контейнерлеу үшін.
2. 
  Жобаны орнату
   Қажетті құралдарды орнату
   Go: https://golang.org/dl/

Docker: https://www.docker.com/get-started

PostgreSQL: Docker контейнерімен бірге орнатылады.

Жоба құрылымы
go
Копировать
Редактировать
crudproject/
│
├── internal/
│   ├── handler/
│   │   └── car_handler.go
│   ├── models/
│   │   └── car.go
│   ├── repository/
│   │   └── car_repository.go
│   ├── routes/
│   │   └── routes.go
│   └── service/
│       └── car_service.go
│
├── Dockerfile
├── docker-compose.yml
└── main.go
3. main.go файлы
   Go кодын орындау
   main.go — жобаның басты файлы. Онда дерекқормен қосылу, маршрутизация, және API серверін іске қосу әрекеттері жүзеге асырылады.
   go
   Копировать
   Редактировать
   package main

import (
"crudproject/internal/handler"
"crudproject/internal/models"
"crudproject/internal/repository"
"crudproject/internal/routes"
"crudproject/internal/service"
"github.com/gin-gonic/gin"
"gorm.io/driver/postgres"
"gorm.io/gorm"
"log"
)

func main() {
// PostgreSQL-ге қосылу
dsn := "postgres://postgres:7982@postgres:5432/postgres?sslmode=disable"
db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
if err != nil {
log.Fatal("DB connection error:", err)
}

	// Миграция жасау
	if err := db.AutoMigrate(&models.Car{}); err != nil {
		log.Fatal("Migration error:", err)
	}

	// Репозиторий, қызмет және хендлер жасау
	carRepo := repository.NewCarRepository(db)
	carService := service.NewCarService(carRepo)
	carHandler := handler.NewCarHandler(carService)

	// Gin серверін орнату
	r := gin.Default()
	routes.SetupRoutes(r, carHandler)
	r.Run(":8080") // Серверді 8080 портында іске қосу
}
Бұл кодта:

PostgreSQL дерекқорымен қосылып, байланыс орнату.

Car моделін дерекқорында миграциялау.

Handler, Service, және Repository қабаттарын орнату.

API маршрутарын орнату және серверді іске қосу.

4. Жобаға арналған Docker конфигурациясы
   Жобаны Docker арқылы контейнерлеу үшін Dockerfile және docker-compose.yml файлдарын құру қажет.

Dockerfile
dockerfile
Копировать
Редактировать
# Go ортасын орнату
FROM golang:1.24.2-alpine as builder

WORKDIR /app
COPY . .

# Зависимостерді орнату
RUN go mod tidy

# Жобаны құру
RUN go build -o main main.go

# Alpine бейнесін қолдана отырып жеңіл бейне жасау
FROM alpine:latest
WORKDIR /root/
COPY --from=builder /app/main .

CMD ["./main"]
docker-compose.yml
yaml
Копировать
Редактировать
version: "3.9"
services:
go-app:
build: .
ports:
- "8080:8080"
depends_on:
- go-postgres
go-postgres:
image: postgres:latest
environment:
POSTGRES_USER: postgres
POSTGRES_PASSWORD: 7982
POSTGRES_DB: postgres
volumes:
- postgres_data:/var/lib/postgresql/data
ports:
- "5432:5432"
volumes:
postgres_data:
Бұл конфигурация:

Go қосымшасын контейнерлеу үшін Dockerfile қолданады.

docker-compose.yml екі контейнерді біріктіреді: біріншісі Go қосымшасы, екіншісі PostgreSQL дерекқоры.

5. Қосымша түсініктемелер
   Моделдер мен CRUD операциялары
   Car моделін GORM арқылы анықтаймыз:

go
Копировать
Редактировать
package models

type Car struct {
ID        uint   `gorm:"primaryKey"`
Make      string `json:"make"`
Model     string `json:"model"`
Year      int    `json:"year"`
Price     float64 `json:"price"`
}
Бұл модель көліктің маркасын, моделін, жылын және бағасын сипаттайды.

Репозиторий
CarRepository дерекқорға сұраныстарды жіберу үшін:

go
Копировать
Редактировать
package repository

import (
"crudproject/internal/models"
"gorm.io/gorm"
)

type CarRepository struct {
DB *gorm.DB
}

func NewCarRepository(db *gorm.DB) *CarRepository {
return &CarRepository{DB: db}
}

func (r *CarRepository) CreateCar(car *models.Car) error {
return r.DB.Create(car).Error
}

func (r *CarRepository) GetAllCars() ([]models.Car, error) {
var cars []models.Car
err := r.DB.Find(&cars).Error
return cars, err
}
Бұл репозиторий дерекқордан көлік мәліметтерін алады, жаңа көлік қосады, және басқа CRUD операцияларын жүзеге асырады.

6. API маршруты
   Routes файлында API маршруты орнатылады:

go
Копировать
Редактировать
package routes

import (
"crudproject/internal/handler"
"github.com/gin-gonic/gin"
)

func SetupRoutes(r *gin.Engine, h *handler.CarHandler) {
r.GET("/cars", h.GetAllCars)
r.POST("/cars", h.CreateCar)
}
Бұл маршрутар келесі функционалдарды қамтамасыз етеді:

GET /cars — барлық көліктерді алу.

POST /cars — жаңа көлік қосу.

7. Қорытынды
   Бұл жоба Go, Gin және GORM арқылы құрылған қарапайым REST API үлгісін көрсетеді, онда PostgreSQL дерекқоры мен Docker контейнерлеу жүйесі пайдаланылады. Әрбір компонент өзара үйлесімді жұмыс істейді, бұл сіздің API-дің жоғары өнімділігін қамтамасыз етеді.

