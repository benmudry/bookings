package forms

import "testing"

func TestErrors_Add(t *testing.T) {
	e := errors(map[string][]string{})

	a, ok := e["a"]
	if ok {
		t.Error("field shouldn't exist")
	}

	e = errors(map[string][]string{})
	e.Add("a", "a")

	a, ok = e["a"]
	if !ok {
		t.Error("field doesn't exist")
	}
	if a[0] != "a" {
		t.Error("field doesn't equal what it's supposed to")
	}

}

func TestErrors_Get(t *testing.T) {
	e := errors(map[string][]string{})
	errorMsg := e.Get("a")

	if errorMsg != "" {
		t.Error("error message exists when it shouldn't")
	}

	e = errors(map[string][]string{})
	e.Add("a", "error")
	errorMsg = e.Get("a")

	if errorMsg == "" {
		t.Error("error message doesn't  when it shouldn't")
	}
}
