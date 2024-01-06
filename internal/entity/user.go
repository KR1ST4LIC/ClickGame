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

type AddBal struct {
	UserID []int64
	AddBal []int64
	Count  int
}

type CantAdd struct {
}
