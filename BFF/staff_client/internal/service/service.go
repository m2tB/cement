package service

import (
	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/wire"
	v1 "staff_client/api/staff_client/v1"
)

// ProviderSet is service providers.
var ProviderSet = wire.NewSet(NewStaffClientService)

// StaffService is a staff service.
type StaffService struct {
	v1.UnimplementedStaffServer

	uc  *biz.StaffUsecase
	log *log.Helper
}

// NewStaffClientService new a staff service.
func NewStaffClientService(uc *biz.StaffUsecase, logger log.Logger) *StaffService {
	return &StaffService{
		uc:  uc,
		log: log.NewHelper(log.With(logger, "module", "service/staff_client")),
	}
}
