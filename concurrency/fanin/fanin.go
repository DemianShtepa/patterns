package fanin

import "sync"

func FanIn(sources ...<-chan interface{}) <-chan interface{} {
	result := make(chan interface{}, len(sources))
	var wg sync.WaitGroup
	wg.Add(len(sources))

	for _, source := range sources {
		go func(source <-chan interface{}) {
			defer wg.Done()

			for el := range source {
				result <- el
			}
		}(source)
	}

	go func() {
		wg.Wait()
		close(result)
	}()

	return result
}
