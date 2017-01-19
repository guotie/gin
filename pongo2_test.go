package gin

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// test pongo2 template
func TestPongo2(t *testing.T) {
	ginMode = debugCode

	c, w, router := CreateTestContext()
	router.LoadHTMLGlob("./examples/pongo2/")
	c.HTML(200, "index.html", H{
		"name": "gin+pongo2",
	})
	assert.Equal(t, w.Code, 200)
	assert.Equal(t, w.Body.String(), `Hello Gin+pongo2!`)
}
