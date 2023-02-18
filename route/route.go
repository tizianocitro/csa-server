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
	incidentsWithId := incidents.Group("/:incidentId")
	incidentsWithId.Get("/", func(c *fiber.Ctx) error {
		log.Printf("/incidents/:incidentId called")
		return controller.GetIncident(c)
	})
	incidentsWithId.Get("/graph", func(c *fiber.Ctx) error {
		log.Printf("/incidents/:incidentId/graph called")
		return controller.GetIncidentGraph(c)
	})
	incidentsWithId.Get("/table", func(c *fiber.Ctx) error {
		log.Printf("/incidents/:incidentId/table called")
		return controller.GetIncidentTable(c)
	})
	incidentsWithId.Get("/text_box", func(c *fiber.Ctx) error {
		log.Printf("/incidents/:incidentId/text_box called")
		return controller.GetIncidentTextBox(c)
	})
}

func useStories(app *fiber.App) {
	stories := app.Group("/stories")
	stories.Get("/", func(c *fiber.Ctx) error {
		log.Printf("/stories called")
		return controller.GetStories(c)
	})
	storiesWithId := stories.Group("/:storyId")
	storiesWithId.Get("/", func(c *fiber.Ctx) error {
		log.Printf("/stories/:storyId called")
		return controller.GetStory(c)
	})
	storiesWithId.Get("/graph", func(c *fiber.Ctx) error {
		log.Printf("/stories/:storyId/graph called")
		return controller.GetStoryGraph(c)
	})
	storiesWithId.Get("/table", func(c *fiber.Ctx) error {
		log.Printf("/stories/:storyId/table called")
		return controller.GetStoryTable(c)
	})
	storiesWithId.Get("/text_box", func(c *fiber.Ctx) error {
		log.Printf("/stories/:storyId/text_box called")
		return controller.GetStoryTextBox(c)
	})
}

func usePolicies(app *fiber.App) {
	policies := app.Group("/policies")
	policies.Get("/", func(c *fiber.Ctx) error {
		log.Printf("/policies called")
		return controller.GetPolicies(c)
	})
	policesWithId := policies.Group("/:policyId")
	policesWithId.Get("/", func(c *fiber.Ctx) error {
		log.Printf("/policies/:policyId called")
		return controller.GetPolicy(c)
	})
	policesWithId.Get("/graph", func(c *fiber.Ctx) error {
		log.Printf("/policies/:policyId/graph called")
		return controller.GetPolicyGraph(c)
	})
	policesWithId.Get("/table", func(c *fiber.Ctx) error {
		log.Printf("/policies/:policyId/table called")
		return controller.GetPolicyTable(c)
	})
	policesWithId.Get("/text_box", func(c *fiber.Ctx) error {
		log.Printf("/policies/:policyId/text_box called")
		return controller.GetPolicyTextBox(c)
	})
}
