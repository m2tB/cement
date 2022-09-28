package data

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"staff/constant"
	"staff/internal/biz"
	"time"
)

// TeamStaff 定义数据表结构体
type TeamStaff struct {
	ID        int64      `gorm:"primaryKey"`
	TID       int64      `gorm:"column:tid;not null"`
	SID       int64      `gorm:"column:sid;not null"`
	CreatedAt *time.Time `gorm:"column:created_at"`
}

type ListTeamStaff struct {
	SID       int64
	SName     string
	TID       int64
	TName     string
	CreatedAt *time.Time
}

type teamStaffRepo struct {
	data *Data
	log  *log.Helper
}

// NewTeamStaffRepo .
func NewTeamStaffRepo(data *Data, logger log.Logger) biz.TeamStaffRepo {
	return &teamStaffRepo{
		data: data,
		log:  log.NewHelper(logger),
	}
}

func (r *teamStaffRepo) Save(_ context.Context, t *biz.TeamStaff) (bool, error) {
	if len(t.SIDs) <= 0 {
		return false, status.Errorf(codes.InvalidArgument, "sids is empty.")
	}
	var teamStaffs []TeamStaff
	for _, sid := range t.SIDs {
		teamStaffs = append(teamStaffs, TeamStaff{
			TID: t.TID,
			SID: sid,
		})
	}
	exec := r.data.db.CreateInBatches(&teamStaffs, constant.ExecBatchSize)
	if exec.Error != nil {
		return false, status.Errorf(codes.Internal, exec.Error.Error())
	}
	return true, nil
}

func (r *teamStaffRepo) Delete(_ context.Context, t *biz.TeamStaff) (bool, error) {
	var total int64
	// 验证是否已经创建
	r.data.db.Model(&TeamStaff{}).Where("tid = ? and sid = ?", t.TID, t.SID).Count(&total)
	if total == 0 {
		return false, status.Errorf(codes.NotFound, "access is not exist.")
	}
	exec := r.data.db.Where("tid = ? and sid = ?", t.TID, t.SID).Unscoped().Delete(&TeamStaff{})
	if exec.Error != nil {
		return false, status.Errorf(codes.Internal, exec.Error.Error())
	}
	return true, nil
}

func (r *teamStaffRepo) List(_ context.Context, t *biz.TeamStaff, pn int, pSize int) ([]*biz.DataListTeamStaff, int64, error) {
	var total int64
	tx := r.data.db.Table("team_staff")
	var listTeamStaff []ListTeamStaff
	if t.TID != 0 {
		tx.Select("s.id as sid, s.name as s_name, team_staff.created_at as created_at")
		tx.Joins("left join staff s on team_staff.sid = s.id")
		tx.Where("team_staff.tid = ?", t.TID)
	}
	if t.SID != 0 {
		tx.Select("t.id as tid, t.c_name as t_name, team_staff.created_at as created_at")
		tx.Joins("left join team t on team_staff.tid = t.id")
		tx.Where("team_staff.sid = ?", t.SID)
	}
	exec := tx.Scan(&listTeamStaff).Count(&total)
	if exec.Error != nil || total == 0 {
		return nil, 0, status.Errorf(codes.NotFound, exec.Error.Error())
	}
	exec = tx.Order("created_at desc").Scopes(paginate(pn, pSize)).Find(&listTeamStaff)
	if exec.Error != nil {
		return nil, 0, status.Errorf(codes.Internal, exec.Error.Error())
	}
	result := make([]*biz.DataListTeamStaff, 0)
	for _, teamStaff := range listTeamStaff {
		result = append(result, &biz.DataListTeamStaff{
			SID:       teamStaff.SID,
			SName:     teamStaff.SName,
			TID:       teamStaff.TID,
			TName:     teamStaff.TName,
			CreatedAt: teamStaff.CreatedAt,
		})
	}
	return result, total, nil
}
