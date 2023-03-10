package controller

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/tizianocitro/csa-server/model"
	"github.com/tizianocitro/csa-server/util"
)

func GetIncidents(c *fiber.Ctx) error {
	organizationId := c.Params("organizationId")
	tableData := model.PaginatedTableData{
		Columns: incidentsPaginatedTableData.Columns,
		Rows:    []model.PaginatedTableRow{},
	}
	for _, incident := range incidentsMap[organizationId] {
		tableData.Rows = append(tableData.Rows, model.PaginatedTableRow{
			ID:          incident.ID,
			Name:        incident.Name,
			Description: incident.Description,
		})
	}
	return c.JSON(tableData)
}

func GetIncident(c *fiber.Ctx) error {
	return c.JSON(getIncidentByID(c))
}

func GetIncidentGraph(c *fiber.Ctx) error {
	return GetGraph(c)
}

func GetIncidentTable(c *fiber.Ctx) error {
	incident := getIncidentByID(c)
	incidentTableData := model.TableData{
		Caption: incidentsTableData.Caption,
		Headers: incidentsTableData.Headers,
		Rows:    []model.TableRow{},
	}
	for i := 0; i < 3; i++ {
		name, err := util.BuildStringFromTemplate(
			"incidentsTableValuesName",
			model.TableValueNameTemplate,
			model.TableValuePlaceholder{Name: incident.Name, Index: i + 1})
		if err != nil {
			return c.JSON(model.Incident{})
		}
		incidentTableData.Rows = append(incidentTableData.Rows, model.TableRow{
			ID:   incidentsTableDataIds[i],
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
	return c.JSON(incidentTableData)
}

func GetIncidentTextBox(c *fiber.Ctx) error {
	return c.JSON(fiber.Map{"text": getIncidentByID(c).Description})
}

func getIncidentByID(c *fiber.Ctx) model.Incident {
	organizationId := c.Params("organizationId")
	incidentId := c.Params("incidentId")
	for _, incident := range incidentsMap[organizationId] {
		if incident.ID == incidentId {
			return incident
		}
	}
	return model.Incident{}
}

var incidentsMap = map[string][]model.Incident{
	"1": {
		{
			ID:          "2ce53d5c-4bd4-4f02-89cc-d5b8f551770c",
			Name:        "Incident X",
			Description: "Incident X description",
		},
		{
			ID:          "be4fcd12-cb96-40f4-96a8-4eed6b61e414",
			Name:        "Incident Y",
			Description: "Incident Y description",
		},
		{
			ID:          "82077bc3-54ff-45ee-b43b-645485319ed1",
			Name:        "Incident Z",
			Description: "Incident Z description",
		},
	},
	"2": {
		{
			ID:          "39b1441b-2b36-4cdc-a1f7-aa38c25bc13b",
			Name:        "Incident 1",
			Description: "Incident 1 description",
		},
		{
			ID:          "ac1b2a79-69ce-4e83-a6cf-203fe7af194d",
			Name:        "Incident 2",
			Description: "Incident 2 description",
		},
	},
	"3": {
		{
			ID:          "7defe83a-acbf-4784-9bc2-eb3447a0a545",
			Name:        "Incident I",
			Description: "Incident I description",
		},
		{
			ID:          "e83f0631-1fec-49c0-805d-223622b1c138",
			Name:        "Incident II",
			Description: "Incident II description",
		},
		{
			ID:          "8a4324b8-c822-4f79-92d9-517039c1887d",
			Name:        "Incident III",
			Description: "Incident III description",
		},
		{
			ID:          "d26b179f-cafd-4908-ba4e-7bf27fa3c925",
			Name:        "Incident IV",
			Description: "Incident IV description",
		},
	},
}

var incidentsTableData = model.TableData{
	Caption: "Incident Elements",
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

var incidentsTableDataIds = []string{
	"8ea68a25-8527-4c0d-8a25-4827f0869bc3",
	"43b94ea3-9630-4fe9-a3a6-8ad7c6fd4c59",
	"3dcdab7c-8dbc-4572-88f3-3d76b00cc211",
}

var incidentsPaginatedTableData = model.PaginatedTableData{
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
