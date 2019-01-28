package member

import (
	"context"

	m "github.com/robertotambunan/knowsmore/member/model"
)

// Usecase : represent what you can do from repository in member module
type Usecase interface {
	GetByAutocomplete(context.Context, string) (m.Response, error)
	CheckBlacklistedKeyword(context.Context, string) (m.Response, error)
}
