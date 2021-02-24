package safemypanic

import "log"

type GoDoer func()

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
