package helpers

import "github.com/sirupsen/logrus"

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
