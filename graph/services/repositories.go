package services

import (
	"context"

	"github.com/kkoji/graphql-sample/graph/db"
	"github.com/kkoji/graphql-sample/graph/model"

	"github.com/volatiletech/sqlboiler/v4/boil"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
)

type repositoryService struct {
	exec boil.ContextExecutor
}

func convertRepository(repo *db.Repository) *model.Repository {
	return &model.Repository{
		ID:        repo.ID,
		Owner:     &model.User{ID: repo.Owner},
		Name:      repo.Name,
		CreatedAt: repo.CreatedAt,
	}
}

func (u *repositoryService) GetRepositoryByID(ctx context.Context, id string) (*model.Repository, error) {
	repo, err := db.Repositories(
		qm.Select(
			db.RepositoryTableColumns.ID,
			db.RepositoryTableColumns.Name,
			db.RepositoryTableColumns.Owner,
			db.RepositoryTableColumns.CreatedAt,
		),
		db.RepositoryWhere.ID.EQ(id),
	).One(ctx, u.exec)
	if err != nil {
		return nil, err
	}

	return convertRepository(repo), nil
}

func (rs *repositoryService) GetRepositoryByNameAndOwner(ctx context.Context, name, owner string) (*model.Repository, error) {
	repo, err := db.Repositories(
		qm.Select(
			db.RepositoryTableColumns.ID,
			db.RepositoryTableColumns.Name,
			db.RepositoryTableColumns.Owner,
			db.RepositoryTableColumns.CreatedAt,
		),
		db.RepositoryWhere.Name.EQ(name),
		db.RepositoryWhere.Owner.EQ(owner),
	).One(ctx, rs.exec)
	if err != nil {
		return nil, err
	}

	return convertRepository(repo), nil
}
