package structs

type NationCapital struct {
	Name		string
	X			int
	Z			int
}

type Nation struct {
	Name        string
	Leader      string
	Towns      	[]string
	Residents 	[]string
	Area		int
	Wiki		string
	Capital		NationCapital
}