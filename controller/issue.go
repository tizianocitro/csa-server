package controller

import (
	"encoding/json"
	"errors"
	"strconv"
	"sync/atomic"

	"github.com/gofiber/fiber/v2"
	"github.com/tizianocitro/csa-server/model"
)

func GetIssues(c *fiber.Ctx) error {
	rows := []model.PaginatedTableRow{}
	for _, issue := range issues {
		rows = append(rows, model.PaginatedTableRow{
			ID:          issue.ID,
			Name:        issue.Name,
			Description: issue.Description,
		})
	}
	return c.JSON(model.PaginatedTableData{
		Columns: columns,
		Rows:    rows,
	})
}

func GetIssue(c *fiber.Ctx) error {
	id := c.Params("issueId")
	if issue, err := getIssueByID(id); err == nil {
		return c.JSON(issue)
	}
	return c.JSON(model.Issue{})
}

func SaveIssue(c *fiber.Ctx) error {
	var issue model.Issue
	err := json.Unmarshal(c.Body(), &issue)
	if err != nil {
		return c.JSON(fiber.Map{
			"error": "Not a valid issue provided",
		})
	}
	id := getNextIssueID()
	issue.ID = strconv.Itoa(id)
	issues = append(issues, issue)
	return c.JSON(fiber.Map{
		"id":   issue.ID,
		"name": issue.Name,
	})
}

func getIssueByID(id string) (model.Issue, error) {
	for _, issue := range issues {
		if issue.ID == id {
			return issue, nil
		}
	}
	return model.Issue{}, errors.New("not found")
}

func getNextIssueID() int {
	id := atomic.LoadInt32(&nextIssueID)
	atomic.AddInt32(&nextIssueID, 1)
	return int(id)
}

var nextIssueID int32 = 0

var issues = []model.Issue{}

var columns = []model.PaginatedTableColumn{
	{
		Title: "Name",
	},
	{
		Title: "Description",
	},
}
