package bird

type Penguin struct  {
	Name string
}

func (b Penguin) Flys() bool {
	return false
}

func (b Penguin) Swims() bool {
	return true
}

func (b Penguin) GetName() string {
 	return b.Name
}
