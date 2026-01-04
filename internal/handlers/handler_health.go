package handlers

import (
	"github.com/BlackRoad-OS/authelia/v4/internal/middlewares"
)

// HealthGET can be used by health checks.
func HealthGET(ctx *middlewares.AutheliaCtx) {
	ctx.ReplyOK()
}
