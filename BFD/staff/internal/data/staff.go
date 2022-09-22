package data

import (
	"context"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"gorm.io/gorm"
	"staff/constant"
	"time"

	"github.com/go-kratos/kratos/v2/log"
	"staff/internal/biz"
)

// Staff 定义数据表结构体
type Staff struct {
	ID        int64      `gorm:"primaryKey"`
	Mobile    string     `gorm:"index:idx_mobile;unique;type:varchar(11) comment '手机号码,用户唯一标识';not null"`
	Name      string     `gorm:"type:varchar(4) comment '用户昵称';not null"`
	CreatedAt *time.Time `gorm:"column:created_at"`
	UpdatedAt *time.Time `gorm:"column:updated_at"`
	IsDeleted int        `gorm:"column:is_deleted"`
}

type staffRepo struct {
	data *Data
	log  *log.Helper
}

// NewStaffRepo .
func NewStaffRepo(data *Data, logger log.Logger) biz.StaffRepo {
	return &staffRepo{
		data: data,
		log:  log.NewHelper(logger),
	}
}

func (r *staffRepo) Save(_ context.Context, s *biz.Staff) (bool, error) {
	var total int64
	// 验证是否已经创建
	r.data.db.Model(&Staff{}).Where("mobile = ?", s.Mobile).Count(&total)
	if total > 0 {
		return false, status.Errorf(codes.AlreadyExists, "access is already exist.")
	}
	var staff Staff
	staff.Mobile = s.Mobile
	staff.Name = s.Name
	res := r.data.db.Create(&staff)
	if res.Error != nil {
		return false, status.Errorf(codes.Internal, res.Error.Error())
	}
	return true, nil
}

func (r *staffRepo) Update(_ context.Context, s *biz.Staff) (bool, error) {
	var total int64
	// 验证是否存在
	r.data.db.Model(&Staff{}).Where("id = ? and is_deleted = ?", s.ID, constant.False).Count(&total)
	if total == 0 {
		return false, status.Errorf(codes.NotFound, "access is not exist.")
	}
	var staff Staff
	staff.ID = s.ID
	staff.Mobile = s.Mobile
	staff.Name = s.Name
	res := r.data.db.Updates(&staff)
	if res.Error != nil {
		return false, status.Errorf(codes.Internal, res.Error.Error())
	}
	return true, nil
}

func (r *staffRepo) Delete(_ context.Context, s *biz.Staff) (bool, error) {
	var total int64
	// 验证是否已经创建
	r.data.db.Model(&Staff{}).Where("id = ? and is_deleted != ?", s.ID, s.IsDeleted).Count(&total)
	if total == 0 {
		return false, status.Errorf(codes.NotFound, "access is not exist.")
	}
	values := map[string]interface{}{
		"is_deleted": s.IsDeleted,
	}
	res := r.data.db.Model(&Staff{}).Where("id = ?", s.ID).Updates(&values)
	if res.Error != nil {
		return false, status.Errorf(codes.Internal, res.Error.Error())
	}
	return true, nil
}

func (r *staffRepo) List(_ context.Context, s *biz.Staff, pn int, pSize int) ([]*biz.Staff, int64, error) {
	var total int64
	tx := r.data.db.Model(&Staff{}).Where("is_deleted", s.IsDeleted)
	if s.Mobile != "" {
		tx = tx.Where("mobile = ?", s.Mobile)
	}
	if s.Name != "" {
		tx = tx.Where("name LIKE %?%", s.Name)
	}
	exec := tx.Count(&total)
	if exec.Error != nil || total == 0 {
		return nil, 0, exec.Error
	}
	var staffs []Staff
	result := make([]*biz.Staff, 0)
	tx.Order("created_at desc").Scopes(paginate(pn, pSize)).Find(&staffs)
	for _, staff := range staffs {
		result = append(result, &biz.Staff{
			ID:        staff.ID,
			Mobile:    staff.Mobile,
			Name:      staff.Name,
			CreatedAt: staff.CreatedAt,
			UpdatedAt: staff.UpdatedAt,
		})
	}
	return result, total, nil
}

// paginate 分页参数调整
func paginate(page, pageSize int) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		if page == 0 {
			page = 1
		}
		switch {
		case pageSize > 100:
			pageSize = 100
		case pageSize <= 0:
			pageSize = 10
		}
		offset := (page - 1) * pageSize
		return db.Offset(offset).Limit(pageSize)
	}
}
