package biz

import "github.com/google/wire"

// ProviderSet is biz providers.
var ProviderSet = wire.NewSet(NewStaffUsecase, NewTeamUsecase, NewTeamStaffUsecase, NewTeamTeamUsecase)
