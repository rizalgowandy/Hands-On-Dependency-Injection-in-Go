//+build wireinject

package main

import (
	"github.com/google/wire"
)

// The build tag makes sure the stub is not built in the final build.

func initializeDeps() *GetPersonHandler {
	wire.Build(wireSet)
	return nil
}
