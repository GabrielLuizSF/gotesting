package main

import (
	"testing"
)

func TestHello(test *testing.T){ 
  checkExpectedResult := func(test *testing.T,result,expected string){
	test.Helper();
	if result != expected{
		test.Errorf("\nResult:%s\nExpected:%s",result,expected);
	}
  }

  test.Run("Says Hello To a Person",func(test *testing.T){
	result := Hello("Yourname")
	expected := "Hello Yourname"
	checkExpectedResult(test,result,expected)
  })
  

}