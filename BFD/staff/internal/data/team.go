package data

import (
	"context"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"staff/constant"
	"time"

	"github.com/go-kratos/kratos/v2/log"
	"staff/internal/biz"
)

// Team 定义数据表结构体
type Team struct {
	ID          int64      `gorm:"primaryKey"`
	CName       string     `gorm:"column:c_name;type:varchar(32) comment '中文名称';not null"`
	EName       string     `gorm:"column:e_name;type:varchar(64) comment '英文名称'"`
	PreTeamCode string     `gorm:"column:pre_team_code;type:varchar(4) comment '前缀代码';not null"`
	CreatedAt   *time.Time `gorm:"column:created_at"`
	UpdatedAt   *time.Time `gorm:"column:updated_at"`
	IsDeleted   int        `gorm:"column:is_deleted"`
}

type teamRepo struct {
	data *Data
	log  *log.Helper
}

// NewTeamRepo .
func NewTeamRepo(data *Data, logger log.Logger) biz.TeamRepo {
	return &teamRepo{
		data: data,
		log:  log.NewHelper(logger),
	}
}

func (r *teamRepo) Save(_ context.Context, t *biz.Team) (bool, error) {
	var team Team
	team.CName = t.CName
	team.EName = t.EName
	team.PreTeamCode = t.PreTeamCode
	exec := r.data.db.Create(&team)
	if exec.Error != nil {
		return false, status.Errorf(codes.Internal, exec.Error.Error())
	}
	return true, nil
}

func (r *teamRepo) Update(_ context.Context, t *biz.Team) (bool, error) {
	var total int64
	// 验证是否存在
	r.data.db.Model(&Team{}).Where("id = ? and is_deleted = ?", t.ID, constant.False).Count(&total)
	if total == 0 {
		return false, status.Errorf(codes.NotFound, "access is not exist.")
	}
	var team Team
	team.ID = t.ID
	team.CName = t.CName
	team.EName = t.EName
	team.PreTeamCode = t.PreTeamCode
	exec := r.data.db.Updates(&team)
	if exec.Error != nil {
		return false, status.Errorf(codes.Internal, exec.Error.Error())
	}
	return true, nil
}

func (r *teamRepo) Delete(_ context.Context, t *biz.Team) (bool, error) {
	var total int64
	// 验证是否已经创建
	r.data.db.Model(&Team{}).Where("id = ? and is_deleted = ?", t.ID, constant.False).Count(&total)
	if total == 0 {
		return false, status.Errorf(codes.NotFound, "access is not exist.")
	}
	values := map[string]interface{}{
		"is_deleted": constant.True,
	}
	exec := r.data.db.Model(&Team{}).Where("id = ?", t.ID).Updates(&values)
	if exec.Error != nil {
		return false, status.Errorf(codes.Internal, exec.Error.Error())
	}
	return true, nil
}

func (r *teamRepo) Read(_ context.Context, t *biz.Team) (*biz.DataTeam, error) {
	var team Team
	exec := r.data.db.Model(&Team{}).Where("id = ? and is_deleted = ?", t.ID, constant.False).First(&team)
	if exec.Error != nil {
		return nil, status.Errorf(codes.NotFound, exec.Error.Error())
	}
	return &biz.DataTeam{
		ID:          team.ID,
		CName:       team.CName,
		EName:       team.EName,
		PreTeamCode: team.PreTeamCode,
		CreatedAt:   team.CreatedAt,
		UpdatedAt:   team.UpdatedAt,
		IsDeleted:   team.IsDeleted,
	}, nil
}

func (r *teamRepo) List(_ context.Context, t *biz.Team, pn int, pSize int) ([]*biz.DataTeam, int64, error) {
	var total int64
	tx := r.data.db.Model(&Team{}).Where("is_deleted = ?", constant.False)
	if t.CName != "" {
		tx = tx.Where("c_name = ?", t.CName)
	}
	if t.EName != "" {
		tx = tx.Where("e_name = ?", t.EName)
	}
	if t.PreTeamCode != "" {
		tx = tx.Where("pre_team_code = ?", t.PreTeamCode)
	}
	exec := tx.Count(&total)
	if exec.Error != nil {
		return nil, 0, status.Errorf(codes.Internal, exec.Error.Error())
	}
	if total == 0 {
		return nil, 0, nil
	}
	var teams []Team
	exec = tx.Order("created_at desc").Scopes(paginate(pn, pSize)).Find(&teams)
	if exec.Error != nil {
		return nil, 0, status.Errorf(codes.Internal, exec.Error.Error())
	}
	result := make([]*biz.DataTeam, 0)
	for _, team := range teams {
		result = append(result, &biz.DataTeam{
			ID:          team.ID,
			CName:       team.CName,
			EName:       team.EName,
			PreTeamCode: team.PreTeamCode,
			CreatedAt:   team.CreatedAt,
			UpdatedAt:   team.UpdatedAt,
			IsDeleted:   team.IsDeleted,
		})
	}
	return result, total, nil
}
