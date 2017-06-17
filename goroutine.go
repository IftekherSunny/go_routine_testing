package example

import "sync"

////////////////////////////////////////////
//  Make message function
////////////////////////////////////////////
func MakeMessage(content chan string) chan string {
	message := make(chan string)
	wg := new(sync.WaitGroup)
	wg.Add(1)

	go generate(message, content, wg)

	go func() {
		wg.Wait()
	}()

	return message
}

func generate(message chan string, content chan string, wg *sync.WaitGroup) {
	message <- "Your message is " + <-content
	wg.Done()
}
