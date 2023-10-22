package graph

import (
	"context"

	"github.com/friendsofgo/errors"
	"github.com/graph-gophers/dataloader/v7"

	"github.com/kkoji/graphql-sample/graph/model"
	"github.com/kkoji/graphql-sample/graph/services"
)

type Loaders struct {
	UserLoader dataloader.Interface[string, *model.User]
}

func NewLoaders(Srv services.Services) *Loaders {
	ub := &userBatcher{Srv: Srv}

	return &Loaders{
		UserLoader: dataloader.NewBatchedLoader[string, *model.User](ub.BatchGetUsers),
	}
}

type userBatcher struct {
	Srv services.Services
}

func (u *userBatcher) BatchGetUsers(ctx context.Context, IDs []string) []*dataloader.Result[*model.User] {
	results := make([]*dataloader.Result[*model.User], len(IDs))
	for i := range results {
		results[i] = &dataloader.Result[*model.User]{
			Error: errors.New("not found"),
		}
	}

	// 結果をID順にするためにidとインデックス番号のマップを作成
	indexes := make(map[string]int, len(IDs))
	for i, ID := range IDs {
		indexes[ID] = i
	}

	users, err := u.Srv.ListUsersByID(ctx, IDs)
	for _, user := range users {
		if err != nil {
			results[indexes[user.ID]] = &dataloader.Result[*model.User]{Error: err}
			continue
		}

		results[indexes[user.ID]] = &dataloader.Result[*model.User]{Data: user}
	}
	return results
}
