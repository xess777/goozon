package spaces

import "testing"

func getCases() map[string]string {
	return map[string]string{
		" On  my   home world":   "On my home world",
		"   On my   home world ": "On my home world",
		"":                       "",
		" ":                      "",
		"    ":                   "",
		" a ":                    "a",
	}
}

func TestRemoveExtraSpaces(t *testing.T) {
	for original, expected := range getCases() {
		s := []rune(original)
		actual := string(*RemoveExtraSpaces(&s))
		if actual != expected {
			t.Errorf("\nActual:%s\nExpect:%s", actual, expected)
		}
	}
}
