package commands

import "flag"

// FSet is a structure that stores all data needed for flag.FlagSet
// FSet has two field: name, flags. They represent flag set name
// and an array of flags needed to create a flagset
type FSet struct {
	name  string
	flags []*Flag
}

// Build is the method that creates a flagset
// out of FSet name and array of dedicated flags
func (fs *FSet) Build() *flag.FlagSet {
	nfs := flag.NewFlagSet(fs.name, flag.ContinueOnError)
	for _, f := range fs.flags {
		val := f.Bind(nfs)
		f.value = val
	}
	return nfs
}

// CheckActive is the method that checks what flags were
// passed as part of a command
func (fs *FSet) CheckActive() (name string) {
	for _, f := range fs.flags {
		if *f.value {
			return f.name
		}
	}
	return
}

// Flag represents an actual command flag
type Flag struct {
	name  string
	value *bool
	usage string
}

// Bind method binds a flag with a flagset with respect
// to the flag's name, value and usage field
func (f *Flag) Bind(fs *flag.FlagSet) *bool {
	return fs.Bool(f.name, *f.value, f.usage)
}

func NewFlag(name string, value bool, usage string) *Flag {
	return &Flag{name, &value, usage}
}

func BuildFlagSet(name string, flags []*Flag) (flagset *flag.FlagSet, wrapper *FSet) {
	fs := FSet{name, flags}
	return fs.Build(), &fs
}
