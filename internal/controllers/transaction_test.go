package controllers

import (
	"bytes"
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"

	"github.com/gin-gonic/gin"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"github.com/transaction-manager/internal/domain"
	"github.com/transaction-manager/internal/usecases"
	"github.com/transaction-manager/tests/mocks"
	"go.uber.org/zap"
)

var _ = Describe("TransactionHandler", func() {
	var (
		createTransactionUseCase      usecases.CreateTransactionUseCase
		logger                        *zap.Logger
		transactionHandler            *TransactionHandler
		router                        *gin.Engine
		accountRepository             *mocks.AccountRepository
		operationTypeRepository       *mocks.OperationTypeRepository
		transactionRepository         *mocks.TransactionRepository
		accountID                     int64
		invalidAccountID              int64
		invalidOperationTypeID        int
		debitOperationTypeID          int
		creditOperationTypeID         int
		validAccount                  *domain.Account
		validCreditOperationType      *domain.OperationType
		validDebitOperationType       *domain.OperationType
		validCreditTransactionRequest CreateTransactionRequest
		validDebitTransactionRequest  CreateTransactionRequest
	)

	BeforeEach(func() {
		accountRepository = &mocks.AccountRepository{}
		operationTypeRepository = &mocks.OperationTypeRepository{}
		transactionRepository = &mocks.TransactionRepository{}
		logger, _ = zap.NewDevelopment()
		createTransactionUseCase = usecases.NewCreateTransactionUseCase(
			transactionRepository,
			accountRepository,
			operationTypeRepository,
			logger)
		transactionHandler = NewTransactionHandler(createTransactionUseCase, logger)
		router = gin.Default()
		router.POST("/v1/transactions", transactionHandler.CreateTransactionHandler)

		accountID = int64(1)
		debitOperationTypeID = 1
		creditOperationTypeID = 4
		validCreditTransactionRequest = CreateTransactionRequest{
			AccountID:       accountID,
			Amount:          100.0,
			OperationTypeID: creditOperationTypeID,
		}
		validDebitTransactionRequest = CreateTransactionRequest{
			AccountID:       accountID,
			Amount:          -100.0,
			OperationTypeID: debitOperationTypeID,
		}
		invalidAccountID = int64(10000)
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

	Context("CreateTransactionHandler", func() {
		Context("when the request body is valid for a credit transaction", func() {
			It("should return status 200", func() {
				accountRepository.On("FindByID", accountID).Return(validAccount, nil)
				operationTypeRepository.On("FindByID", creditOperationTypeID).Return(validCreditOperationType, nil)
				transactionRequested := &domain.Transaction{
					AccountID:       accountID,
					OperationTypeID: creditOperationTypeID,
					Amount:          100.0,
				}
				transactionRepository.On("Create", transactionRequested).Return(int64(1), nil)
				reqBody, _ := json.Marshal(validCreditTransactionRequest)

				req, _ := http.NewRequest("POST", "/v1/transactions", bytes.NewBuffer(reqBody))
				req.Header.Set("Content-Type", "application/json")
				resp := httptest.NewRecorder()
				router.ServeHTTP(resp, req)

				var response CreateTransactionResponse
				err := json.Unmarshal(resp.Body.Bytes(), &response)

				Expect(resp.Code).To(Equal(http.StatusOK))
				Expect(err).To(BeNil())
				Expect(response.TransactionID).To(Equal(int64(1)))
			})
		})
		Context("when the request body is valid for a debit transaction", func() {
			It("should return status 200", func() {
				accountRepository.On("FindByID", accountID).Return(validAccount, nil)
				operationTypeRepository.On("FindByID", debitOperationTypeID).Return(validDebitOperationType, nil)
				transactionRequested := &domain.Transaction{
					AccountID:       accountID,
					OperationTypeID: debitOperationTypeID,
					Amount:          -100.0,
				}
				transactionRepository.On("Create", transactionRequested).Return(int64(1), nil)
				reqBody, _ := json.Marshal(validDebitTransactionRequest)

				req, _ := http.NewRequest("POST", "/v1/transactions", bytes.NewBuffer(reqBody))
				req.Header.Set("Content-Type", "application/json")
				resp := httptest.NewRecorder()
				router.ServeHTTP(resp, req)

				var response CreateTransactionResponse
				err := json.Unmarshal(resp.Body.Bytes(), &response)

				Expect(resp.Code).To(Equal(http.StatusOK))
				Expect(err).To(BeNil())
				Expect(response.TransactionID).To(Equal(int64(1)))
			})
		})

		Context("when the request body is invalid", func() {
			It("should return status 400", func() {
				reqBody, _ := json.Marshal(`{"document_number": "1234567890"`)

				req, _ := http.NewRequest("POST", "/v1/transactions", bytes.NewBuffer(reqBody))
				req.Header.Set("Content-Type", "application/json")
				resp := httptest.NewRecorder()

				router.ServeHTTP(resp, req)

				Expect(resp.Code).To(Equal(http.StatusBadRequest))
			})
		})

		Context("when an error occurs during transaction creation", func() {
			It("should return status 500", func() {
				accountRepository.On("FindByID", accountID).Return(validAccount, nil)
				operationTypeRepository.On("FindByID", debitOperationTypeID).Return(validDebitOperationType, nil)
				transactionRequested := &domain.Transaction{
					AccountID:       accountID,
					OperationTypeID: debitOperationTypeID,
					Amount:          -100.0,
				}
				transactionRepository.On("Create", transactionRequested).Return(int64(0), sql.ErrConnDone)
				reqBody, _ := json.Marshal(validDebitTransactionRequest)

				req, _ := http.NewRequest("POST", "/v1/transactions", bytes.NewBuffer(reqBody))
				req.Header.Set("Content-Type", "application/json")
				resp := httptest.NewRecorder()

				router.ServeHTTP(resp, req)

				Expect(resp.Code).To(Equal(http.StatusInternalServerError))
			})
		})

		Context("when the account informed is invalid", func() {
			It("should return status 400", func() {
				accountRepository.On("FindByID", accountID).Return(nil, nil)
				operationTypeRepository.On("FindByID", debitOperationTypeID).Return(nil, nil)
				reqBody, _ := json.Marshal(validDebitTransactionRequest)

				req, _ := http.NewRequest("POST", "/v1/transactions", bytes.NewBuffer(reqBody))
				req.Header.Set("Content-Type", "application/json")
				resp := httptest.NewRecorder()

				router.ServeHTTP(resp, req)

				Expect(resp.Code).To(Equal(http.StatusBadRequest))
			})
		})

		Context("when the amount informed is invalid", func() {
			It("should return status 400", func() {
				accountRepository.On("FindByID", accountID).Return(validAccount, nil)
				operationTypeRepository.On("FindByID", debitOperationTypeID).Return(validCreditOperationType, nil)
				reqBody, _ := json.Marshal(validDebitTransactionRequest)

				req, _ := http.NewRequest("POST", "/v1/transactions", bytes.NewBuffer(reqBody))
				req.Header.Set("Content-Type", "application/json")
				resp := httptest.NewRecorder()

				router.ServeHTTP(resp, req)

				Expect(resp.Code).To(Equal(http.StatusBadRequest))
			})
		})

		Context("when the operation type informed is invalid", func() {
			It("should return status 400", func() {
				accountRepository.On("FindByID", accountID).Return(validAccount, nil)
				operationTypeRepository.On("FindByID", debitOperationTypeID).Return(nil, nil)
				reqBody, _ := json.Marshal(validDebitTransactionRequest)

				req, _ := http.NewRequest("POST", "/v1/transactions", bytes.NewBuffer(reqBody))
				req.Header.Set("Content-Type", "application/json")
				resp := httptest.NewRecorder()

				router.ServeHTTP(resp, req)

				Expect(resp.Code).To(Equal(http.StatusBadRequest))
			})
		})
		Context("when an error occurs during transaction creation", func() {
			It("should return status 500", func() {
				fmt.Println(validDebitTransactionRequest)
				fmt.Println(accountID)
				fmt.Println(invalidAccountID)
				fmt.Println(invalidOperationTypeID)
				fmt.Println(debitOperationTypeID)
				fmt.Println(creditOperationTypeID)
				fmt.Println(validAccount)
				fmt.Println(validCreditOperationType)
				fmt.Println(validDebitOperationType)
			})
		})
	})
})
