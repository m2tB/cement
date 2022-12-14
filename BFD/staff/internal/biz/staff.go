package biz

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	"time"
)

// Staff is a Staff model.
type Staff struct {
	ID        int64
	Mobile    string
	Name      string
	IsDeleted int
}

type DataStaff struct {
	ID        int64
	Mobile    string
	Name      string
	CreatedAt *time.Time
	UpdatedAt *time.Time
	IsDeleted int
}

// StaffRepo is a Staff repo.
type StaffRepo interface {
	Save(context.Context, *Staff) (bool, error)
	Update(context.Context, *Staff) (bool, error)
	Delete(context.Context, *Staff) (bool, error)
	Read(context.Context, *Staff) (*DataStaff, error)

	List(context.Context, *Staff, int, int) ([]*DataStaff, int64, error)
}

// StaffUsecase is a Staff usecase.
type StaffUsecase struct {
	repo StaffRepo
	log  *log.Helper
}

// NewStaffUsecase new a Staff usecase.
func NewStaffUsecase(repo StaffRepo, logger log.Logger) *StaffUsecase {
	return &StaffUsecase{repo: repo, log: log.NewHelper(log.With(logger, "module", "usecase/staff"))}
}

// CreateStaff creates a Staff, and returns the exec result.
func (uc *StaffUsecase) CreateStaff(ctx context.Context, s *Staff) (bool, error) {
	return uc.repo.Save(ctx, s)
}

// UpdateStaff update a Staff, and returns the exec result.
func (uc *StaffUsecase) UpdateStaff(ctx context.Context, s *Staff) (bool, error) {
	return uc.repo.Update(ctx, s)
}

// DeleteStaff delete a Staff, and returns the exec result.
func (uc *StaffUsecase) DeleteStaff(ctx context.Context, s *Staff) (bool, error) {
	return uc.repo.Delete(ctx, s)
}

// ReadStaff read a Staff, and returns the result.
func (uc *StaffUsecase) ReadStaff(ctx context.Context, s *Staff) (*DataStaff, error) {
	return uc.repo.Read(ctx, s)
}

// ListStaff returns the Staff List.
func (uc *StaffUsecase) ListStaff(ctx context.Context, s *Staff, pn int, pSize int) ([]*DataStaff, int64, error) {
	return uc.repo.List(ctx, s, pn, pSize)
}
