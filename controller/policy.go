package controller

import (
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/tizianocitro/csa-server/model"
)

func GetPolicies(c *fiber.Ctx) error {
	return c.JSON(policies)
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
