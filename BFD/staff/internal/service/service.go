package service

import (
	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/wire"
	v1 "staff/api/staff/v1"
	"staff/internal/biz"
)

// ProviderSet is service providers.
var ProviderSet = wire.NewSet(NewStaffService)

type StaffService struct {
	v1.UnimplementedStaffServer

	staff *biz.StaffUsecase
	team  *biz.TeamUsecase
	log   *log.Helper
}

func NewStaffService(staff *biz.StaffUsecase, team *biz.TeamUsecase, logger log.Logger) *StaffService {
	return &StaffService{
		staff: staff,
		team:  team,
		log:   log.NewHelper(log.With(logger, "module", "service/server-service")),
	}
}
