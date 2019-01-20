package member

import (
	"context"

	m "github.com/robertotambunan/knowsmore/member/model"
)

//go:generate moq -out repository_test.go . Repository

// Repository : represent what you can do in member module
type Repository interface {
	GetByID(ctx context.Context, id string) (member m.Member, err error)
	GetByAutocomplete(ctx context.Context, keyword string) (members []m.Member, err error)
}
