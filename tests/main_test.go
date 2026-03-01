package tests

import (
	//"errors"
	"errors"
	"testing"
)

/*
Every test function must:

1. Start with the word "Test"
2. Take a single parameter: t *testing.T
3. Be in a file ending with _test.go

The testing package automatically discovers and runs them.
*/

func TestSum(t *testing.T) {

	/*
		Table-driven test pattern.
		We define a slice of test cases (inputs + expected outputs),
		then loop over them.

		This avoids repetitive test code.
	*/

	tests := []struct {
		name string
		in   []int
		want int
	}{
		{"nil", nil, 0},
		{"empty", []int{}, 0},
		{"one", []int{5}, 5},
		{"many", []int{5, 15, 30}, 50},
		{"negatives", []int{-5, 2, 1}, -2},
	}

	for _, tt := range tests {

		t.Run(tt.name, func(t *testing.T) {

			got := Sum(tt.in)

			/*
				If the result does not match expected value,
				we fail the test immediately.
			*/

			if got != tt.want {
				t.Fatalf("Sum(%v) = %d, wand %d", tt.in, got, tt.want)
			}
		})

	}

}

func TestDivide_Success(t *testing.T) {

	got, err := Divide(10, 2)

	/*
		If err is not nil, something failed unexpectedly.
	*/

	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	if got != 5 {
		t.Fatalf("expected 5, got %d", got)
	}

}

func TestDivide_ByZero_TypedError(t *testing.T) {
	_, err := Divide(10, 0)

	if err == nil {
		t.Fatalf("expected error, got nil")
	}

	/*
		errors.As checks whether the error chain
		contains a specific type (ValidationError).
	*/

	var ve ValidationError

	if !errors.As(err, &ve) {
		t.Fatalf("expected ValidationError via errors.As, got %T: %v", err, err)
	}

	if ve.Field != "b" {
		t.Fatalf("expected Field=b, got %s", ve.Field)
	}
}

func TestFindUserEmail_NotFound_Sentinel(t *testing.T) {
	_, err := FindUserEmail("100")

	if err == nil {
		t.Fatalf("expected error, got nil")
	}

	/*
		errors.Is checks whether the error chain
		contains a specific sentinel error.
	*/

	if !errors.Is(err, ErrNotFound) {
		t.Fatalf("expected ErrNotFound via errors.Is, got %v", err)
	}
}

func TestFindUserEmail_Success(t *testing.T) {
	email, err := FindUserEmail("42")

	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	if email != "sz@example.com" {
		t.Fatalf("expected sz@example.com, got %s", email)
	}
}
