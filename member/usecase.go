package member

import (
	"context"

	m "github.com/robertotambunan/knowsmore/member/model"
)

// Usecase : represent what you can do from repository in member module
type Usecase interface {
	GetByAutocomplete(ctx context.Context, keyword string) (members m.Response, err error)
}
