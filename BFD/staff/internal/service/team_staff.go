package service

import (
	"context"
	v1 "staff/api/staff/v1"
	"staff/constant"
	"staff/internal/biz"
)

func (s *StaffService) InviteStaff(ctx context.Context, req *v1.InviteStaffRequest) (*v1.InviteStaffReply, error) {
	_, err := s.teamStaff.InviteStaff(ctx, &biz.TeamStaff{
		TID:  req.Tid,
		SIDs: req.Sids,
	})
	if err != nil {
		return nil, err
	}
	inviteStaffReply := v1.InviteStaffReply{
		Exec: true,
	}
	return &inviteStaffReply, nil
}

func (s *StaffService) ExpelStaff(ctx context.Context, req *v1.ExpelStaffRequest) (*v1.ExpelStaffReply, error) {
	_, err := s.teamStaff.ExpelStaff(ctx, &biz.TeamStaff{
		TID: req.Tid,
		SID: req.Sid,
	})
	if err != nil {
		return nil, err
	}
	expelStaffReply := v1.ExpelStaffReply{
		Exec: true,
	}
	return &expelStaffReply, nil
}

func (s *StaffService) ListTeamStaff(ctx context.Context, req *v1.ListTeamStaffRequest) (*v1.ListTeamStaffReply, error) {
	_, err := s.team.ReadTeam(ctx, &biz.Team{
		ID: req.Tid,
	})
	if err != nil {
		return nil, err
	}
	list, total, err := s.teamStaff.ListTeamStaff(ctx, &biz.TeamStaff{
		TID: req.Tid,
	}, int(req.Pn), int(req.PSize))
	if err != nil {
		return nil, err
	}
	rsp := &v1.ListTeamStaffReply{}
	rsp.Total = int32(total)
	for _, teamStaff := range list {
		teamStaffInfoRsp := v1.TeamStaffReply{
			Sid:       teamStaff.SID,
			SName:     teamStaff.SName,
			CreatedAt: teamStaff.CreatedAt.Format(constant.CSTLayout),
		}
		rsp.Data = append(rsp.Data, &teamStaffInfoRsp)
	}

	return rsp, nil
}

func (s *StaffService) ListStaffTeam(ctx context.Context, req *v1.ListStaffTeamRequest) (*v1.ListStaffTeamReply, error) {
	_, err := s.staff.ReadStaff(ctx, &biz.Staff{
		ID: req.Sid,
	})
	if err != nil {
		return nil, err
	}
	list, total, err := s.teamStaff.ListStaffTeam(ctx, &biz.TeamStaff{
		SID: req.Sid,
	}, int(req.Pn), int(req.PSize))
	if err != nil {
		return nil, err
	}
	rsp := &v1.ListStaffTeamReply{}
	rsp.Total = int32(total)
	for _, staffTeam := range list {
		staffTeamInfoRsp := v1.StaffTeamReply{
			Tid:       staffTeam.TID,
			TName:     staffTeam.TName,
			CreatedAt: staffTeam.CreatedAt.Format(constant.CSTLayout),
		}
		rsp.Data = append(rsp.Data, &staffTeamInfoRsp)
	}

	return rsp, nil
}
