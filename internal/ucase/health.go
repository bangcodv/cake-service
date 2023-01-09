package ucase

import (
	"github.com/bangcodv/cake-service/internal/appctx"
	"github.com/bangcodv/cake-service/internal/consts"
	"github.com/bangcodv/cake-service/internal/ucase/contract"
)

type healthCheck struct {
}

func NewHealthCheck() contract.UseCase {
	return &healthCheck{}
}

func (u *healthCheck) Serve(*appctx.Data) appctx.Response {
	return *appctx.NewResponse().WithCode(consts.CodeSuccess).WithMessage("ok")
}
