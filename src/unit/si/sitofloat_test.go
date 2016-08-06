package si

import "testing"

func TestItShouldReturn1000WhenEnter1k(t *testing.T){
	input  := "1k"
	expected := 1000.0
	result, _ := SIToFloat(input)
	if expected != result{
		t.Error("expect", expected, "but return", result)
	}
}

func TestItShouldReturn2000WhenEnter2k(t *testing.T){
	input  := "2k"
	expected := 2000.0
	result, _ := SIToFloat(input)
	if expected != result{
		t.Error("expect", expected, "but return", result)
	}
}

func TestItShouldReturn2200WhenEnter2dot2k(t *testing.T){
	input  := "2.2k"
	expected := 2200.0
	result, _ := SIToFloat(input)
	if expected != result{
		t.Error("expect", expected, "but return", result)
	}
}

func TestItShouldReturn2200000WhenEnter2dot2M(t *testing.T){
	input  := "2.2M"
	expected := 2200000.0
	result, _ := SIToFloat(input)
	if expected != result{
		t.Error("expect", expected, "but return", result)
	}
}

func TestItShouldReturn2WhenEnter2(t *testing.T){
	input  := "2"
	expected := 2.0
	result, _ := SIToFloat(input)
	if expected != result{
		t.Error("expect", expected, "but return", result)
	}
}

func TestItShouldReturn2200000WhenEnter2dot2MB(t *testing.T){
	input  := "2.2MB"
	expected := 2200000.0
	result, _ := SIToFloat(input)
	if expected != result{
		t.Error("expect", expected, "but return", result)
	}
}