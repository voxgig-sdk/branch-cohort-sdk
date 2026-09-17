package core

type BranchCohortError struct {
	IsBranchCohortError bool
	Sdk              string
	Code             string
	Msg              string
	Ctx              *Context
	Result           any
	Spec             any
}

func NewBranchCohortError(code string, msg string, ctx *Context) *BranchCohortError {
	return &BranchCohortError{
		IsBranchCohortError: true,
		Sdk:              "BranchCohort",
		Code:             code,
		Msg:              msg,
		Ctx:              ctx,
	}
}

func (e *BranchCohortError) Error() string {
	return e.Msg
}
