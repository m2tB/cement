package biz

import (
	"context"
	"time"

	"github.com/go-kratos/kratos/v2/log"
)

// Team is a Team model.
type Team struct {
	ID          int64
	CName       string
	EName       string
	PreTeamCode string
	CreatedAt   *time.Time
	UpdatedAt   *time.Time
	IsDeleted   int
}

// TeamRepo is a Team repo.
type TeamRepo interface {
	Save(context.Context, *Team) (bool, error)
	Update(context.Context, *Team) (bool, error)
	Delete(context.Context, *Team) (bool, error)
	Read(context.Context, *Team) (*Team, error)

	List(context.Context, *Team, int, int) ([]*Team, int64, error)
}

// TeamUsecase is a Team usecase.
type TeamUsecase struct {
	repo TeamRepo
	log  *log.Helper
}

// NewTeamUsecase new a Team usecase.
func NewTeamUsecase(repo TeamRepo, logger log.Logger) *TeamUsecase {
	return &TeamUsecase{repo: repo, log: log.NewHelper(log.With(logger, "module", "usecase/team"))}
}

// CreateTeam creates a Team, and returns the exec result.
func (uc *TeamUsecase) CreateTeam(ctx context.Context, t *Team) (bool, error) {
	return uc.repo.Save(ctx, t)
}

// UpdateTeam update a Team, and returns the exec result.
func (uc *TeamUsecase) UpdateTeam(ctx context.Context, t *Team) (bool, error) {
	return uc.repo.Update(ctx, t)
}

// DeleteTeam delete a Team, and returns the exec result.
func (uc *TeamUsecase) DeleteTeam(ctx context.Context, t *Team) (bool, error) {
	return uc.repo.Delete(ctx, t)
}

// ReadTeam read a Team, and returns the result.
func (uc *TeamUsecase) ReadTeam(ctx context.Context, t *Team) (*Team, error) {
	return uc.repo.Read(ctx, t)
}

// ListTeam returns the Team List.
func (uc *TeamUsecase) ListTeam(ctx context.Context, t *Team, pn int, pSize int) ([]*Team, int64, error) {
	return uc.repo.List(ctx, t, pn, pSize)
}
