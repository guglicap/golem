package modules

type Update struct {
	Slot    Slot
	Color   Colors
	Content string
}

type Colors struct {
	Background string
	Foreground string
}

type Slot struct {
	Position int
	Index    int
}
