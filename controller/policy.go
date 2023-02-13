package controller

import (
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/tizianocitro/csa-server/model"
)

func GetPolicies(c *fiber.Ctx) error {
	return c.JSON(policiesTableData)
}

func GetPolicy(c *fiber.Ctx) error {
	id := c.Params("policyId")
	index, err := strconv.Atoi(id)
	if err != nil {
		return c.JSON(model.Policy{})
	}
	policy := policies[index]
	return c.JSON(policy)
}

var policies = []model.Policy{
	{
		ID:          "0",
		Name:        "pX",
		Description: "X policy",
	},
	{
		ID:          "1",
		Name:        "pY",
		Description: "Y policy",
	},
	{
		ID:          "2",
		Name:        "pZ",
		Description: "Z policy",
	},
}

var policiesTableData = model.TableData{
	Caption: "Policies",
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
			ID:   policies[0].ID,
			Name: policies[0].Name,
			Values: []model.TableValue{
				{
					Dim:   4,
					Value: policies[0].Name,
				},
				{
					Dim:   8,
					Value: policies[0].Description,
				},
			},
		},
		{
			ID:   policies[1].ID,
			Name: policies[1].Name,
			Values: []model.TableValue{
				{
					Dim:   4,
					Value: policies[1].Name,
				},
				{
					Dim:   8,
					Value: policies[1].Description,
				},
			},
		},
		{
			ID:   policies[2].ID,
			Name: policies[2].Name,
			Values: []model.TableValue{
				{
					Dim:   4,
					Value: policies[2].Name,
				},
				{
					Dim:   8,
					Value: policies[2].Description,
				},
			},
		},
	},
}