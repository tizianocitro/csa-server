package controller

import (
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/tizianocitro/csa-server/model"
)

func GetStories(c *fiber.Ctx) error {
	return c.JSON(stories)
}

func GetStory(c *fiber.Ctx) error {
	id := c.Params("storyId")
	index, err := strconv.Atoi(id)
	if err != nil {
		return c.JSON(model.Story{})
	}
	story := stories[index]
	return c.JSON(story)
}

var stories = []model.Story{
	{
		ID:          "0",
		Name:        "sX",
		Description: "X story",
	},
	{
		ID:          "1",
		Name:        "sY",
		Description: "Y story",
	},
	{
		ID:          "2",
		Name:        "sZ",
		Description: "Z story",
	},
}
