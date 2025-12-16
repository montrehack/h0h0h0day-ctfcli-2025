package xmazegift

type Elf struct {
	Name      string `db:"name" json:"name"`
	BirthDate string `db:"birth_date" json:"birthDate"`
	DeathDate string `db:"death_date,omitempty" json:"deathDate"`
	Moves     string `db:"moves,omitempty" json:"moves,omitempty"`
}

type newElfPayload struct {
	Name string `json:"name"`
}

type Meteo string

const (
	MeteoSunny Meteo = "sunny"
	MeteoFoggy Meteo = "foggy"
)

func (s Meteo) Valid() bool {
	switch s {
	case MeteoFoggy, MeteoSunny:
		return true
	}
	return false
}

type explorePayload struct {
	Name  string `json:"name"`
	Meteo Meteo  `json:"meteo"`
}

type movePayload struct {
	Name string `json:"name"`
	Move string `json:"move"`
}

type moveOutput struct {
	Name    string `json:"name"`
	Move    string `json:"move"`
	Moves   string `json:"moves"`
	Message string `json:"message"`
}

type Item string

type Maze struct {
	Grid [GridWidth][GridHeight]Item
}

type ExploreOutput struct {
	Maze    Maze   `json:"maze"`
	Message string `json:"message"`
}

type ValidationResult struct {
	IsValid    bool
	Reason     string
	FinalX     int
	FinalY     int
	TotalMoves int
}

type ItemOutput struct {
	Message string `json:"message"`
}
