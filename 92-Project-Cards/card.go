package main

type card struct {
	value string
	house string
}

// (card) to_s(): return string
func (c card) toStr() string {
	return c.value + " of " + c.house
}
