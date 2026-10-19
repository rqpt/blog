package handlers

import "testing"

func Test_getTitleAndIDFromSlug(t *testing.T) {
	tests := []struct {
		name string // description of this test case
		// Named input parameters for target function.
		slug  string
		want  string
		want2 int64
	}{
		{"Slugs with no hyphens returns a blank string and 0", "nohyphen", "", 0},
		{"Slugs where the last segment is non-numeric returns a blank string and 0", "hyphen-sixteen", "", 0},
		{"Slugs where the last segment is numeric returns the title and ID", "hyphen-myphen-16", "hyphen-myphen", 16},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got2 := getTitleAndIDFromSlug(tt.slug)
			if got != tt.want {
				t.Errorf("getTitleAndIDFromSlug() = %v, want %v", got, tt.want)
			}
			if got2 != tt.want2 {
				t.Errorf("getTitleAndIDFromSlug() = %v, want %v", got2, tt.want2)
			}
		})
	}
}
