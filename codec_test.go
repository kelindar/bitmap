// Copyright (c) Roman Atachiants and contributors. All rights reserved.
// Licensed under the MIT license. See LICENSE file in the project root for details.

package bitmap

import (
	"bytes"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSaveLoad(t *testing.T) {
	m := Bitmap{}
	for i := 0; i <= 5000; i += 10 {
		m.Set(uint32(i))
	}

	// Save the map
	enc := new(bytes.Buffer)
	cloned := m.Clone(nil)
	n, err := cloned.WriteTo(enc)
	assert.NoError(t, err)
	assert.Equal(t, int64(636), n)

	// Load the map back
	out, err := ReadFrom(enc)
	assert.NoError(t, err)
	assert.Equal(t, len(m), len(out))
}
