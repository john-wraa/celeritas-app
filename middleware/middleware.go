package middleware

import (
	"myapp/data"

	"github.com/john-wraa/celeritas"
)

type Middleware struct {
	App    *celeritas.Celeritas
	Models data.Models
}
