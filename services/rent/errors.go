package rent

import (
	"github.com/army4d/housing-break-even-calculator/app"
)

type Errorer app.Errorer
type ErrorCode app.ErrorCode

const (
	errCodeNotImplemented = app.ErrorCode(ErrorRange) + iota
	errCodeCaculation
	errCodeUnsupportedValue
)

var (
	ErrNotImplemented = app.NewErrorTranslated(map[app.ErrorTranslation]string{
		app.ERR_EN: "not implemented error",
		app.ERR_NL: "niet geïmplementeerde fout",
	}, errCodeNotImplemented, ServiceName)
	ErrCalculation = app.NewErrorTranslated(map[app.ErrorTranslation]string{
		app.ERR_EN: "calculation error",
		app.ERR_NL: "rekenfout",
	}, errCodeCaculation, ServiceName)
	ErrUnsupportedValue = app.NewErrorTranslated(map[app.ErrorTranslation]string{
		app.ERR_EN: "unsupported value error",
		app.ERR_NL: "niet-ondersteunde waardefout",
	}, errCodeUnsupportedValue, ServiceName)
)
