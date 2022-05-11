package handler

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/intrntsrfr/ghost/db"
	"github.com/intrntsrfr/ghost/model"
	"go.uber.org/zap"
	"math/rand"
	"net/http"
	"path/filepath"
	"time"
)

type Handler struct {
	G      *gin.Engine
	d      db.DB
	logger *zap.Logger
}

func NewHandler(d db.DB, logger *zap.Logger) *Handler {
	r := gin.Default()
	h := &Handler{r, d, logger}

	r.POST("/upload", h.upload())
	r.GET("/:imageID", h.findFile())

	return h
}

func (h *Handler) upload() gin.HandlerFunc {
	return func(c *gin.Context) {
		file, err := c.FormFile("file")
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"message": "no file present in multipart/form-data body",
			})
			return
		}

		ext := filepath.Ext(file.Filename)
		hash := generate(10)
		for h.d.GetUploadByHash(hash) != nil {
			hash = generate(10)
		}

		u := &model.Upload{
			Hash:      hash,
			Extension: ext,
			Created:   time.Now(),
			AccountID: 1,
		}

		s, err := h.d.CreateUpload(u)
		if err != nil {
			fmt.Println(err)
			c.JSON(http.StatusInternalServerError, gin.H{"message": "something happened"})
			return
		}

		err = c.SaveUploadedFile(file, fmt.Sprintf("./_storage/%v%v", u.Hash, u.Extension))
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"message": "something happened"})
			return
		}

		c.JSON(http.StatusCreated, gin.H{
			"url":  s.String(),
			"size": file.Size,
			"hash": file.Filename,
		})
	}
}

func (h *Handler) findFile() gin.HandlerFunc {
	return func(c *gin.Context) {
		weed := c.Params.ByName("imageID")
		c.File("./_storage/" + weed)
	}
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
