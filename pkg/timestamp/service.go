package timestamp

import (
	"errors"
	"strconv"
	"time"
)

type Service interface {
	ParseTime(string) (int64, string, error)
}

type service struct {}

func NewService() Service {
	return &service{}
}

func (s *service) ParseTime(t string) (int64, string, error) {
	timeObj, err := generateTimeObj(t)
	if err != nil {
		return 0, "", errors.New(err.Error())
	}

	unixSecs := timeObj.Unix()
	dateUTC := timeObj.Format("2020-12-25")

	return unixSecs, dateUTC, nil
}

func generateTimeObj(s string) (time.Time, error) {
	unixSecs, err := strconv.ParseInt(s, 10, 64)
	if err == nil {
		return time.Unix(unixSecs, 0), nil
	}

	timeObj, err := time.Parse("2020-05-17", s)
	if err != nil {
		return time.Now(), errors.New("String not valid time value")
	}

	return timeObj, nil
}
