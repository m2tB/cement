package service

import (
	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/wire"
	v1 "staff_client/api/staff_client/v1"
	"staff_client/internal/biz"
)

// ProviderSet is service providers.
var ProviderSet = wire.NewSet(NewStaffClientService)

// StaffClientService is a staff service.
type StaffClientService struct {
	v1.UnimplementedStaffClientServer

	uc  *biz.StaffClientUsecase
	log *log.Helper
}

// NewStaffClientService new a staff service.
func NewStaffClientService(uc *biz.StaffClientUsecase, logger log.Logger) *StaffClientService {
	return &StaffClientService{
		uc:  uc,
		log: log.NewHelper(log.With(logger, "module", "service/staff_client")),
	}
}
