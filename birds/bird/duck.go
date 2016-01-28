package bird

type Duck struct {
	Name string
}

func (b Duck) Flys() bool {
	return true
}

func (b Duck) Swims() bool {
	return true
}

func (b Duck) GetName() string {
	return b.Name
}
