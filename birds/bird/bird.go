package bird

import "fmt"

type Bird interface {
	Flys() bool
	Name() string
	/*
		There's nothing wrong with providing getters and setters yourself, and it's often appropriate to do so, but it's
		neither idiomatic nor necessary to put Get into the getter's name. If you have a field called owner (lower case,
		unexported), the getter method should be called Owner (upper case, exported), not GetOwner.
	*/
	Swims() bool
}

func Info(b ...Bird){
	for _, bird  := range b {
		var swims string
		var flies string
		if bird.Flys() {
			flies = fmt.Sprintf("%v flys,", bird.Name())
		} else {
			flies = fmt.Sprintf("%v does not fly,", bird.Name())
		}

		if bird.Swims() {
			swims = fmt.Sprintf("and swims!!! %T \n", bird)
		} else {
			swims = fmt.Sprintf("and does not swim!!! %T \n", bird)
		}
		fmt.Printf("%v %v", flies, swims)
	}

}
