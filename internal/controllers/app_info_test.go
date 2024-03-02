package controllers

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"runtime"

	"github.com/gin-gonic/gin"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"go.uber.org/zap"
)

var _ = Describe("AppInfo", func() {
	var (
		appInfoHandler *AppInfo
		router         *gin.Engine
	)

	BeforeEach(func() {
		logger, _ := zap.NewDevelopment()
		appInfoHandler = NewAppInfoHandler("v1.0.0", "abc123", "build123", logger)
		router = gin.Default()
		router.GET("/info", appInfoHandler.Handler)
	})

	Context("GET /info", func() {
		It("should return status 200", func() {
			req, _ := http.NewRequest("GET", "/info", nil)
			resp := httptest.NewRecorder()
			router.ServeHTTP(resp, req)
			Expect(resp.Code).To(Equal(http.StatusOK))
		})

		It("should return correct JSON response", func() {
			req, _ := http.NewRequest("GET", "/info", nil)
			resp := httptest.NewRecorder()
			router.ServeHTTP(resp, req)

			expected := gin.H{
				"version":    "v1.0.0",
				"git_commit": "abc123",
				"build_id":   "build123",
				"go_version": runtime.Version(),
				"compiler":   runtime.Compiler,
				"platform":   runtime.GOOS + "/" + runtime.GOARCH,
			}
			jsonResponse, err := json.Marshal(expected)
			Expect(err).ToNot(HaveOccurred())
			Expect(resp.Code).To(Equal(http.StatusOK))
			Expect(resp.Body).To(MatchJSON(jsonResponse))
		})
	})
})
