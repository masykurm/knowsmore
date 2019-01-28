package usecase

import (
	"context"
	"time"

	"github.com/robertotambunan/knowsmore/member"
	m "github.com/robertotambunan/knowsmore/member/model"
)

type memberUsecase struct {
	memberRepo member.Repository
	timeout    time.Duration
}

// NewMemberUsecase create object representation from new member usecase
func NewMemberUsecase(repo member.Repository, timeout time.Duration) member.Usecase {
	return &memberUsecase{repo, timeout}
}

func (mu *memberUsecase) GetByAutocomplete(ctx context.Context, keyword string) (resp m.Response, err error) {
	ctx, cancel := context.WithTimeout(ctx, mu.timeout)
	defer cancel()

	var members []m.Member
	members, err = mu.memberRepo.GetByAutocomplete(ctx, keyword)
	if err != nil {
		resp.Error = err.Error()
		return
	}

	// add more operation here, process something, get another required data, etc
	resp.Members = members

	return
}

// CheckBlacklistedKeyword TODO
func (mu *memberUsecase) CheckBlacklistedKeyword(context.Context, string) (m.Response, error) {
	return m.Response{}, nil
}
