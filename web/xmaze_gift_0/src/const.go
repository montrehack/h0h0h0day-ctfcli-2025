package xmazegift

const (
	MOVE_UP    string = "w"
	MOVE_DOWN  string = "s"
	MOVE_LEFT  string = "a"
	MOVE_RIGHT string = "d"
	INTERACT   string = "e"
	LETTER_B   string = "b" // Not sure if this one is useful
)

var VALID_MOVES = []string{
	MOVE_UP,
	MOVE_DOWN,
	MOVE_LEFT,
	MOVE_RIGHT,
	INTERACT,
	LETTER_B,
}

const (
	MoveUp    = 'w'
	MoveDown  = 's'
	MoveLeft  = 'a'
	MoveRight = 'd'
	GrabItem  = 'q'
)

const (
	ItemSnow   Item = "snow"
	ItemTree   Item = "tree"
	ItemGift   Item = "gift"
	ItemScroll Item = "scroll"
	ItemHat    Item = "hat"
)

const (
	startX, startY   = 5, 9
	hatX, hatY       = 3, 5
	giftX, giftY     = 5, 0
	scrollX, scrollY = 0, 9
	maximumHatMove   = 5
	maximumTotalMove = 10
)

const (
	GridSize        = 10
	GridWidth       = GridSize
	GridHeight      = GridSize
	TreeProbability = 0.75
)
