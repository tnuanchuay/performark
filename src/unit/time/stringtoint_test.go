package time

import "testing"

func TestItShouldReturn1WhenEnter1s(t *testing.T){
	input := "1s"
	expect := 1.0
	result, _ := StringToInt(input)
	if result != expect{
		t.Error("expect",expect,"but return", result)
	}
}

func TestItShouldReturn60WhenEnter1m(t *testing.T){
	input := "1m"
	expect := 60.0
	result, _ := StringToInt(input)
	if result != expect{
		t.Error("expect",expect,"but return", result)
	}
}

func TestItSholdReturn3600WhenEnter1h(t *testing.T){
	input := "1h"
	expect := 3600.0
	result, _ := StringToInt(input)
	if result != expect{
		t.Error("expect",expect,"but return", result)
	}
}

func TestItSholdReturn61WhenEnter1dot02(t *testing.T){
	input := "1.02m"
	expect := 61.0
	result, _ := StringToInt(input)
	if result != expect{
		t.Error("expect",expect,"but return", result)
	}
}