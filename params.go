package checkers

func DefaultParams() Params {
	return Params{}
}

// validate does the sanity check on the params

func (p Params) Validate() error {
	return nil
}
