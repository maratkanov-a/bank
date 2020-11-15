package accounts

import (
	"github.com/maratkanov-a/bank/internal/pkg/repository/mock"
	"testing"
)

type testImplementation struct {
	*Implementation

	arMock *mock.AccountRepositoryMock
}

func newTestImplementation(t *testing.T) *testImplementation {
	arMock := mock.NewAccountRepositoryMock(t)

	return &testImplementation{
		Implementation: NewAccounts(arMock),
		arMock:         arMock,
	}
}
