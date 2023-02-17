package main

import "fmt"

func my_swap(a *int, b *int) {
	// Temporary variable to store the value of a
	temp := *a

	// Swap the values
	*a = *b
	*b = temp
}
