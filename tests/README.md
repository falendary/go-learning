
# About

How to run tests, how to make comparisons and fail tests

# How to run

```go test ./tests -v```


# Notes


Every test function must:

1. Start with the word "Test"
2. Take a single parameter: t *testing.T
3. Be in a file ending with _test.go

The testing package automatically discovers and runs them.
