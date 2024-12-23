package challenge

import (
	"fmt"
	"math"
	"slices"
	"strconv"

	"github.com/wesleyfebarretos/advent-of-code/2024/utils"
)

type Blocks struct {
	ID        int
	Qty       int
	FreeSpace int
}

func Pt1() {
	puzzle := parsePuzzle(utils.GetPuzzle())

	diskMap := arrangeDiskMap(puzzle)

	freeSpaceIdxs, fileIdxs := findFreeSpaceAndRevertedFileIdxs(diskMap)

	removeFreeSpacesGap(freeSpaceIdxs, fileIdxs, diskMap)

	total := fileSystemCheckSum(diskMap)

	fmt.Printf("Part 1 -> %d", total)
}

func fileSystemCheckSum(diskMap []string) int {
	total := 0

	for i, v := range diskMap {
		if v == "." {
			continue
		}

		n, _ := strconv.Atoi(v)

		total += i * n
	}

	return total
}

func removeFreeSpacesGap(freeSpaceIdxs []int, fileIdxs []int, diskMap []string) {
	lastSwappedFileIdx := math.MaxInt

	for i, freeSpaceIdx := range freeSpaceIdxs {
		if freeSpaceIdx >= lastSwappedFileIdx {
			break
		}

		diskMap[freeSpaceIdx] = diskMap[fileIdxs[i]]
		diskMap[fileIdxs[i]] = "."

		lastSwappedFileIdx = fileIdxs[i]
	}
}

func arrangeDiskMap(blocks []Blocks) []string {
	diskMap := []string{}

	for _, block := range blocks {
		for range block.Qty {
			diskMap = append(diskMap, fmt.Sprintf("%d", block.ID))
		}

		for range block.FreeSpace {
			diskMap = append(diskMap, ".")
		}
	}

	return diskMap
}

func findFreeSpaceAndRevertedFileIdxs(diskMap []string) ([]int, []int) {
	freeSpaceIdxs := []int{}
	fileIdxs := []int{}

	for i, v := range diskMap {
		if v == "." {
			freeSpaceIdxs = append(freeSpaceIdxs, i)
		} else {
			fileIdxs = append(fileIdxs, i)
		}
	}

	slices.Reverse(fileIdxs)

	return freeSpaceIdxs, fileIdxs
}

func parsePuzzle(puzzle string) []Blocks {
	blocks := []Blocks{}

	for i := 0; i < len(puzzle); i += 2 {
		n, _ := strconv.Atoi(string(puzzle[i]))

		freeSpace := 0

		if i+2 < len(puzzle) {
			n2, _ := strconv.Atoi(string(puzzle[i+1]))
			freeSpace = n2
		}

		blocks = append(blocks, Blocks{
			Qty:       n,
			FreeSpace: freeSpace,
		})
	}

	for i := range blocks {
		blocks[i].ID = i
	}

	return blocks
}
