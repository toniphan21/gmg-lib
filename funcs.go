package gmg

type MyID struct {
	ID string
}

func MyIDToString(id MyID) string {
	return id.ID
}

func StringToMyID(id string) MyID {
	return MyID{ID: id}
}

type converter struct{}

func (c *converter) MyIDToString(id MyID) string {
	return id.ID
}

func (c *converter) StringToMyID(id string) MyID {
	return MyID{ID: id}
}

var Converter = &converter{}
