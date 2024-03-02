package usecases

import (
	"context"
	"database/sql"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"github.com/transaction-manager/internal/domain"
	"github.com/transaction-manager/tests/mocks"
	"go.uber.org/zap"
)

var _ = Describe("GetAccountUseCase", func() {

	var (
		useCase           *getAccountUseCase
		accountRepository *mocks.AccountRepository
		logger            *zap.Logger
	)
	BeforeEach(func() {
		accountRepository = &mocks.AccountRepository{}
		logger = zap.NewNop()
		useCase = NewGetAccountUseCase(accountRepository, logger)
	})
	Context("GetAccountUseCase", func() {
		When("The account ID is from one existent account", func() {
			It("Return the account information", func() {
				accountID := int64(1)
				accountRepository.On("FindByID", accountID).Return(&domain.Account{
					ID:             1,
					DocumentNumber: "12345678900",
				}, nil)
				account, err := useCase.Run(context.Background(), accountID)
				Expect(account).ToNot(BeNil())
				Expect(err).ToNot(HaveOccurred())
			})
		})

		When("The account ID are not related to a account", func() {
			It("Returns a domain error", func() {
				accountID := int64(1)
				accountRepository.On("FindByID", accountID).Return(nil, nil)
				account, err := useCase.Run(context.Background(), accountID)
				Expect(account).To(BeNil())
				Expect(err).To(HaveOccurred())
				Expect(err).To(Equal(domain.ErrAccountNotFound))
			})
		})

		When("Fail to search on database", func() {
			It("Returns an error", func() {
				accountID := int64(1)
				accountRepository.On("FindByID", accountID).Return(nil, sql.ErrConnDone)
				account, err := useCase.Run(context.Background(), accountID)
				Expect(account).To(BeNil())
				Expect(err).To(HaveOccurred())
				Expect(err).ToNot(Equal(domain.ErrAccountNotFound))
			})
		})
	})
})
