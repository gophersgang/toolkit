package retry_test

import (
	"context"
	"fmt"
	"github.com/donutloop/toolkit/retry"
)

func ExampleRetrier() {
	r := retry.NewRetrier()
	err := r.Retry(context.Background(), func() (bool, error) {
		fmt.Println("fire request")
		return true, nil
	})

	if err != nil {
		fmt.Println(fmt.Sprintf("error: (%v)", err))
	}

	// Output: fire request
}