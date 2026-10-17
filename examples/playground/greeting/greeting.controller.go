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
	r.Get("/{name}", c.Hello)
	r.Get("/{name}/{age}", c.WithAge)
}

func (c *GreetingController) BasicGreeting(ctx *gest.RouteContext) string {
	return c.svc.BasicGreeting()
}

func (c *GreetingController) Hello(ctx *gest.RouteContext) string {
	name, err := ctx.GetParam("name")
	if err != nil {
		panic(gest.NewHTTPBadRequestException("Name not provided"))
	}

	return c.svc.Hello(name)
}

func (c *GreetingController) WithAge(ctx *gest.RouteContext) string {
	name, err := ctx.GetParam("name")
	if err != nil {
		panic(gest.NewHTTPBadRequestException("Name not provided"))
	}

	age, err := ctx.GetParamInt("age")
	if err != nil {
		panic(gest.NewHTTPBadRequestException("Age not provided or is not an integer"))
	}

	return c.svc.WithAge(name, age)
}
