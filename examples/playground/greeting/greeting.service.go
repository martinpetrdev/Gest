package greeting

type GreetingService struct {
}

func newGreetingService() *GreetingService {
	return &GreetingService{}
}

func (s *GreetingService) BasicGreeting() string {
	return "Hello, world!"
}
