package datasource

import "github.com/yukiouma/provider-maker/examples/demo/usecase"

type repo struct{}

// +provider
func NewRepo() usecase.Repo {
	return &repo{}
}

func (r *repo) Demo() {}
