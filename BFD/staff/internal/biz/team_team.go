package biz

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	"time"
)

// TeamTeam is a Relationship table between Team and Team model.
type TeamTeam struct {
	TID int64
	PID int64

	OpID int64
}

type DataListTeamTeam struct {
	TID       int64
	TName     string
	PID       int64
	PName     string
	CreatedAt *time.Time
}

// TeamTeamRepo is a TeamTeam repo.
type TeamTeamRepo interface {
	Save(context.Context, *TeamTeam) (bool, error)
	Update(context.Context, *TeamTeam) (bool, error)
	Delete(context.Context, *TeamTeam) (bool, error)

	List(context.Context, *TeamTeam, int, int) ([]*DataListTeamTeam, int64, error)
}

// TeamTeamUsecase is a TeamTeam usecase.
type TeamTeamUsecase struct {
	repo TeamTeamRepo
	log  *log.Helper
}

// NewTeamTeamUsecase new a TeamTeam usecase.
func NewTeamTeamUsecase(repo TeamTeamRepo, logger log.Logger) *TeamTeamUsecase {
	return &TeamTeamUsecase{repo: repo, log: log.NewHelper(log.With(logger, "module", "usecase/team_team"))}
}

// CreateSubTeam creates a TeamTeam, and returns the exec result.
func (uc *TeamTeamUsecase) CreateSubTeam(ctx context.Context, t *TeamTeam) (bool, error) {
	return uc.repo.Save(ctx, t)
}

// UpdateSubTeam delete a TeamTeam, and returns the exec result.
func (uc *TeamTeamUsecase) UpdateSubTeam(ctx context.Context, t *TeamTeam) (bool, error) {
	return uc.repo.Update(ctx, t)
}

// DeleteSubTeam delete a TeamTeam, and returns the exec result.
func (uc *TeamTeamUsecase) DeleteSubTeam(ctx context.Context, t *TeamTeam) (bool, error) {
	return uc.repo.Delete(ctx, t)
}

// ListTeamSubTeam returns the TeamTeam List.
func (uc *TeamTeamUsecase) ListTeamSubTeam(ctx context.Context, t *TeamTeam, pn int, pSize int) ([]*DataListTeamTeam, int64, error) {
	return uc.repo.List(ctx, t, pn, pSize)
}
