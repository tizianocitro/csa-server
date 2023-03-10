package controller

import (
	"encoding/json"
	"fmt"

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

func GetStoryGraph(c *fiber.Ctx) error {
	return GetGraph(c)
}

func GetStoryTable(c *fiber.Ctx) error {
	story := getStoryByID(c)
	storyTableData := model.TableData{
		Caption: storiesTableData.Caption,
		Headers: storiesTableData.Headers,
		Rows:    []model.TableRow{},
	}
	for i := 0; i < 3; i++ {
		name, err := util.BuildStringFromTemplate(
			"storiesTableValuesName",
			model.TableValueNameTemplate,
			model.TableValuePlaceholder{Name: story.Name, Index: i + 1})
		if err != nil {
			return c.JSON(model.Policy{})
		}
		storyTableData.Rows = append(storyTableData.Rows, model.TableRow{
			ID:   storiesTableDataIds[i],
			Name: name,
			Values: []model.TableValue{
				{
					Dim:   4,
					Value: name,
				},
				{
					Dim:   8,
					Value: fmt.Sprintf("This the description for the %s", name),
				},
			},
		})
	}
	return c.JSON(storyTableData)
}

func GetStoryTextBox(c *fiber.Ctx) error {
	return c.JSON(fiber.Map{"text": getStoryByID(c).Description})
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
			Name:        "Story X",
			Description: "Story X description",
		},
		{
			ID:          "6d2a4472-69ef-45e8-abce-17a12d433dce",
			Name:        "Story Y",
			Description: "Story Y description",
		},
		{
			ID:          "c8995f27-8419-43a9-bdba-c69dfdcd2853",
			Name:        "Story Z",
			Description: "Story Z description",
		},
	},
	"2": {
		{
			ID:          "d8cd62a4-eb17-4922-b773-738112a714b8",
			Name:        "Story 1",
			Description: "Story 1 description",
		},
		{
			ID:          "d26ffddd-673e-4526-8c9f-48a5814cd314",
			Name:        "Story 2",
			Description: "Story 2 description",
		},
	},
	"3": {
		{
			ID:          "2d5b607a-2204-49b1-b640-239cf586c962",
			Name:        "Story I",
			Description: "Story I description",
		},
		{
			ID:          "f731ebca-1d63-4470-ab40-71a8b89d59c1",
			Name:        "Story II",
			Description: "Story II description",
		},
		{
			ID:          "0ee9646b-8ce2-4831-b355-e1fb640cf29f",
			Name:        "Story III",
			Description: "Story III description",
		},
		{
			ID:          "b739b350-e975-4716-be6c-388a7f1c0ff9",
			Name:        "Story IV",
			Description: "Story IV description",
		},
	},
}

var storiesTableData = model.TableData{
	Caption: "Story Elements",
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
	Rows: []model.TableRow{},
}

var storiesTableDataIds = []string{
	"495c21dc-9f4c-4bd9-922f-5fc5475cafb7",
	"5214998c-a9be-4f9c-b89c-118999cc1dfe",
	"b569dc93-2e4d-4a9a-9ce6-12a646deb808",
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
