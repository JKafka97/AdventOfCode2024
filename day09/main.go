package main

import (
	"errors"
	"strconv"
)

// Filesystem represents a virtual file system
type Filesystem struct {
	Blocks []int
	Size   []int
	Loc    []int
}

// SolutionForPart1 computes the answer for part 1
func SolutionForPart1(input string) (int, error) {
	if input == "" {
		return 0, errors.New("input cannot be empty")
	}

	fs := NewFilesystem(input, false)
	fs.MoveSimple()
	return fs.Checksum(), nil
}

// SolutionForPart2 computes the answer for part 2
func SolutionForPart2(input string) (int, error) {
	if input == "" {
		return 0, errors.New("input cannot be empty")
	}

	fs := NewFilesystem(input, true)
	fs.MoveAdvanced()
	return fs.Checksum(), nil
}

// NewFilesystem initializes a filesystem from the input string
func NewFilesystem(diskmap string, trackFiles bool) *Filesystem {
	fs := &Filesystem{
		Blocks: []int{},
	}

	isFile := true
	id := 0

	if trackFiles {
		fs.Size = make([]int, len(diskmap))
		fs.Loc = make([]int, len(diskmap))
	}

	for _, ch := range diskmap {
		x, err := strconv.Atoi(string(ch))
		if err != nil {
			continue
		}

		if isFile {
			if trackFiles {
				fs.Loc[id] = len(fs.Blocks)
				fs.Size[id] = x
			}
			for i := 0; i < x; i++ {
				fs.Blocks = append(fs.Blocks, id)
			}
			id++
			isFile = false
		} else {
			for i := 0; i < x; i++ {
				fs.Blocks = append(fs.Blocks, -1)
			}
			isFile = true
		}
	}

	return fs
}

// MoveSimple performs basic file movement
func (fs *Filesystem) MoveSimple() {
	arr := fs.Blocks
	firstFree := 0

	for arr[firstFree] != -1 {
		firstFree++
	}

	i := len(arr) - 1
	for arr[i] == -1 {
		i--
	}

	for i > firstFree {
		arr[firstFree] = arr[i]
		arr[i] = -1

		for i >= 0 && arr[i] == -1 {
			i--
		}
		for firstFree < len(arr) && arr[firstFree] != -1 {
			firstFree++
		}
	}
}

// MoveAdvanced performs more complex file movement
func (fs *Filesystem) MoveAdvanced() {
	arr := fs.Blocks
	big := 0

	for fs.Size[big] > 0 {
		big++
	}
	big--

	for toMove := big; toMove >= 0; toMove-- {
		firstFree := 0
		freeSpace := 0

		for firstFree < fs.Loc[toMove] && freeSpace < fs.Size[toMove] {
			firstFree += freeSpace
			freeSpace = 0

			for firstFree < len(arr) && arr[firstFree] != -1 {
				firstFree++
			}

			for firstFree+freeSpace < len(arr) && arr[firstFree+freeSpace] == -1 {
				freeSpace++
			}
		}

		if firstFree >= fs.Loc[toMove] {
			continue
		}

		for idx := 0; idx < fs.Size[toMove]; idx++ {
			arr[firstFree+idx] = toMove
			arr[fs.Loc[toMove]+idx] = -1
		}
	}
}

// Checksum calculates the checksum of the filesystem
func (fs *Filesystem) Checksum() int {
	ans := 0
	for i, x := range fs.Blocks {
		if x != -1 {
			ans += i * x
		}
	}
	return ans
}
