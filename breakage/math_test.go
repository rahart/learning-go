package breakage

import ( 
	"testing"
)

func TestSum(t *testing.T){
	result := Sum(1,2,3)

	if result != 6 {
 		t.Error("Failed the sum function.")
	}
}
