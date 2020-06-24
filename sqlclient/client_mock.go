package sqlclient

import "errors"

var (
	isMocked bool
)

func StartMockServer() {
	isMocked = true
}

func StopMockServer() {
	isMocked = false
}

type clientMock struct {
	mocks map[string]Mock
}

type Mock struct {
	Query   string
	Args    []interface{}
	Error   error
	Columns []string
	Rows    [][]interface{}
}

func AddMock(mock Mock) {
	if dbClient == nil {
		return
	}
	client, okType := dbClient.(*clientMock)
	if !okType {
		return
	}
	if client.mocks == nil {
		client.mocks = make(map[string]Mock, 0)
	}
	client.mocks[mock.Query] = mock
}

func (c *clientMock) Query(query string, args ...interface{}) (rows, error) {
	mock, exists := c.mocks[query]
	if !exists {
		return nil, errors.New("no mock found")
	}

	if mock.Error != nil {
		return nil, mock.Error
	}

	rows := rowsMock{
		Columns: mock.Columns,
		Rows:    mock.Rows,
	}

	return &rows, nil
}
