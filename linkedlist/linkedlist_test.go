package linkedlist

import "testing"

func TestEmptyList(t *testing.T) {
	list := LinkedList{}
	got := list.Head
	var want *Node = nil

	if got != want {
		t.Errorf("got %v wanted nil", got)
	}
}

func TestListHead(t *testing.T) {
	test_string := "Hello World"

	list := LinkedList{}
	list.AppendNode(NewNode(test_string))

	got := list.Head.Value
	want := test_string

	if got != want {
		t.Errorf("got %s wanted %s", got, want)
	}
}

func TestListTail(t *testing.T) {
	test_string := "Hello World"

	list := LinkedList{}
	list.AppendNode(NewNode(test_string))

	got := list.Tail.Value
	want := test_string

	if got != want {
		t.Errorf("got %s wanted %s", got, want)
	}
}
