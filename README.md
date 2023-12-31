# Provider Maker

Help you go generate provider file when using wire

# Installation
```bash
$ go install github.com/yukiouma/provider-maker/cmd/providermaker@latest
```

# Usage
For example, we have a project like this:
```bash
./examples/demo/
├── adapter
│   └── datasource
│       └── datasource.go
├── infra
│   └── constructor
└── usecase
    └── usecase.go
```

When we start the project, we need to create instance in `./adapter/datasource/datasource.go` and `./adapter/usecase/usecase.go`, like this:
```go
// ./adapter/usecase/usecase.go
package usecase

type Repo interface {
	Demo()
}

type Usecase struct {
	repo Repo
}

func NewUsecase(repo Repo) *Usecase {
	return &Usecase{
		repo: repo,
	}
}


// ./adapter/datasource/datasource.go
package datasource

import "github.com/yukiouma/provider-maker/examples/demo/usecase"

type repo struct{}

func NewRepo() usecase.Repo {
	return &repo{}
}

func (r *repo) Demo() {}

```

Tag the factory functions, for example in `./adapter/datasource/datasource.go` 
```go
// +provider
func NewRepo() usecase.Repo {
	return &repo{}
}

```

Then run providermaker, setup your project directory and output file directory:
```bash
$ providermaker --in . --out ./internal/infra/constructor
```

After that, a file named `provider.go` will be generated in `./internal/infra/constructor`, like:
```go
// ./internal/infra/constructor/provider.go


// Code generated by provider-maker. DO NOT EDIT.

package constructor

import (
	"github.com/google/wire"
	"github.com/yukiouma/provider-maker/examples/demo/adapter/datasource"
	"github.com/yukiouma/provider-maker/examples/demo/usecase"
)

var Providers = wire.NewSet(
	datasource.NewRepo,
	usecase.NewUsecase,
)

```