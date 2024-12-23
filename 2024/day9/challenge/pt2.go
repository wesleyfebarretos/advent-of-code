package challenge

import (
	"fmt"
	"slices"

	"github.com/wesleyfebarretos/advent-of-code/utils"
)

func Pt2() {
	puzzle := parsePuzzle(utils.GetPuzzle())

	diskMap := arrangeDiskMap(puzzle)

	freeSpaceIdxsRanges, fileIdxsRanges := findFreeSpaceAndRevertedFileIdxsRange(diskMap)

	removeFreeSpacesRangeGap(freeSpaceIdxsRanges, fileIdxsRanges, diskMap)

	total := fileSystemCheckSum(diskMap)

	fmt.Printf("Part 2 -> %d", total)
}

func removeFreeSpacesRangeGap(freeSpaceIdxsRanges, fileIdxsRanges [][2]int, diskMap []string) {
	for _, fileIdxsRange := range fileIdxsRanges {
		filesRange := fileIdxsRange[1] - fileIdxsRange[0] + 1

		for y, freeSpaceIdxsRange := range freeSpaceIdxsRanges {
			if freeSpaceIdxsRange[1] >= fileIdxsRange[1] {
				continue
			}

			spaceRange := freeSpaceIdxsRange[1] - freeSpaceIdxsRange[0] + 1

			if filesRange <= spaceRange {
				initFromSpace := freeSpaceIdxsRange[0]

				for i := fileIdxsRange[0]; i <= fileIdxsRange[1]; i++ {
					diskMap[initFromSpace] = diskMap[i]
					diskMap[i] = "."
					initFromSpace++
				}

				freeSpaceIdxsRanges[y][0] += filesRange

				break
			}
		}
	}
}

func findFreeSpaceAndRevertedFileIdxsRange(diskMap []string) ([][2]int, [][2]int) {
	freeSpaceIdxsRange := [][2]int{}
	fileIdxsRange := [][2]int{}

	for i := 0; i < len(diskMap); {
		v := diskMap[i]
		if v == "." {
			freeSpaceRange := [2]int{i}

			for i2 := i + 1; i2 < len(diskMap); i2++ {

				if i2 == len(diskMap)-1 {
					freeSpaceRange[1] = i2
					freeSpaceIdxsRange = append(freeSpaceIdxsRange, freeSpaceRange)
					i = len(diskMap)
					break
				}

				if diskMap[i2] != "." || i2 == len(diskMap)-1 {
					freeSpaceRange[1] = i2 - 1
					i = i2
					freeSpaceIdxsRange = append(freeSpaceIdxsRange, freeSpaceRange)
					break
				}
			}
		} else {
			fileRange := [2]int{i}

			for i3 := i + 1; i3 < len(diskMap); i3++ {
				if i3 == len(diskMap)-1 {
					fileRange[1] = i3
					fileIdxsRange = append(fileIdxsRange, fileRange)
					i = len(diskMap)
					break
				}

				if diskMap[i3] == "." || diskMap[i3] != v {
					fileRange[1] = i3 - 1
					i = i3
					fileIdxsRange = append(fileIdxsRange, fileRange)
					break
				}
			}
		}
	}

	slices.Reverse(fileIdxsRange)

	return freeSpaceIdxsRange, fileIdxsRange
}
