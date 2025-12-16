package xmazegift

import (
	"fmt"
	"hash/fnv"
	"math/rand"
)

func hashStringToSeed(s string) int64 {
	h := fnv.New64a()
	h.Write([]byte(s))
	return int64(h.Sum64())
}

func compileMaze(name string, meteo Meteo) Maze {
	var maze Maze

	seed := hashStringToSeed(name)
	source := rand.NewSource(seed)
	r := rand.New(source)

	for x := range GridWidth {
		for y := range GridHeight {
			maze.Grid[x][y] = "empty"
		}
	}

	for x := range GridWidth {
		for y := range GridHeight {
			if r.Float64() < TreeProbability {
				maze.Grid[x][y] = ItemTree
			}
		}
	}

	maze.Grid[startX][startY] = "empty"

	if meteo == MeteoSunny {
		maze.Grid[giftX][giftY] = ItemGift
		maze.Grid[scrollX][scrollY] = ItemScroll
	} else {
		maze.Grid[giftX][giftY] = "empty"
		maze.Grid[scrollX][scrollY] = "empty"
	}

	maze.Grid[hatX][hatY] = ItemHat

	maze.Grid[1][9] = ItemTree
	maze.Grid[0][8] = ItemTree

	return maze
}

func validateMazeMoves(name string, moves string, target string, meteo Meteo, collision bool) ValidationResult {

	maze := compileMaze(name, meteo)

	currentX := startX
	currentY := startY
	processedMoves := 0
	var lastMove rune

	for _, moveRune := range moves {
		move := moveRune

		dx, dy := 0, 0

		switch move {
		case MoveUp:
			dy = -1
		case MoveDown:
			dy = 1
		case MoveLeft:
			dx = -1
		case MoveRight:
			dx = 1
		default:
			continue
		}

		// If a valid move character was found, track it as the last move (but only after move is confirmed valid)
		nextX := currentX + dx
		nextY := currentY + dy

		processedMoves++

		// --- 1. Boundary Check ---
		if nextX < 0 || nextX >= GridWidth || nextY < 0 || nextY >= GridHeight {
			return ValidationResult{
				IsValid:    false,
				Reason:     fmt.Sprintf("Out of bounds move to (%d, %d). Final valid position: (%d, %d).", nextX, nextY, currentX, currentY),
				FinalX:     currentX,
				FinalY:     currentY,
				TotalMoves: processedMoves - 1,
			}
		}

		// --- 2. Tree Collision Check ---
		if maze.Grid[nextX][nextY] == ItemTree && collision {
			return ValidationResult{
				IsValid:    false,
				Reason:     fmt.Sprintf("Invalid move: Collided with a Tree at (%d, %d). Final valid position: (%d, %d).", nextX, nextY, currentX, currentY),
				FinalX:     currentX,
				FinalY:     currentY,
				TotalMoves: processedMoves - 1,
			}
		}

		// Move is valid: update the current position and record the move
		currentX = nextX
		currentY = nextY
		lastMove = move
	}

	targetX, targetY := -1, -1

	switch target {
	case string(ItemHat):
		targetX, targetY = hatX, hatY
	case string(ItemGift):
		targetX, targetY = giftX, giftY
	case string(ItemScroll):
		targetX, targetY = scrollX, scrollY
	default:
		return ValidationResult{
			IsValid:    false,
			Reason:     fmt.Sprintf("Invalid target item specified: '%s'.", target),
			FinalX:     currentX,
			FinalY:     currentY,
			TotalMoves: processedMoves,
		}
	}

	// The player's final position must be adjacent to the target
	// and the last move's direction must be TOWARDS the target.

	// Check 1: Final position must be adjacent to the target
	// Distance formula: |targetX - currentX| + |targetY - currentY| == 1
	isAdjacent := (abs(targetX-currentX)+abs(targetY-currentY) == 1)

	if !isAdjacent {
		return ValidationResult{
			IsValid:    false,
			Reason:     fmt.Sprintf("Path complete, but final position (%d, %d) is not adjacent to the target %s (%d, %d).", currentX, currentY, target, targetX, targetY),
			FinalX:     currentX,
			FinalY:     currentY,
			TotalMoves: processedMoves,
		}
	}

	// Check 2: Final facing direction must be TOWARDS the target
	// The player is facing the tile (targetX, targetY) if:
	// - Last move was UP ('w') AND target is directly above (currentY - 1 = targetY)
	// - Last move was DOWN ('s') AND target is directly below (currentY + 1 = targetY)
	// - Last move was LEFT ('a') AND target is directly left (currentX - 1 = targetX)
	// - Last move was RIGHT ('d') AND target is directly right (currentX + 1 = targetX)

	isFacingTarget := false

	switch lastMove {
	case MoveUp:
		isFacingTarget = (currentY-1 == targetY) && (currentX == targetX)
	case MoveDown:
		isFacingTarget = (currentY+1 == targetY) && (currentX == targetX)
	case MoveLeft:
		isFacingTarget = (currentX-1 == targetX) && (currentY == targetY)
	case MoveRight:
		isFacingTarget = (currentX+1 == targetX) && (currentY == targetY)
	}

	if !isFacingTarget {
		return ValidationResult{
			IsValid:    false,
			Reason:     fmt.Sprintf("Path complete and adjacent, but facing direction ('%c') is not towards the target %s.", lastMove, target),
			FinalX:     currentX,
			FinalY:     currentY,
			TotalMoves: processedMoves,
		}
	}

	if target == string(ItemHat) && processedMoves > maximumHatMove {
		return ValidationResult{
			IsValid:    false,
			Reason:     fmt.Sprintf("All moves are valid and successfully facing the %s at (%d, %d).", target, targetX, targetY),
			FinalX:     currentX,
			FinalY:     currentY,
			TotalMoves: processedMoves,
		}
	}

	// If adjacent AND facing the target, the sequence is fully valid
	return ValidationResult{
		IsValid:    true,
		Reason:     fmt.Sprintf("All moves are valid and successfully facing the %s at (%d, %d).", target, targetX, targetY),
		FinalX:     currentX,
		FinalY:     currentY,
		TotalMoves: processedMoves,
	}
}

// Simple helper function for absolute value
func abs(i int) int {
	if i < 0 {
		return -i
	}
	return i
}
