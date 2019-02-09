/*
Copyright 2018 Google LLC

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package files

import (
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"

	"github.com/pkg/errors"
)

var Root = "."

type NodePath struct {
	Name string
	IsDir bool
	Children *[]NodePath
}

// Read reads a file in .demoit folder.
func Read(path ...string) ([]byte, error) {
	content, err := ioutil.ReadFile(fullpath(path))
	if err != nil {
		return nil, errors.Wrap(err, "Unable to read "+fullpath(path))
	}

	return content, nil
}

// Get tree repository in folder.
func Tree(rootPath string) (NodePath, error) {
	var folder = fullpath([]string{rootPath})
	var root = NodePath{
		Name: rootPath,
		IsDir: true,
		Children: &[]NodePath{},
	}
	var err = filepath.Walk(folder, func(path string, info os.FileInfo, err error) error {
		splitPath := strings.Split(path, string(filepath.Separator))
		var currentFolder = &root
		for _, path := range splitPath {
			if path == Root || path ==  rootPath {
				continue
			}
			currentFolder = getNodePath(currentFolder, path, info)
		}
		return nil
	})
	return root, err
}



// Exists tests if a file exists.
func Exists(path ...string) bool {
	_, err := os.Stat(fullpath(path))
	return err == nil
}

func fullpath(path []string) string {
	return filepath.Join(Root, filepath.Join(path...))
}

func getNodePath(parent *NodePath, currentPath string, info os.FileInfo) *NodePath {
	if parent.Children != nil {
		for _, child := range *parent.Children {
			if child.Name == currentPath {
				return &child
			}
		}
	}
	var nodePath = &NodePath{
		Name: info.Name(),
		IsDir: info.IsDir(),
		Children: &[]NodePath{},
	}
	*parent.Children = append(*parent.Children, *nodePath)
	return nodePath
}