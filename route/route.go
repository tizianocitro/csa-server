package route

import (
	"github.com/gofiber/fiber/v2"

	"github.com/tizianocitro/csa-server/controller"
)

func UseRoutes(app *fiber.App) {
	useOrganizations(app)
}

func useOrganizations(app *fiber.App) {
	organizations := app.Group("/organizations")
	organizations.Get("/", func(c *fiber.Ctx) error {
		return controller.GetOrganizations(c)
	})
	organizations.Get("/no_page", func(c *fiber.Ctx) error {
		return controller.GetOrganizationsNoPage(c)
	})
	organizations.Get("/:organizationId", func(c *fiber.Ctx) error {
		return controller.GetOrganization(c)
	})
	useIncidents(organizations)
	useStories(organizations)
	usePolicies(organizations)
}

func useIncidents(organizations fiber.Router) {
	incidents := organizations.Group("/:organizationId/incidents")
	incidents.Get("/", func(c *fiber.Ctx) error {
		return controller.GetIncidents(c)
	})
	incidents.Get("/:incidentId", func(c *fiber.Ctx) error {
		return controller.GetIncident(c)
	})
}

func useStories(organizations fiber.Router) {
	stories := organizations.Group("/:organizationId/stories")
	stories.Get("/", func(c *fiber.Ctx) error {
		return controller.GetStories(c)
	})
	stories.Get("/:storyId", func(c *fiber.Ctx) error {
		return controller.GetStory(c)
	})
}

func usePolicies(organizations fiber.Router) {
	policies := organizations.Group("/:organizationId/policies")
	policies.Get("/", func(c *fiber.Ctx) error {
		return controller.GetPolicies(c)
	})
	policies.Get("/:policyId", func(c *fiber.Ctx) error {
		return controller.GetPolicy(c)
	})
}
