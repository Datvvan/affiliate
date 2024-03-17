package util

type CustomErr struct {
	Code int
	Msg  string
}

func (c CustomErr) Error() string {
	return c.Msg
}

func NewCustomErr(code int, msg string) CustomErr {
	return CustomErr{
		code,
		msg,
	}
}
