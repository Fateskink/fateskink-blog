package fateskinkResolvers

import (
	"context"
	"server/graphql/gqlTypes/globalTypes"
	fateskinkInputs "server/graphql/inputs/feateskinkInputs"
	"server/services"
)

func (r *Resolver) SignIn(ctx context.Context, args fateskinkInputs.SignInInput) (*globalTypes.SignInType, error) {
	service := services.AuthService{
		Email:    args.Email,
		Password: args.Password,
		Ctx:      &ctx,
		Db:       r.Db,
	}

	err := service.Execute()

	if err != nil {
		return nil, err
	}

	return &globalTypes.SignInType{
		Token: service.Token,
	}, nil
}
