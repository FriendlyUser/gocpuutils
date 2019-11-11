package main
import (
  "github.com/leaanthony/mewn"
  "github.com/wailsapp/wails"
  "github.com/FriendlyUser/gocpuutils/pkg/sys"
	"github.com/FriendlyUser/gocpuutils/pkg/files"
)

func basic() string {
  return "Hello World!"
}

func main() {

  js := mewn.String("./frontend/dist/app.js")
  css := mewn.String("./frontend/dist/app.css")

  myFileList, err := files.NewFiles()
  stats := &sys.Stats{}
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
  app.Bind(myFileList)
  app.Bind(stats)
  app.Run()
}
