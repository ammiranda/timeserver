package timestamp

import (
	"errors"
	"strconv"
	"time"
)

const (
	dateTimeLayoutInput = "2006-01-02"
	dateTimeLayoutOutput = "Mon, 2 Jan 2006 15:04:05 MST"
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
	dateUTC := timeObj.Format(dateTimeLayoutOutput)

	return unixSecs, dateUTC, nil
}

func generateTimeObj(s string) (time.Time, error) {
	if s == "" {
		return time.Now(), nil
	}
	unixSecs, err := strconv.ParseInt(s, 10, 64)
	if err == nil {
		return time.Unix(unixSecs, 0), nil
	}

	timeObj, err := time.Parse(dateTimeLayoutInput, s)
	if err != nil {
		return time.Now(), errors.New("String not valid time value")
	}

	return timeObj, nil
}
