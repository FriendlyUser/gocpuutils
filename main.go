package main
import (
  "github.com/leaanthony/mewn"
  "github.com/wailsapp/wails"
)

func basic() string {
  return "Hello World!"
}

func main() {

  js := mewn.String("./frontend/dist/app.js")
  css := mewn.String("./frontend/dist/app.css")

  myTodoList, err := NewTodos()
	if err != nil {
		// log.Fatal(err)
	}
  app := wails.CreateApp(&wails.AppConfig{
    Width:  1024,
    Height: 768,
    Title:  "comp_util",
    JS:     js,
    CSS:    css,
    Colour: "#131313",
  })
  // app.Bind(basic)
  app.Bind(myTodoList)
  app.Run()
}
