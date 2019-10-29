// contains logic for accessing files from directory
package main

import (
	"encoding/json"
  // "fmt"
  "os"
	"path/filepath"
  "time"
	"io/ioutil"
	"github.com/wailsapp/wails"
)

type Todos struct {
	filename string
	runtime  *wails.Runtime
	logger   *wails.CustomLogger
}

// NewTodos attempts to create a new Todo list
func NewTodos() (*Todos, error) {
	// Create new Todos instance
	result := &Todos{}
	// Return it
	return result, nil
}

// file picker seems broken for now
func (t *Todos) GetFiles() (string, error) {
  t.logger.Infof("This is fine")
	// filename := t.runtime.Dialog.SelectFile()
	// call is directory IsDirectory
	values := iterateJSON("frontend/src")
	// if len(filename) > 0 {
	// 	// t.setFilename(filename)
	// 	t.runtime.Events.Emit("filemodified")
  // }
  dir, _ := os.Getwd()
  t.logger.Infof(dir)
  t.logger.Infof(dir)
  return string(values), nil
}
func (t *Todos) WailsInit(runtime *wails.Runtime) error {
	t.runtime = runtime
	t.logger = t.runtime.Log.New("Todos")

	// Set the default filename to $HOMEDIR/mylist.json
	// homedir, err := runtime.FileSystem.HomeDir()
	// if err != nil {
	// 	return err
	// }
	// t.filename = path.Join(homedir, "mylist.json")
	t.runtime.Window.SetTitle(t.filename)
	// t.ensureFileExists()
	return nil
}

// get file logic
type File struct {
  ModifiedTime time.Time `json:"ModifiedTime"`
  IsLink       bool      `json:"IsLink"`
  IsDir        bool      `json:"IsDir"`
  LinksTo      string    `json:"LinksTo"`
  Size         int64     `json:"Size"`
  Name         string    `json:"Name"`
  Path         string    `json:"Path"`
  Children     []*File   `json:"Children"`
}

func iterateJSON(path string) []byte {
  rootOSFile, _ := os.Stat(path)
  rootFile := toFile(rootOSFile, path) //start with root file
  stack := []*File{rootFile}
  for len(stack) > 0 { //until stack is empty,
      file := stack[len(stack)-1] //pop entry from stack
      stack = stack[:len(stack)-1]
      children, _ := ioutil.ReadDir(file.Path) //get the children of entry
      for _, chld := range children {          //for each child
          child := toFile(chld, filepath.Join(file.Path, chld.Name())) //turn it into a File object
          file.Children = append(file.Children, child)                 //append it to the children of the current file popped
          stack = append(stack, child)                                 //append the child to the stack, so the same process can be run again
      }
  }
  output, _ := json.MarshalIndent(rootFile, "", "     ")
  return output
}

func toFile(file os.FileInfo, path string) *File {
  JSONFile := File{ModifiedTime: file.ModTime(),
      IsDir:    file.IsDir(),
      Size:     file.Size(),
      Name:     file.Name(),
      Path:     path,
      Children: []*File{},
  }
  if file.Mode()&os.ModeSymlink == os.ModeSymlink {
      JSONFile.IsLink = true
      JSONFile.LinksTo, _ = filepath.EvalSymlinks(filepath.Join(path, file.Name()))
  } // Else case is the zero values of the fields
  return &JSONFile
}


func IsDirectory(path string) (bool, error) {
  fileInfo, err := os.Stat(path)
  if err != nil{
    return false, err
  }
  return fileInfo.IsDir(), err
}

// FileExists reports whether the named file exists as a boolean
func FileExists(name string) bool {
  if fi, err := os.Stat(name); err == nil {
      if fi.Mode().IsRegular() {
          return true
      }
  }
  return false
}
