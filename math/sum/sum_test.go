package sum

import (
	"testing"
)


func TestSum(test *testing.T){
  sum := Addition(50,50);
  expected := 100;

  if sum != expected{
	test.Errorf("\n[ERROR]\n\nResult:%d Expected:%d",sum,expected)
  }

}

