# Exercise 2

## Question
Write a function called `fileLen` that has an input parameter of type `string` and returns an `int` and an `error`. 
The function takes in a file name and returns the number of bytes in the file. If there is an error reading the file, 
return the error. Use `defer` to make sure the file is closed properly.

## Solution
The `fileLen` function should look a lot like the sample code for the _simple_cat_ program. Rather than being in the
body of `main`, it is a separate function. Use the `count` returned from calls to `Read` and add them up. It is
important to get the `count` before checking if `Read` returned an error, since it returns an (ignorable) `io.EOF` 
error when it reaches the end of a file. 