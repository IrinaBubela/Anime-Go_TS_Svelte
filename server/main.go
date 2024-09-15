package main

import (
	"log/slog"
	"net/http"
	"os"
	"task/services"
	"task/system"
	"time"
	"fmt"
	"github.com/labstack/echo/v4"
	"github.com/lmittmann/tint"
)

type ApiResponse struct {
    Success bool   `json:"success"`
    Error   string `json:"error"`
}

type Favorite struct {
	ID string `json:"id"` 
    AnimeID      string `json:"anime_id"`
	AnimeTitle   string `json:"anime_title"`
	AnimeImageURL string `json:"anime_image_url"`
	Created string `json:"created"`
	Updated string `json:"updated"`
}

func main() {
	// Set up the logger
	w := os.Stderr
	var log slog.Level
	if system.LOG_LEVEL == "info" {
		log = slog.LevelInfo
	} else {
		log = slog.LevelDebug
	}
	slog.SetDefault(slog.New(
		tint.NewHandler(w, &tint.Options{
			AddSource:  true,
			Level:      log,
			TimeFormat: time.Kitchen,
		}),
	))

	// Connect to the database
	err := system.Connect()
	if err != nil {
		slog.Error("Error opening database", "db.Connect", err)
		panic(err)
	}
	slog.Info("Database connected")

	// Run migrations
	err = system.Migrations()
	if err != nil {
		slog.Error("Error running migrations", "db.Migrations", err)
		panic(err)
	}
	slog.Info("Migrations completed")

	// Run the HTTP server
	e := echo.New()
	e.GET("/", func(c echo.Context) error {
		id := 0
		err := system.Db.QueryRow("SELECT 1").Scan(&id)
		if err != nil {
			slog.Error("Error pinging database", "Db.QueryRow", err)
			return c.String(http.StatusInternalServerError, "Error pinging database")
		}
		return c.String(http.StatusOK, "Hello, World!")
	})

	e.GET("/favorites", func(c echo.Context) error {
		return services.GetFavorites(c)
	})

	e.POST("/favorites", AddFavorite)

	e.DELETE("/favorites/:id", RemoveFavorite)

	slog.Info("HTTP server listening on", "port", system.HTTP_PORT)
	err = e.Start(":" + system.HTTP_PORT)
	if err != nil {
		slog.Error("Error serving HTTP", "e.Start", err)
		panic(err)
	}
}

func AddFavorite(c echo.Context) error {
	var fav services.Favorite
	
	if err := c.Bind(&fav); err != nil {
		fmt.Println(err, " err 1 cheeeeeeeeeeeck")
		return c.String(http.StatusBadRequest, "Invalid input")
	}

	// Check if the user has reached the limit of 5 favorites
	count, err := services.CountFavorites()
	if err != nil {
		slog.Error("Error counting favorites", "error", err)
		return c.JSON(http.StatusInternalServerError, ApiResponse{
			Success: false,
			Error:   "Error checking favorite limit",
		})
	}

	if count >= 5 {
		slog.Error("Favorite limit reached", "count", count)
		return c.String(http.StatusBadRequest,  "Maximum limit of 5 favorites reached")
	}

	// Insert the new favorite
	if err := services.InsertFavorite(&fav); err != nil {
		fmt.Println(err, "Error inserting favorite")
		return c.String(http.StatusInternalServerError, "Error adding favorite")
	}

	return c.JSON(http.StatusOK, fav)
}

func RemoveFavorite(c echo.Context) error {
    id := c.Param("id")
    if id == "" {
        return c.JSON(http.StatusBadRequest, ApiResponse{
            Success: false,
            Error:   "Missing favorite id",
        })
    }

    err := services.DeleteFavorite(id)
    if err != nil {
        slog.Error("Error removing favorite", "id", id, "error", err)
        return c.JSON(http.StatusInternalServerError, ApiResponse{
            Success: false,
            Error:   "Error removing favorite",
        })
    }

    return c.JSON(http.StatusOK, ApiResponse{
        Success: true,
        Error:   "",
    })
}