package c9errors

type ErrorInvalidID struct{}

func (e ErrorInvalidID) String() string {
	return "ID is invalid"
}
