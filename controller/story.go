package controller

import (
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/tizianocitro/csa-server/model"
)

func GetStories(c *fiber.Ctx) error {
	return c.JSON(storiesTableData)
}

func GetStory(c *fiber.Ctx) error {
	id := c.Params("storyId")
	index, err := strconv.Atoi(id)
	if err != nil {
		return c.JSON(model.Story{})
	}
	story := stories[index%len(stories)]
	return c.JSON(story)
}

func GetStoryTable(c *fiber.Ctx) error {
	return c.JSON(storiesTableData)
}

func GetStoryTextBox(c *fiber.Ctx) error {
	id := c.Params("storyId")
	index, err := strconv.Atoi(id)
	if err != nil {
		return c.JSON(model.Story{})
	}
	description := stories[index%len(stories)].Description
	return c.JSON(fiber.Map{"text": description})
}

var stories = []model.Story{
	{
		ID:          "0",
		Name:        "Story X",
		Description: "Story X description",
	},
	{
		ID:          "1",
		Name:        "Story Y",
		Description: "Story Y description",
	},
	{
		ID:          "2",
		Name:        "Story Z",
		Description: "Story Z description",
	},
}

var storiesTableData = model.TableData{
	Caption: "Stories",
	Headers: []model.TableHeader{
		{
			Dim:  4,
			Name: "Name",
		},
		{
			Dim:  8,
			Name: "Description",
		},
	},
	Rows: []model.TableRow{
		{
			ID:   stories[0].ID,
			Name: stories[0].Name,
			Values: []model.TableValue{
				{
					Dim:   4,
					Value: stories[0].Name,
				},
				{
					Dim:   8,
					Value: stories[0].Description,
				},
			},
		},
		{
			ID:   stories[1].ID,
			Name: stories[1].Name,
			Values: []model.TableValue{
				{
					Dim:   4,
					Value: stories[1].Name,
				},
				{
					Dim:   8,
					Value: stories[1].Description,
				},
			},
		},
		{
			ID:   stories[2].ID,
			Name: stories[2].Name,
			Values: []model.TableValue{
				{
					Dim:   4,
					Value: stories[2].Name,
				},
				{
					Dim:   8,
					Value: stories[2].Description,
				},
			},
		},
	},
}
