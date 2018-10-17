package main

import (
	"./model"
	"flag"
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/julienschmidt/httprouter"
	"log"
	"net/http"
)

var (
	addr = flag.String("addr", "postgresql://productuser@localhost:26257/product?sslmode=disable", "database connection string")
)

func main() {
	flag.Parse()

	db := setupDB(*addr)
	defer db.Close()

	router := httprouter.New()

	server := NewServer(db)
	server.RegisterRouter(router)

	log.Fatal(http.ListenAndServe(":6543", router))
}

func setupDB(addr string) *gorm.DB {
	db, err := gorm.Open("postgres", addr)
	if err != nil {
		panic(fmt.Sprintf("failed to connect to database: %v", err))
	}
	db.Exec("SET search_path TO public")
	// Migrate the schema
	db.AutoMigrate(&model.Customer{}, &model.Order{}, &model.Product{})

	return db
}
