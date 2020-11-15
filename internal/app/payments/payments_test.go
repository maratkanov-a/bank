package payments

import (
	"github.com/maratkanov-a/bank/internal/pkg/repository/mock"
	"testing"
)

type testImplementation struct {
	*Implementation

	prMock *mock.PaymentRepositoryMock
}

func newTestImplementation(t *testing.T) *testImplementation {
	prMock := mock.NewPaymentRepositoryMock(t)

	return &testImplementation{
		Implementation: NewPayments(prMock),
		prMock:         prMock,
	}
}
