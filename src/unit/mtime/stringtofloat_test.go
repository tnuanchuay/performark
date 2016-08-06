package mtime

import "testing"

func TestItShouldReturn1WhenEnter1s(t *testing.T){
	input := "1s"
	expect := 1.0
	result, _ := StringToFloat(input)
	if result != expect{
		t.Error("expect",expect,"but return", result)
	}
}

func TestItShouldReturn60WhenEnter1m(t *testing.T){
	input := "1m"
	expect := 60.0
	result, _ := StringToFloat(input)
	if result != expect{
		t.Error("expect",expect,"but return", result)
	}
}

func TestItSholdReturn3600WhenEnter1h(t *testing.T){
	input := "1h"
	expect := 3600.0
	result, _ := StringToFloat(input)
	if result != expect{
		t.Error("expect",expect,"but return", result)
	}
}

func TestItSholdReturn61WhenEnter1dot02(t *testing.T){
	input := "1.02m"
	expect := 61.2
	result, _ := StringToFloat(input)
	if result != expect{
		t.Error("expect",expect,"but return", result)
	}
}

func TestItSholdReturn61WhenEnter463dot50us(t *testing.T){
	input := "463.50us"
	expect := 0.0004635
	result, _ := StringToFloat(input)
	if result != expect{
		t.Error("expect",expect,"but return", result)
	}
}

func TestItSholdReturn61WhenEnter1dot72s(t *testing.T){
	input := "1.72s"
	expect := 1.72
	result, _ := StringToFloat(input)
	if result != expect{
		t.Error("expect",expect,"but return", result)
	}
}