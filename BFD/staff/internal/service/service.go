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

	staff     *biz.StaffUsecase
	team      *biz.TeamUsecase
	teamStaff *biz.TeamStaffUsecase
	teamTeam  *biz.TeamTeamUsecase
	log       *log.Helper
}

func NewStaffService(staff *biz.StaffUsecase, team *biz.TeamUsecase, teamStaff *biz.TeamStaffUsecase, teamTeam *biz.TeamTeamUsecase, logger log.Logger) *StaffService {
	return &StaffService{
		staff:     staff,
		team:      team,
		teamStaff: teamStaff,
		teamTeam:  teamTeam,
		log:       log.NewHelper(log.With(logger, "module", "service/server-service")),
	}
}
