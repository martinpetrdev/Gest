package module

import (
	"fmt"
	"net/http"

	"github.com/martinpetrdev/Gest/internal/DI"
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
}

func (m *Module) Register(c *DI.Container, mux *http.ServeMux, seen map[*Module]bool) error {
	if seen[m] {
		return nil
	}
	seen[m] = true

	for _, imp := range m.Imports {
		if err := imp.Register(c, mux, seen); err != nil {
			return err
		}
	}

	for _, provider := range m.Providers {
		if err := c.Provide(provider); err != nil {
			return fmt.Errorf("module %s: %w", m.Name, err)
		}
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
	}

	return nil
}
