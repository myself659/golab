package grifts

import (
	"github.com/gobuffalo/buffalo"
	"github.com/myself659/golab/buffalo/business-card/actions"
)

func init() {
	buffalo.Grifts(actions.App())
}
