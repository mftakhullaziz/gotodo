package helpers

import (
	"github.com/sirupsen/logrus"
)
import "gorm.io/gorm"

func PanicIfError(err error) {
	if err != nil {
		panic(err)
	}
}

func PanicIfErrorWithCustomMessage(err error, str string) {
	if err != nil {
		panic(str)
	}
}

func FatalIfErrorWithCustomMessage(err error, log *logrus.Logger, str string) {
	if err != nil {
		log.Fatalf(str, err)
	}
}

func ErrorStructJoinUserAccountRecord(gdb *gorm.DB) {
	log := LoggerParent()
	var emptyInterface interface{}
	if gdb.Error != nil {
		log.Errorln("Error fetch gorm record: ", emptyInterface, gdb.Error)
	} else if gdb.RowsAffected == 0 {
		log.Errorln("Row affected record is zero: ", emptyInterface)
	}
}

func LoggerIfError(err error) {
	log := LoggerParent()
	if err != nil {
		log.Errorln("Logger : ", err.Error())
	}
}

func LoggerIfErrorWithCustomMessage(err error, log *logrus.Logger, str string) {
	if err != nil {
		log.Errorln(err.Error(), str)
	}
}
