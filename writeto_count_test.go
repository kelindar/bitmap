// Copyright (c) Roman Atachiants and contributors. All rights reserved.
// Licensed under the MIT license. See LICENSE file in the project root for details.

package bitmap_test

import (
	"errors"
	"fmt"
	"io"
	"testing"

	"github.com/kelindar/bitmap"
)

func TestWriteToErrorCountIncludesHeader(t *testing.T) {
	for _, limit := range []int{0, 2, 4, 5, 8, 11, 12} {
		t.Run(fmt.Sprintf("peer_reads_%d", limit), func(t *testing.T) {
			reader, writer := io.Pipe()
			defer reader.Close()
			defer writer.Close()
			failure := errors.New("peer stopped reading")
			received := make(chan int, 1)
			go func() {
				n, _ := io.ReadFull(reader, make([]byte, limit))
				reader.CloseWithError(failure)
				received <- n
			}()
			source := bitmap.Bitmap{0x1234}
			n, err := source.WriteTo(writer)
			actual := <-received
			if n != int64(actual) {
				t.Fatalf("WriteTo reports %d bytes, peer read %d", n, actual)
			}
			if limit < 12 && !errors.Is(err, failure) {
				t.Fatalf("error=%v, want peer error", err)
			}
			if limit == 12 && err != nil {
				t.Fatalf("successful write error=%v", err)
			}
		})
	}
}
