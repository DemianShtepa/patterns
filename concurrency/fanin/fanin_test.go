package fanin

import (
	"testing"
)

func TestFanIn(t *testing.T) {
	sources := make([]<-chan interface{}, 3)
	for i := 0; i < cap(sources); i++ {
		source := make(chan interface{}, 1)
		source <- i
		close(source)
		sources[i] = source
	}

	var resultCount int
	result := FanIn(sources...)

	for range result {
		resultCount++
	}

	if resultCount != len(sources) {
		t.Errorf("expected result count to be %d, got %d", len(sources), resultCount)
	}
}
