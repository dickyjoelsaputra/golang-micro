package entity

type Laptop struct {
	ID    int    `json:"id" form:"form:id;primaryKey;autoIncrement"`
	Nama  string `json:"nama" form:"form:nama"`
	Harga int    `json:"harga" form:"form:harga"`
	Merk  string `json:"merk" form:"form:merk"`
	Os    string `json:"os" form:"form:os"`
}
