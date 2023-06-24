package main

import (
	"errors"

	"github.com/labstack/echo/v4"
	"github.com/nlebedevinc/postfeed/internal/models"
	"github.com/nlebedevinc/postfeed/internal/services"
	"github.com/redis/go-redis/v9"
)

func main() {
	rdb := redis.NewClient(&redis.Options{
		Addr: "localhost: 6379",
	})

	postService := services.NewPostService(services.NewRedis[models.Post](rdb))
	e := echo.New()
	e.POST("/post", func(c echo.Context) error {
		content := c.Request().PostFormValue("post")
		author := c.Request().PostFormValue("author")
		post := models.Post{Post: content, Author: author}
		post, err := postService.Save(c.Request().Context(), post)
		if err != nil {
			return c.String(500, err.Error())
		}

		return c.String(201, post.Id)
	})

	e.GET("/post/:uid", func(c echo.Context) error {
		uid := c.Param("uid")
		post, err := postService.Get(c.Request().Context(), uid)
		if errors.Is(err, redis.Nil) {
			return c.String(404, "post not found")
		} else if err != nil {
			return c.String(500, err.Error())
		}

		return c.String(200, post.Author+" : "+post.Post)
	})
	e.Logger.Fatal(e.Start(":8000"))
}
