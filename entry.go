package grs

type ID string

type Entry struct {
	id ID
	pairs map[string]interface{}
}

func (e *Entry) GetID() ID {
	return e.id
}

func (e *Entry) GetPairs() map[string]interface{} {
	return e.pairs
}

func (e *Entry) Ack() bool {
	return true
}

