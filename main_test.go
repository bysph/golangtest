package main

import "testing"

func Test_1(t *testing.T) {
	t.Error("fail")
}

func Test_2(t *testing.T) {
	t.Log("pass")
}
