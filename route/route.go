package route

import (
	"log"

	"github.com/gofiber/fiber/v2"

	"github.com/tizianocitro/csa-server/controller"
)

func UseRoutes(app *fiber.App) {
	useOrganizations(app)
	useIncidents(app)
	useStories(app)
	usePolicies(app)
}

func useOrganizations(app *fiber.App) {
	organizations := app.Group("/organizations")
	organizations.Get("/", func(c *fiber.Ctx) error {
		log.Printf("/organizations called")
		return controller.GetOrganizations(c)
	})
	organizations.Get("/no_page", func(c *fiber.Ctx) error {
		log.Printf("/organizations/no_page called")
		return controller.GetOrganizationsNoPage(c)
	})
	organizations.Get("/:organizationId", func(c *fiber.Ctx) error {
		log.Printf("/organizations/:organizationId called")
		return controller.GetOrganization(c)
	})
	useOrganizationsIncidents(organizations)
	useOrganizationsStories(organizations)
	useOrganizationsPolicies(organizations)
}

func useOrganizationsIncidents(organizations fiber.Router) {
	incidents := organizations.Group("/:organizationId/incidents")
	incidents.Get("/", func(c *fiber.Ctx) error {
		log.Printf("/:organizationId/incidents called")
		return controller.GetIncidents(c)
	})
	incidents.Get("/:incidentId", func(c *fiber.Ctx) error {
		log.Printf("/:organizationId/incidents/:incidentId called")
		return controller.GetIncident(c)
	})
}

func useOrganizationsStories(organizations fiber.Router) {
	stories := organizations.Group("/:organizationId/stories")
	stories.Get("/", func(c *fiber.Ctx) error {
		log.Printf("/:organizationId/stories called")
		return controller.GetStories(c)
	})
	stories.Get("/:storyId", func(c *fiber.Ctx) error {
		log.Printf("/:organizationId/stories/:storyId called")
		return controller.GetStory(c)
	})
}

func useOrganizationsPolicies(organizations fiber.Router) {
	policies := organizations.Group("/:organizationId/policies")
	policies.Get("/", func(c *fiber.Ctx) error {
		log.Printf("/:organizationId/policies called")
		return controller.GetPolicies(c)
	})
	policies.Get("/:policyId", func(c *fiber.Ctx) error {
		log.Printf("/:organizationId/policies/:policyId called")
		return controller.GetPolicy(c)
	})
}

func useIncidents(app *fiber.App) {
	incidents := app.Group("/incidents")
	incidents.Get("/", func(c *fiber.Ctx) error {
		log.Printf("/incidents called")
		return controller.GetIncidents(c)
	})
	incidents.Get("/:incidentId", func(c *fiber.Ctx) error {
		log.Printf("/incidents/:incidentId called")
		return controller.GetIncident(c)
	})
}

func useStories(app *fiber.App) {
	stories := app.Group("/stories")
	stories.Get("/", func(c *fiber.Ctx) error {
		log.Printf("/stories called")
		return controller.GetStories(c)
	})
	stories.Get("/:storyId", func(c *fiber.Ctx) error {
		log.Printf("/stories/:storyId called")
		return controller.GetStory(c)
	})
}

func usePolicies(app *fiber.App) {
	policies := app.Group("/policies")
	policies.Get("/", func(c *fiber.Ctx) error {
		log.Printf("/policies called")
		return controller.GetPolicies(c)
	})
	policies.Get("/:policyId", func(c *fiber.Ctx) error {
		log.Printf("/policies/:policyId called")
		return controller.GetPolicy(c)
	})
}
