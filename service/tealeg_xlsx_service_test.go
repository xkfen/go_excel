package service

import "testing"

func TestCreateTealegXlsx(t *testing.T) {
	oneStepTest(func() {
		CreateTealegXlsx()
	})
}