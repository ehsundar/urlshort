package composer

import "testing"
import "github.com/stretchr/testify/assert"

func TestMd5Base64_Compose(t *testing.T) {
	s := assert.New(t)
	c := NewMd5Base64()
	short := c.Compose("www.test.com", "amir")

	s.Equal("Yls7W9MvYvBSndtmcTdG", short)
}
