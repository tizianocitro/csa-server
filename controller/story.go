package controller

import (
	"encoding/json"
	"strconv"
	"sync/atomic"

	"github.com/gofiber/fiber/v2"
	"github.com/tizianocitro/csa-server/model"
)

func GetStories(c *fiber.Ctx) error {
	return c.JSON(storiesPaginatedTableData)
}

func GetStory(c *fiber.Ctx) error {
	id := c.Params("storyId")
	index, err := strconv.Atoi(id)
	if err != nil {
		return c.JSON(model.Story{})
	}
	if story := getStoryByID(id); story != (model.Story{}) {
		return c.JSON(story)
	}
	story := stories[index%len(stories)]
	return c.JSON(story)
}

func GetStoryGraph(c *fiber.Ctx) error {
	return GetGraph(c)
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
	if story := getStoryByID(id); story != (model.Story{}) {
		return c.JSON(fiber.Map{"text": story.Description})
	}
	description := stories[index%len(stories)].Description
	return c.JSON(fiber.Map{"text": description})
}

func SaveStory(c *fiber.Ctx) error {
	var story model.Story
	err := json.Unmarshal(c.Body(), &story)
	if err != nil {
		return c.JSON(fiber.Map{
			"error": "Not a valid story provided",
		})
	}
	id := getNextID()
	story.ID = strconv.Itoa(id)
	stories = append(stories, story)
	storiesPaginatedTableData.Rows = append(storiesPaginatedTableData.Rows, model.PaginatedTableRow{
		ID:          stories[len(stories)-1].ID,
		Name:        stories[len(stories)-1].Name,
		Description: stories[len(stories)-1].Description,
	})
	return c.JSON(fiber.Map{
		"id":   story.ID,
		"name": story.Name,
	})
}

func getStoryByID(id string) model.Story {
	for _, story := range stories {
		if story.ID == id {
			return story
		}
	}
	return model.Story{}
}

func getNextID() int {
	id := atomic.LoadInt32(&nextID)
	atomic.AddInt32(&nextID, 1)
	return int(id)
}

var nextID int32 = 12

var stories = []model.Story{
	{
		ID:          "6",
		Name:        "Story X",
		Description: "Story X description",
	},
	{
		ID:          "7",
		Name:        "Story Y",
		Description: "Story Y description",
	},
	{
		ID:          "8",
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

var storiesPaginatedTableData = model.PaginatedTableData{
	Columns: []model.PaginatedTableColumn{
		{
			Title: "Name",
		},
		{
			Title: "Description",
		},
	},
	Rows: []model.PaginatedTableRow{
		{
			ID:          stories[0].ID,
			Name:        stories[0].Name,
			Description: stories[0].Description,
		},
		{
			ID:          stories[1].ID,
			Name:        stories[1].Name,
			Description: stories[1].Description,
		},
		{
			ID:          stories[2].ID,
			Name:        stories[2].Name,
			Description: stories[2].Description,
		},
	},
}
