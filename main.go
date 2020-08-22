package main

import (
	"fmt"
)

func main() {
	// インターフェイスに実装を注入します.
	presenter := NewPresenter()
	interactor := NewInteractorUsecase(presenter)
	controller := NewController(interactor)

	for {
		var line string
		fmt.Scanln(&line)
		if line == "" {
			break
		}
		vreq := &ViewRequest{Message: line}
		controller.Control(vreq)
	}
	fmt.Println("Good bye...")
}
