package service

import (
	"context"
	v1 "staff/api/staff/v1"
	"staff/constant"
	"staff/internal/biz"
)

func (s *StaffService) CreateSubTeam(ctx context.Context, req *v1.CreateSubTeamRequest) (*v1.CreateSubTeamReply, error) {
	_, err := s.teamTeam.CreateSubTeam(ctx, &biz.TeamTeam{
		TID: req.Tid,
		PID: req.Pid,
	})
	if err != nil {
		return nil, err
	}
	createSubTeamReply := v1.CreateSubTeamReply{
		Exec: true,
	}
	return &createSubTeamReply, nil
}

func (s *StaffService) UpdateSubTeam(ctx context.Context, req *v1.UpdateSubTeamRequest) (*v1.UpdateSubTeamReply, error) {
	_, err := s.teamTeam.UpdateSubTeam(ctx, &biz.TeamTeam{
		TID:  req.Tid,
		PID:  req.Pid,
		OpID: req.OPid,
	})
	if err != nil {
		return nil, err
	}
	updateSubTeamReply := v1.UpdateSubTeamReply{
		Exec: true,
	}
	return &updateSubTeamReply, nil
}

func (s *StaffService) DeleteSubTeam(ctx context.Context, req *v1.DeleteSubTeamRequest) (*v1.DeleteSubTeamReply, error) {
	_, err := s.teamTeam.DeleteSubTeam(ctx, &biz.TeamTeam{
		TID: req.Tid,
	})
	if err != nil {
		return nil, err
	}
	deleteSubTeamReply := v1.DeleteSubTeamReply{
		Exec: true,
	}
	return &deleteSubTeamReply, nil
}

func (s *StaffService) ListTeamSubTeam(ctx context.Context, req *v1.ListTeamSubTeamRequest) (*v1.ListTeamSubTeamReply, error) {
	list, total, err := s.teamTeam.ListTeamSubTeam(ctx, &biz.TeamTeam{
		PID: req.Pid,
	}, int(req.Pn), int(req.PSize))
	if err != nil {
		return nil, err
	}
	rsp := &v1.ListTeamSubTeamReply{}
	rsp.Total = int32(total)
	for _, teamTeam := range list {
		teamTeamInfoRsp := v1.TeamSubTeamReply{
			Tid:       teamTeam.TID,
			TName:     teamTeam.TName,
			CreatedAt: teamTeam.CreatedAt.Format(constant.CSTLayout),
		}
		rsp.Data = append(rsp.Data, &teamTeamInfoRsp)
	}
	return rsp, nil
}
