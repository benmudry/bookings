package forms

import (
	"net/url"
	"testing"
)

func TestForm_Valid(t *testing.T) {
	form := New(url.Values{})

	isValid := form.Valid()
	if !isValid {
		t.Error("got invalid when should have been valid")
	}
}

func TestForm_Required(t *testing.T) {
	form := New(url.Values{})

	form.Required("a", "b", "c")
	if form.Valid() {
		t.Error("form shows valid when required fields are missing")
	}

	postedData := url.Values{}
	postedData.Add("a", "a")
	postedData.Add("b", "a")
	postedData.Add("c", "a")
	form = New(postedData)

	form.Required("a", "b", "c")
	if !form.Valid() {
		t.Error("shows does not have required fields when it does")
	}
}

func TestForm_Has(t *testing.T) {
	form := New(url.Values{})

	if form.Has("a") {
		t.Error("form shows has when field is missing")
	}

	postedData := url.Values{}
	postedData.Add("a", "a")
	postedData.Add("b", "a")
	form = New(postedData)

	if !form.Has("a") {
		t.Error("form does not have field when it does")
	}
}

func TestForm_MinLength(t *testing.T) {
	form := New(url.Values{})

	minLength := form.MinLength("a", 2)
	if minLength {
		t.Error("form minlength passes when value is nil")
	}
	if form.Valid() {
		t.Error("form shows valid when field is too short")
	}

	postedData := url.Values{}
	postedData.Add("a", "abc")
	postedData.Add("b", "abc")
	form = New(postedData)

	minLength = form.MinLength("a", 200)
	if minLength {
		t.Errorf("form minlength of %s is %d, passes when value is less than 200 length", form.Get("a"), len(form.Get("a")))
	}
	if form.Valid() {
		t.Error(form.Errors.Get("a"))
		t.Error("form shows valid when field passes less than 200 length")
	}

	postedData = url.Values{}
	postedData.Add("a", "abc")
	postedData.Add("b", "abc")
	form = New(postedData)

	minLength = form.MinLength("a", 2)
	if !minLength {
		t.Error("form minlength fails when value is longer than minlength")
	}
	if !form.Valid() {
		t.Error(form.Errors.Get("a"))
		t.Error("form invalid when field passes minlength")
	}
}

func TestForm_IsEmail(t *testing.T) {
	postedData := url.Values{}
	form := New(postedData)

	form.IsEmail("a")
	if form.Valid() {
		t.Error("form shows nil is email")
	}

	postedData = url.Values{}
	postedData.Add("a", "a")
	form = New(postedData)

	form.IsEmail("a")
	if form.Valid() {
		t.Error("form shows invalid email is email")
	}

	postedData = url.Values{}
	postedData.Add("a", "a@a.com")
	form = New(postedData)

	form.IsEmail("a")
	if !form.Valid() {
		t.Error("form shows valid email as invalid")
	}
}
