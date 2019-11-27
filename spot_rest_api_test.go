package gateio

import (
	"fmt"
	"testing"
)

func TestGetPairs(t *testing.T) {
	c := NewTestClient()
	ac := c.GetPairs()

	fmt.Printf("%+v", ac)

	jstr, _ := Struct2JsonString(ac)
	println(jstr)
}
