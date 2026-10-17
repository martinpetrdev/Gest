package greeting

import gest "github.com/martinpetrdev/Gest"

var GreetingModule = gest.Module{
	Name:        "greeting",
	Providers:   []any{newGreetingService},
	Controllers: []any{newGreetingController},
}
