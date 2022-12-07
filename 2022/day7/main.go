package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

type File struct {
	name   string
	size   uint64
	parent *Directory
}

type Directory struct {
	name        string
	files       []File
	directories []*Directory
	parent      *Directory
	size        uint64
}

func getSize(dir *Directory, sum uint64) (uint64, uint64) {
	size := uint64(0)
	for _, file := range dir.files {
		size += file.size
	}
	for _, dir := range dir.directories {
		newSize := uint64(0)
		newSize, sum = getSize(dir, sum)
		size += newSize
	}
	dir.size = size
	if size <= uint64(100000) {
		sum += size
	}

	return size, sum
}

func findMax(parent *Directory, required uint64, candidate *Directory) {
	if parent.size >= required && parent.size < candidate.size {
		*candidate = *parent
	}
	for _, dir := range parent.directories {
		if dir.size < required {
			continue
		}
		findMax(dir, required, candidate)
	}
}

const TOTAL_SIZE = uint64(70000000)

func main() {
	sum := uint64(0)

	root := Directory{name: "/", files: []File{}, directories: []*Directory{}, parent: nil}

	current := &root

	input, _ := ioutil.ReadFile("input.txt")
	lines := strings.Split(string(input), "\n")

	for _, line := range lines {
		command := strings.Split(line, " ")

		if command[0] == "$" {
			if command[1] == "cd" {
				dirName := command[2]
				if dirName == ".." {
					current = current.parent
					continue
				} else {
					for _, dir := range current.directories {
						if dir.name == dirName {
							current = dir
							break
						}
					}
				}
				continue
			} else {
				continue
			}
		}
		if command[0] == "dir" {
			dirName := command[1]
			dd := Directory{name: dirName, files: []File{}, directories: []*Directory{}, parent: current}
			for _, dir := range current.directories {
				if dir.name == dirName {
					dd = *dir
				}
			}
			current.directories = append(current.directories, &dd)
			continue
		}
		size := uint64(0)
		fmt.Sscanf(command[0], "%d", &size)
		file := File{name: command[1], size: size, parent: current}
		current.files = append(current.files, file)
		continue
	}
	root.size, sum = getSize(&root, sum)

	required := uint64(30000000) - (TOTAL_SIZE - root.size)

	fmt.Println(required)

	candidate := root

	findMax(&root, required, &candidate)

	fmt.Println(candidate)
}
