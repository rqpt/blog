package slug

import "testing"

func Test_Parse(t *testing.T) {
	tests := []struct {
		name string // description of this test case
		// Named input parameters for target function.
		slug  string
		want  string
		want2 int64
	}{
		{"Slugs with no hyphens returns a blank string and 0", "nohyphen", "", 0},
		{"Slugs where the last segment is non-numeric returns a blank title and 0", "hyphen-sixteen", "", 0},
		{"Slugs where the last segment is numeric returns the title and ID", "hyphen-myphen-16", "hyphen-myphen", 16},
		{"Slugs with only a numeric segment returns a blank title and the ID", "16", "", 16},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got2 := Parse(tt.slug)
			if got != tt.want {
				t.Errorf("getTitleAndIDFromSlug() = %v, want %v", got, tt.want)
			}
			if got2 != tt.want2 {
				t.Errorf("getTitleAndIDFromSlug() = %v, want %v", got2, tt.want2)
			}
		})
	}
}

func Test_Make(t *testing.T) {
	tests := []struct {
		name string // description of this test case
		// Named input parameters for target function.
		input   string
		want    string
		wantErr bool
	}{
		{
			"Can make slug from a string",
			"  Water is wet? $#*  ",
			"water-is-wet",
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, gotErr := Make(tt.input)
			if gotErr != nil {
				if !tt.wantErr {
					t.Errorf("makeSlug() failed: %v", gotErr)
				}
				return
			}
			if tt.wantErr {
				t.Fatal("makeSlug() succeeded unexpectedly")
			}
			if got != tt.want {
				t.Errorf("makeSlug() = %v, want %v", got, tt.want)
			}
		})
	}
}
