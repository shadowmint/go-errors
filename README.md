# Errors

A common error wrapper.

## Usage

    import "ntoolkit/errors"

    const (
    	ErrCode1 int = iota
    	ErrCode2
      ...
    )

    func dummy(value int) (int, error) {
    	if value == 0 {
    		return 0, errors.Fail(ErrCode1, nil, "Invalid value: %d", value)
    	}
    	return value + 1, nil
    }
