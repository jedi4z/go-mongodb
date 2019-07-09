package registry

import (
	"github.com/jedi4z/go-mongodb/app/domain/service"
	"github.com/jedi4z/go-mongodb/app/interface/persistence/mongodb"
	"github.com/jedi4z/go-mongodb/app/usecase"
	"github.com/sarulabs/di"
)

type Container struct {
	ctn di.Container
}

func NewContainer() (*Container, error) {
	builder, err := di.NewBuilder()
	if err != nil {
		return nil, err
	}

	if err := builder.Add([]di.Def{
		{
			Name:  "user-usecase",
			Build: buildUserUsecase,
		},
	}...); err != nil {
		return nil, err
	}

	return &Container{
		ctn: builder.Build(),
	}, nil
}

func (c *Container) Resolve(name string) interface{} {
	return c.ctn.Get(name)
}

func (c *Container) Clean() error {
	return c.ctn.Clean()
}

func buildUserUsecase(ctn di.Container) (interface{}, error) {
	repo := mongodb.NewUserRepository()
	svc := service.NewUserService(repo)

	return usecase.NewUserUsecase(repo, svc), nil
}
