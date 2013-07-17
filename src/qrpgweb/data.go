package qrpgweb

type Data struct {
	Name string
	What string
}

func NewData(name string, what string) *Data {
	data := new(Data)
	data.Name = name
	data.What = what
	return data
}

var noData *Data = &Data{Name:"Anonymous",What:"nothing"}
