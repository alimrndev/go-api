# go-api
go api crud


docker-compose up --build -d
docker build -t go-api:latest .
docker run -p 8080:8080 --name go_backend_container go-api:latest 

module github.com/alimrndev/go-api

go 1.22.1

require (
    github.com/felixge/httpsnoop v1.0.4 // indirect
    github.com/gorilla/handlers v1.5.2
    github.com/gorilla/mux v1.8.1
    github.com/jinzhu/gorm v1.9.16
    github.com/jinzhu/inflection v1.0.0 // indirect
    github.com/joho/godotenv v1.5.1
    github.com/lib/pq v1.10.9 // indirect
    golang.org/x/crypto v0.22.0
)
