// Copyright (c) 2020. All rights reserved.
// Use of this source code is governed by an MIT-style license that can be
// found in the LICENSE file.

package gog_types

type Media int

const (
	UnknownMedia Media = iota
	Game
	Movie
)

func (mt Media) String() string {
	switch mt {
	case Game:
		return "game"
	case Movie:
		return "movie"
	default:
		return ""
	}
}

func ParseMedia(mt string) Media {
	switch mt {
	case "game":
		return Game
	case "movie":
		return Movie
	default:
		return UnknownMedia
	}
}

func ValidMedia(mt Media) bool {
	switch mt {
	case Game:
		fallthrough
	case Movie:
		return true
	default:
		return false
	}
}
