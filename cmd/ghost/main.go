package main

import (
	"fmt"
	"github.com/intrntsrfr/ghost/db"
	"github.com/intrntsrfr/ghost/handler"
	"github.com/jmoiron/sqlx"
	"go.uber.org/zap"
	"math/rand"
	"net/http"
	"os"
	"time"
)

func main() {
	rand.Seed(time.Now().Unix())

	DbUser, DbPass := os.Getenv("POSTGRES_USER"), os.Getenv("POSTGRES_PASSWORD")

	fmt.Println(DbUser, DbPass)

	psql := sqlx.MustConnect("postgres", fmt.Sprintf("host=database user=%s password=%s dbname=%s sslmode=disable",
		DbUser, DbPass, DbUser))

	logger, err := zap.NewDevelopment()
	if err != nil {
		panic(err)
	}

	d := db.NewPsqlDB(psql, logger.Named("database"))

	if _, err := os.Stat("./_storage"); os.IsNotExist(err) {
		err := os.Mkdir("./_storage", os.ModeDir)
		if err != nil {
			panic(err)
		}
	}

	r := handler.NewHandler(d, logger.Named("handler"))

	http.ListenAndServe(":8080", r.G)
}
