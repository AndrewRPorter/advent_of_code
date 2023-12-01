package main

import "testing"
import "fmt"

func TestDoSectionsFullyOverlap(t *testing.T) {
	got := DoSectionsFullyOverlap("1-3", "2-2")
	fmt.Println(got)
	if !got {
		t.Errorf("DoSectionsFullyOverlap = %v; want True", got)
	}
}
