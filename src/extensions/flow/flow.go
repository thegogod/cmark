package flow

type Flow struct{}

func New() *Flow {
	return &Flow{}
}

func (self Flow) Name() string {
	return "flow"
}
