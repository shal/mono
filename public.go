package mono

import "context"

// Public gives access to public methods.
type Public struct {
	core
}

// NewPublic returns new client of MonoBank Public API.
func NewPublic() *Public {
	return &Public{
		core: *newCore(nil),
	}
}

func (p *Public) WithContext(context context.Context) Public {
	newP := *p
	newP.context = context
	return newP
}
