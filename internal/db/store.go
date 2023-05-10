package db

import "from_scratch_wep_api/types"

type Storer interface {
	GetTests() ([]*types.TestRequest, error)
	CreateTest(test types.TestRequest) (types.TestRequest, error)
}
