package router

import (
	"rest-api-go/internal/handler"

	"github.com/gin-gonic/gin"
)

func SetupRouter(filmHandler *handler.FilmHandler) *gin.Engine {
	// inisiasi Gin
	r := gin.Default()

	api := r.Group("/api")
	{
		films := api.Group("/films")
		{
			films.POST("/", filmHandler.CreateFilm)
			films.GET("/", filmHandler.GetAllFilms)
			films.GET("/:id", filmHandler.GetFilmByID)
			films.PUT("/:id", filmHandler.UpdateFilm)
			films.DELETE("/:id", filmHandler.DeleteFilm)
		}
	}

	return r
}
