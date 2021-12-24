package expand

import (
	"testing"
)

func TestExpand(t *testing.T) {
	expander := NewExpander()
	for _, test := range expandTests {
		v, e := expander.Expand(test.Data)
		if e != nil && !test.Errors {
			t.Fatalf("Errored on %s, err: %v", test.Data, e)
		}

		if len(v) != len(test.Result) {
			t.Fatalf("%v (%d) != %v (%d)", v, len(v), test.Result, len(test.Result))
		}
		for i, s := range test.Result {
			if v[i] != s {
				t.Fatalf("%v != %v", v, test.Result)
			}
		}
	}
}

func BenchmarkExpand(b *testing.B) {
	expander := NewExpander()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		for _, test := range expandTests {
			expander.Expand(test.Data)
		}
	}
	b.StopTimer()
}

type expandTest struct {
	Data   string
	Result []string
	Errors bool
}

var expandTests = []expandTest{
	expandTest{"foo", []string{"foo"}, false},
	expandTest{"[a]", []string{"a"}, false},
	expandTest{"[a,b,c]", []string{"a", "b", "c"}, false},
	expandTest{"[a,b,c]1", []string{"a1", "b1", "c1"}, false},
	expandTest{"[a,b,c]-1", []string{"a-1", "b-1", "c-1"}, false},
	expandTest{"[1]", []string{"1"}, false},
	expandTest{"[1-2]", []string{"1", "2"}, false},
	expandTest{"[1-3]", []string{"1", "2", "3"}, false},
	expandTest{"[a1]", []string{"a1"}, false},
	expandTest{"foo-[a]", []string{"foo-a"}, false},
	expandTest{"foo-[1]", []string{"foo-1"}, false},
	expandTest{"foo[1]", []string{"foo1"}, false},
	expandTest{"foo[1,2]", []string{"foo1", "foo2"}, false},
	expandTest{"foo[1-2]", []string{"foo1", "foo2"}, false},
	expandTest{"foo[1-3]", []string{"foo1", "foo2", "foo3"}, false},
	expandTest{"foo[1-2,3]", []string{"foo1", "foo2", "foo3"}, false},
	expandTest{"foo[1-2,3]", []string{"foo1", "foo2", "foo3"}, false},
	expandTest{"foo[a,1-2,4,d]", []string{"fooa", "foo1", "foo2", "foo4", "food"}, false},

	expandTest{"[a1-2]", []string{"a1-2"}, false},
	expandTest{"[a,b,c]-1-[d,e]", []string{"a-1-d", "b-1-d", "c-1-d", "a-1-e", "b-1-e", "c-1-e"}, false},
	expandTest{"[1-2][3-4]", []string{"13", "23", "14", "24"}, false},

	expandTest{"[1-2]]", []string{}, true},
	expandTest{"[[1-2]]", []string{}, true},
	expandTest{"[[1-2]", []string{}, true},
	expandTest{",", []string{}, true},
}
