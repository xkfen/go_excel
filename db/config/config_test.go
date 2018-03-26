package config

import "testing"

func TestClearAllData(t *testing.T) {
	GetUserInfoDbConfig("test")
	GetDb()
	ClearAllData()
}
