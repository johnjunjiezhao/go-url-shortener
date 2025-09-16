package handler

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis/v8"
	"github.com/johnjunjiezhao/go-url-shortener/shortener"
	"github.com/johnjunjiezhao/go-url-shortener/store"
)

// Request model definition
type URLCreationRequest struct {
	LongURL string `json:"long_url" binding:"required"`
	UserID  string `json:"user_id" binding:"required"`
}

type errorResponse struct {
	Message string `json:"message"`
}

func writeError(c *gin.Context, status int, message string) {
	c.AbortWithStatusJSON(status, errorResponse{Message: message})
}

var errShortURLNotFound = errors.New("short url not found")

func saveURLMapping(shortURL, longURL, userID string) error {
	return capturePanic(func() {
		store.SaveURLMapping(shortURL, longURL, userID)
	})
}

func retrieveOriginalURL(shortURL string) (string, error) {
	var (
		result string
		err    error
	)
	err = capturePanic(func() {
		result = store.RetrieveOriginalURL(shortURL)
	})
	if err != nil {
		if errors.Is(err, redis.Nil) {
			return "", errShortURLNotFound
		}
		return "", err
	}
	return result, nil
}

func capturePanic(fn func()) (err error) {
	defer func() {
		if r := recover(); r != nil {
			err = normalizeRecovered(r)
		}
	}()
	fn()
	return err
}

func normalizeRecovered(r interface{}) error {
	switch v := r.(type) {
	case error:
		return v
	case string:
		return errors.New(v)
	default:
		return fmt.Errorf("%v", v)
	}
}

func CreateShortURL(c *gin.Context) {
	var creationRequest URLCreationRequest
	if err := c.ShouldBindJSON(&creationRequest); err != nil {
		writeError(c, http.StatusBadRequest, err.Error())
		return
	}

	shortURL := shortener.GenerateShortLink(creationRequest.LongURL, creationRequest.UserID)
	if err := saveURLMapping(shortURL, creationRequest.LongURL, creationRequest.UserID); err != nil {
		writeError(c, http.StatusInternalServerError, "failed to store short url")
		return
	}

	host := "http://localhost:9808/"
	c.JSON(200, gin.H{
		"message":   "short url created successfully",
		"short_url": host + shortURL,
	})

}

func HandleShortURLRedirect(c *gin.Context) {
	shortURL := c.Param("short-url")
	initialURL, err := retrieveOriginalURL(shortURL)
	if err != nil {
		if errors.Is(err, errShortURLNotFound) {
			writeError(c, http.StatusNotFound, "short url not found")
			return
		}
		writeError(c, http.StatusInternalServerError, "failed to retrieve url")
		return
	}

	c.Redirect(http.StatusFound, initialURL)
}
