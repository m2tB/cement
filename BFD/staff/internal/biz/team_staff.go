package biz

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	"time"
)

// TeamStaff is a Relationship table between Team and Staff model.
type TeamStaff struct {
	TID  int64
	SID  int64
	SIDs []int64
}

type DataListTeamStaff struct {
	SID       int64
	SName     string
	TID       int64
	TName     string
	CreatedAt *time.Time
}

// TeamStaffRepo is a TeamStaff repo.
type TeamStaffRepo interface {
	Save(context.Context, *TeamStaff) (bool, error)
	Delete(context.Context, *TeamStaff) (bool, error)

	List(context.Context, *TeamStaff, int, int) ([]*DataListTeamStaff, int64, error)
}

// TeamStaffUsecase is a TeamStaff usecase.
type TeamStaffUsecase struct {
	repo TeamStaffRepo
	log  *log.Helper
}

// NewTeamStaffUsecase new a TeamStaff usecase.
func NewTeamStaffUsecase(repo TeamStaffRepo, logger log.Logger) *TeamStaffUsecase {
	return &TeamStaffUsecase{repo: repo, log: log.NewHelper(log.With(logger, "module", "usecase/team_staff"))}
}

// InviteStaff creates a TeamStaff, and returns the exec result.
func (uc *TeamStaffUsecase) InviteStaff(ctx context.Context, t *TeamStaff) (bool, error) {
	return uc.repo.Save(ctx, t)
}

// ExpelStaff delete a TeamStaff, and returns the exec result.
func (uc *TeamStaffUsecase) ExpelStaff(ctx context.Context, t *TeamStaff) (bool, error) {
	return uc.repo.Delete(ctx, t)
}

// ListTeamStaff returns the TeamStaff List.
func (uc *TeamStaffUsecase) ListTeamStaff(ctx context.Context, t *TeamStaff, pn int, pSize int) ([]*DataListTeamStaff, int64, error) {
	return uc.repo.List(ctx, t, pn, pSize)
}

// ListStaffTeam returns the TeamStaff List.
func (uc *TeamStaffUsecase) ListStaffTeam(ctx context.Context, t *TeamStaff, pn int, pSize int) ([]*DataListTeamStaff, int64, error) {
	return uc.repo.List(ctx, t, pn, pSize)
}
