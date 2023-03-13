package controller

import (
	"github.com/gofiber/fiber/v2"
	"github.com/tizianocitro/csa-server/model"
)

func GetPolicies(c *fiber.Ctx) error {
	organizationId := c.Params("organizationId")
	tableData := model.PaginatedTableData{
		Columns: policiesPaginatedTableData.Columns,
		Rows:    []model.PaginatedTableRow{},
	}
	for _, policy := range policiesMap[organizationId] {
		tableData.Rows = append(tableData.Rows, model.PaginatedTableRow{
			ID:          policy.ID,
			Name:        policy.Name,
			Description: policy.Description,
		})
	}
	return c.JSON(tableData)
}

func GetPolicy(c *fiber.Ctx) error {
	return c.JSON(getPolicyByID(c))
}

func GetPolicyDos(c *fiber.Ctx) error {
	policyId := c.Params("policyId")
	return c.JSON(fiber.Map{"items": policiesDosMap[policyId]})
}

func GetPolicyDonts(c *fiber.Ctx) error {
	policyId := c.Params("policyId")
	return c.JSON(fiber.Map{"items": policiesDontsMap[policyId]})
}

func getPolicyByID(c *fiber.Ctx) model.Policy {
	organizationId := c.Params("organizationId")
	policyId := c.Params("policyId")
	for _, policy := range policiesMap[organizationId] {
		if policy.ID == policyId {
			return policy
		}
	}
	return model.Policy{}
}

var policiesMap = map[string][]model.Policy{
	"1": {
		{
			ID:          "e39edc4b-5f19-4210-a576-a8e679717a86",
			Name:        "Policy X",
			Description: "Policy X description",
		},
		{
			ID:          "00b1ce5e-95f6-4466-952e-754efbbc4224",
			Name:        "Policy Y",
			Description: "Policy Y description",
		},
		{
			ID:          "454a6288-7c0c-4fc7-9ccd-97e7f293eb19",
			Name:        "Policy Z",
			Description: "Policy Z description",
		},
	},
	"2": {
		{
			ID:          "7d33bc3d-dbd9-40a7-8274-670400aa9ba7",
			Name:        "Policy 1",
			Description: "Policy 1 description",
		},
	},
	"3": {
		{
			ID:          "8b74f14d-a7cb-407b-9098-ebced0ee018b",
			Name:        "Policy I",
			Description: "Policy I description",
		},
		{
			ID:          "b6178e38-1cfa-4a26-bb14-ec2d57cf55e6",
			Name:        "Policy II",
			Description: "Policy II description",
		},
	},
}

var policiesDosMap = map[string][]model.ListData{
	"e39edc4b-5f19-4210-a576-a8e679717a86": {
		{
			ID:   "b86f19da-0bda-44c2-8be3-5396a78f273f",
			Text: "Do 1",
		},
		{
			ID:   "cad1fd77-2bdd-44c0-8791-14819a90d9ba",
			Text: "Do 2",
		},
		{
			ID:   "00a24dd6-38c4-4ae6-8e42-6f4e2773470b",
			Text: "Do 3",
		},
	},
	"00b1ce5e-95f6-4466-952e-754efbbc4224": {
		{
			ID:   "772dd04c-fdae-4809-a2a2-dfc0edbc7670",
			Text: "Do 4",
		},
	},
	"454a6288-7c0c-4fc7-9ccd-97e7f293eb19": {
		{
			ID:   "fd63b9a3-53bb-4c1f-8ea6-a5505b4d26bc",
			Text: "Do 5",
		},
		{
			ID:   "dbeebb13-7a2e-4566-b7e8-4364ec51d118",
			Text: "Do 6",
		},
	},
	"7d33bc3d-dbd9-40a7-8274-670400aa9ba7": {
		{
			ID:   "fb6d2f7e-84a8-4fda-ad40-dd4b9c65b9b8",
			Text: "Do 7",
		},
		{
			ID:   "4823ebed-9e27-42f1-93da-9dad21af2780",
			Text: "Do 8",
		},
		{
			ID:   "4a941a25-3460-4fc8-98cd-1e0b1402ca4b",
			Text: "Do 9",
		},
	},
	"8b74f14d-a7cb-407b-9098-ebced0ee018b": {
		{
			ID:   "0be62676-2f05-46a2-8db7-78ead21ceefe",
			Text: "Do 10",
		},
	},
	"b6178e38-1cfa-4a26-bb14-ec2d57cf55e6": {
		{
			ID:   "3003fba0-cace-4203-b23b-989004149d49",
			Text: "Do 11",
		},
		{
			ID:   "bf5f204c-a3cf-4af6-8111-5f2b16c34bc2",
			Text: "Do 12",
		},
	},
}

var policiesDontsMap = map[string][]model.ListData{
	"e39edc4b-5f19-4210-a576-a8e679717a86": {
		{
			ID:   "723c8d4d-a222-44fa-9de4-cb5f19371135",
			Text: "Don't 1",
		},
		{
			ID:   "216e8f2e-339c-44e0-b16c-983e6308fe6e",
			Text: "Don't 2",
		},
		{
			ID:   "845b699f-d2f5-4806-a2c2-b09e7b32036c",
			Text: "Don't 3",
		},
	},
	"00b1ce5e-95f6-4466-952e-754efbbc4224": {
		{
			ID:   "88e182c9-aff0-4356-a2e9-6b3ad98e271b",
			Text: "Don't 4",
		},
	},
	"454a6288-7c0c-4fc7-9ccd-97e7f293eb19": {
		{
			ID:   "0a4caeaa-fcb0-4ef3-a0d2-730462329606",
			Text: "Don't 5",
		},
		{
			ID:   "c456806f-e67c-4c07-b0ee-f47be7f461e5",
			Text: "Don't 6",
		},
	},
	"7d33bc3d-dbd9-40a7-8274-670400aa9ba7": {
		{
			ID:   "9611ce8e-db13-4205-bc60-56f998f41426",
			Text: "Don't 7",
		},
		{
			ID:   "821dbef5-4714-4b4b-abad-4e98b8a4d6da",
			Text: "Don't 8",
		},
		{
			ID:   "a3e72a7e-27cb-46b3-ae5c-3417a3a2afaa",
			Text: "Don't 9",
		},
	},
	"8b74f14d-a7cb-407b-9098-ebced0ee018b": {
		{
			ID:   "7f8eb74a-02f4-4a63-bd03-715a4dbc0caf",
			Text: "Don't 10",
		},
	},
	"b6178e38-1cfa-4a26-bb14-ec2d57cf55e6": {
		{
			ID:   "749bd99f-bcef-4239-a7d7-8756dfbc61e2",
			Text: "Don't 11",
		},
		{
			ID:   "791e5824-c537-46a5-b9ca-ef6753fbc1c6",
			Text: "Don't 12",
		},
	},
}

var policiesPaginatedTableData = model.PaginatedTableData{
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
