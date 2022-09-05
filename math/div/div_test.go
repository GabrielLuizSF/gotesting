package div

import "testing"

func TestDivision(test *testing.T){

  division := Division(100,2)
  expected := 50.00

	if division != expected{
		test.Errorf("\n[ERROR]\n\nResult:%f\nExpected:%f",division,expected)
	}
	

}