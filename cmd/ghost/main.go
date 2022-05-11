package main

import (
	"fmt"
	"github.com/intrntsrfr/ghost/db"
	"github.com/jmoiron/sqlx"
	"go.uber.org/zap"
	"math/rand"
	"net/http"
	"os"
	"path/filepath"
	"time"

	"github.com/gin-gonic/gin"
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

	r := gin.Default()

	r.POST("/upload", upload(d))
	r.GET("/:imageID", findFile())

	fmt.Println(generate(10))

	http.ListenAndServe(":8080", r)
}

var letterRunes = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

func generate(n int) string {
	b := make([]rune, n)
	for i := range b {
		b[i] = letterRunes[rand.Intn(len(letterRunes))]
	}
	return string(b)
}

func Auth() gin.HandlerFunc {
	return func(c *gin.Context) {
		key := c.GetHeader("Authorization")
		if key == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"message": "Unauthorized"})
			c.Abort()
		}
		c.Set("auth", key)
		c.Next()
	}
}

func upload(d *db.PsqlDB) gin.HandlerFunc {
	return func(c *gin.Context) {
		file, err := c.FormFile("file")
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"message": "no file provided",
			})
			return
		}

		ext := filepath.Ext(file.Filename)
		name := generate(10)
		/*
			s, err := d.CreateUpload(name, ext, 1)
			if err != nil {
				fmt.Println(err)
				c.JSON(http.StatusInternalServerError, gin.H{"message": "something happened"})
				return
			}*/

		s := name + ext

		fmt.Println(s)

		err = c.SaveUploadedFile(file, "./_storage/"+s)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"message": "something happened"})
			return
		}

		c.JSON(http.StatusCreated, gin.H{
			"file": s,
		})
	}
}

func findFile() gin.HandlerFunc {
	return func(c *gin.Context) {
		weed := c.Params.ByName("imageID")
		c.File("./_storage/" + weed)
	}
}
