package module

import (
	"fmt"
	"net/http"

	"github.com/martinpetrdev/Gest/internal/DI"
	"github.com/martinpetrdev/Gest/internal/format"
	"github.com/martinpetrdev/Gest/internal/logger"
	"github.com/martinpetrdev/Gest/internal/router"
)

type Controller interface {
	Routes(r *router.Router)
}

type Module struct {
	Name        string
	Imports     []*Module
	Providers   []any
	Controllers []any

	logger *logger.Logger
}

func (m *Module) Register(c *DI.Container, mux *http.ServeMux, seen map[*Module]bool) error {
	m.logger = logger.New(format.PascalCase([]string{m.Name, "module"}))

	var totalDeps int = 0

	if seen[m] {
		return nil
	}
	seen[m] = true

	for _, imp := range m.Imports {
		if err := imp.Register(c, mux, seen); err != nil {
			return err
		}

		totalDeps++
	}

	for _, provider := range m.Providers {
		if err := c.Provide(provider); err != nil {
			return fmt.Errorf("module %s: %w", m.Name, err)
		}

		totalDeps++
	}

	for _, constructor := range m.Controllers {
		value, err := c.Invoke(constructor)
		if err != nil {
			return fmt.Errorf("module %s: %w", m.Name, err)
		}

		controller, ok := value.Interface().(Controller)
		if !ok {
			return fmt.Errorf("module %s: %s has no Routes method", m.Name, value.Type())
		}

		controller.Routes(router.New(mux))

		totalDeps++
	}

	m.logger.Log(fmt.Sprintf("%d dependencies registered", totalDeps))

	return nil
}
