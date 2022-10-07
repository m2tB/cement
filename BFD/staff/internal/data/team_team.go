package data

import (
	"context"
	"fmt"
	"github.com/go-kratos/kratos/v2/log"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"gorm.io/gorm"
	"staff/internal/biz"
	"strings"
	"time"
)

// TeamTeam 定义数据表结构体
type TeamTeam struct {
	ID        int64      `gorm:"primaryKey"`
	TID       int64      `gorm:"column:tid;not null"`
	PID       int64      `gorm:"column:pid;not null"`
	Path      string     `gorm:"column:path;not null"`
	CreatedAt *time.Time `gorm:"column:created_at"`
}

type ListTeamTeam struct {
	TID       int64
	TName     string
	PID       int64
	PName     string
	CreatedAt *time.Time
}

type teamTeamRepo struct {
	data *Data
	log  *log.Helper
}

// NewTeamTeamRepo .
func NewTeamTeamRepo(data *Data, logger log.Logger) biz.TeamTeamRepo {
	return &teamTeamRepo{
		data: data,
		log:  log.NewHelper(logger),
	}
}

func (r *teamTeamRepo) Save(_ context.Context, t *biz.TeamTeam) (bool, error) {
	var total int64
	// 验证是否已经创建
	r.data.db.Model(&TeamTeam{}).Where("tid = ?", t.TID).Count(&total)
	if total > 0 {
		return false, status.Errorf(codes.AlreadyExists, "access is already exist.")
	}
	var pNode TeamTeam
	var path string
	// 验证父节点是否存在
	exec := r.data.db.Model(&TeamTeam{}).Where("tid = ?", t.PID).First(&pNode)
	if exec.Error != nil {
		if exec.Error == gorm.ErrRecordNotFound {
			path = "-"
		}
		return false, status.Errorf(codes.Internal, exec.Error.Error())
	} else {
		path = pNode.Path + "-"
	}
	var teamTeam TeamTeam
	teamTeam.PID = t.PID
	teamTeam.TID = t.TID
	teamTeam.Path = fmt.Sprintf("%s%v", path, pNode.TID)
	exec = r.data.db.Create(&teamTeam)
	if exec.Error != nil {
		return false, status.Errorf(codes.Internal, exec.Error.Error())
	}
	return true, nil
}

func (r *teamTeamRepo) Update(_ context.Context, t *biz.TeamTeam) (bool, error) {
	var fromTeam TeamTeam
	// 验证原父节点是否已经创建
	exec := r.data.db.Model(&TeamTeam{}).Where("tid = ?", t.OpID).First(&fromTeam)
	if exec.Error != nil {
		if exec.Error == gorm.ErrRecordNotFound {
			return false, status.Errorf(codes.NotFound, "access is not exist.")
		}
		return false, status.Errorf(codes.Internal, exec.Error.Error())
	}
	var toTeam TeamTeam
	// 验证新父节点是否已经创建
	exec = r.data.db.Model(&TeamTeam{}).Where("tid = ?", t.PID).First(&toTeam)
	if exec.Error != nil {
		if exec.Error == gorm.ErrRecordNotFound {
			return false, status.Errorf(codes.NotFound, "access is not exist.")
		}
		return false, status.Errorf(codes.Internal, exec.Error.Error())
	}
	tx := r.data.db.Begin()
	var teamTeam TeamTeam
	teamTeam.PID = t.PID
	exec = tx.Where("tid = ? and pid = ?", t.TID, t.OpID).Updates(&teamTeam)
	if exec.Error != nil {
		tx.Rollback()
		return false, status.Errorf(codes.Internal, exec.Error.Error())
	}
	needEdit := fmt.Sprintf("-%v-%v-", t.OpID, t.TID)
	expectEdit := fmt.Sprintf("-%v-%v-", t.PID, t.TID)
	exec = tx.Exec("Update team_team Set path = REPLACE(path, ?, ?) Where path Like ?", teamTeam.Path, strings.Replace(teamTeam.Path, needEdit, expectEdit, -1), teamTeam.Path+"%")
	if exec.Error != nil {
		tx.Rollback()
		return false, status.Errorf(codes.Internal, exec.Error.Error())
	}
	tx.Commit()
	return true, nil
}

func (r *teamTeamRepo) Delete(_ context.Context, t *biz.TeamTeam) (bool, error) {
	var team TeamTeam
	// 验证是否已经创建
	exec := r.data.db.Model(&TeamTeam{}).Where("tid = ?", t.TID).First(&team)
	if exec.Error != nil {
		if exec.Error == gorm.ErrRecordNotFound {
			return false, status.Errorf(codes.NotFound, "access is not exist.")
		}
		return false, status.Errorf(codes.Internal, exec.Error.Error())
	}
	var total int64
	// 验证是否存在下级节点 - 存在不允许进行删除,需要先进行迁移
	r.data.db.Model(&TeamTeam{}).Where("pid = ?", t.TID).Count(&total)
	if total > 0 {
		return false, status.Errorf(codes.AlreadyExists, "sub nodes are exist.")
	}
	exec = r.data.db.Where("tid = ?", t.TID).Unscoped().Delete(&TeamTeam{})
	if exec.Error != nil {
		return false, status.Errorf(codes.Internal, exec.Error.Error())
	}
	return true, nil
}

func (r *teamTeamRepo) List(_ context.Context, t *biz.TeamTeam, pn int, pSize int) ([]*biz.DataListTeamTeam, int64, error) {
	var total int64
	var listTeamTeam []ListTeamTeam
	// 只获取下一级节点, -- 待考虑
	tx := r.data.db.Table("team_team")
	tx.Select("t.id as tid, t.c_name as t_name, team_team.created_at as created_at")
	tx.Joins("left join team t on team_team.tid = t.id")
	tx.Where("team_team.pid = ?", t.PID)
	exec := tx.Scan(&listTeamTeam).Count(&total)
	if exec.Error != nil {
		return nil, 0, status.Errorf(codes.Internal, exec.Error.Error())
	}
	if total == 0 {
		return nil, 0, nil
	}
	exec = tx.Order("created_at asc").Scopes(paginate(pn, pSize)).Find(&listTeamTeam)
	if exec.Error != nil {
		return nil, 0, status.Errorf(codes.Internal, exec.Error.Error())
	}
	result := make([]*biz.DataListTeamTeam, 0)
	for _, teamTeam := range listTeamTeam {
		result = append(result, &biz.DataListTeamTeam{
			TID:       teamTeam.TID,
			TName:     teamTeam.TName,
			CreatedAt: teamTeam.CreatedAt,
		})
	}
	return result, total, nil
}
