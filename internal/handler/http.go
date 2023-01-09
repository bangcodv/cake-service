// Package handler
package handler

import (
	"net/http"
	"context"

	"github.com/bangcodv/cake-service/internal/appctx"
	"github.com/bangcodv/cake-service/internal/consts"
	"github.com/bangcodv/cake-service/internal/ucase/contract"
)


// HttpRequest handler func wrapper
func HttpRequest(request *http.Request, svc contract.UseCase, conf *appctx.Config) appctx.Response {
	ctx := context.WithValue(request.Context(), consts.CtxLang, request.Header.Get(consts.HeaderLanguageKey))

	req := request.WithContext(ctx)

	data := &appctx.Data{
		Request:     req,
		Config:      conf,
		ServiceType: consts.ServiceTypeHTTP,
	}

	return svc.Serve(data)
}
