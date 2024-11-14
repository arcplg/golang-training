package hourshelper

import (
	"reflect"
	"testing"
)

// Hello returns a greeting for the named person.
func TestCreate(t *testing.T) {
	want := []string{"06:00", "06:05", "06:10", "06:15", "06:20", "06:25", "06:30", "06:35", "06:40", "06:45", "06:50", "06:55", "07:00"}
	rev := Create("06:00", "07:00", 5, "15:04")
	if !reflect.DeepEqual(rev, want) {
		t.Errorf("Create: %q, want %q", rev, want)
	}
}
