package controller

import (
	"encoding/json"

	"github.com/gofiber/fiber/v2"
	"github.com/tizianocitro/csa-server/model"
	"github.com/tizianocitro/csa-server/util"
)

func GetStories(c *fiber.Ctx) error {
	organizationId := c.Params("organizationId")
	tableData := model.PaginatedTableData{
		Columns: storiesPaginatedTableData.Columns,
		Rows:    []model.PaginatedTableRow{},
	}
	for _, story := range storiesMap[organizationId] {
		tableData.Rows = append(tableData.Rows, model.PaginatedTableRow{
			ID:          story.ID,
			Name:        story.Name,
			Description: story.Description,
		})
	}
	return c.JSON(tableData)
}

func GetStory(c *fiber.Ctx) error {
	return c.JSON(getStoryByID(c))
}

func GetStoryTimeline(c *fiber.Ctx) error {
	storyId := c.Params("storyId")
	return c.JSON(model.TimelineData{
		Items: storiesTimelineDataMap[storyId],
	})
}

func SaveStory(c *fiber.Ctx) error {
	organizationId := c.Params("organizationId")
	var story model.Story
	err := json.Unmarshal(c.Body(), &story)
	if err != nil {
		return c.JSON(fiber.Map{
			"error": "Not a valid story provided",
		})
	}
	story.ID = util.GenerateUUID()
	storiesMap[organizationId] = append(storiesMap[organizationId], story)
	storiesTimelineDataMap[story.ID] = storiesTimelineItems
	return c.JSON(fiber.Map{
		"id":   story.ID,
		"name": story.Name,
	})
}

func getStoryByID(c *fiber.Ctx) model.Story {
	organizationId := c.Params("organizationId")
	storyId := c.Params("storyId")
	for _, story := range storiesMap[organizationId] {
		if story.ID == storyId {
			return story
		}
	}
	return model.Story{}
}

var storiesMap = map[string][]model.Story{
	"1": {
		{
			ID:          "d5454bed-4e1a-4f5b-b71c-2993e157ba6f",
			Name:        "The Sudden Incident",
			Description: "An incident was detected on February 25th.",
		},
		{
			ID:          "6d2a4472-69ef-45e8-abce-17a12d433dce",
			Name:        "The Investigation",
			Description: "How to conduct an investigation.",
		},
	},
	"2": {
		{
			ID:          "d8cd62a4-eb17-4922-b773-738112a714b8",
			Name:        "The Report",
			Description: "Planning how to share results.",
		},
	},
	"3": {
		{
			ID:          "d26ffddd-673e-4526-8c9f-48a5814cd314",
			Name:        "The Unexplained Incident",
			Description: "An incident occurred on February 25th.",
		},
		{
			ID:          "2d5b607a-2204-49b1-b640-239cf586c962",
			Name:        "The Aftermath",
			Description: "How to manage consequences of the incident.",
		},
	},
}

var storiesTimelineDataMap = map[string][]model.TimelineItem{
	"d5454bed-4e1a-4f5b-b71c-2993e157ba6f": storiesTimelineItems,
	"6d2a4472-69ef-45e8-abce-17a12d433dce": storiesTimelineItems,
	"d8cd62a4-eb17-4922-b773-738112a714b8": storiesTimelineItems,
	"d26ffddd-673e-4526-8c9f-48a5814cd314": storiesTimelineItems,
	"2d5b607a-2204-49b1-b640-239cf586c962": storiesTimelineItems,
}

var storiesTimelineItems = []model.TimelineItem{
	{
		ID:    "2ea54075-da38-44cf-afb9-cbe31a3d9513",
		Label: "2023-02-25 09:12:11",
		Text:  "An event was detected",
	},
	{
		ID:    "e3484f88-482b-457a-9f74-cff900c81ce4",
		Label: "2023-03-09 11:33:48",
		Text:  "The event was investigated",
	},
	{
		ID:    "3626ed66-4ce3-4a56-a4d9-18ad1fb80f9a",
		Label: "2023-03-12 16:51:31",
		Text:  "The event was reported",
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
	Rows: []model.PaginatedTableRow{},
}
