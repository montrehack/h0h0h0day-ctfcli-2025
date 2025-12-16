package xmazegift

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"net/http"
	"slices"
	"strings"
	"time"
	"unicode"
)

func enforceMethod(w http.ResponseWriter, r *http.Request, method string) {
	if r.Method != method {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

var allowedSpecialMap = map[rune]bool{
	'/':  true,
	'\'': true,
	'(':  true,
	'*':  true,
	')':  true,
	'=':  true,
	'-':  true,
	' ':  true,
}

func validateElfName(s string) error {
	for _, char := range s {
		if unicode.IsLower(char) && char >= 'a' && char <= 'z' {
			continue
		}

		if char >= '0' && char <= '9' {
			continue
		}

		if allowedSpecialMap[char] {
			continue
		}

		return fmt.Errorf("invalid character '%c' found in string: %s", char, s)
	}
	return nil
}

func validateElfMove(s string) error {
	if len(s) == 0 {
		return fmt.Errorf("standing still is not a smart move")
	}

	if len(s) > 1 {
		return fmt.Errorf("you are trying to move to fast here, slow down")
	}

	if slices.Contains(VALID_MOVES, s) {
		return nil
	}

	return fmt.Errorf("this move is not in my repertoire")
}

func getGameState(r *http.Request, w http.ResponseWriter) (explorePayload, Elf, bool) {
	var data explorePayload
	err := json.NewDecoder(r.Body).Decode(&data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return explorePayload{}, Elf{}, true
	}

	if !data.Meteo.Valid() {
		http.Error(w, "invalid meteo status", http.StatusBadRequest)
		return explorePayload{}, Elf{}, true
	}

	data.Name = strings.ToLower(data.Name)
	err = validateElfName(data.Name)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return explorePayload{}, Elf{}, true
	}

	elf, err := getElfByName(db, data.Name)
	if err != nil {
		http.Error(w, "Oh my dear, I don't know you, are you sure your 'name' is correct?", http.StatusBadRequest)
		return explorePayload{}, Elf{}, true
	}

	return data, elf, false
}

func randomSleep() {
	minMs := 500
	maxMs := 1000

	randomMs := rand.Intn(maxMs-minMs) + minMs

	fmt.Printf("Sleeping for %d milliseconds...\n", randomMs)

	sleepDuration := time.Duration(randomMs) * time.Millisecond

	time.Sleep(sleepDuration)
}
