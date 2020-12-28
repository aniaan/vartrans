package trans

import "testing"

func TestCamelCase1(t *testing.T) {
	origin := "How are you"
	expect := "howAreYou"
	acutual := camelCase(origin)

	if acutual != expect {
		t.Errorf("expect result: %s, but got acutal result %s", expect, acutual)
	}
}

func TestCamelCase2(t *testing.T) {
	origin := "hello world"
	expect := "helloWorld"
	acutual := camelCase(origin)
	if acutual != expect {
		t.Errorf("expect result: %s, but got acutal result %s", expect, acutual)
	}
}
