// Code generated by moq; DO NOT EDIT.
// github.com/matryer/moq

package member

import (
	"context"
	"github.com/robertotambunan/knowsmore/member/model"
	"sync"
)

var (
	lockRepositoryMockGetByAutocomplete sync.RWMutex
)

// RepositoryMock is a mock implementation of Repository.
//
//     func TestSomethingThatUsesRepository(t *testing.T) {
//
//         // make and configure a mocked Repository
//         mockedRepository := &RepositoryMock{
//             GetByAutocompleteFunc: func(ctx context.Context, keyword string) ([]member.Member, error) {
// 	               panic("mock out the GetByAutocomplete method")
//             },
//         }
//
//         // use mockedRepository in code that requires Repository
//         // and then make assertions.
//
//     }
type RepositoryMock struct {
	// GetByAutocompleteFunc mocks the GetByAutocomplete method.
	GetByAutocompleteFunc func(ctx context.Context, keyword string) ([]member.Member, error)

	// calls tracks calls to the methods.
	calls struct {
		// GetByAutocomplete holds details about calls to the GetByAutocomplete method.
		GetByAutocomplete []struct {
			// Ctx is the ctx argument value.
			Ctx context.Context
			// Keyword is the keyword argument value.
			Keyword string
		}
	}
}

// GetByAutocomplete calls GetByAutocompleteFunc.
func (mock *RepositoryMock) GetByAutocomplete(ctx context.Context, keyword string) ([]member.Member, error) {
	if mock.GetByAutocompleteFunc == nil {
		panic("RepositoryMock.GetByAutocompleteFunc: method is nil but Repository.GetByAutocomplete was just called")
	}
	callInfo := struct {
		Ctx     context.Context
		Keyword string
	}{
		Ctx:     ctx,
		Keyword: keyword,
	}
	lockRepositoryMockGetByAutocomplete.Lock()
	mock.calls.GetByAutocomplete = append(mock.calls.GetByAutocomplete, callInfo)
	lockRepositoryMockGetByAutocomplete.Unlock()
	return mock.GetByAutocompleteFunc(ctx, keyword)
}

// GetByAutocompleteCalls gets all the calls that were made to GetByAutocomplete.
// Check the length with:
//     len(mockedRepository.GetByAutocompleteCalls())
func (mock *RepositoryMock) GetByAutocompleteCalls() []struct {
	Ctx     context.Context
	Keyword string
} {
	var calls []struct {
		Ctx     context.Context
		Keyword string
	}
	lockRepositoryMockGetByAutocomplete.RLock()
	calls = mock.calls.GetByAutocomplete
	lockRepositoryMockGetByAutocomplete.RUnlock()
	return calls
}