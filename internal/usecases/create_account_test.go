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

var _ = Describe("CreateAccountUseCase", func() {

	var (
		useCase               *createAccountUseCase
		accountRepository     *mocks.AccountRepository
		logger                *zap.Logger
		accountID             int64
		validDocumentNumber   string
		inValidDocumentNumber string
		validAccount          *domain.Account
		invalidAccount        *domain.Account
	)
	BeforeEach(func() {
		accountRepository = &mocks.AccountRepository{}
		logger = zap.NewNop()
		useCase = NewCreateAccountUseCase(accountRepository, logger)
		accountID = int64(1)
		validDocumentNumber = "76793495097"
		inValidDocumentNumber = "76793495098"
		validAccount = &domain.Account{
			ID:             accountID,
			DocumentNumber: validDocumentNumber,
		}
		invalidAccount = &domain.Account{
			DocumentNumber: inValidDocumentNumber,
		}
	})
	Context("CreateAccountUseCase", func() {
		When("Fail on verify account ID is from an existent account", func() {
			It("Returns an error", func() {
				accountRepository.On("FindByDocumentNumber", validDocumentNumber).Return(nil, sql.ErrConnDone)
				accountID, err := useCase.Run(context.Background(), validAccount)
				Expect(accountID).To(BeZero())
				Expect(err).To(HaveOccurred())
			})
		})
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
		When("The account ID are not related to a account, but fail on save.", func() {
			It("Returns an error", func() {
				accountRepository.On("FindByDocumentNumber", validDocumentNumber).Return(nil, nil)
				accountRepository.On("Create", validAccount).Return(int64(0), sql.ErrConnDone)
				accountID, err := useCase.Run(context.Background(), validAccount)
				Expect(accountID).To(BeZero())
				Expect(err).To(HaveOccurred())
			})
		})

		When("The account ID are not related to a account, but the document number ir invalid", func() {
			It("Returns a domain error ErrInvalidDocumentNumber", func() {
				accountRepository.On("FindByDocumentNumber", inValidDocumentNumber).Return(nil, nil)
				accountRepository.On("Create", validAccount).Return(accountID, nil)
				accountID, err := useCase.Run(context.Background(), invalidAccount)
				Expect(accountID).To(BeZero())
				Expect(err).To(HaveOccurred())
				Expect(err).To(Equal(domain.ErrInvalidDocumentNumber))
			})
		})
	})
})
