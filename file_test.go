package crare

import (
	"io"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFile(t *testing.T) {
	f := FromDisk("crare.go")
	g := FromURL("http://")

	assert.True(t, f.OnDisk())
	assert.True(t, (&File{FileID: "1"}).InCloud())
	assert.Equal(t, File{FileLocal: "crare.go"}, f)
	assert.Equal(t, File{FileURL: "http://"}, g)
	assert.Equal(t, File{FileReader: io.Reader(nil)}, FromReader(io.Reader(nil)))

	g.stealRef(&f)
	f.stealRef(&g)
	assert.Equal(t, g.FileLocal, f.FileLocal)
	assert.Equal(t, f.FileURL, g.FileURL)
}
