package do

import (
	"fmt"
	"time"
)

// TimedExecution represents the result of a timed function execution
type TimedExecution struct {
	Duration time.Duration
	Error    error
}

// Time executes the provided closure and returns the execution duration and any error
func Time(fn func() error) TimedExecution {
	start := time.Now()
	err := fn()
	duration := time.Since(start)

	return TimedExecution{
		Duration: duration,
		Error:    err,
	}
}

// TimeWithResult executes a closure that returns a result and error, timing the execution
func TimeWithResult[T any](fn func() (T, error)) (T, TimedExecution) {
	start := time.Now()
	result, err := fn()
	duration := time.Since(start)

	return result, TimedExecution{
		Duration: duration,
		Error:    err,
	}
}

// TimeAndPrint executes the closure, prints the duration, and returns the result
func TimeAndPrint(name string, fn func() error) error {
	result := Time(fn)
	fmt.Printf("%s took: %v\n", name, result.Duration)
	return result.Error
}

// TimeAndPrintWithResult executes the closure, prints the duration, and returns the result
func TimeAndPrintWithResult[T any](name string, fn func() (T, error)) (T, error) {
	result, timing := TimeWithResult(fn)
	fmt.Printf("%s took: %v\n", name, timing.Duration)
	return result, timing.Error
}
