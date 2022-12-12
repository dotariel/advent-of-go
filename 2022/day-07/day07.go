package main

import (
	"fmt"
	"strconv"
	"strings"
)

type Directory struct {
	Name        string
	Parent      *Directory
	Directories []*Directory
	Files       []*File
}

type File struct {
	Name string
	Size int
}

type DirectoryBuilder struct {
}

type BuildFunction func(string)

func NewDirectory(name string) *Directory {
	return &Directory{
		Name:        name,
		Parent:      nil,
		Directories: make([]*Directory, 0),
		Files:       make([]*File, 0),
	}
}

func NewFile(name string, size int) *File {
	return &File{
		Name: name,
		Size: size,
	}
}

func (d *Directory) AddEntries(entries []string) {
	for _, entry := range entries {
		parts := strings.Fields(entry)

		if parts[0] == "dir" {
			d.AddSubDirectory(NewDirectory(parts[1]))
			continue
		}

		size, _ := strconv.Atoi(parts[0])
		name := parts[1]

		d.Files = append(d.Files, NewFile(name, size))
	}
}

type Navigator struct {
	CurrentDir *Directory
}

func NewNavigator() Navigator {
	return Navigator{
		CurrentDir: NewDirectory("/"),
	}
}

func (n *Navigator) AddEntries(entries []string) {
	n.CurrentDir.AddEntries(entries)
}

func (n *Navigator) Go(path string) {
	if path == "/" {
		for {
			if n.CurrentDir.Parent == nil {
				return
			}
			n.CurrentDir = n.CurrentDir.Parent
		}
	}

	if path == ".." {
		n.CurrentDir = n.CurrentDir.Parent
		return
	}

	newDir := n.CurrentDir.GetSubDirectory(path)

	n.CurrentDir, newDir.Parent = newDir, n.CurrentDir
}

func (n *Navigator) Execute(command string) {
	parts := strings.Fields(command)

	if cmd, path := parts[0], parts[1]; cmd == "cd" {
		n.Go(path)
		return
	}
}

func (d *Directory) GetSubDirectory(name string) *Directory {
	for _, dir := range d.Directories {
		if dir.Name == name {
			return dir
		}
	}

	return nil
}

func (n *Navigator) ParseHistory(entries []string) {
	commandMarker := "$ "

	for i := 0; i < len(entries); i++ {
		entry := entries[i]

		if strings.HasPrefix(entry, commandMarker) {
			cmdAndArgs := entry[len(commandMarker):]

			if cmdAndArgs == "ls" {
				listings := make([]string, 0)

				for _, next := range entries[i+1:] {
					if strings.HasPrefix(next, commandMarker) {
						break
					}
					listings = append(listings, next)
				}
				n.AddEntries(listings)
			}

			if strings.HasPrefix(cmdAndArgs, "cd") {
				n.Execute(cmdAndArgs)
			}
		}
	}
}

func (n *Navigator) Print() {
	n.Go("/")
	n.CurrentDir.Print(0)
}

func (d *Directory) Print(indent int) {
	for _, file := range d.Files {
		file.Print(indent)
	}
	for _, dir := range d.Directories {
		fmt.Printf("%v %v (size=%v)\n", strings.Repeat("-", indent), dir.Name, dir.Size())
		dir.Print(indent + 2)
	}
}

func (f File) Print(indent int) {
	fmt.Printf("%v %v (%v)\n", strings.Repeat("-", indent), f.Name, f.Size)
}

func (d *Directory) AddSubDirectory(sub *Directory) {
	sub.Parent = d
	d.Directories = append(d.Directories, sub)
}

func (d *Directory) AddFile(f *File) {
	d.Files = append(d.Files, f)
}

func (d *Directory) Size() int {
	size := 0

	for _, file := range d.Files {
		size += file.Size
	}

	for _, subdir := range d.Directories {
		size += subdir.Size()
	}

	return size
}

func (d *Directory) Collect(limit int) []*Directory {
	dirs := make([]*Directory, 0)
	if d.Size() <= limit {
		dirs = append(dirs, d)
	}

	for _, dir := range d.Directories {
		if dir.Size() <= limit {
			dirs = append(dirs, dir)
		}

		for _, sub := range dir.Directories {
			dirs = append(dirs, sub.Collect(limit)...)
		}
	}

	return dirs
}

type BySize []*Directory

func (b BySize) Len() int           { return len(b) }
func (b BySize) Less(i, j int) bool { return b[i].Size() < b[j].Size() }
func (b BySize) Swap(i, j int)      { b[i], b[j] = b[j], b[i] }
