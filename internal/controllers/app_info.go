package controllers

import (
	"fmt"
	"net/http"
	"runtime"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type AppInfo struct {
	version   string
	gitCommit string
	buildID   string
	logger    *zap.Logger
}

func NewAppInfoHandler(version, gitCommit, buildID string, logger *zap.Logger) *AppInfo {
	return &AppInfo{
		version:   version,
		gitCommit: gitCommit,
		buildID:   buildID,
		logger:    logger,
	}
}

func (a *AppInfo) Handler(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"version":    a.version,
		"git_commit": a.gitCommit,
		"build_id":   a.buildID,
		"go_version": runtime.Version(),
		"compiler":   runtime.Compiler,
		"platform":   fmt.Sprintf("%s/%s", runtime.GOOS, runtime.GOARCH),
	})
}
