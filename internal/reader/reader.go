package reader

type Reader interface {
	ReadFromCSV() ([]Order, error)
}

type reader struct {
	csvName string
}

func NewReader(csvName string) Reader {
	return &reader{
		csvName: csvName,
	}
}
