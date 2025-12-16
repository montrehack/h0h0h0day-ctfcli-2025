package main

import (
	"fmt"
	"hash/fnv"
	"math/rand"
)

var GridHeight, GridWidth = 10, 10

func hashStringToSeed(s string) int64 {
	h := fnv.New64a()
	h.Write([]byte(s))
	return int64(h.Sum64())
}

func checkMazeRow(name string, rowY int) bool {
	seed := hashStringToSeed(name)
	source := rand.NewSource(seed)
	r := rand.New(source)

	const treeProbability = 0.75

	for y := range GridHeight {
		for x := range GridWidth {

			randValue := r.Float64()

			if y == rowY {
				if randValue < treeProbability {
					return false
				}
			}
			if x == 5 && y == 4 {
				if randValue < treeProbability {
					return false
				}
			}
		}
	}
	return true
}

// bruteForceNamesForClearRow systematically checks names starting from length 1
// up to a specified maximum length, and stops after finding the target number of names
// that have no trees in the specified row (y=5).
func BruteForceNamesForClearRow(targetRowY int, targetCount int, maxLength int) []string {
	var foundNames []string
	alphabet := "abcdefghijklmnopqrstuvwxyz"

	queue := []string{""}

	fmt.Printf("Starting deterministic brute-force search for %d names with clear row y=%d...\n", targetCount, targetRowY)

	for len(queue) > 0 {
		currentName := queue[0]
		queue = queue[1:]

		if len(currentName) >= maxLength {
			continue
		}

		isTestable := len(currentName) > 0

		for _, char := range alphabet {
			newName := currentName + string(char)

			if isTestable || len(newName) == 1 {
				if checkMazeRow(newName, targetRowY) {
					foundNames = append(foundNames, newName)
					fmt.Printf("✅ Found Name #%d: '%s' (Length %d)\n", len(foundNames), newName, len(newName))

					if len(foundNames) >= targetCount {
						return foundNames
					}
				}

				queue = append(queue, newName)
			}
		}
	}

	fmt.Println("Warning: Brute force completed but could not find the target number of names.")
	return foundNames
}
