package main

import (
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

func TestIsBusinessDay(t *testing.T) {
	testedDate := time.Date(2022, 11, 2, 11, 11, 11, 11, time.Now().Location())
	require.Equal(t, IsBusinessDay(testedDate), true)
}

func TestNextBusinessDay1(t *testing.T) {
	testedDate := time.Date(2022, 11, 2, 11, 11, 11, 11, time.Now().Location())
	expected := time.Date(2022, 11, 3, 11, 11, 11, 11, time.Now().Location())
	require.Equal(t, NextBusinessDay(testedDate), expected)
}

func TestNextBusinessDay2(t *testing.T) {
	testedDate := time.Date(2022, 11, 3, 11, 11, 11, 11, time.Now().Location())
	expected := time.Date(2022, 11, 6, 11, 11, 11, 11, time.Now().Location())
	require.Equal(t, NextBusinessDay(testedDate), expected)
}
