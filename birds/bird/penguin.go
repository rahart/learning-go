package bird

type Penguin struct  {
	name string
}

func (b Penguin) Flys() bool {
	return false
}

func (b Penguin) Swims() bool {
	return true
}

func (b Penguin) GetName() string {
 	return b.name
}
