package commands

import (
	"flag"
	"reflect"
	"testing"
)

// TestBuild is a test for Build function
// Test scenarious:
//
//	"simlpe": building a flag set with one flag
func TestBuild(t *testing.T) {
	type test struct {
		input FSet
		want  flag.FlagSet
	}

	fs := flag.NewFlagSet("test", flag.ContinueOnError)
	fs.Bool("t1", false, "t1")

	tests := map[string]test{
		"simlpe": {FSet{"test", []*Flag{NewFlag("t1", "t1")}}, *fs},
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			got := tc.input.Build()
			flag := got.Lookup("t1")
			wflag := tc.want.Lookup("t1")
			if flag == nil || flag.Name != wflag.Name {
				t.Fatalf("%s: expected: %v, got: %v", name, tc.want, got)
			}
		})
	}
}

// TestCheckActive is a test for CheckActive function.
// Check if there is any flag that was passed with the command
// Test scenarious:
//
//	"simple": test on flag set with one flag that was passed
//	"empty flags": test on flag set with no flags
//	"nil flags": test on nil flag set
func TestCheckActive(t *testing.T) {

	type test struct {
		input FSet
		want  string
	}

	fSets := []FSet{
		{name: "test", flags: []*Flag{{"t1", true, "t1"}}},
		{name: "empty", flags: []*Flag{}},
		{name: "nil", flags: nil},
	}

	tests := map[string]test{
		"simple":      {input: FSet{"test", []*Flag{{"t1", true, "t1"}}}, want: "t1"},
		"empty flags": {input: fSets[1], want: ""},
		"nil flags":   {input: fSets[2], want: ""},
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			got := tc.input.CheckActive()
			if got != tc.want {
				t.Fatalf("%s: expected: \"t2\", got: %v", name, got)
			}
		})
	}
}

// TestBind is a test for Bind fuction
// Test scenarious:
//
//	"simple": adding a flag to an empty flagset
//	"empty flag": adding empty flag to an empty flag set
func TestBind(t *testing.T) {
	type test struct {
		flag  Flag
		input *flag.FlagSet
		want  bool
	}

	tests := map[string]test{
		"simple":     {flag: *NewFlag("t1", "t1"), input: &flag.FlagSet{}, want: false},
		"empty flag": {flag: Flag{}, input: &flag.FlagSet{}, want: false},
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			got := tc.flag.Bind(tc.input)
			if *got != tc.want {
				t.Fatalf("expected: %v, got: %v", tc.want, *got)
			}
		})
	}
}

// TestNewFlag is a test for NewFlag function
// Test scenarious:
//
//	"simple": creating a flag with passed arguments
//	"empty args": creating a flag with empty arguments
func TestNewFlag(t *testing.T) {
	type arg struct {
		name  string
		usage string
	}

	type test struct {
		input arg
		want  *Flag
	}

	tests := map[string]test{
		"simple":     {input: arg{"t1", "t1"}, want: &Flag{"t1", false, "t1"}},
		"empty args": {input: arg{"", ""}, want: &Flag{"", false, ""}},
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			got := NewFlag(tc.input.name, tc.input.usage)
			if !reflect.DeepEqual(tc.want, got) {
				t.Fatalf("expected: %v, got: %v", tc.want, got)
			}
		})
	}
}
