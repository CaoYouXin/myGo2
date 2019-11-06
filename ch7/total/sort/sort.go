package sort

import "time"

// Track represents a digit book entity
type Track struct {
	Title  string
	Artist string
	Album  string
	Year   int
	Length time.Duration
}

// MuliFieldsSortable represents a sortable book list
// witch sorts by multiple fields
type MuliFieldsSortable struct {
	Fields []string
	Tracks []*Track
}

// Len xxx
func (s MuliFieldsSortable) Len() int { return len(s.Tracks) }

// Less xxx
func (s MuliFieldsSortable) Less(i, j int) bool {
	for _, f := range s.Fields {
		switch f {
		case "Title":
			if s.Tracks[i].Title != s.Tracks[j].Title {
				return s.Tracks[i].Title < s.Tracks[j].Title
			}
		case "Artist":
			if s.Tracks[i].Artist != s.Tracks[j].Artist {
				return s.Tracks[i].Artist < s.Tracks[j].Artist
			}
		case "Album":
			if s.Tracks[i].Album != s.Tracks[j].Album {
				return s.Tracks[i].Album < s.Tracks[j].Album
			}
		case "Year":
			if s.Tracks[i].Year != s.Tracks[j].Year {
				return s.Tracks[i].Year < s.Tracks[j].Year
			}
		case "Length":
			if s.Tracks[i].Length != s.Tracks[j].Length {
				return s.Tracks[i].Length < s.Tracks[j].Length
			}
		}
	}
	return false
}

// Swap xxx
func (s MuliFieldsSortable) Swap(i, j int) {
	s.Tracks[i], s.Tracks[j] = s.Tracks[j], s.Tracks[i]
}
