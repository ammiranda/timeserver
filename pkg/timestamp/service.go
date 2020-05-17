package timestamp

import (
	"errors"
	"regexp"
	"time"
)

type Service interface {
	ParseTime(string) (int, string, error)
}

type service struct {}

func NewService() Service {
	return &service{}
}

func (s *service) ParseTime(t string) (int, string, error) {
	isDate, _ := checkDate(t)
	if isDate {

	}

	isUnix, _ := checkUnix(t)
	if isUnix {
		int64(time.Nanosecond) * timeT.UnixNano() / int64(time.Millisecond)
	}
}

func checkDate(s string) (bool, error) {
	re, err := regexp.Compile("(0?[1-9]|[12][0-9]|3[01])/(0?[1-9]|1[012])/((19|20)\\d\\d)")
	if err != nil {
		return false, errors.New("regex compile failed")
	}

	return re.MatchString(s), nil
}

func checkUnix(s string) (bool, error) {
	re, err := regexp.Compile("(\\d{13})?")
	if err != nil {
		return false, errors.New("regex compile failed")
	}

	return re.MatchString(s), nil
}
