package dto

type LaptopRequest struct {
	Nama  string `json:"nama" form:"form:nama"`
	Harga int    `json:"harga" form:"form:harga"`
	Merk  string `json:"merk" form:"form:merk"`
	Os    string `json:"os" form:"form:os"`
}
