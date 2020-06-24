package sqlclient

import "errors"

type rowsMock struct {
	Columns []string
	Rows    [][]interface{}

	currentIndex int
}

func (m *rowsMock) HasNext() bool {
	return m.currentIndex < len(m.Rows)
}
func (m *rowsMock) Close() error {
	return nil
}
func (m *rowsMock) Scan(destinations ...interface{}) error {
	mockedRow := m.Rows[m.currentIndex]
	if len(mockedRow) != len(destinations) {
		return errors.New("invalid destination len")
	}
	for index, value := range mockedRow {
			destinations[index] = value
	}
	return nil
}
