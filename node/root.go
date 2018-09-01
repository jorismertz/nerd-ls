package node

import (
  "os"
  "log"
  "fmt"
  "io/ioutil"
)

// Node - File or directory helper methods
type Node interface {
  Name() string
  Size() int
}

type node struct {
  file os.FileInfo
  name string
  size int
}

// Fetch - Fetch nodes in currently directory
func Fetch() []Node {
  files, err := ioutil.ReadDir(".")
  if err != nil { log.Fatal(err) }

  nodes := make([]Node, len(files))

  for i:=0; i<len(files); i++ {
    nodes[i] = node{files[i], "", 0}
  }

  return nodes
}

func (n node) Name() string {
  if n.name != "" { return n.name }

  name := n.file.Name()
  if n.file.IsDir() {
    n.name = fmt.Sprintf(" %s/ ", name)
  } else {
    n.name = fmt.Sprintf("  %s ", name)
  }

  return n.name
}

func (n node) Size() int {
  if n.size != 0 { return n.size }
  n.size = len([]rune(n.Name()))
  return n.size
}
