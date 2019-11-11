// contains logic for accessing files from directory
package files

import (
	"encoding/json"
  // "fmt"
  "os"
	"path/filepath"
  "time"
	"io/ioutil"
  "github.com/wailsapp/wails"
  "strings"
)

type Files struct {
  filename string
  SelectedDir string
	runtime  *wails.Runtime
  logger   *wails.CustomLogger
}

// NewTodos attempts to create a new Todo list
func NewFiles() (*Files, error) {
	// Create new Todos instance
	result := &Files{}
	// Return it
	return result, nil
}

// file picker seems broken for now
func (t *Files) GetFiles() (string, error) {
  t.logger.Infof("This is fine")
  filename := t.runtime.Dialog.SelectDirectory()
  directory := "src/frontend"
  t.logger.Infof(directory)
  // call is directory IsDirectory
  values := iterateJSON(directory)
  // values := iterateJSON(filename)
	// if len(filename) > 0 {
	// 	// t.setFilename(filename)
	// 	t.runtime.Events.Emit("filemodified")
  // }
  t.SelectedDir = filename
  return string(values), nil
}
func (t *Files) GetDir() (string, error) {
  t.logger.Infof(t.SelectedDir)
  return string(t.SelectedDir), nil
}

func (t *Files) WailsInit(runtime *wails.Runtime) error {
	t.runtime = runtime
	t.logger = t.runtime.Log.New("Files")

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
  Size         int64     `json:"value"`
  Name         string    `json:"label"`
  Path         string    `json:"Path"`
  Children     []*File   `json:"children"`
}

func stringInSlice(a string, list []string) bool {
  for _, b := range list {
      if b == a {
          return true
      }
  }
  return false
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
        if stringInSlice(chld.Name(), []string{"node_modules", "dist", "package-lock.json", "assets/images/logo.png", "logo.png"}) == false {
          child := toFile(chld, filepath.Join(file.Path, chld.Name())) //turn it into a File object
          file.Children = append(file.Children, child)                 //append it to the children of the current file popped
          stack = append(stack, child)                                 //append the child to the stack, so the same process can be run again
        }
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
