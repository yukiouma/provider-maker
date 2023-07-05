package usecase

type Repo interface {
	Demo()
}

type Usecase struct {
	repo Repo
}

// +provider
func NewUsecase(repo Repo) *Usecase {
	return &Usecase{
		repo: repo,
	}
}
