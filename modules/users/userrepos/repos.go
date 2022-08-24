package userrepos

import "template/components/appcontext"

type UserRepos struct {
	ctx appcontext.AppContext
}

func NewUserRepos(ctx appcontext.AppContext) *UserRepos {
	return &UserRepos{
		ctx: ctx,
	}
}
