package bird

type Duck struct {
	name string
}

func (b Duck) Flys() bool {
	return true
}

func (b Duck) Swims() bool {
	return true
}

func (b Duck) GetName() string {
	return b.name
}
