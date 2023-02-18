package controller

import (
	"github.com/gofiber/fiber/v2"
	"github.com/tizianocitro/csa-server/model"
)

func GetGraph(c *fiber.Ctx) error {
	return c.JSON(graphData)
}

var graphData = model.GraphData{
	Nodes: []model.GraphNode{
		{
			ID: "n1",
			Position: model.GraphNodePosition{
				X: 0,
				Y: 0,
			},
			Data: model.GraphNodeData{
				Label: "Node 1",
			},
		},
		{
			ID: "n2",
			Position: model.GraphNodePosition{
				X: 100,
				Y: 100,
			},
			Data: model.GraphNodeData{
				Label: "Node 2",
			},
		},
	},
	Edges: []model.GraphEdge{
		{
			ID:     "n1-n2",
			Source: "n1",
			Target: "n2",
		},
	},
}
