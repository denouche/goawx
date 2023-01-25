package awx

import (
	"fmt"
	"testing"
)

type MockAWX struct {
	AWX
}

func (m *MockAWX) Ping() (*Ping, error) {
	return nil, nil
}

func Test_cast(t *testing.T) {
	mockAWX := &MockAWX{}
	fmt.Println(dest(mockAWX))
}

func dest(awx AWX) error {
	awx.ListApplication(nil)
	return nil
}
