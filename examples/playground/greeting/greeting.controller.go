package greeting

import gest "github.com/martinpetrdev/Gest"

type GreetingController struct {
	svc *GreetingService
}

func newGreetingController(svc *GreetingService) *GreetingController {
	return &GreetingController{
		svc: svc,
	}
}

func (c *GreetingController) Routes(r *gest.Router) {
	r.Prefix("/greeting")

	r.Get("/", c.BasicGreeting)
}

func (c *GreetingController) BasicGreeting() string {
	return c.svc.BasicGreeting()
}
