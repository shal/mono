package mono

type Public struct {
	core
}

func NewPublic() *Public {
	return &Public{
		core: *newCore(),
	}
}
