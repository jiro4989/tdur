package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCalcTimeDuration(t *testing.T) {
	assert.Equal(t, 9*60*60, calcTimeDuration("10:00", "19:00"))
	assert.Equal(t, 0, calcTimeDuration("10:00", "10:00"))
	assert.Equal(t, 9*60*60, calcTimeDuration("0:00", "9:00"))
	assert.Equal(t, 25*60*60, calcTimeDuration("0:00", "25:00"))
	assert.Equal(t, 23*60*60, calcTimeDuration("1:00", "24:00"))
	assert.Equal(t, 9*60*60+30*60, calcTimeDuration("0:00", "9:30"))
	assert.Equal(t, -1, calcTimeDuration("1", "9:30"))
	assert.Equal(t, -1, calcTimeDuration("10:00", "1"))
	assert.Equal(t, -1, calcTimeDuration("", ""))
	assert.Equal(t, -1, calcTimeDuration("a", "b"))
	assert.Equal(t, -1, calcTimeDuration("10:00", "11:xa"))
}

func TestTimeToSecond(t *testing.T) {
	assert.Equal(t, 0, timeToSecond("00:00"))
	assert.Equal(t, 30*60, timeToSecond("00:30"))
	assert.Equal(t, 1*60*60+30*60, timeToSecond("1:30"))
	assert.Equal(t, 10*60*60, timeToSecond("10:00"))
	assert.Equal(t, 24*60*60, timeToSecond("24:00"))
	assert.Equal(t, 25*60*60, timeToSecond("25:00"))
	assert.Equal(t, -1, timeToSecond("-1"))
	assert.Equal(t, -1, timeToSecond("1"))
	assert.Equal(t, -1, timeToSecond("a"))
	assert.Equal(t, -1, timeToSecond(""))
}

func TestSecondToMinute(t *testing.T) {
	assert.Equal(t, 0.0, secondToMinute(0))
	assert.Equal(t, 1.0, secondToMinute(60))
	assert.Equal(t, 1.5, secondToMinute(90))
	assert.Equal(t, 60.0, secondToMinute(3600))
	assert.Equal(t, -1.0, secondToMinute(-1))
}

func TestSecondToHour(t *testing.T) {
	assert.Equal(t, 0.0, secondToHour(0))
	assert.Equal(t, 0.5, secondToHour(1800))
	assert.Equal(t, 1.0, secondToHour(3600))
	assert.Equal(t, 1.5, secondToHour(5400))
	assert.Equal(t, 24.0, secondToHour(24*60*60))
	assert.Equal(t, -1.0, secondToHour(-1))
}
