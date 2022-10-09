package data

import (
	"context"

	"staff_client/internal/biz"

	"github.com/go-kratos/kratos/v2/log"

	staff_v1 "staff_client/api/staff/v1"
)

type staffClientRepo struct {
	data *Data
	log  *log.Helper
}

// NewStaffClientRepo .
func NewStaffClientRepo(data *Data, logger log.Logger) biz.StaffClientRepo {
	return &staffClientRepo{
		data: data,
		log:  log.NewHelper(log.With(logger, "module", "repo/staff_client")),
	}
}

func (r *staffClientRepo) Save(ctx context.Context, staff *biz.StaffClient) (bool, error) {
	_, err := r.data.sc.CreateStaff(ctx, &staff_v1.CreateStaffRequest{
		Mobile: staff.Mobile,
		Name:   staff.Name,
	})
	if err != nil {
		return false, err
	}
	return true, nil
}

func (r *staffClientRepo) Read(ctx context.Context, mobile string) (*biz.StaffClient, error) {
	staffs, err := r.data.sc.ListStaff(ctx, &staff_v1.ListStaffRequest{
		Mobile: mobile,
	})
	if err != nil || staffs.Total <= 0 {
		return nil, err
	}
	staff := staffs.GetData()[0]
	return &biz.StaffClient{
		ID:        staff.Id,
		Mobile:    staff.Mobile,
		Name:      staff.Name,
		CreatedAt: staff.CreatedAt,
		UpdatedAt: staff.UpdatedAt,
	}, nil
}
