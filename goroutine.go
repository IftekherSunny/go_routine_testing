package go_routine_testing

import "sync"

////////////////////////////////////////////
//  Make message function
////////////////////////////////////////////
func MakeMessage(content chan string) chan string {
	message := make(chan string)
	wg := new(sync.WaitGroup)
	wg.Add(1)

	go func() {
		message <- "Your message is " + <-content
		wg.Done()
	}()

	go func() {
		wg.Wait()
	}()

	return message
}
