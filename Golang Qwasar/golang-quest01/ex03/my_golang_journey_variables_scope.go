package main
import "fmt"

func main() {
    {
        var my_nested_variable = 30;
        fmt.Printf("30 -> %d\n", my_nested_variable);
        // end for my_nested_variable    
    }
    {
        var my_nested_variable = 1;

        fmt.Printf("1 -> %d\n", my_nested_variable);
        // end for my_nested_variable    
    }
    {
        var my_nested_variable   = 30;
        var my_nested_variable_2 = 1;

        fmt.Printf("30 && 1 -> %d - %d\n", my_nested_variable_2, my_nested_variable);
        // end for my_nested_variable and my_nested_variable_2
    }
    {
        var my_nested_variable   = 1_000;
        var my_nested_variable_2 = -1;

        fmt.Printf("1_000 && -1 -> %d - %d\n", my_nested_variable, my_nested_variable_2);
        // end for my_nested_variable and my_nested_variable_2
    }
}
