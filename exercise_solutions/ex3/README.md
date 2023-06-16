# Exercise 3

## Question
Write a function called `prefixer` that has an input parameter of type +string+ and returns a function that has an 
input parameter of type `string` and returns a `string`. The returned function should prefix its input with the 
string passed into `prefixer`. Use the following `main` function to test `prefixer`:

```go
func main() {
    helloPrefix := prefixer("Hello")
    fmt.Println(helloPrefix("Bob")) // should print Hello Bob
    fmt.Println(helloPrefix("Maria")) // should print Hello Maria
}
```

## Solution
The `prefixer` function looks like this:

```go
func prefixer(prefix string) func(string) string {
    return func (body string) string {
        return prefix + " " + body
    }
}
```
The return type is a function signature, which has the keyword `func`, the input parameter for the returned function (`(string)`)
and the return type for the returned function (`string`). In the body of the function, you return an anonymous function. This
function is a closure; it refers to the `prefix` input parameter to `prefixer`.