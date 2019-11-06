package sort

import (
	"fmt"
	"os"
	"sort"
	"testing"
	"text/tabwriter"
	"time"
)

func length(s string) time.Duration {
	d, err := time.ParseDuration(s)
	if err != nil {
		panic(s)
	}
	return d
}

func printTracks(tracks []*Track) {
	const format = "%v\t%v\t%v\t%v\t%v\t\n"
	tw := new(tabwriter.Writer).Init(os.Stdout, 0, 8, 2, ' ', 0)
	fmt.Fprintf(tw, format, "Title", "Artist", "Album", "Year", "Length")
	fmt.Fprintf(tw, format, "-----", "------", "-----", "----", "------")
	for _, t := range tracks {
		fmt.Fprintf(tw, format, t.Title, t.Artist, t.Album, t.Year, t.Length)
	}
	tw.Flush() // calculate column widths and print table")
	fmt.Println()
}

var tracks = []*Track{
	{"Go", "Delilah", "From the Roots Up", 2012, length("3m38s")},
	{"Go", "Moby", "Moby", 1992, length("3m37s")},
	{"Go Ahead", "Alicia Keys", "As I Am", 2007, length("4m36s")},
	{"Ready 2 Go", "Martin Solveig", "Smash", 2011, length("4m24s")},
}

var sortable = MuliFieldsSortable{Fields: []string{}, Tracks: tracks}

func sortBy(s string) {
	sortable.Fields = append([]string{s}, sortable.Fields...)
	if len(sortable.Fields) > 5 {
		sortable.Fields = sortable.Fields[:5]
	}

	fmt.Println(sortable.Fields)

	sort.Sort(sortable)
	printTracks(sortable.Tracks)
}

// TestSort xxx
func TestSort(t *testing.T) {
	sortBy("Title")
	sortBy("Artist")
	sortBy("Album")
	sortBy("Year")
	sortBy("Length")
}
