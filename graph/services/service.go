package services

import (
	"context"
	"github.com/kkoji/graphql-sample/graph/model"
	"github.com/volatiletech/sqlboiler/v4/boil"
)

type Services interface {
	UserService
	RepositoryService
	IssueService
	UserService
}

type UserService interface {
	GetUserByName(ctx context.Context, name string) (*model.User, error)
	GetUserByID(ctx context.Context, id string) (*model.User, error)
	ListUsersByID(ctx context.Context, IDs []string) ([]*model.User, error)
}

type RepositoryService interface {
	GetRepositoryByID(ctx context.Context, id string) (*model.Repository, error)
	GetRepositoryByNameAndOwner(ctx context.Context, name, owner string) (*model.Repository, error)
}

type IssueService interface {
	GetIssueByRepoAndNumber(ctx context.Context, repoID string, number int) (*model.Issue, error)
	ListIssueInRepository(ctx context.Context, repoID string, after *string, before *string, first *int, last *int) (*model.IssueConnection, error)
}

type services struct {
	*userService
	*repositoryService
	*issueService
}

func New(exec boil.ContextExecutor) Services {
	return &services{
		userService:       &userService{exec: exec},
		repositoryService: &repositoryService{exec: exec},
		issueService:      &issueService{exec: exec},
	}
}
