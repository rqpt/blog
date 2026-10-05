package converter

import (
	"bytes"
	"testing"
)

func Test_mdToHtml(t *testing.T) {
	md := "# header\n\n Sample text.\n\n - apple\n - pear\n - banana"
	html := `<h1 id="header">header</h1>

<p>Sample text.</p>

<ul>
<li>apple</li>
<li>pear</li>
<li>banana</li>
</ul>
`

	tests := []struct {
		name string // description of this test case
		// Named input parameters for target function.
		md   []byte
		want []byte
	}{
		{name: "Tests markdown converts to HTML", md: []byte(md), want: []byte(html)},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := MdToHtml(tt.md)
			if !bytes.Equal(got, tt.want) {
				t.Errorf("mdToHtml() = %v, want %v", string(got), string(tt.want))
			}
		})
	}
}
