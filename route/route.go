package route

import (
	"log"

	"github.com/gofiber/fiber/v2"

	"github.com/tizianocitro/csa-server/controller"
)

func UseRoutes(app *fiber.App) {
	useOrganizations(app)
	useEcosystem(app)
	useIncidents(app)
	useStories(app)
	usePolicies(app)
}

func useOrganizations(app *fiber.App) {
	organizations := app.Group("/organizations")
	organizations.Get("/", func(c *fiber.Ctx) error {
		log.Printf("GET /organizations called")
		return controller.GetOrganizations(c)
	})
	organizations.Get("/no_page", func(c *fiber.Ctx) error {
		log.Printf("GET /organizations/no_page called")
		return controller.GetOrganizationsNoPage(c)
	})
	organizations.Get("/:organizationId", func(c *fiber.Ctx) error {
		log.Printf("GET /organizations/:organizationId called")
		return controller.GetOrganization(c)
	})
	useOrganizationsIncidents(organizations)
	useOrganizationsStories(organizations)
	useOrganizationsPolicies(organizations)
}

func useOrganizationsIncidents(organizations fiber.Router) {
	incidents := organizations.Group("/:organizationId/incidents")
	incidents.Get("/", func(c *fiber.Ctx) error {
		log.Printf("GET /:organizationId/incidents called")
		return controller.GetIncidents(c)
	})
	incidentsWithId := incidents.Group("/:incidentId")
	incidentsWithId.Get("/", func(c *fiber.Ctx) error {
		log.Printf("GET /:organizationId/incidents/:incidentId called")
		return controller.GetIncident(c)
	})
	incidentsWithId.Get("/graph", func(c *fiber.Ctx) error {
		log.Printf("GET /:organizationId/incidents/:incidentId/graph called")
		return controller.GetIncidentGraph(c)
	})
	incidentsWithId.Get("/table", func(c *fiber.Ctx) error {
		log.Printf("GET /:organizationId/incidents/:incidentId/table called")
		return controller.GetIncidentTable(c)
	})
	incidentsWithId.Get("/text_box", func(c *fiber.Ctx) error {
		log.Printf("GET /:organizationId/incidents/:incidentId/text_box called")
		return controller.GetIncidentTextBox(c)
	})
}

func useOrganizationsStories(organizations fiber.Router) {
	stories := organizations.Group("/:organizationId/stories")
	stories.Get("/", func(c *fiber.Ctx) error {
		log.Printf("GET /:organizationId/stories called")
		return controller.GetStories(c)
	})
	stories.Post("/", func(c *fiber.Ctx) error {
		log.Printf("POST /:organizationId/stories called")
		return controller.SaveStory(c)
	})
	storiesWithId := stories.Group("/:storyId")
	storiesWithId.Get("/", func(c *fiber.Ctx) error {
		log.Printf("GET /:organizationId/stories/:storyId called")
		return controller.GetStory(c)
	})
	storiesWithId.Get("/graph", func(c *fiber.Ctx) error {
		log.Printf("GET /:organizationId/stories/:storyId/graph called")
		return controller.GetStoryGraph(c)
	})
	storiesWithId.Get("/table", func(c *fiber.Ctx) error {
		log.Printf("GET /:organizationId/stories/:storyId/table called")
		return controller.GetStoryTable(c)
	})
	storiesWithId.Get("/text_box", func(c *fiber.Ctx) error {
		log.Printf("GET /:organizationId/stories/:storyId/text_box called")
		return controller.GetStoryTextBox(c)
	})
}

func useOrganizationsPolicies(organizations fiber.Router) {
	policies := organizations.Group("/:organizationId/policies")
	policies.Get("/", func(c *fiber.Ctx) error {
		log.Printf("GET /:organizationId/policies called")
		return controller.GetPolicies(c)
	})
	policiesWithId := policies.Group("/:policyId")
	policiesWithId.Get("/", func(c *fiber.Ctx) error {
		log.Printf("GET /:organizationId/policies/:policyId called")
		return controller.GetPolicy(c)
	})
	policiesWithId.Get("/graph", func(c *fiber.Ctx) error {
		log.Printf("GET /:organizationId/policies/:policyId/graph called")
		return controller.GetPolicyGraph(c)
	})
	policiesWithId.Get("/table", func(c *fiber.Ctx) error {
		log.Printf("GET /:organizationId/policies/:policyId/table called")
		return controller.GetPolicyTable(c)
	})
	policiesWithId.Get("/text_box", func(c *fiber.Ctx) error {
		log.Printf("GET /:organizationId/policies/:policyId/text_box called")
		return controller.GetPolicyTextBox(c)
	})
}

func useEcosystem(app *fiber.App) {
	ecosystem := app.Group("/issues")
	ecosystem.Get("/", func(c *fiber.Ctx) error {
		log.Printf("GET /issues called")
		return controller.GetIssues(c)
	})
	ecosystem.Get("/:issueId", func(c *fiber.Ctx) error {
		log.Printf("GET /issues/:issueId called")
		return controller.GetIssue(c)
	})
	ecosystem.Post("/", func(c *fiber.Ctx) error {
		log.Printf("POST /issues called")
		return controller.SaveIssue(c)
	})
}

func useIncidents(app *fiber.App) {
	incidents := app.Group("/incidents")
	incidents.Get("/", func(c *fiber.Ctx) error {
		log.Printf("GET /incidents called")
		return controller.GetIncidents(c)
	})
	incidentsWithId := incidents.Group("/:incidentId")
	incidentsWithId.Get("/", func(c *fiber.Ctx) error {
		log.Printf("GET /incidents/:incidentId called")
		return controller.GetIncident(c)
	})
	incidentsWithId.Get("/graph", func(c *fiber.Ctx) error {
		log.Printf("GET /incidents/:incidentId/graph called")
		return controller.GetIncidentGraph(c)
	})
	incidentsWithId.Get("/table", func(c *fiber.Ctx) error {
		log.Printf("GET /incidents/:incidentId/table called")
		return controller.GetIncidentTable(c)
	})
	incidentsWithId.Get("/text_box", func(c *fiber.Ctx) error {
		log.Printf("GET /incidents/:incidentId/text_box called")
		return controller.GetIncidentTextBox(c)
	})
}

func useStories(app *fiber.App) {
	stories := app.Group("/stories")
	stories.Get("/", func(c *fiber.Ctx) error {
		log.Printf("GET /stories called")
		return controller.GetStories(c)
	})
	stories.Post("/", func(c *fiber.Ctx) error {
		log.Printf("POST /stories called")
		return controller.SaveStory(c)
	})
	storiesWithId := stories.Group("/:storyId")
	storiesWithId.Get("/", func(c *fiber.Ctx) error {
		log.Printf("GET /stories/:storyId called")
		return controller.GetStory(c)
	})
	storiesWithId.Get("/graph", func(c *fiber.Ctx) error {
		log.Printf("GET /stories/:storyId/graph called")
		return controller.GetStoryGraph(c)
	})
	storiesWithId.Get("/table", func(c *fiber.Ctx) error {
		log.Printf("GET /stories/:storyId/table called")
		return controller.GetStoryTable(c)
	})
	storiesWithId.Get("/text_box", func(c *fiber.Ctx) error {
		log.Printf("GET /stories/:storyId/text_box called")
		return controller.GetStoryTextBox(c)
	})
}

func usePolicies(app *fiber.App) {
	policies := app.Group("/policies")
	policies.Get("/", func(c *fiber.Ctx) error {
		log.Printf("GET /policies called")
		return controller.GetPolicies(c)
	})
	policesWithId := policies.Group("/:policyId")
	policesWithId.Get("/", func(c *fiber.Ctx) error {
		log.Printf("GET /policies/:policyId called")
		return controller.GetPolicy(c)
	})
	policesWithId.Get("/graph", func(c *fiber.Ctx) error {
		log.Printf("GET /policies/:policyId/graph called")
		return controller.GetPolicyGraph(c)
	})
	policesWithId.Get("/table", func(c *fiber.Ctx) error {
		log.Printf("GET /policies/:policyId/table called")
		return controller.GetPolicyTable(c)
	})
	policesWithId.Get("/text_box", func(c *fiber.Ctx) error {
		log.Printf("GET /policies/:policyId/text_box called")
		return controller.GetPolicyTextBox(c)
	})
}
