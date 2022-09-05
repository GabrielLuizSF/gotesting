package sub

import "testing"

func TestSubtraction(test *testing.T){
	subtraction := Subtraction(50,5)
	expected := 45.00

	if subtraction != expected{
		test.Errorf("\n[ERROR]\nResult:%f\nExpected:%f",subtraction,expected)
	}
}