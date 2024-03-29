package timestamp

import (
	"fmt"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

func TestGenerateTimeObj_Unix_Valid(t *testing.T) {
	unixTime := int64(1589735673000)
	unixStr := fmt.Sprintf("%d", unixTime)

	expectedTimeObj := time.Unix(unixTime/1000, 0)
	actualTimeObj, err := generateTimeObj(unixStr)

	require.NoError(t, err)
	require.Equal(t, expectedTimeObj, actualTimeObj)
}

func TestGenerateTimeObj_DateStr_Valid(t *testing.T) {
	dateTime := "2020-12-25"

	expectedTimeObj, err := time.Parse(dateTimeLayoutInput, dateTime)
	require.NoError(t, err)

	actualTimeObj, err := generateTimeObj(dateTime)

	require.NoError(t, err)
	require.Equal(t, expectedTimeObj, actualTimeObj)
}

func TestGenerateTimeObj_NoInput_Valid(t *testing.T) {
	expectedTime := time.Now()
	actualTime, err := generateTimeObj("")

	require.NoError(t, err)

	require.Equal(t, expectedTime.Format(dateTimeLayoutOutput), actualTime.Format(dateTimeLayoutOutput))
}

func TestGenerateTimeObj_ParseError(t *testing.T) {
	dateTime := "2030040f-f9a9a"

	_, err := generateTimeObj(dateTime)

	require.Error(t, err)
}

func TestParseTime_InvalidDatetime_Errors(t *testing.T) {
	timeService := NewService()
	dateTime := "kdkafkjak"
	_, _, err := timeService.ParseTime(dateTime)
	require.Error(t, err)
}

func TestParseTime_ValidUnix_OK(t *testing.T) {
	timeService := NewService()
	unixTime := int64(10001664000)
	unixTimeStr := fmt.Sprintf("%d", unixTime)
	expectedTime := time.Unix(unixTime/1000, 0)
	expectedDateTimeStr := expectedTime.Format(dateTimeLayoutOutput)

	unixMS, dateTimeUTC, err := timeService.ParseTime(unixTimeStr)
	require.NoError(t, err)
	require.Equal(t, unixTime, unixMS)
	require.Equal(t, dateTimeUTC, expectedDateTimeStr)
}
