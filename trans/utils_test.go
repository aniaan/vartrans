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

func TestCamelCase3(t *testing.T) {
	origin := "you and me"
	expect := "youMe"
	acutual := camelCase(origin)
	if acutual != expect {
		t.Errorf("expect result: %s, but got acutal result %s", expect, acutual)
	}
}

func TestCamelCase4(t *testing.T) {
	origin := "you  "
	expect := "you"
	acutual := camelCase(origin)
	if acutual != expect {
		t.Errorf("expect result: %s, but got acutal result %s", expect, acutual)
	}
}


func TestBigCamelCase1(t *testing.T) {
	origin := "How are you"
	expect := "howAreYou"
	acutual := camelCase(origin)

	if acutual != expect {
		t.Errorf("expect result: %s, but got acutal result %s", expect, acutual)
	}
}

func TestBigCamelCase2(t *testing.T) {
	origin := "hello world"
	expect := "HelloWorld"
	acutual := bigCamelCase(origin)
	if acutual != expect {
		t.Errorf("expect result: %s, but got acutal result %s", expect, acutual)
	}
}

func TestBigCamelCase3(t *testing.T) {
	origin := "you and me"
	expect := "YouMe"
	acutual := bigCamelCase(origin)
	if acutual != expect {
		t.Errorf("expect result: %s, but got acutal result %s", expect, acutual)
	}
}

func TestBigCamelCase4(t *testing.T) {
	origin := "you  "
	expect := "You"
	acutual := bigCamelCase(origin)
	if acutual != expect {
		t.Errorf("expect result: %s, but got acutal result %s", expect, acutual)
	}
}

func TestConstant1(t *testing.T) {
	origin := "How are you"
	expect := "HOW_ARE_YOU"
	acutual := constant(origin)

	if acutual != expect {
		t.Errorf("expect result: %s, but got acutal result %s", expect, acutual)
	}
}

func TestConstant2(t *testing.T) {
	origin := "hello world"
	expect := "HELLO_WORLD"
	acutual := constant(origin)
	if acutual != expect {
		t.Errorf("expect result: %s, but got acutal result %s", expect, acutual)
	}
}

func TestConstant3(t *testing.T) {
	origin := "you and me"
	expect := "YOU_ME"
	acutual := constant(origin)
	if acutual != expect {
		t.Errorf("expect result: %s, but got acutal result %s", expect, acutual)
	}
}

func TestConstant4(t *testing.T) {
	origin := "you  "
	expect := "YOU"
	acutual := constant(origin)
	if acutual != expect {
		t.Errorf("expect result: %s, but got acutal result %s", expect, acutual)
	}
}

func TestSnakeCase1(t *testing.T) {
	origin := "How are you"
	expect := "how_are_you"
	acutual := snakeCase(origin)

	if acutual != expect {
		t.Errorf("expect result: %s, but got acutal result %s", expect, acutual)
	}
}

func TestSnakeCase2(t *testing.T) {
	origin := "hello world"
	expect := "hello_world"
	acutual := snakeCase(origin)
	if acutual != expect {
		t.Errorf("expect result: %s, but got acutal result %s", expect, acutual)
	}
}

func TestSnakeCase3(t *testing.T) {
	origin := "you and me"
	expect := "you_me"
	acutual := snakeCase(origin)
	if acutual != expect {
		t.Errorf("expect result: %s, but got acutal result %s", expect, acutual)
	}
}

func TestSnakeCase4(t *testing.T) {
	origin := "you  "
	expect := "you"
	acutual := snakeCase(origin)
	if acutual != expect {
		t.Errorf("expect result: %s, but got acutal result %s", expect, acutual)
	}
}

func TestHyphen1(t *testing.T) {
	origin := "How are you"
	expect := "how-are-you"
	acutual := hyphen(origin)

	if acutual != expect {
		t.Errorf("expect result: %s, but got acutal result %s", expect, acutual)
	}
}

func TestHyphen2(t *testing.T) {
	origin := "hello world"
	expect := "hello-world"
	acutual := hyphen(origin)
	if acutual != expect {
		t.Errorf("expect result: %s, but got acutal result %s", expect, acutual)
	}
}

func TestHyphen3(t *testing.T) {
	origin := "you and me"
	expect := "you-me"
	acutual := hyphen(origin)
	if acutual != expect {
		t.Errorf("expect result: %s, but got acutal result %s", expect, acutual)
	}
}

func TestHyphen4(t *testing.T) {
	origin := "you  "
	expect := "you"
	acutual := hyphen(origin)
	if acutual != expect {
		t.Errorf("expect result: %s, but got acutal result %s", expect, acutual)
	}
}