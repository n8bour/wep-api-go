package types

import "strconv"

type TestRequest struct {
	ID       string `json:"id"`
	Name     string `json:"name"`
	Quantity uint   `json:"quantity"`
}

// Test Specific for sql
type Test struct {
	ID       uint `gorm:"primarykey"`
	Name     string
	Quantity uint
}

func (t *Test) ToTestRequest() TestRequest {
	return TestRequest{
		ID:       strconv.Itoa(int(t.ID)),
		Name:     t.Name,
		Quantity: t.Quantity,
	}
}

func (t *TestRequest) ToTest() Test {
	return Test{
		Name:     t.Name,
		Quantity: t.Quantity,
	}
}
