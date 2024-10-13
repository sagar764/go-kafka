package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gofiber/fiber"
	"github.com/sagar764/go-kafka/producer/entities"
	"github.com/sagar764/go-kafka/producer/helper"
)

func main() {
	app := fiber.New()
	api := app.Group("/api/v1")
	api.Post("/comments", createComment)
	app.Listen(":3000")
}

func createComment(c *fiber.Ctx) {
	cmt := new(entities.Comment)

	if err := c.BodyParser(cmt); err != nil {
		log.Println(err)
		c.Status(http.StatusBadRequest).JSON(&fiber.Map{
			"success": false,
			"message": err,
		})
		return
	}

	cmtInBytes, err := json.Marshal(cmt)
	if err != nil {
		log.Println(err)
		c.Status(http.StatusBadRequest).JSON(&fiber.Map{
			"success": false,
			"message": err.Error(),
		})
		return
	}

	if err := helper.PushCommnetToQueue("comments", cmtInBytes); err != nil {
		c.Status(http.StatusInternalServerError).JSON(&fiber.Map{
			"success": false,
			"message": err.Error(),
		})
	}

	c.Status(http.StatusOK).JSON(&fiber.Map{
		"success": true,
		"message": "Comment pushed successfully",
		"comment": cmt,
	})

}
