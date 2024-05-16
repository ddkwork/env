package main

import (
	"env"

	"github.com/ddkwork/app"
	"github.com/richardwilkes/unison"
)

func main() {
	app.Run("env", func(w *unison.Window) {
		env.New().Layout(w.Content())
	})
}
