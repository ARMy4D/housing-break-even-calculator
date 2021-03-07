package calculator

import (
	"github.com/army4d/housing-break-even-calculator/app"
)

type Errorer app.Errorer
type ErrorCode app.ErrorCode

const (
	errCodeNotImplemented = app.ErrorCode(ErrorRange) + iota
	errCodeClient
	errCodeTransformer
	errCodeConstraintReached
	errCodeParameter
	errCodeMinReside
	errCodeMinMortgageTerm
	errCodeMaxMortgageTerm
)

var (
	ErrNotImplemented = app.NewErrorTranslated(map[app.ErrorTranslation]string{
		app.ERR_EN: "not implemented error",
		app.ERR_NL: "niet geïmplementeerde fout",
	}, errCodeNotImplemented, ServiceName)
	ErrClient = app.NewErrorTranslated(map[app.ErrorTranslation]string{
		app.ERR_EN: "client error",
		app.ERR_NL: "client fout",
	}, errCodeClient, ServiceName)
	ErrConstraintReached = app.NewErrorTranslated(map[app.ErrorTranslation]string{
		app.ERR_EN: "constraint error",
		app.ERR_NL: "beperking fout",
	}, errCodeConstraintReached, ServiceName)
	ErrTransformer = app.NewErrorTranslated(map[app.ErrorTranslation]string{
		app.ERR_EN: "transformer error",
		app.ERR_NL: "transformatorfout",
	}, errCodeTransformer, ServiceName)
	ErrParameter = app.NewErrorTranslated(map[app.ErrorTranslation]string{
		app.ERR_EN: "wrong parameter error",
		app.ERR_NL: "verkeerde parameterfout",
	}, errCodeParameter, ServiceName)
	ErrMinReside = app.NewErrorTranslated(map[app.ErrorTranslation]string{
		app.ERR_EN: "minimum reside constraint reached",
		app.ERR_NL: "minimale residentiebeperking bereikt",
	}, errCodeMinReside, ServiceName)
	ErrMinMortgageTerm = app.NewErrorTranslated(map[app.ErrorTranslation]string{
		app.ERR_EN: "minimum mortgage term constraint reached",
		app.ERR_NL: "minimale hypotheektermijn bereikt",
	}, errCodeMinMortgageTerm, ServiceName)
	ErrMaxMortgageTerm = app.NewErrorTranslated(map[app.ErrorTranslation]string{
		app.ERR_EN: "maximum mortgage term constraint reached",
		app.ERR_NL: "maximale hypotheektermijn bereikt",
	}, errCodeMaxMortgageTerm, ServiceName)
)
