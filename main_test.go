package main

import (
	"testing"
	"time"
)

func Test_main(t *testing.T) {
	//variables
	eatTime = 0 * time.Second
	thinkTime = 0 * time.Second
	sleepTime = 0 * time.Second

	main()
	if len(orderFinished) != 5 {
		t.Errorf("incorrect number of entries")
	}
}
