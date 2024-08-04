package main

import (
	"env"

	"github.com/ddkwork/unison"
	"github.com/ddkwork/unison/app"
)

func main() {
	app.Run("env", func(w *unison.Window) {
		w.Content().AddChild(env.New().Layout())
	})
}
