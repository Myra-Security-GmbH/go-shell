// +build !testing

package authentication

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestBuild(t *testing.T) {
	s := &MyraSignature{
		Secret: "freeSecret",
	}

	ret, err := s.Build(
		"",
		"GET",
		"/en/rapi/redirects/www.example.com/1",
		"application/json",
		"2013-08-11T00:00:00 +0100",
	)

	require.Nil(t, err)
	require.Equal(
		t,
		"uLcrwObs7b0osHlejSPr1Cz4Lv97vobb2hGref2HWhh"+
			"C+2YIpswpXwViFcEaEJqo8L4Nte9XKZhvSddsuJIRMw==",
		ret,
		"Invalid",
	)
}

func TestBuildContent(t *testing.T) {
	s := &MyraSignature{
		Secret: "freeSecret",
	}

	ret, err := s.Build(
		"{ \"id\": 1, \"modified\" : \"2014-08-11T10:00:00 +0100\" }",
		"POST",
		"/en/rapi/redirects/www.example.com/1",
		"application/json",
		"2013-08-11T00:00:00 +0100",
	)

	require.Nil(t, err)
	require.Equal(
		t,
		"lTx/27qrrsqOJXzgS3VqdEYbCHRqExAE9jP+Ujgj"+
			"XZ4GxI6vxrLeB4GV6ZfnJvDG23g4cNPAnlVCxzICWilQkw==",
		ret,
		"Invalid",
	)
}
