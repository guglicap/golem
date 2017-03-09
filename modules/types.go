package modules

//Update is what we send to the main goroutine
type Update struct {
	Slot    Slot
	Color   Colors
	Content string
}

//Colors holds like, you know, the colors
type Colors struct {
	Background string
	Foreground string
}

//Slot tells the main goroutine where to put the Content it receives
type Slot struct {
	Position int
	Index    int
}
