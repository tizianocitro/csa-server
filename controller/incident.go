package controller

import (
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/tizianocitro/csa-server/model"
)

func GetIncidents(c *fiber.Ctx) error {
	return c.JSON(incidents)
}

func GetIncident(c *fiber.Ctx) error {
	id := c.Params("incidentId")
	index, err := strconv.Atoi(id)
	if err != nil {
		return c.JSON(model.Incident{})
	}
	incident := incidents[index]
	return c.JSON(incident)
}

var incidents = []model.Incident{
	{
		ID:          "0",
		Name:        "iX",
		Description: "X incident",
	},
	{
		ID:          "1",
		Name:        "iY",
		Description: "Y incident",
	},
	{
		ID:          "2",
		Name:        "iZ",
		Description: "Z incident",
	},
}
