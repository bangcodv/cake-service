// Package middleware
package middleware

import (
	"net/http"
	"strings"
	"fmt"

	"github.com/bangcodv/cake-service/internal/appctx"
	"github.com/bangcodv/cake-service/internal/consts"
	"github.com/bangcodv/cake-service/pkg/logger"
)

// ValidateContentType header
func ValidateContentType(r *http.Request, conf *appctx.Config) int {

	if ct := strings.ToLower(r.Header.Get(`Content-Type`)) ; ct != `application/json` {
		logger.Warn(fmt.Sprintf("[middleware] invalid content-type %s", ct ))

		return consts.CodeBadRequest
	}


	return consts.CodeSuccess
}