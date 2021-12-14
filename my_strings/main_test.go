package my_strings

import (
	"fmt"
	"reflect"
	"testing"
)

func TestCompare(t *testing.T) {
	type args struct {
		a string
		b string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "a, b",
			args: args{"a", "b"},
			want: -1,
		},
		{
			name: "a, a",
			args: args{"a", "a"},
			want: 0,
		},
		{
			name: "b, a",
			args: args{"b", "a"},
			want: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Compare(tt.args.a, tt.args.b); got != tt.want {
				t.Errorf("Compare() = %v, want %v", got, tt.want)
			}
		})
	}
}

func ExampleCompare() {
	fmt.Println(Compare("a", "b"))
	fmt.Println(Compare("a", "a"))
	fmt.Println(Compare("b", "a"))
	//	output: -1, 0, 1
}

func TestContains(t *testing.T) {
	type args struct {
		s      string
		substr string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "contains found",
			args: args{"seafood", "foo"},
			want: true,
		},
		{
			name: "equal",
			args: args{"seafood", "seafood"},
			want: true,
		},
		{
			name: "not found",
			args: args{"seafood", "bar"},
			want: false,
		},
		{
			name: "args nothing",
			args: args{"seafood", ""},
			want: false,
		},
		{
			name: "substr is longer",
			args: args{"seafood", "seafoods"},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Contains(tt.args.s, tt.args.substr); got != tt.want {
				t.Errorf("Contains() = %v, want %v", got, tt.want)
			}
		})
	}
}

func ExampleContains() {
	fmt.Println(Contains("seafood", "foo"))
	fmt.Println(Contains("seafood", "bar"))
	fmt.Println(Contains("seafood", ""))
	fmt.Println(Contains("", ""))
	// Output:
	// true
	// false
	// true
	// true
}

func TestCount(t *testing.T) {
	type args struct {
		s      string
		substr string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "cheese, e",
			args: args{"cheese", "e"},
			want: 3,
		}, // TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Count(tt.args.s, tt.args.substr); got != tt.want {
				t.Errorf("Count() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFields(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		{
			name: "1",
			args: args{"  foo bar  baz   "},
			want: []string{"foo", "bar", "baz"},
		},
		{
			name: "2",
			args: args{"   "},
			want: []string{""},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Fields(tt.args.s); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Fields() = %v, want %v", got, tt.want)
			}
		})
	}
}
