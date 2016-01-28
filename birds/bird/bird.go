package bird

import "fmt"

type Bird interface {
	Flys() bool
	GetName() string
	Swims() bool
}

func Info(b ...Bird){
	for _, bird  := range b {
		var swims string
		var flies string
		if bird.Flys() {
			flies = fmt.Sprintf("%v flys,", bird.GetName())
		} else {
			flies = fmt.Sprintf("%v does not fly,", bird.GetName())
		}

		if bird.Swims() {
			swims = fmt.Sprintf("and swims!!! %T \n", bird)
		} else {
			swims = fmt.Sprintf("and does not swim!!! %T \n", bird)
		}
		fmt.Printf("%v %v", flies, swims)
	}

}
