package monster

import "testing"

/*
Test command:
1. Go to path /home/xmzhao/mygit/GoLang/practice/example/test/monster
2. Run go test -v. This will run all test files
  2.1 If want to test only one test file, then you can run go test -v monster_test.go monster.go
  2.2 If want to test only one function, then you can run
*/

func TestStore(t *testing.T) {
	m1 := &Monster{
		Name:  "WuKong",
		Age:   18,
		Skill: "Zhua Zi Attack",
	}
	res := m1.Store()
	if !res {
		t.Fatalf("Monster Store() failed, expectation=%v, actual=%v", true, res)
	}
	t.Logf("Monster Store() testing successful...")
}

func TestReStore(t *testing.T) {
	m1 := &Monster{}
	res := m1.Restore()
	if !res {
		t.Fatalf("Monster ReStore() failed, expectation=%v, actual=%v", true, res)
	}
	if m1.Name != "WuKong" {
		t.Fatalf("Monster ReStore() failed, expectation=%v, actual=%v", "WuKong", m1.Name)
	}
	t.Logf("Monster ReStore() testing successful...")
}
