package pi

type Wrapper struct {
}

func NewPIWrapper() *Wrapper {
	return &Wrapper{}
}

func (w *Wrapper) ToggleScreen() error {
	// TODO: add impl
	return nil
}
