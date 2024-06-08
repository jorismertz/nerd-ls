package list

import (
	"log"
	"os"
	"sort"

	"github.com/drn/nerd-ls/git"
	"github.com/drn/nerd-ls/node"
)

// Fetch - Fetch List representing current directory
func Fetch(dir string, options map[string]interface{}) []node.Node {
	// files, err := ioutil.ReadDir(dir)
	files, err := os.ReadDir(dir)
	if err != nil {
		log.Fatal(err)
	}

	if options["git-ignore"].(bool) {
		patterns := git.ReadIgnoreFile()
		files = git.FilterFiles(files, patterns)
	}

	if options["size-sort"].(bool) {
		sort.Slice(files, func(i, j int) bool {
			infoI, _ := files[i].Info()
			infoJ, _ := files[j].Info()
			return infoI.Size() > infoJ.Size()
		})
	}

	nodes := make([]node.Node, len(files)+2)

	index := 0
	if options["all"].(bool) && !options["no-relative"].(bool) {
		file, _ := os.Stat(".")
		nodes[0] = node.New(dir, file)
		file, _ = os.Stat("..")
		nodes[1] = node.New(dir, file)
		index += 2
	}

	for i := 0; i < len(files); i++ {
		if !options["all"].(bool) && []rune(files[i].Name())[0] == '.' {
			continue
		}

		info, err := files[i].Info()
		if err != nil {
			log.Fatal(err)
		}
		nodes[index] = node.New(dir, info)
		index++
	}

	return nodes[:index]
}
