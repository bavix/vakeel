package app

import (
	"context"
)

type RegisterUseCase interface {
	Register() error
}

func AgentRegister(
	_ context.Context,
	register RegisterUseCase,
) error {
	return register.Register()
}
