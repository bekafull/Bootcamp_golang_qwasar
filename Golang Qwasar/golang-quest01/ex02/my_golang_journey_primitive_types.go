package main
import "fmt"

func main() {
    uinteger_value_400 := 400
    integer_value_14 := 14
    uinteger_value_guess := 0x406
    true_boolean := true
    false_boolean := false
    float_variable := 3.14
    string_variable := "Scotland has 421 words for 'snow'"
    char_variable := 'z'
    array_variable := [5]int{0, 1, 2, 3, 4}

    fmt.Printf("Unsigned integer %d\n", uinteger_value_400 + 2)
    fmt.Printf("Signed integer = %d\n", integer_value_14 - 1000)

    fmt.Printf("Hex integer = %x\n", uinteger_value_guess + 0)

    fmt.Printf("Boolean #1 %t\n", true_boolean)
    fmt.Printf("Boolean #2 %t\n", false_boolean)

    fmt.Printf("Float -> %f\n", float_variable)
    fmt.Printf("String -> %s\n", string_variable)
    fmt.Printf("Character -> %c\n", char_variable)
    fmt.Printf("Array -> %v\n", array_variable)
}
