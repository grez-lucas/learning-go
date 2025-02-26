package c6

import "fmt"

// UpdateSlice updates the last value of a string slice to "s"
func UpdateSlice(slice []string, s string) {
	slice[len(slice)-1] = s
	fmt.Printf("UpdateSlice: slice value inside function: %v\n", slice)
}

// GrowSlice attempts to grow a slice by appending a string "s" onto the slice
// It fails to do so because len and capacity values are not modified in GoLang's
// implementation of slice attributes being passed to functions.
func GrowSlice(slice []string, s string) {
	slice = append(slice, s)
	fmt.Printf("GrowSlice: slice value inside function: %v\n", slice)
}
