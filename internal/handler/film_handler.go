package handler

import (
	"net/http"
	"rest-api-go/internal/model"
	"rest-api-go/internal/service"
	"strconv"

	"github.com/gin-gonic/gin"
)

// struct handler memegang service
type FilmHandler struct {
	service service.FilmService
}

// constructor
func NewFilmHandler(s service.FilmService) *FilmHandler {
	return &FilmHandler{service: s}
}

// handler create POST
func (h *FilmHandler) CreateFilm(c *gin.Context) {
	var film model.Film

	// Binding JSON dari request body ke struct
	if err := c.ShouldBindJSON(&film); err != nil {
		c.JSON(http.StatusBadGateway, gin.H{
			"error": "format json tidak valid",
		})
		return
	}

	// panggil service
	if err := h.service.CreateFilm(&film); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"message": "Film berhasil ditambahkan",
		"data":    film,
	})
}

// handler GetAll (GET)
func (h *FilmHandler) GetAllFilms(c *gin.Context) {
	films, err := h.service.GetAllFilms()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Daftar film",
		"data":    films,
	})
}

// handler GetByID(GET)
func (h *FilmHandler) GetFilmByID(c *gin.Context) {
	//ambil id dari url parameter
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "ID harus berupa angka",
		})
	}

	film, err := h.service.GetFilmByID(uint(id))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "film tidak ditemukan",
		})
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Detail film",
		"data":    film,
	})
}

// handler update (PUT)
func (h *FilmHandler) UpdateFilm(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "ID harus berupa angka",
		})
	}

	var film model.Film
	if err := c.ShouldBindJSON(&film); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Format JSON tidak valid",
		})
		return
	}

	// set ID dari URL
	film.ID = uint(id)
	if err := h.service.UpdateFilm(&film); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Film berhasil di update",
		"data":    film,
	})

}

// delete (DELETE)
func (h *FilmHandler) DeleteFilm(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "ID harus berupa angka",
		})
	}

	if err := h.service.DeleteFilm(uint(id)); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Film berhasil dihapus",
	})
}
