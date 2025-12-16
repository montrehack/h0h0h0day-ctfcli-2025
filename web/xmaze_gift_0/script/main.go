package main

import (
	"fmt"
	"math/rand"
	"time"
)

// Elf struct definition for completeness
type Elf struct {
	Name      string
	BirthDate time.Time
	DeathDate time.Time
}

// Helper function to create a new Elf
func NewElf(name string, birth time.Time, death time.Time) *Elf {
	return &Elf{
		Name:      name,
		BirthDate: birth,
		DeathDate: death,
	}
}

// generateRandomLifespan assigns random BirthDate and DeathDate to an Elf
func generateRandomLifespan(elf *Elf, start time.Time, end time.Time) {
	// 1. Calculate the total available duration in seconds
	totalDuration := end.Sub(start)
	totalSeconds := int64(totalDuration.Seconds())

	// 2. Pick a random BirthDate offset within the total duration
	// Ensures birth date is before the maximum end time
	randBirthOffsetSeconds := rand.Int63n(totalSeconds - 1800) // subtract 30 minutes (1800s) to guarantee room for life span
	birthTime := start.Add(time.Duration(randBirthOffsetSeconds) * time.Second)

	// 3. Calculate the maximum life span (30 minutes in seconds)
	maxLifespanSeconds := int64(30 * time.Minute / time.Second)

	// 4. Calculate a random life span (1 second to 30 minutes)
	// We want a minimum of 1 second, so add 1 to the random number.
	randLifespanSeconds := rand.Int63n(maxLifespanSeconds) + 1

	// 5. Calculate the DeathDate
	deathTime := birthTime.Add(time.Duration(randLifespanSeconds) * time.Second)

	// Check if death time exceeds the absolute end time (Dec 16, 23:59:00)
	// This check is a safeguard, but steps 1 & 2 should prevent it.
	if deathTime.After(end) {
		deathTime = end
		// You could recalculate birthTime here, but truncating death is simpler
	}

	elf.BirthDate = birthTime
	elf.DeathDate = deathTime
}

func main() {
	// --- Time Range Definition ---
	// 1st December 2025, 00:00:00
	startTime := time.Date(2025, time.December, 1, 0, 0, 0, 0, time.UTC)
	// 16th December 2025, 23:59:00
	endTime := time.Date(2025, time.December, 16, 23, 59, 0, 0, time.UTC)

	// --- Elf List Initialization ---
	elves := []*Elf{
		NewElf("Jingle", time.Time{}, time.Time{}),
		NewElf("Tinsel", time.Time{}, time.Time{}),
		NewElf("Sparkle", time.Time{}, time.Time{}),
		NewElf("Twinkle", time.Time{}, time.Time{}),
		NewElf("Holly", time.Time{}, time.Time{}),
		NewElf("Noel", time.Time{}, time.Time{}),

		// --- Toy Production & Design ---
		NewElf("Flicker", time.Time{}, time.Time{}),
		NewElf("Gizmo", time.Time{}, time.Time{}),
		NewElf("Pip", time.Time{}, time.Time{}),
		NewElf("Crackle", time.Time{}, time.Time{}),
		NewElf("Wren", time.Time{}, time.Time{}),
		NewElf("Mistletoe", time.Time{}, time.Time{}),

		// --- Sleigh & Logistics Team ---
		NewElf("Frosty", time.Time{}, time.Time{}),
		NewElf("Snowflake", time.Time{}, time.Time{}),
		NewElf("Cringle", time.Time{}, time.Time{}),
		NewElf("Blizzard", time.Time{}, time.Time{}),

		// --- Miscellaneous Helpers ---
		NewElf("Pepper", time.Time{}, time.Time{}),
		NewElf("Cookie", time.Time{}, time.Time{}),
		NewElf("Glimmer", time.Time{}, time.Time{}),
		NewElf("Chipper", time.Time{}, time.Time{}),
	}

	// --- Assign Random Times ---
	fmt.Println("🎄 Generating Elf Lifespans (Max 30 minutes, Dec 1-16, 2025)")
	fmt.Println("---------------------------------------------------------------")

	// Seed the random number generator
	rand.Seed(time.Now().UnixNano())

	for _, elf := range elves {
		generateRandomLifespan(elf, startTime, endTime)
		fmt.Printf("*NewElf(\"%s\", time.Date(%d, time.December, %d, %d, %d, %d, 0, time.UTC), time.Date(%d, time.December, %d, %d, %d, %d, 0, time.UTC)),\n",
			elf.Name,
			elf.BirthDate.Year(), elf.BirthDate.Day(), elf.BirthDate.Hour(), elf.BirthDate.Minute(), elf.BirthDate.Second(),
			elf.DeathDate.Year(), elf.DeathDate.Day(), elf.DeathDate.Hour(), elf.DeathDate.Minute(), elf.DeathDate.Second(),
		)
	}
	fmt.Println("---------------------------------------------------------------")
}
