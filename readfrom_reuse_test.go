// Copyright (c) Roman Atachiants and contributors. All rights reserved.
// Licensed under the MIT license. See LICENSE file in the project root for details.

package bitmap_test

import (
	"bytes"
	"fmt"
	"io"
	"reflect"
	"testing"

	"github.com/kelindar/bitmap"
)

func TestReadFromReusedDestinationFrame(t *testing.T) {
	for _, initialWords := range []int{0, 1, 4, 8} {
		for _, sourceWords := range []int{0, 1, 2} {
			t.Run(fmt.Sprintf("destination_%d/source_%d", initialWords, sourceWords), func(t *testing.T) {
				initial := make(bitmap.Bitmap, initialWords)
				for i := range initial {
					initial[i] = ^uint64(0)
				}
				source := make(bitmap.Bitmap, sourceWords)
				for i := range source {
					source[i] = uint64(i + 1)
				}
				next := bitmap.Bitmap{0x12, 0x34}
				var stream bytes.Buffer
				n, err := source.WriteTo(&stream)
				if err != nil {
					t.Fatal(err)
				}
				if _, err := next.WriteTo(&stream); err != nil {
					t.Fatal(err)
				}
				gotN, err := initial.ReadFrom(&stream)
				if err != nil {
					t.Fatalf("ReadFrom first frame: %v", err)
				}
				if gotN != n {
					t.Fatalf("ReadFrom consumed %d bytes, want %d", gotN, n)
				}
				if !reflect.DeepEqual(append(bitmap.Bitmap{}, initial...), append(bitmap.Bitmap{}, source...)) {
					t.Fatalf("first value=%v, want %v", initial, source)
				}
				gotNext, err := bitmap.ReadFrom(&stream)
				if err != nil {
					t.Fatalf("ReadFrom next frame: %v", err)
				}
				if !reflect.DeepEqual(gotNext, next) {
					t.Fatalf("next value=%v, want %v", gotNext, next)
				}
				if _, err := stream.ReadByte(); err != io.EOF {
					t.Fatalf("stream not exhausted: %v", err)
				}
			})
		}
	}
}
