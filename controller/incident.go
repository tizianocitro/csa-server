package controller

import (
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/tizianocitro/csa-server/model"
)

func GetIncidents(c *fiber.Ctx) error {
	return c.JSON(incidentsTableData)
}

func GetIncident(c *fiber.Ctx) error {
	id := c.Params("incidentId")
	index, err := strconv.Atoi(id)
	if err != nil {
		return c.JSON(model.Incident{})
	}
	incident := incidents[index%len(incidents)]
	return c.JSON(incident)
}

func GetIncidentTable(c *fiber.Ctx) error {
	return c.JSON(incidentsTableData)
}

func GetIncidentTextBox(c *fiber.Ctx) error {
	id := c.Params("incidentId")
	index, err := strconv.Atoi(id)
	if err != nil {
		return c.JSON(model.Incident{})
	}
	description := incidents[index%len(incidents)].Description
	return c.JSON(fiber.Map{"text": description})
}

var incidents = []model.Incident{
	{
		ID:          "0",
		Name:        "Incident X",
		Description: "Incident X description",
	},
	{
		ID:          "1",
		Name:        "Incident Y",
		Description: "Incident Y description",
	},
	{
		ID:          "2",
		Name:        "Incident Z",
		Description: "Incident Z description",
	},
}

var incidentsTableData = model.TableData{
	Caption: "Incidents",
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
			ID:   incidents[0].ID,
			Name: incidents[0].Name,
			Values: []model.TableValue{
				{
					Dim:   4,
					Value: incidents[0].Name,
				},
				{
					Dim:   8,
					Value: incidents[0].Description,
				},
			},
		},
		{
			ID:   incidents[1].ID,
			Name: incidents[1].Name,
			Values: []model.TableValue{
				{
					Dim:   4,
					Value: incidents[1].Name,
				},
				{
					Dim:   8,
					Value: incidents[1].Description,
				},
			},
		},
		{
			ID:   incidents[2].ID,
			Name: incidents[2].Name,
			Values: []model.TableValue{
				{
					Dim:   4,
					Value: incidents[2].Name,
				},
				{
					Dim:   8,
					Value: incidents[2].Description,
				},
			},
		},
	},
}
