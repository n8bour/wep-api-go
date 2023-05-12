package db

import "from_scratch_wep_api/types"

type Local struct {
	store []*types.TestRequest
}

func NewLocalDB() Local {
	return Local{
		store: make([]*types.TestRequest, 0),
	}
}

func (l *Local) GetTests() ([]*types.TestRequest, error) {
	return l.store, nil
}
func (l *Local) CreateTest(test types.TestRequest) (types.TestRequest, error) {
	l.store = append(l.store, &test)

	return test, nil
}
