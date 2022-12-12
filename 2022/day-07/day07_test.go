package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDirectory_AddEntries(t *testing.T) {
	entries := []string{
		"dir a",
		"dir b",
		"29116 f",
		"2557 g",
	}

	dir := NewDirectory("/")
	dir.AddEntries(entries)

	expected := NewDirectory("/")
	expected.AddSubDirectory(NewDirectory("a"))
	expected.AddSubDirectory(NewDirectory("b"))
	expected.AddFile(NewFile("f", 29116))
	expected.AddFile(NewFile("g", 2557))

	assert.Equal(t, expected, dir)
}

func TestDirectory_GetSubDirectory(t *testing.T) {
	dir := Directory{
		Name: "/",
		Directories: []*Directory{
			{Name: "a"},
			{Name: "b"},
		},
		Files: []*File{
			{Name: "f", Size: 29116},
			{Name: "g", Size: 2557},
		},
	}

	assert.Nil(t, dir.GetSubDirectory("foo"))
	assert.Equal(t, "a", dir.GetSubDirectory("a").Name)
}

func TestNavigator_Execute(t *testing.T) {
	n := NewNavigator()

	n.Execute("cd /")
	assert.Equal(t, n.CurrentDir.Name, "/")
	assert.Empty(t, n.CurrentDir.Directories)
	assert.Empty(t, n.CurrentDir.Files)

	n.CurrentDir.AddEntries([]string{
		"dir a",
		"dir b",
	})
	assert.Equal(t, 2, len(n.CurrentDir.Directories))
	assert.Equal(t, 0, len(n.CurrentDir.Files))

	n.Execute("cd a")
	assert.Equal(t, "a", n.CurrentDir.Name)

	n.Execute("cd ..")
	assert.Equal(t, "/", n.CurrentDir.Name)
	assert.Equal(t, 2, len(n.CurrentDir.Directories))
}

func TestParseHistory(t *testing.T) {
	input := []string{
		"$ cd /",
		"$ ls",
		"dir a",
		"14848514 b.txt",
		"8504156 c.dat",
		"dir d",
		"$ cd a",
		"$ ls",
		"dir e",
		"29116 f",
		"2557 g",
		"62596 h.lst",
		"$ cd e",
		"$ ls",
		"584 i",
		"$ cd ..",
		"$ cd ..",
		"$ cd d",
		"$ ls",
		"4060174 j",
		"8033020 d.log",
		"5626152 d.ext",
		"7214296 k",
	}

	n := NewNavigator()
	n.ParseHistory(input)
	n.Print()
}

func TestDirectory_Size(t *testing.T) {
	d1 := NewDirectory("/")
	d2 := NewDirectory("a")
	d3 := NewDirectory("b")

	d1.AddSubDirectory(d2)
	d2.AddSubDirectory(d3)

	d2.AddFile(NewFile("foo.txt", 100))
	d3.AddFile(NewFile("bar.txt", 200))

	assert.Equal(t, 300, d1.Size())
}

func TestDirectory_Collect(t *testing.T) {
	d1 := NewDirectory("/")
	d1.AddFile(NewFile("root.txt", 100))

	d2 := NewDirectory("a")
	d2.AddFile(NewFile("a.txt", 200))
	d1.AddSubDirectory(d2)

	d3 := NewDirectory("b")
	d3.AddFile(NewFile("b.txt", 300))
	d2.AddSubDirectory(d3)

	assert.Equal(t, 600, d1.Size())
	assert.Len(t, d1.Collect(100), 0)
	assert.Len(t, d1.Collect(300), 1)
	assert.Len(t, d1.Collect(600), 3)
}
