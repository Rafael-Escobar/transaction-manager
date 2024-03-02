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

var _ = Describe("AccountHandler", func() {
	var (
		createAccountUseCase  usecases.CreateAccountUseCase
		getAccountUseCase     usecases.GetAccountUseCase
		logger                *zap.Logger
		accountController     *AccountHandler
		router                *gin.Engine
		accountRepository     *mocks.AccountRepository
		accountID             int64
		validDocumentNumber   string
		inValidDocumentNumber string
		validAccount          *domain.Account
		invalidAccount        *domain.Account
		validAccountRequest   CreateAccountRequest
		inValidAccountRequest CreateAccountRequest
	)

	BeforeEach(func() {
		logger, _ = zap.NewDevelopment()
		accountRepository = &mocks.AccountRepository{}
		createAccountUseCase = usecases.NewCreateAccountUseCase(accountRepository, logger)
		getAccountUseCase = usecases.NewGetAccountUseCase(accountRepository, logger)
		accountController = NewAccountHandler(createAccountUseCase, getAccountUseCase, logger)
		router = gin.Default()
		router.POST("/v1/accounts", accountController.CreateAccountHandler)
		router.GET("/v1/accounts/:id", accountController.GetAccountHandler)

		accountID = int64(1)
		validDocumentNumber = "79754271011"
		inValidDocumentNumber = "76793495098"
		validAccount = &domain.Account{
			DocumentNumber: validDocumentNumber,
		}
		invalidAccount = &domain.Account{
			DocumentNumber: inValidDocumentNumber,
		}
		validAccountRequest = CreateAccountRequest{DocumentNumber: validDocumentNumber}
		inValidAccountRequest = CreateAccountRequest{DocumentNumber: inValidDocumentNumber}
	})

	Describe("CreateAccountHandler", func() {
		Context("when the request body is valid", func() {
			It("should return status 200", func() {
				accountRepository.On("FindByDocumentNumber", validDocumentNumber).Return(nil, nil)
				accountRepository.On("Create", validAccount).Return(accountID, nil)
				reqBody, _ := json.Marshal(validAccountRequest)
				req, _ := http.NewRequest("POST", "/v1/accounts", bytes.NewBuffer(reqBody))
				req.Header.Set("Content-Type", "application/json")
				resp := httptest.NewRecorder()
				router.ServeHTTP(resp, req)
				Expect(resp.Code).To(Equal(http.StatusOK))
			})
		})

		Context("when the request body is valid, but fail on save account", func() {
			It("should return status 500", func() {
				accountRepository.On("FindByDocumentNumber", validDocumentNumber).Return(nil, nil)
				accountRepository.On("Create", validAccount).Return(int64(0), sql.ErrConnDone)
				reqBody, _ := json.Marshal(validAccountRequest)
				req, _ := http.NewRequest("POST", "/v1/accounts", bytes.NewBuffer(reqBody))
				req.Header.Set("Content-Type", "application/json")
				resp := httptest.NewRecorder()
				router.ServeHTTP(resp, req)
				Expect(resp.Code).To(Equal(http.StatusInternalServerError))
			})
		})

		Context("when the request body is valid, but the account already exists", func() {
			It("should return status 409", func() {
				accountRepository.On("FindByDocumentNumber", validDocumentNumber).Return(validAccount, nil)
				reqBody, _ := json.Marshal(validAccountRequest)
				req, _ := http.NewRequest("POST", "/v1/accounts", bytes.NewBuffer(reqBody))
				req.Header.Set("Content-Type", "application/json")
				resp := httptest.NewRecorder()
				router.ServeHTTP(resp, req)
				Expect(resp.Code).To(Equal(http.StatusConflict))
			})
		})

		Context("when the request body is valid, but the the document number isn't", func() {
			It("should return status 400", func() {
				reqBody, _ := json.Marshal(inValidAccountRequest)
				req, _ := http.NewRequest("POST", "/v1/accounts", bytes.NewBuffer(reqBody))
				req.Header.Set("Content-Type", "application/json")
				resp := httptest.NewRecorder()
				router.ServeHTTP(resp, req)
				Expect(resp.Code).To(Equal(http.StatusBadRequest))
			})
		})

		Context("when the request body is invalid", func() {
			It("should return status 400", func() {
				reqBody, _ := json.Marshal(`{"document_number": "1234567890"`)
				req, _ := http.NewRequest("POST", "/v1/accounts", bytes.NewBuffer(reqBody))
				req.Header.Set("Content-Type", "application/json")
				resp := httptest.NewRecorder()
				router.ServeHTTP(resp, req)
				Expect(resp.Code).To(Equal(http.StatusBadRequest))
			})
		})
	})

	Describe("GetAccountHandler", func() {
		Context("when the account ID is valid", func() {
			It("should return status 200", func() {
				accountRepository.On("FindByID", accountID).Return(&domain.Account{ID: accountID, DocumentNumber: validDocumentNumber}, nil)
				req, _ := http.NewRequest("GET", "/v1/accounts/1", nil)
				resp := httptest.NewRecorder()
				router.ServeHTTP(resp, req)
				Expect(resp.Code).To(Equal(http.StatusOK))
			})
		})

		Context("when the account ID is valid, but fail on find", func() {
			It("should return status 200", func() {
				accountRepository.On("FindByID", accountID).Return(nil, sql.ErrConnDone)
				req, _ := http.NewRequest("GET", "/v1/accounts/1", nil)
				resp := httptest.NewRecorder()
				router.ServeHTTP(resp, req)
				Expect(resp.Code).To(Equal(http.StatusInternalServerError))
			})
		})

		Context("when the account ID is invalid", func() {
			It("should return status 400", func() {
				req, _ := http.NewRequest("GET", "/v1/accounts/invalid", nil)
				resp := httptest.NewRecorder()
				router.ServeHTTP(resp, req)
				Expect(resp.Code).To(Equal(http.StatusBadRequest))
			})
		})

		Context("when the account is not found", func() {
			It("should return status 404", func() {
				accountRepository.On("FindByID", accountID).Return(nil, domain.ErrAccountNotFound)
				req, _ := http.NewRequest("GET", "/v1/accounts/1", nil)
				resp := httptest.NewRecorder()
				router.ServeHTTP(resp, req)
				Expect(resp.Code).To(Equal(http.StatusNotFound))
			})
		})
		Context("when the account is not found", func() {
			It("should return status 404", func() {
				fmt.Println(invalidAccount, validAccount, accountID)
			})
		})
	})
})

type WrongCreateAccountRequest struct {
	DocumentNumber string `json:"docuent_number"`
}
