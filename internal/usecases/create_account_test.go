package usecases

import (
	"context"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"github.com/transaction-manager/internal/domain"
	"github.com/transaction-manager/tests/mocks"
	"go.uber.org/zap"
)

var _ = Describe("CreateAccountUseCase", func() {

	var (
		useCase             *CreateAccountUseCase
		accountRepository   *mocks.AccountRepository
		logger              *zap.Logger
		accountID           int64
		validDocumentNumber string
		// inValidDocumentNumber string
		validAccount *domain.Account
	)
	BeforeEach(func() {
		accountRepository = &mocks.AccountRepository{}
		logger = zap.NewNop()
		useCase = NewCreateAccountUseCase(accountRepository, logger)
		accountID = int64(1)
		validDocumentNumber = "76793495097"
		validAccount = &domain.Account{
			ID:             accountID,
			DocumentNumber: validDocumentNumber,
		}
	})
	Context("CreateAccountUseCase", func() {
		When("The account ID is from one existent account", func() {
			It("Returns a domain error ErrAccountAlreadyExists", func() {
				accountRepository.On("FindByDocumentNumber", validDocumentNumber).Return(validAccount, nil)
				accountID, err := useCase.Run(context.Background(), validAccount)
				Expect(accountID).To(BeZero())
				Expect(err).To(HaveOccurred())
				Expect(err).To(Equal(domain.ErrAccountAlreadyExists))
			})
		})

		When("The account ID are not related to a account", func() {
			It("Returns no error and account ID", func() {
				accountRepository.On("FindByDocumentNumber", validDocumentNumber).Return(nil, nil)
				accountRepository.On("Create", validAccount).Return(accountID, nil)
				accountID, err := useCase.Run(context.Background(), validAccount)
				Expect(accountID).ToNot(BeZero())
				Expect(err).ToNot(HaveOccurred())
			})
		})

	})
})
