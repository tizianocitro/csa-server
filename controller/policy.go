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
	policy := policies[index%len(policies)]
	return c.JSON(policy)
}

func GetPolicyTable(c *fiber.Ctx) error {
	return c.JSON(policiesTableData)
}

func GetPolicyTextBox(c *fiber.Ctx) error {
	id := c.Params("policyId")
	index, err := strconv.Atoi(id)
	if err != nil {
		return c.JSON(model.Policy{})
	}
	description := policies[index%len(policies)].Description
	return c.JSON(fiber.Map{"text": description})
}

var policies = []model.Policy{
	{
		ID:          "0",
		Name:        "Policy X",
		Description: "Policy X description",
	},
	{
		ID:          "1",
		Name:        "Policy Y",
		Description: "Policy Y description",
	},
	{
		ID:          "2",
		Name:        "Policy Z",
		Description: "Policy Z description",
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
