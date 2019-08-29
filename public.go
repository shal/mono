package mono

// Public gives access to public methods.
type Public struct {
	core
}

// NewPublic returns new client of MonoBank Public API.
func NewPublic() *Public {
	return &Public{
		core: *newCore(),
	}
}
