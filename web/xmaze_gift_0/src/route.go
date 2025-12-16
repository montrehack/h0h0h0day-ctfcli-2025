package xmazegift

import (
	"database/sql"
	"encoding/json"
	"net/http"
	"strconv"
	"strings"
	"time"
)

var db *sql.DB

func init() {
	var err error
	db, err = InitDB()
	if err != nil {
		panic(err)
	}
}

func Static() http.Handler {
	return http.FileServer(http.Dir("static"))
}

func GetAncestors(w http.ResponseWriter, r *http.Request) {
	enforceMethod(w, r, "GET")
	randomSleep()

	name := r.URL.Query().Get("name")
	if name == "" {
		http.Error(w, "Oh, you didn't tell me who you are, what is your 'name' my dear?", http.StatusBadGateway)
		return
	}

	ancestors, err := getAncestors(db, name)

	if err != nil {
		http.Error(w, "Something seems strange, I coudn't get your ancestors", http.StatusBadRequest)
		return
	}

	err = json.NewEncoder(w).Encode(ancestors)
	if err != nil {
		http.Error(w, "oopsie, this elf had an issue", http.StatusBadRequest)
		return
	}
}

func Explore(w http.ResponseWriter, r *http.Request) {
	enforceMethod(w, r, "POST")
	randomSleep()

	data, elf, shouldReturn := getGameState(r, w)
	if shouldReturn {
		err := json.NewEncoder(w).Encode(ExploreOutput{})
		if err != nil {
			http.Error(w, "Encoding Error, this elf had an issue", http.StatusBadRequest)
			return
		}
	}

	maze := compileMaze(elf.Name, data.Meteo)

	exploreOutput := ExploreOutput{
		Maze: maze,
	}

	if data.Meteo == MeteoSunny {
		exploreOutput.Message = FLAG_0
		err := json.NewEncoder(w).Encode(exploreOutput)
		if err != nil {
			http.Error(w, "Encoding Error, this elf had an issue", http.StatusBadRequest)
			return
		}
	} else {
		err := json.NewEncoder(w).Encode(exploreOutput)
		if err != nil {
			http.Error(w, "Encoding Error, this elf had an issue", http.StatusBadRequest)
			return
		}
	}
}

func Gift(w http.ResponseWriter, r *http.Request) {
	enforceMethod(w, r, "POST")
	randomSleep()

	data, elf, shouldReturn := getGameState(r, w)
	if shouldReturn {
		return
	}

	if data.Meteo != MeteoSunny {
		return
	}

	var err error
	moveToHat := elf.Moves[:5]
	// 1. Grab Hat
	validationResult := validateMazeMoves(elf.Name, moveToHat, string(ItemHat), data.Meteo, true)
	if !validationResult.IsValid || validationResult.TotalMoves > maximumHatMove {
		err = json.NewEncoder(w).Encode(validationResult)
		if err != nil {
			http.Error(w, "Encoding Error, this elf had an issue", http.StatusBadRequest)
			return
		}
		return
	}

	// 2. Grab Gift
	validationResult = validateMazeMoves(elf.Name, elf.Moves, string(ItemGift), data.Meteo, true)

	if validationResult.IsValid && validationResult.TotalMoves <= maximumTotalMove {
		ItemOutput := ItemOutput{
			Message: "The lost Xmaze gift: " + FLAG_2,
		}
		err = json.NewEncoder(w).Encode(ItemOutput)
	} else {
		err = json.NewEncoder(w).Encode(validationResult)
	}

	if err != nil {
		http.Error(w, "Encoding Error, this elf had an issue", http.StatusBadRequest)
		return
	}

}

func Hat(w http.ResponseWriter, r *http.Request) {
	enforceMethod(w, r, "POST")

	data, elf, shouldReturn := getGameState(r, w)

	if shouldReturn {
		return
	}

	validationResult := validateMazeMoves(data.Name, elf.Moves, string(ItemHat), data.Meteo, true)

	if validationResult.IsValid {
		ItemOutput := ItemOutput{
			Message: string(ItemHat),
		}
		err := json.NewEncoder(w).Encode(ItemOutput)
		if err != nil {
			http.Error(w, "Encoding Error, this elf had an issue", http.StatusBadRequest)
			return
		}
	}

}
func Scroll(w http.ResponseWriter, r *http.Request) {
	enforceMethod(w, r, "POST")

	data, elf, shouldReturn := getGameState(r, w)
	if shouldReturn {
		return
	}

	if data.Meteo != MeteoSunny {
		return
	}

	validationResult := validateMazeMoves(data.Name, elf.Moves, string(ItemScroll), data.Meteo, false)
	if validationResult.IsValid {
		ItemOutput := ItemOutput{
			Message: "The Maze secret is at /api/b4ckup_d0wnl04D_P4tH",
		}
		err := json.NewEncoder(w).Encode(ItemOutput)
		if err != nil {
			http.Error(w, "Encoding Error, this elf had an issue", http.StatusBadRequest)
			return
		}
	}

}

func RegisterNewElf(w http.ResponseWriter, r *http.Request) {
	enforceMethod(w, r, "POST")
	randomSleep()

	var data newElfPayload
	err := json.NewDecoder(r.Body).Decode(&data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	data.Name = strings.ToLower(data.Name)
	err = validateElfName(data.Name)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	newElf := Elf{
		Name:      data.Name,
		BirthDate: time.Now().UTC().String(),
		DeathDate: "",
	}

	err = insertElf(db, &newElf)
	if err != nil {
		err = json.NewEncoder(w).Encode(Elf{})
		if err != nil {
			http.Error(w, "Encoding Error, this elf had an issue", http.StatusBadRequest)
			return
		}
		return
	}

	err = json.NewEncoder(w).Encode(newElf)
	if err != nil {
		http.Error(w, "Encoding Error, this elf had an issue", http.StatusBadRequest)
		return
	}
}

func KillElf(w http.ResponseWriter, r *http.Request) {
	enforceMethod(w, r, "POST")
	randomSleep()

	var data newElfPayload
	err := json.NewDecoder(r.Body).Decode(&data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	data.Name = strings.ToLower(data.Name)
	err = validateElfName(data.Name)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err = killElf(db, data.Name)
	if err != nil {
		message := "DB Error, this elf already exists, or died..."
		http.Error(w, message, http.StatusBadRequest)
		return
	}

	elf, err := getElfByName(db, data.Name)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err = json.NewEncoder(w).Encode(elf)
	if err != nil {
		http.Error(w, "Encoding Error, this elf had an issue", http.StatusBadRequest)
		return
	}
}

func MoveElf(w http.ResponseWriter, r *http.Request) {
	enforceMethod(w, r, "POST")

	var data movePayload
	err := json.NewDecoder(r.Body).Decode(&data)
	if err != nil {
		http.Error(w, "you want to move? please tell me who your 'name' and 'move'", http.StatusBadRequest)
		return
	}

	data.Move = strings.ToLower(data.Move)
	err = validateElfMove(data.Move)
	if err != nil {
		err = json.NewEncoder(w).Encode(moveOutput{})
		if err != nil {
			http.Error(w, "oopsie, this elf had an issue", http.StatusBadRequest)
			return
		}
		return
	}

	moves, err := moveElf(db, data.Name, data.Move)
	if err != nil {
		err = json.NewEncoder(w).Encode(moveOutput{
			Message: err.Error(),
		})
		if err != nil {
			http.Error(w, "oopsie, this elf had an issue", http.StatusBadRequest)
			return
		}
		return
	}

	containsEasterEgg := checkEasterEgg(moves)

	outputData := moveOutput{
		Name:    data.Name,
		Move:    data.Move,
		Moves:   moves,
		Message: "",
	}
	if containsEasterEgg {
		outputData.Message = FLAG_4
	}

	err = json.NewEncoder(w).Encode(outputData)
	if err != nil {
		http.Error(w, "oopsie, this elf had an issue", http.StatusBadRequest)
		return
	}
}

func DownloadSrc(w http.ResponseWriter, r *http.Request) {
	enforceMethod(w, r, "GET")
	randomSleep()

	fileName := "FLAG_1_c0ll1sion_f4ilure.zip"
	filePath := "./backup/" + fileName

	w.Header().Set("Content-Disposition", "attachment; filename="+strconv.Quote(fileName))
	w.Header().Set("Content-Type", "application/octet-stream")

	http.ServeFile(w, r, filePath)
}
