package bird

type Hawk struct  {
	name string
}

func (b Hawk) Flys() bool {
	return true
}

func (b Hawk) Swims() bool {
	return false
}

func (b Hawk) GetName() string {
 	return b.name
}

