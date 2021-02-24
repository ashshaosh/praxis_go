package safemypanic

import "log"

// GoDoer func w/o params
type GoDoer func()

// Go recovers with usual
func Go(todo GoDoer) {
	go func() {
		defer func() {
			if err := recover(); err != nil {
				log.Printf("Panic in: %s", err)
			}
		}()
		todo()
	}()
}
