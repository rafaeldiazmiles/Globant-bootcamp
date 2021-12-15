package errors

type error interface {
	Error() string
}

type errorType struct {
	etype string
}

func (e *errorType) Error() string {
	return e.etype
}

func new(s string) error {
	return &errorType{s}
}

func NewInternal() error {
	return new("Internal")
}

func NewThirdParty() error {
	return new("ThirdParty")
}

func NewOther() error {
	return new("Other")
}

func ErrorChecker(e error, s string) bool {
	if e.Error() == s {
		return true
	}
	return false
}
