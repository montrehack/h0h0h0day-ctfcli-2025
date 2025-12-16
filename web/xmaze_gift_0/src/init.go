package xmazegift

import "time"

var elves = []Elf{
	// --- The Main Workshop Crew ---
	{"Jingle", time.Date(2025, time.December, 3, 21, 27, 47, 0, time.UTC).String(), time.Date(2025, time.December, 3, 21, 46, 5, 0, time.UTC).String(), ""},
	{"Tinsel", time.Date(2025, time.December, 2, 7, 8, 10, 0, time.UTC).String(), time.Date(2025, time.December, 2, 7, 24, 13, 0, time.UTC).String(), ""},
	{"Sparkle", time.Date(2025, time.December, 4, 19, 33, 52, 0, time.UTC).String(), time.Date(2025, time.December, 4, 19, 39, 59, 0, time.UTC).String(), ""},
	{"Twinkle", time.Date(2025, time.December, 15, 10, 35, 55, 0, time.UTC).String(), time.Date(2025, time.December, 15, 10, 56, 33, 0, time.UTC).String(), ""},
	{"Holly", time.Date(2025, time.December, 11, 14, 35, 25, 0, time.UTC).String(), time.Date(2025, time.December, 11, 14, 40, 4, 0, time.UTC).String(), ""},
	{"Noel", time.Date(2025, time.December, 9, 6, 41, 56, 0, time.UTC).String(), time.Date(2025, time.December, 9, 6, 46, 34, 0, time.UTC).String(), ""},

	// --- Toy Production & Design ---
	{"Flicker", time.Date(2025, time.December, 14, 21, 15, 1, 0, time.UTC).String(), time.Date(2025, time.December, 14, 21, 19, 2, 0, time.UTC).String(), ""},
	{"Gizmo", time.Date(2025, time.December, 4, 1, 10, 9, 0, time.UTC).String(), time.Date(2025, time.December, 4, 1, 38, 51, 0, time.UTC).String(), ""},
	{"Pip", time.Date(2025, time.December, 10, 8, 5, 7, 0, time.UTC).String(), time.Date(2025, time.December, 10, 8, 34, 59, 0, time.UTC).String(), ""},
	{"Crackle", time.Date(2025, time.December, 3, 1, 10, 26, 0, time.UTC).String(), time.Date(2025, time.December, 3, 1, 31, 13, 0, time.UTC).String(), ""},
	{"Wren", time.Date(2025, time.December, 5, 13, 21, 38, 0, time.UTC).String(), time.Date(2025, time.December, 5, 13, 35, 40, 0, time.UTC).String(), ""},
	{"Mistletoe", time.Date(2025, time.December, 16, 0, 58, 32, 0, time.UTC).String(), time.Date(2025, time.December, 16, 0, 58, 38, 0, time.UTC).String(), ""},

	// --- Sleigh & Logistics Team ---
	{"Frosty", time.Date(2025, time.December, 10, 5, 38, 55, 0, time.UTC).String(), time.Date(2025, time.December, 10, 5, 56, 14, 0, time.UTC).String(), ""},
	{"Snowflake", time.Date(2025, time.December, 9, 4, 11, 59, 0, time.UTC).String(), time.Date(2025, time.December, 9, 4, 16, 48, 0, time.UTC).String(), ""},
	{"Cringle", time.Date(2025, time.December, 5, 4, 4, 44, 0, time.UTC).String(), time.Date(2025, time.December, 5, 4, 25, 59, 0, time.UTC).String(), ""},
	{"Blizzard", time.Date(2025, time.December, 8, 3, 57, 12, 0, time.UTC).String(), time.Date(2025, time.December, 8, 4, 5, 30, 0, time.UTC).String(), ""},

	// --- Miscellaneous Helpers ---
	{"Pepper", time.Date(2025, time.December, 6, 13, 45, 31, 0, time.UTC).String(), time.Date(2025, time.December, 6, 13, 48, 20, 0, time.UTC).String(), ""},
	{"Cookie", time.Date(2025, time.December, 5, 3, 18, 48, 0, time.UTC).String(), time.Date(2025, time.December, 5, 3, 38, 57, 0, time.UTC).String(), ""},
	{"Glimmer", time.Date(2025, time.December, 4, 0, 45, 34, 0, time.UTC).String(), time.Date(2025, time.December, 4, 1, 7, 9, 0, time.UTC).String(), ""},
	{"Chipper", time.Date(2025, time.December, 14, 20, 33, 41, 0, time.UTC).String(), time.Date(2025, time.December, 14, 20, 39, 58, 0, time.UTC).String(), ""},

	// Santa
	{FLAG_3, time.Date(101, time.January, 01, 01, 01, 01, 01, time.UTC).String(), time.Date(1010, time.October, 10, 10, 10, 10, 10, time.UTC).String(), ""},
}
