package err

import "errors"

var (
	StoveErrorLength = errors.New("10001") //errors.New("Length cannot be null")
	StoveErrorWidth  = errors.New("10002") //errors.New("Width cannot be null")
	StoveErrorHeight = errors.New("10003") //errors.New("Height cannot be null")
	StoveErrorNumber = errors.New("10004") //errors.New("Number cannot be null")
)
