package service

import (
	"gfin/db/config"
	"gcoresys/common/logger"
	"testing"
)

// 单步测试
func oneStepTest(testMethods ...func()) {
	logger.InitLogger(logger.LvlDebug, nil)
	config.GetFinDbConfig("test")
	config.GetDb().LogMode(false)
	defer config.ClearAllData()

	for _, testMethods := range testMethods {
		config.ClearAllData()
		testMethods()
		config.ClearAllData()
	}
}

func TestCreateExcel(t *testing.T) {
	oneStepTest(func() {
		CreateExcel()
	})
}

func TestReadExcel(t *testing.T) {
	oneStepTest(func() {
		ReadExcel()
	})
}