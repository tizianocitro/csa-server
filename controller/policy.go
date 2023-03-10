package controller

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/tizianocitro/csa-server/model"
	"github.com/tizianocitro/csa-server/util"
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

func GetPolicyGraph(c *fiber.Ctx) error {
	return GetGraph(c)
}

func GetPolicyTable(c *fiber.Ctx) error {
	policy := getPolicyByID(c)
	policyTableData := model.TableData{
		Caption: policiesTableData.Caption,
		Headers: policiesTableData.Headers,
		Rows:    []model.TableRow{},
	}
	for i := 0; i < 3; i++ {
		name, err := util.BuildStringFromTemplate(
			"policiesTableValuesName",
			model.TableValueNameTemplate,
			model.TableValuePlaceholder{Name: policy.Name, Index: i + 1})
		if err != nil {
			return c.JSON(model.Policy{})
		}
		policyTableData.Rows = append(policyTableData.Rows, model.TableRow{
			ID:   policiesTableDataIds[i],
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
	return c.JSON(policyTableData)
}

func GetPolicyTextBox(c *fiber.Ctx) error {
	return c.JSON(fiber.Map{"text": getPolicyByID(c).Description})
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
		{
			ID:          "95102ca1-193d-4ea9-b4ee-9c358a0d9af7",
			Name:        "Policy 2",
			Description: "Policy 2 description",
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
		{
			ID:          "2b39b9ad-0094-44a3-969b-1dc591a12a70",
			Name:        "Policy III",
			Description: "Policy III description",
		},
		{
			ID:          "3025ed0f-0fd1-4c7d-9509-11e6b4cad962",
			Name:        "Policy IV",
			Description: "Policy IV description",
		},
	},
}

var policiesTableData = model.TableData{
	Caption: "Policy Elements",
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

var policiesTableDataIds = []string{
	"c86852fb-7c64-496c-8d5f-555d061c28eb",
	"cd71258f-8085-4f0c-ad61-5e097317239c",
	"8f6322d3-489f-4a0b-9790-085cdef8335a",
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
