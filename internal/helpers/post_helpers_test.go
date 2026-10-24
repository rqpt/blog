package helpers

import "testing"

func Test_ParseID(t *testing.T) {
	tests := []struct {
		name string // description of this test case
		// Named input parameters for target function.
		slug    string
		want    int64
		wantErr bool
	}{
		{"Slugs with no hyphens returns a 0 and an error", "nohyphen", 0, true},
		{"Slugs where the last segment is non-numeric returns 0 and an error", "hyphen-sixteen", 0, true},
		{"Slugs where the last segment is numeric returns ID", "hyphen-myphen-16", 16, false},
		{"Slugs with only a numeric segment returns the ID", "16", 16, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, gotErr := ParseID(tt.slug)
			if gotErr != nil {
				if !tt.wantErr {
					t.Errorf("ParseID() failed: %v", gotErr)
				}
				return
			}
			if tt.wantErr {
				t.Fatal("ParseID() succeeded unexpectedly")
			}
			if got != tt.want {
				t.Errorf("ParseID() = %v, want %v", got, tt.want)
			}
		})
	}
}
