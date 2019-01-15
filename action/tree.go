package action

import (
	"fmt"
	"os"
	"path"
	"sort"

	"github.com/fatih/color"
	"github.com/urfave/cli"
)

const defaultPath = "."
const levelIndent = "  "

var c = color.New(color.FgBlue)

// tree struct
type tree struct {
	c           *cli.Context
	files       uint64
	directories uint64
}

func newTree(c *cli.Context) *tree {
	return &tree{c: c}
}

func (t *tree) isDirOnly() bool {
	return t.c.Bool("d")
}

// TreeAction tree command action
func TreeAction(c *cli.Context) error {
	searchRootPath := c.Args().First()
	if searchRootPath == "" {
		searchRootPath = defaultPath
	}

	if !fileExist(searchRootPath) {
		return fmt.Errorf("%s [error opening dir]\n", searchRootPath)
	}

	tree := newTree(c)
	fmt.Println(searchRootPath)

	err := fileSearch(tree, searchRootPath, "")
	if err != nil {
		return err
	}

	fmt.Printf("\n%d directories, %d Files\n", tree.directories, tree.files)

	return nil
}

func fileSearch(tree *tree, filePath, prefix string) error {
	names, err := acquireFilesUnder(filePath)
	if err != nil {
		return err
	}

	if len(names) == 0 {
		return nil
	}

	fInfoList, err := makeFileInfoList(tree, filePath, names)
	if err != nil {
		return err
	}

	lastIndex := len(fInfoList) - 1
	for i, fInfo := range fInfoList {
		name := fInfo.Name()
		joinedPath := path.Join(filePath, name)

		// directory or files counter
		switch {
		case fInfo.IsDir():
			// for directories, add a color
			name = c.Sprint(name)
			tree.directories++
		default:
			tree.files++
		}

		var outputCtx, newPrefix string
		// create output content & new prefix
		switch lastIndex {
		case i:
			outputCtx = fmt.Sprintf("%s%s%s", prefix, "└", name)
			newPrefix = fmt.Sprintf("%s%s", prefix, levelIndent)
		default:
			outputCtx = fmt.Sprintf("%s%s%s", prefix, "├", name)
			newPrefix = fmt.Sprintf("%s%s%s", prefix, "│", levelIndent)
		}

		fmt.Println(outputCtx)
		err = fileSearch(tree, joinedPath, newPrefix)
		if err != nil {
			return err
		}
	}

	return nil
}

func fileExist(path string) bool {
	_, err := os.Stat(path)

	return err == nil
}

func acquireFilesUnder(filePath string) ([]string, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}

	names, _ := file.Readdirnames(0)
	sort.Strings(names)

	err = file.Close()
	if err != nil {
		return nil, err
	}

	return names, nil
}

func makeFileInfoList(tree *tree, filePath string, names []string) ([]os.FileInfo, error) {
	fInfoList := make([]os.FileInfo, 0)

	for _, name := range names {
		joinedPath := path.Join(filePath, name)

		fInfo, err := os.Stat(joinedPath)
		if err != nil {
			return nil, err
		}

		// search only directories
		if tree.isDirOnly() && !fInfo.IsDir() {
			continue
		}

		fInfoList = append(fInfoList, fInfo)
	}

	return fInfoList, nil
}
