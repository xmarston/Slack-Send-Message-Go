package slack

import (
	"testing"
	"time"
)

func TestInit(t *testing.T) {
	slackInitialization := Init("./config")
	var expectedResult error = nil
	if slackInitialization != expectedResult {
		t.Fatalf("Expected %s but got %s", expectedResult, slackInitialization)
	}
}

func TestColorDefaults(t *testing.T) {
	f, _ := MakeField("", "")
	a, _ := MakeAttachment("#ff000", "", "", []Field{f}, time.Now().Unix())
	var expectedResult = "#ff000"

	if a.Color != expectedResult {
		t.Fatalf("Expected %s but got %s", expectedResult, a.Color)
	}
}

func TestTitleDefaults(t *testing.T) {
	f, _ := MakeField("", "")
	a, _ := MakeAttachment("", "Title", "", []Field{f}, time.Now().Unix())
	var expectedResult = "Title"

	if a.Title != expectedResult {
		t.Fatalf("Expected %s but got %s", expectedResult, a.Title)
	}
}

func TestTextDefaults(t *testing.T) {
	f, _ := MakeField("", "")
	a, _ := MakeAttachment("", "", "Text", []Field{f}, time.Now().Unix())
	var expectedResult = "Text"

	if a.Text != expectedResult {
		t.Fatalf("Expected %s but got %s", expectedResult, a.Title)
	}
}

func TestFieldsDefaults(t *testing.T) {
	f, _ := MakeField("", "")
	a, _ := MakeAttachment("", "", "", []Field{f}, time.Now().Unix())
	var expectedResult = 1

	if len(a.Fields) != 1 {
		t.Fatalf("Expected %d but got %d", expectedResult, len(a.Fields))
	}
}

func TestFooterDefaults(t *testing.T) {
	f, _ := MakeField("", "")
	a, _ := MakeAttachment("", "", "", []Field{f}, time.Now().Unix())
	var expectedResult = ""

	if a.Footer != expectedResult {
		t.Fatalf("Expected %s but got %s", expectedResult, a.Footer)
	}
}

func TestTsDefaults(t *testing.T) {
	f, _ := MakeField("", "")
	a, _ := MakeAttachment("", "", "", []Field{f}, time.Now().Unix())
	var expectedResult = time.Now().Unix()

	if a.Ts != expectedResult {
		t.Fatalf("Expected %s but got %s", expectedResult, a.Ts)
	}
}
