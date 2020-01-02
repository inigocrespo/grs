package grs

type Entries struct {
	entries []Entry
}

func (e *Entries) Ack() int {
	return 0
}

