package main

import (
	"errors"
	"fmt"

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
	followService := models.NewFollow()
	timelineService := models.NewTimeline(rdb)

	e := echo.New()
	e.POST("/post", func(c echo.Context) error {
		ctx := c.Request().Context()
		content := c.Request().PostFormValue("post")
		author := c.Request().PostFormValue("author")
		post := models.Post{Post: content, Author: author}
		post, err := postService.Save(c.Request().Context(), post)
		followers, _ := followService.Followers(ctx, author)

		for _, follower := range followers {
			if err := timelineService.Push(ctx, follower, post.Id); err != nil {
				return c.String(500, err.Error())
			}
		}

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

	e.GET("/timeline/:user", func(c echo.Context) error {
		ctx := c.Request().Context()
		user := c.Param("user")
		postIDs, err := timelineService.Latest(ctx, user, 10)

		if errors.Is(err, redis.Nil) {
			return c.String(404, "timeline not found")
		} else if err != nil {
			return c.String(500, err.Error())
		}

		// implement this
		posts, err := postService.MGet(ctx, postIDs...)
		if err != nil {
			return c.String(500, err.Error())
		}

		// form result
		timeline := ""
		for i := len(posts) - 1; i >= 0; i-- {
			post := posts[i]
			timeline += fmt.Sprintf("%s: %s\n__________________\n", post.Author, post.Post)
		}

		return c.String(200, timeline)
	})
	e.Logger.Fatal(e.Start(":8000"))
}
