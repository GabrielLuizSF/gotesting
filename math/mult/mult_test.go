package mult

import "testing"

func TestMult(test *testing.T){

	multiplication := Muliplication(400,5)
	expected := 2000.00

	if multiplication != expected{
		test.Errorf("\n[ERROR]\n\nResult:%f\nExpected:%f",multiplication,expected)
	}
}