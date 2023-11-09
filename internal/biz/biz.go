package biz

import (
	"github.com/google/wire"
	"slacker/internal/biz/record"
	"slacker/internal/biz/user"
)

// ProviderSet is biz providers.
var ProviderSet = wire.NewSet(
	user.NewUseCase,
	record.NewUseCase,
)
