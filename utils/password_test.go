package utils

import "testing"

func TestPassword(t *testing.T) {
	var testCase = []struct {
		plainPassword string
	}{
		{""},
		{"1"},
		{"123"},
		{"a"},
		{"abc"},
		{"abc_&@*($&)"},
		{"~p2J0SAH@$\":AJALSD"},
		{"somepassword"},
		{"somepasswordsomepasswordsomepasswordsomepasswordsomepasswordsomepasswordsomepasswordsomepasswordsomepasswordsomepassword"}, // length: 120
		{"somepasswordsomepasswordsomepasswordsomepasswordsomepasswordsomepasswordsomepasswordsomepasswordsomepasswordsomepasswordsomepasswordsomepasswordsomepasswordsomepasswordsomepasswordsomepasswordsomepasswordsomepasswordsomepasswordsomepasswordsomepasswordsomepasswordsomepasswordsomepasswordsomepasswordsomepasswordsomepasswordsomepasswordsomepasswordsomepasswordsomepasswordsomepasswordsomepasswordsomepasswordsomepasswordsomepasswordsomepasswordsomepasswordsomepasswordsomepasswordsomepasswordsomepasswordsomepasswordsomepasswordsomepasswordsomepasswordsomepasswordsomepasswordsomepasswordsomepasswordsomepasswordsomepasswordsomepasswordsomepasswordsomepasswordsomepasswordsomepasswordsomepasswordsomepasswordsomepasswordsomepasswordsomepasswordsomepasswordsomepasswordsomepasswordsomepasswordsomepasswordsomepasswordsomepasswordsomepasswordsomepasswordsomepasswordsomepasswordsomepasswordsomepasswordsomepasswordsomepasswordsomepasswordsomepasswordsomepasswordsomepasswordsomepasswordsomepasswordsomepasswordsomepasswordsomepasswordsomepasswordsomepasswordsomepasswordsomepasswordsomepasswordsomepasswordsomepasswordsomepasswordsomepasswordsomepasswordsomepasswordsomepasswordsomepasswordsomepassword"}, // length: 1200
	}

	for _, tC := range testCase {
		hashedPassword := Hash(tC.plainPassword)
		if len(hashedPassword) != 120 {
			t.Error("hashed password's length has change!!!")
		}
		if !IsCorrect(hashedPassword, tC.plainPassword) {
			t.Error("password hashed and validate error")
		}
	}
}
