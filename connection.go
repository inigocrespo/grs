package grs

type Connection struct {

}

func NewConnection() *Connection {
	return &Connection{}
}

func (*Connection) GetStream(key string) *Stream {
	return &Stream{key: key}
}

func (*Connection) GetStreams(key string, keys... string) *Streams {
	return nil;
}

func GetGroup(name string) *Group {
	return &Group{name: name}
}


