package main

import (
    "fmt"
)

// Account interface defines the component with Name() and IsActive() functions.
// TODO: 1. Implement 'Account'.
type Account interface {
	Name() string
	IsActive() bool
}

func Name() string {
    return "a"
}

func IsActive() bool {
    return true
}

// GetAllAccounts returns 3 items, as follows:
// | Name | Active |
// | Foo  | Yes    |
// | Bar  | No     |
// | Baz  | Yes    |
func GetAllAccounts() []Account {
	// TODO: 2. Implement GetAllAccounts method.
	// return slice with values which satisfy the description above
    r := {Account.Name("Foo"), Account.IsActive("false")}
    var ar []Account
    ar = append(ar, r)
    return ar 
}

// GetAllAccounts should be called in main function and returned slice iterated over to produce following output:
// Foo is active
// Bar is inactive
// Baz is active
func main() {
	// TODO: 3. Call GetAllAccounts() and produce the output from description above.
    for idx, el := range GetAllAccounts() {
        fmt.Println(idx, el)
    }
}

