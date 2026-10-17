package greeting

import "fmt"

type GreetingService struct {
}

func newGreetingService() *GreetingService {
	return &GreetingService{}
}

func (s *GreetingService) BasicGreeting() string {
	return "Hello, world!"
}

func (s *GreetingService) Hello(name string) string {
	return fmt.Sprintf("Hello, %s!", name)
}

func (s *GreetingService) WithAge(name string, age int) string {
	return fmt.Sprintf("Hello, %s! You are %d years old.", name, age)
}
