package main

import (
	"fmt"
	"testing"
)

func TestRain(t *testing.T) {
	rs := NewRainService(config.RainService)
	res, err := rs.GetRow("")
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(res)
}
