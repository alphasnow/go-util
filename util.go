package util

import "github.com/alphasnow/go-util/v2/internal"

func Version() string {
	return "v" + internal.Version()
}
