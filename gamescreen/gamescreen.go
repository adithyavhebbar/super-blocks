package gamescreen

const (
	Menu     = iota
	Game     = iota
	Credits  = iota
	Controls = iota
	End      = iota
)

type screenType struct {
	Menu     int
	Game     int
	Credits  int
	Controls int
	End      int
}

var ScreenType = getAllScreenType()

var CurrentScreenType int

func getAllScreenType() *screenType {
	return &screenType{
		Menu:     Menu,
		Game:     Game,
		Credits:  Credits,
		Controls: Controls,
		End:      End,
	}
}

func GetScreenType() int {
	return CurrentScreenType
}

func SetScreenType(screenType int) {
	if screenType > 4 || screenType < 0 {
		CurrentScreenType = ScreenType.Menu
	} else {
		CurrentScreenType = screenType
	}
}
