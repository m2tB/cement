package service

import (
	"context"
	"staff/constant"
	"staff/internal/biz"

	v1 "staff/api/staff/v1"
)

func (s *StaffService) CreateStaff(ctx context.Context, req *v1.CreateStaffRequest) (*v1.CreateStaffReply, error) {
	_, err := s.staff.CreateStaff(ctx, &biz.Staff{
		Mobile: req.Mobile,
		Name:   req.Name,
	})
	if err != nil {
		return nil, err
	}
	createStaffReply := v1.CreateStaffReply{
		Exec: true,
	}
	return &createStaffReply, nil
}

func (s *StaffService) UpdateStaff(ctx context.Context, req *v1.UpdateStaffRequest) (*v1.UpdateStaffReply, error) {
	_, err := s.staff.UpdateStaff(ctx, &biz.Staff{
		ID:     req.Id,
		Mobile: req.Mobile,
		Name:   req.Name,
	})
	if err != nil {
		return nil, err
	}
	updateStaffReply := v1.UpdateStaffReply{
		Exec: true,
	}
	return &updateStaffReply, nil
}

func (s *StaffService) DeleteStaff(ctx context.Context, req *v1.DeleteStaffRequest) (*v1.DeleteStaffReply, error) {
	_, err := s.staff.DeleteStaff(ctx, &biz.Staff{
		ID:        req.Id,
		IsDeleted: constant.True,
	})
	if err != nil {
		return nil, err
	}
	deleteStaffReply := v1.DeleteStaffReply{
		Exec: true,
	}
	return &deleteStaffReply, nil
}

func (s *StaffService) RecoveryStaff(ctx context.Context, req *v1.RecoveryStaffRequest) (*v1.RecoveryStaffReply, error) {
	_, err := s.staff.DeleteStaff(ctx, &biz.Staff{
		ID:        req.Id,
		IsDeleted: constant.False,
	})
	if err != nil {
		return nil, err
	}
	recoveryStaffReply := v1.RecoveryStaffReply{
		Exec: true,
	}
	return &recoveryStaffReply, nil
}

func (s *StaffService) ListStaff(ctx context.Context, req *v1.ListStaffRequest) (*v1.ListStaffReply, error) {
	// 类型转换 int -> bool
	isDeleted := constant.False
	if req.IsDeleted {
		isDeleted = constant.True
	}
	list, total, err := s.staff.ListStaff(ctx, &biz.Staff{
		Mobile:    req.Mobile,
		Name:      req.Name,
		IsDeleted: isDeleted,
	}, int(req.Pn), int(req.PSize))
	if err != nil {
		return nil, err
	}
	rsp := &v1.ListStaffReply{}
	rsp.Total = int32(total)
	for _, staff := range list {
		staffInfoRsp := v1.StaffReply{
			Id:        staff.ID,
			Name:      staff.Name,
			Mobile:    staff.Mobile,
			CreatedAt: staff.CreatedAt.Format(constant.CSTLayout),
			UpdatedAt: staff.UpdatedAt.Format(constant.CSTLayout),
		}
		rsp.Data = append(rsp.Data, &staffInfoRsp)
	}

	return rsp, nil
}
