package example

import (
	"sync"
	"testing"
)

func TestMakeMessage(t *testing.T) {
	content := make(chan string)
	wg := new(sync.WaitGroup)
	wg.Add(1)

	go func() {
		content <- "Hi, Go routine"
		wg.Done()
	}()

	go func() {
		wg.Wait()
	}()

	message := MakeMessage(content)

	if <-message != "Your message is Hi, Go routine" {
		t.Error("Message must be 'Your message is Hi, go routine'")
	}
}
