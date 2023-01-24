# Exercise 1

## Question
The simple calculator program doesn't handle one error case: division by zero. Change the function signature for the 
math operations to return both an `int` and an `error`. In the `div` function, if the divisor is `0`, 
return `0, errors.New("division by zero")`. In all other cases, return the calculated value and `nil`. Adjust 
the `main` function to check for this error.

## Solution
Change the math operation functions to:

```go
func add(i int, j int) (int, error) { return i + j, nil }

func sub(i int, j int) (int, error) { return i - j, nil }

func mul(i int, j int) (int, error) { return i * j, nil }

func div(i int, j int) (int, error) {
    if j == 0 {
        return 0, errors.New("division by zero")
    }
    return i / j, nil
}
```

Change the `opMap` declaration to:

```go
var opMap = map[string]func(int, int) (int, error){
	"+": add,
	"-": sub,
	"*": mul,
	"/": div,
}
```

Finally, you need to have a variable for each return value when `opFunc` is called:

```go
		result, err := opFunc(p1, p2)
		if err != nil {
			fmt.Println(err)
			continue
		}
        fmt.Println(result)
```

It's a good practice to add a test case to make sure your changes work.