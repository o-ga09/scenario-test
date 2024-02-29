package main

import (
	"github.com/google/uuid"
	"github.com/zoncoen/scenarigo/plugin"
)

func init() {
	plugin.RegisterSetup(setRunID)
}

func setRunID(ctx *plugin.Context) (*plugin.Context, func(*plugin.Context)) {
	return ctx.WithVars(map[string]string{
		"runId": uuid.NewString(),
	}), nil
}
