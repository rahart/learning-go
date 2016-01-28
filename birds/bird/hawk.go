package bird

type Hawk struct  {
	Name string
}

func (b Hawk) Flys() bool {
	return true
}

func (b Hawk) Swims() bool {
	return false
}

func (b Hawk) GetName() string {
 	return b.Name
}

