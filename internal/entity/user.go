package entity

type User struct {
	ID   int
	Name string
}

type Clicks struct {
	Multiplier float32
	Click      int
	Afk        int
}
