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

var _ = Describe("CreateTransactionUseCase", func() {

	var (
		useCase                  *CreateTransactionUseCase
		accountRepository        *mocks.AccountRepository
		operationTypeRepository  *mocks.OperationTypeRepository
		transactionRepository    *mocks.TransactionRepository
		logger                   *zap.Logger
		accountID                int64
		invalidAccountID         int64
		invalidOperationTypeID   int
		debitOperationTypeID     int
		creditOperationTypeID    int
		validAccount             *domain.Account
		validCreditOperationType *domain.OperationType
		validDebitOperationType  *domain.OperationType
	)
	BeforeEach(func() {
		accountRepository = &mocks.AccountRepository{}
		operationTypeRepository = &mocks.OperationTypeRepository{}
		transactionRepository = &mocks.TransactionRepository{}
		logger = zap.NewNop()
		useCase = NewCreateTransactionUseCase(
			transactionRepository,
			accountRepository,
			operationTypeRepository,
			logger)
		accountID = int64(1)
		invalidAccountID = int64(10000)
		debitOperationTypeID = 1
		creditOperationTypeID = 4
		invalidOperationTypeID = 40000
		validAccount = &domain.Account{
			ID:             accountID,
			DocumentNumber: "12345678900",
		}
		validCreditOperationType = &domain.OperationType{
			ID:          creditOperationTypeID,
			Description: "credit",
			IsDebit:     false,
		}
		validDebitOperationType = &domain.OperationType{
			ID:          debitOperationTypeID,
			Description: "debit",
			IsDebit:     true,
		}
	})
	Context("CreateTransactionUseCase", func() {
		When("The account ID is invalid", func() {
			It("Returns a domain error invalid account", func() {
				accountRepository.On("FindByID", invalidAccountID).Return(nil, nil)
				transactionRequested := &domain.Transaction{
					AccountID:       invalidAccountID,
					OperationTypeID: debitOperationTypeID,
					Amount:          100.0,
				}
				account, err := useCase.Run(context.Background(), transactionRequested)
				Expect(account).To(Equal(int64(0)))
				Expect(err).To(HaveOccurred())
				Expect(err).To(Equal(domain.ErrInvalidAccount))
			})
		})
		When("The account ID is valid but fail on verify", func() {
			It("Returns an error", func() {
				accountRepository.On("FindByID", accountID).Return(nil, sql.ErrConnDone)
				transactionRequested := &domain.Transaction{
					AccountID:       accountID,
					OperationTypeID: invalidOperationTypeID,
					Amount:          100.0,
				}
				account, err := useCase.Run(context.Background(), transactionRequested)
				Expect(account).ToNot(BeNil())
				Expect(err).To(HaveOccurred())
			})
		})
		When("The account ID is valid but the operation ID is not", func() {
			It("Returns a domain error invalid operation type", func() {
				accountRepository.On("FindByID", accountID).Return(validAccount, nil)
				transactionRequested := &domain.Transaction{
					AccountID:       accountID,
					OperationTypeID: invalidOperationTypeID,
					Amount:          100.0,
				}
				operationTypeRepository.On("FindByID", invalidOperationTypeID).Return(nil, nil)
				account, err := useCase.Run(context.Background(), transactionRequested)
				Expect(account).ToNot(BeNil())
				Expect(err).To(HaveOccurred())
				Expect(err).To(Equal(domain.ErrInvalidOperationType))
			})
		})

		When("The account are valid but fail on verify operation type", func() {
			It("Returns an error", func() {
				accountRepository.On("FindByID", accountID).Return(validAccount, nil)
				transactionRequested := &domain.Transaction{
					AccountID:       accountID,
					OperationTypeID: debitOperationTypeID,
					Amount:          100.0,
				}
				operationTypeRepository.On("FindByID", debitOperationTypeID).Return(nil, sql.ErrConnDone)
				account, err := useCase.Run(context.Background(), transactionRequested)
				Expect(account).ToNot(BeNil())
				Expect(err).To(HaveOccurred())
			})
		})

		When("The account and debit operation type ID are valid but the amount not", func() {
			It("Returns a domain error invalid amount for the operation type", func() {
				accountRepository.On("FindByID", accountID).Return(validAccount, nil)
				transactionRequested := &domain.Transaction{
					AccountID:       accountID,
					OperationTypeID: debitOperationTypeID,
					Amount:          100.0,
				}
				operationTypeRepository.On("FindByID", debitOperationTypeID).Return(validDebitOperationType, nil)
				account, err := useCase.Run(context.Background(), transactionRequested)
				Expect(account).ToNot(BeNil())
				Expect(err).To(HaveOccurred())
				Expect(err).To(Equal(domain.ErrInvalidAmountForOperationType))
			})
		})

		When("The account and credit operation type ID are valid but the amount not", func() {
			It("Returns a domain error invalid amount for the operation type", func() {
				accountRepository.On("FindByID", accountID).Return(validAccount, nil)
				transactionRequested := &domain.Transaction{
					AccountID:       accountID,
					OperationTypeID: creditOperationTypeID,
					Amount:          -100.0,
				}
				operationTypeRepository.On("FindByID", creditOperationTypeID).Return(validCreditOperationType, nil)
				account, err := useCase.Run(context.Background(), transactionRequested)
				Expect(account).ToNot(BeNil())
				Expect(err).To(HaveOccurred())
				Expect(err).To(Equal(domain.ErrInvalidAmountForOperationType))
			})
		})

		When("The account and credit operation type ID are valid but the amount is zero", func() {
			It("Returns a domain error invalid amount for the operation type", func() {
				accountRepository.On("FindByID", accountID).Return(validAccount, nil)
				transactionRequested := &domain.Transaction{
					AccountID:       accountID,
					OperationTypeID: creditOperationTypeID,
					Amount:          0.0,
				}
				operationTypeRepository.On("FindByID", creditOperationTypeID).Return(validCreditOperationType, nil)
				account, err := useCase.Run(context.Background(), transactionRequested)
				Expect(account).ToNot(BeNil())
				Expect(err).To(HaveOccurred())
				Expect(err).To(Equal(domain.ErrInvalidAmountForOperationType))
			})
		})

		When("All the information is correct but fail on save", func() {
			It("Returns no error and the transaction ID", func() {
				accountRepository.On("FindByID", accountID).Return(validAccount, nil)
				transactionRequested := &domain.Transaction{
					AccountID:       accountID,
					OperationTypeID: creditOperationTypeID,
					Amount:          100.0,
				}
				operationTypeRepository.On("FindByID", creditOperationTypeID).Return(validCreditOperationType, nil)
				transactionRepository.On("Create", transactionRequested).Return(int64(0), sql.ErrConnDone)
				account, err := useCase.Run(context.Background(), transactionRequested)
				Expect(account).To(BeZero())
				Expect(err).To(HaveOccurred())
			})
		})

		When("All the information is correct for a credit operation", func() {
			It("Returns no error and the transaction ID", func() {
				accountRepository.On("FindByID", accountID).Return(validAccount, nil)
				transactionRequested := &domain.Transaction{
					AccountID:       accountID,
					OperationTypeID: creditOperationTypeID,
					Amount:          100.0,
				}
				operationTypeRepository.On("FindByID", creditOperationTypeID).Return(validCreditOperationType, nil)
				transactionRepository.On("Create", transactionRequested).Return(int64(1), nil)
				account, err := useCase.Run(context.Background(), transactionRequested)
				Expect(account).ToNot(BeZero())
				Expect(err).ToNot(HaveOccurred())
			})
		})

		When("All the information is correct for a debit operation", func() {
			It("Returns no error and the transaction ID", func() {
				accountRepository.On("FindByID", accountID).Return(validAccount, nil)
				transactionRequested := &domain.Transaction{
					AccountID:       accountID,
					OperationTypeID: debitOperationTypeID,
					Amount:          -100.0,
				}
				operationTypeRepository.On("FindByID", debitOperationTypeID).Return(validDebitOperationType, nil)
				transactionRepository.On("Create", transactionRequested).Return(int64(1), nil)
				account, err := useCase.Run(context.Background(), transactionRequested)
				Expect(account).ToNot(BeZero())
				Expect(err).ToNot(HaveOccurred())
			})
		})
	})
})
