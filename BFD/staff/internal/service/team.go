package service

import (
	"context"
	"staff/constant"
	"staff/internal/biz"

	v1 "staff/api/staff/v1"
)

func (s *StaffService) CreateTeam(ctx context.Context, req *v1.CreateTeamRequest) (*v1.CreateTeamReply, error) {
	_, err := s.team.CreateTeam(ctx, &biz.Team{
		CName:       req.CName,
		EName:       req.EName,
		PreTeamCode: req.PreTeamCode,
	})
	if err != nil {
		return nil, err
	}
	createTeamReply := v1.CreateTeamReply{
		Exec: true,
	}
	return &createTeamReply, nil
}

func (s *StaffService) UpdateTeam(ctx context.Context, req *v1.UpdateTeamRequest) (*v1.UpdateTeamReply, error) {
	_, err := s.team.UpdateTeam(ctx, &biz.Team{
		ID:          req.Id,
		CName:       req.CName,
		EName:       req.EName,
		PreTeamCode: req.PreTeamCode,
	})
	if err != nil {
		return nil, err
	}
	updateTeamReply := v1.UpdateTeamReply{
		Exec: true,
	}
	return &updateTeamReply, nil
}

func (s *StaffService) DeleteTeam(ctx context.Context, req *v1.DeleteTeamRequest) (*v1.DeleteTeamReply, error) {
	_, err := s.team.DeleteTeam(ctx, &biz.Team{
		ID: req.Id,
	})
	if err != nil {
		return nil, err
	}
	deleteTeamReply := v1.DeleteTeamReply{
		Exec: true,
	}
	return &deleteTeamReply, nil
}

func (s *StaffService) ReadTeam(ctx context.Context, req *v1.ReadTeamRequest) (*v1.ReadTeamReply, error) {
	team, err := s.team.ReadTeam(ctx, &biz.Team{
		ID: req.Id,
	})
	if err != nil {
		return nil, err
	}
	readTeamReply := v1.ReadTeamReply{
		Team: &v1.TeamReply{
			Id:          team.ID,
			CName:       team.CName,
			EName:       team.EName,
			PreTeamCode: team.PreTeamCode,
			CreatedAt:   team.CreatedAt.Format(constant.CSTLayout),
			UpdatedAt:   team.UpdatedAt.Format(constant.CSTLayout),
		},
	}
	return &readTeamReply, nil
}

func (s *StaffService) ListTeam(ctx context.Context, req *v1.ListTeamRequest) (*v1.ListTeamReply, error) {
	list, total, err := s.team.ListTeam(ctx, &biz.Team{
		CName:       req.CName,
		EName:       req.EName,
		PreTeamCode: req.PreTeamCode,
	}, int(req.Pn), int(req.PSize))
	if err != nil {
		return nil, err
	}
	rsp := &v1.ListTeamReply{}
	rsp.Total = int32(total)
	for _, team := range list {
		teamInfoRsp := v1.TeamReply{
			Id:          team.ID,
			CName:       team.CName,
			EName:       team.EName,
			PreTeamCode: team.PreTeamCode,
			CreatedAt:   team.CreatedAt.Format(constant.CSTLayout),
			UpdatedAt:   team.UpdatedAt.Format(constant.CSTLayout),
		}
		rsp.Data = append(rsp.Data, &teamInfoRsp)
	}

	return rsp, nil
}
