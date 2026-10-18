package handlers

import "testing"

func Test_getIDFromSlug(t *testing.T) {
	tests := []struct {
		name string // description of this test case
		// Named input parameters for target function.
		slug string
		want int64
	}{
		{"Slugs with no hyphens return 0", "slug16", 0},
		{"Non numerical IDs return 0", "slug-sixteen", 0},
		{"Numerical IDs return the ID", "slug-16", 16},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := getIDFromSlug(tt.slug)
			if got != tt.want {
				t.Errorf("getIDFromSlug() = %v, want %v", got, tt.want)
			}
		})
	}
}
