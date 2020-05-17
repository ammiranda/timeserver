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

	expectedTimeObj := time.Unix(unixTime, 0)
	actualTimeObj, err := generateTimeObj(unixStr)

	require.NoError(t, err)
	require.Equal(t, expectedTimeObj, actualTimeObj)
}

func TestGenerateTimeObj_DateStr_Valid(t *testing.T) {
	dateTime := "2020-12-25"

	expectedTimeObj, err := time.Parse("2006-01-02", dateTime)
	require.NoError(t, err)

	actualTimeObj, err := generateTimeObj(dateTime)

	require.NoError(t, err)
	require.Equal(t, expectedTimeObj, actualTimeObj)
}