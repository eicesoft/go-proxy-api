package core

import "testing"

type AStruct struct {
	Name  string
	Code  int
	Title string
	Age   int
}

type BStruct struct {
	Name  string
	Code  int
	Title string
}

func TestStructCopy(t *testing.T) {
	a := BStruct{
		"bill gate",
		200,
		"this is title",
	}

	c := AStruct{}
	StructCopy(&c, &a)
	if c.Name == a.Name && c.Code == a.Code && c.Title == a.Title {
		t.Log("StructCopy is success")
	} else {
		t.Fatal("StructCopy is error")
	}
}
