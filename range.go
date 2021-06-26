// Copyright (c) Roman Atachiants and contributors. All rights reserved.
// Licensed under the MIT license. See LICENSE file in the project root for details.

package bitmap

import "unsafe"

// Range iterates over the bitmap elements. If the callback returns false it halts
// the iteration.
func (dst Bitmap) Range(f func(x uint32) bool) {
	for blkAt := 0; blkAt < len(dst); blkAt++ {
		blk := (dst)[blkAt]
		offset := uint32(blkAt << 6)

		if blk&0x1 != 0 && !f(uint32(offset+0)) {
			return
		}
		if blk&0x2 != 0 && !f(uint32(offset+1)) {
			return
		}
		if blk&0x4 != 0 && !f(uint32(offset+2)) {
			return
		}
		if blk&0x8 != 0 && !f(uint32(offset+3)) {
			return
		}
		if blk&0x10 != 0 && !f(uint32(offset+4)) {
			return
		}
		if blk&0x20 != 0 && !f(uint32(offset+5)) {
			return
		}
		if blk&0x40 != 0 && !f(uint32(offset+6)) {
			return
		}
		if blk&0x80 != 0 && !f(uint32(offset+7)) {
			return
		}
		if blk&0x100 != 0 && !f(uint32(offset+8)) {
			return
		}
		if blk&0x200 != 0 && !f(uint32(offset+9)) {
			return
		}
		if blk&0x400 != 0 && !f(uint32(offset+10)) {
			return
		}
		if blk&0x800 != 0 && !f(uint32(offset+11)) {
			return
		}
		if blk&0x1000 != 0 && !f(uint32(offset+12)) {
			return
		}
		if blk&0x2000 != 0 && !f(uint32(offset+13)) {
			return
		}
		if blk&0x4000 != 0 && !f(uint32(offset+14)) {
			return
		}
		if blk&0x8000 != 0 && !f(uint32(offset+15)) {
			return
		}
		if blk&0x10000 != 0 && !f(uint32(offset+16)) {
			return
		}
		if blk&0x20000 != 0 && !f(uint32(offset+17)) {
			return
		}
		if blk&0x40000 != 0 && !f(uint32(offset+18)) {
			return
		}
		if blk&0x80000 != 0 && !f(uint32(offset+19)) {
			return
		}
		if blk&0x100000 != 0 && !f(uint32(offset+20)) {
			return
		}
		if blk&0x200000 != 0 && !f(uint32(offset+21)) {
			return
		}
		if blk&0x400000 != 0 && !f(uint32(offset+22)) {
			return
		}
		if blk&0x800000 != 0 && !f(uint32(offset+23)) {
			return
		}
		if blk&0x1000000 != 0 && !f(uint32(offset+24)) {
			return
		}
		if blk&0x2000000 != 0 && !f(uint32(offset+25)) {
			return
		}
		if blk&0x4000000 != 0 && !f(uint32(offset+26)) {
			return
		}
		if blk&0x8000000 != 0 && !f(uint32(offset+27)) {
			return
		}
		if blk&0x10000000 != 0 && !f(uint32(offset+28)) {
			return
		}
		if blk&0x20000000 != 0 && !f(uint32(offset+29)) {
			return
		}
		if blk&0x40000000 != 0 && !f(uint32(offset+30)) {
			return
		}
		if blk&0x80000000 != 0 && !f(uint32(offset+31)) {
			return
		}
		if blk&0x100000000 != 0 && !f(uint32(offset+32)) {
			return
		}
		if blk&0x200000000 != 0 && !f(uint32(offset+33)) {
			return
		}
		if blk&0x400000000 != 0 && !f(uint32(offset+34)) {
			return
		}
		if blk&0x800000000 != 0 && !f(uint32(offset+35)) {
			return
		}
		if blk&0x1000000000 != 0 && !f(uint32(offset+36)) {
			return
		}
		if blk&0x2000000000 != 0 && !f(uint32(offset+37)) {
			return
		}
		if blk&0x4000000000 != 0 && !f(uint32(offset+38)) {
			return
		}
		if blk&0x8000000000 != 0 && !f(uint32(offset+39)) {
			return
		}
		if blk&0x10000000000 != 0 && !f(uint32(offset+40)) {
			return
		}
		if blk&0x20000000000 != 0 && !f(uint32(offset+41)) {
			return
		}
		if blk&0x40000000000 != 0 && !f(uint32(offset+42)) {
			return
		}
		if blk&0x80000000000 != 0 && !f(uint32(offset+43)) {
			return
		}
		if blk&0x100000000000 != 0 && !f(uint32(offset+44)) {
			return
		}
		if blk&0x200000000000 != 0 && !f(uint32(offset+45)) {
			return
		}
		if blk&0x400000000000 != 0 && !f(uint32(offset+46)) {
			return
		}
		if blk&0x800000000000 != 0 && !f(uint32(offset+47)) {
			return
		}
		if blk&0x1000000000000 != 0 && !f(uint32(offset+48)) {
			return
		}
		if blk&0x2000000000000 != 0 && !f(uint32(offset+49)) {
			return
		}
		if blk&0x4000000000000 != 0 && !f(uint32(offset+50)) {
			return
		}
		if blk&0x8000000000000 != 0 && !f(uint32(offset+51)) {
			return
		}
		if blk&0x10000000000000 != 0 && !f(uint32(offset+52)) {
			return
		}
		if blk&0x20000000000000 != 0 && !f(uint32(offset+53)) {
			return
		}
		if blk&0x40000000000000 != 0 && !f(uint32(offset+54)) {
			return
		}
		if blk&0x80000000000000 != 0 && !f(uint32(offset+55)) {
			return
		}
		if blk&0x100000000000000 != 0 && !f(uint32(offset+56)) {
			return
		}
		if blk&0x200000000000000 != 0 && !f(uint32(offset+57)) {
			return
		}
		if blk&0x400000000000000 != 0 && !f(uint32(offset+58)) {
			return
		}
		if blk&0x800000000000000 != 0 && !f(uint32(offset+59)) {
			return
		}
		if blk&0x1000000000000000 != 0 && !f(uint32(offset+60)) {
			return
		}
		if blk&0x2000000000000000 != 0 && !f(uint32(offset+61)) {
			return
		}
		if blk&0x4000000000000000 != 0 && !f(uint32(offset+62)) {
			return
		}
		if blk&0x8000000000000000 != 0 && !f(uint32(offset+63)) {
			return
		}
	}
}

// Filter iterates over the bitmap elements and calls a predicate provided for each
// containing element. If the predicate returns false, the bitmap at the element's
// position is set to zero.
func (dst *Bitmap) Filter(f func(x uint32) bool) {
	var page [64]bool
	ptr := (*[64]byte)(unsafe.Pointer(&page))

	for blkAt := 0; blkAt < len(*dst); blkAt++ {
		blk := (*dst)[blkAt]
		offset := uint32(blkAt << 6)

		page[0] = blk&0x1 != 0 && !f(offset+0)
		page[1] = blk&0x2 != 0 && !f(offset+1)
		page[2] = blk&0x4 != 0 && !f(offset+2)
		page[3] = blk&0x8 != 0 && !f(offset+3)
		page[4] = blk&0x10 != 0 && !f(offset+4)
		page[5] = blk&0x20 != 0 && !f(offset+5)
		page[6] = blk&0x40 != 0 && !f(offset+6)
		page[7] = blk&0x80 != 0 && !f(offset+7)
		page[8] = blk&0x100 != 0 && !f(offset+8)
		page[9] = blk&0x200 != 0 && !f(offset+9)
		page[10] = blk&0x400 != 0 && !f(offset+10)
		page[11] = blk&0x800 != 0 && !f(offset+11)
		page[12] = blk&0x1000 != 0 && !f(offset+12)
		page[13] = blk&0x2000 != 0 && !f(offset+13)
		page[14] = blk&0x4000 != 0 && !f(offset+14)
		page[15] = blk&0x8000 != 0 && !f(offset+15)
		page[16] = blk&0x10000 != 0 && !f(offset+16)
		page[17] = blk&0x20000 != 0 && !f(offset+17)
		page[18] = blk&0x40000 != 0 && !f(offset+18)
		page[19] = blk&0x80000 != 0 && !f(offset+19)
		page[20] = blk&0x100000 != 0 && !f(offset+20)
		page[21] = blk&0x200000 != 0 && !f(offset+21)
		page[22] = blk&0x400000 != 0 && !f(offset+22)
		page[23] = blk&0x800000 != 0 && !f(offset+23)
		page[24] = blk&0x1000000 != 0 && !f(offset+24)
		page[25] = blk&0x2000000 != 0 && !f(offset+25)
		page[26] = blk&0x4000000 != 0 && !f(offset+26)
		page[27] = blk&0x8000000 != 0 && !f(offset+27)
		page[28] = blk&0x10000000 != 0 && !f(offset+28)
		page[29] = blk&0x20000000 != 0 && !f(offset+29)
		page[30] = blk&0x40000000 != 0 && !f(offset+30)
		page[31] = blk&0x80000000 != 0 && !f(offset+31)
		page[32] = blk&0x100000000 != 0 && !f(offset+32)
		page[33] = blk&0x200000000 != 0 && !f(offset+33)
		page[34] = blk&0x400000000 != 0 && !f(offset+34)
		page[35] = blk&0x800000000 != 0 && !f(offset+35)
		page[36] = blk&0x1000000000 != 0 && !f(offset+36)
		page[37] = blk&0x2000000000 != 0 && !f(offset+37)
		page[38] = blk&0x4000000000 != 0 && !f(offset+38)
		page[39] = blk&0x8000000000 != 0 && !f(offset+39)
		page[40] = blk&0x10000000000 != 0 && !f(offset+40)
		page[41] = blk&0x20000000000 != 0 && !f(offset+41)
		page[42] = blk&0x40000000000 != 0 && !f(offset+42)
		page[43] = blk&0x80000000000 != 0 && !f(offset+43)
		page[44] = blk&0x100000000000 != 0 && !f(offset+44)
		page[45] = blk&0x200000000000 != 0 && !f(offset+45)
		page[46] = blk&0x400000000000 != 0 && !f(offset+46)
		page[47] = blk&0x800000000000 != 0 && !f(offset+47)
		page[48] = blk&0x1000000000000 != 0 && !f(offset+48)
		page[49] = blk&0x2000000000000 != 0 && !f(offset+49)
		page[50] = blk&0x4000000000000 != 0 && !f(offset+50)
		page[51] = blk&0x8000000000000 != 0 && !f(offset+51)
		page[52] = blk&0x10000000000000 != 0 && !f(offset+52)
		page[53] = blk&0x20000000000000 != 0 && !f(offset+53)
		page[54] = blk&0x40000000000000 != 0 && !f(offset+54)
		page[55] = blk&0x80000000000000 != 0 && !f(offset+55)
		page[56] = blk&0x100000000000000 != 0 && !f(offset+56)
		page[57] = blk&0x200000000000000 != 0 && !f(offset+57)
		page[58] = blk&0x400000000000000 != 0 && !f(offset+58)
		page[59] = blk&0x800000000000000 != 0 && !f(offset+59)
		page[60] = blk&0x1000000000000000 != 0 && !f(offset+60)
		page[61] = blk&0x2000000000000000 != 0 && !f(offset+61)
		page[62] = blk&0x4000000000000000 != 0 && !f(offset+62)
		page[63] = blk&0x8000000000000000 != 0 && !f(offset+63)

		var mask uint64
		mask |= uint64((*ptr)[0]) << 0
		mask |= uint64((*ptr)[1]) << 1
		mask |= uint64((*ptr)[2]) << 2
		mask |= uint64((*ptr)[3]) << 3
		mask |= uint64((*ptr)[4]) << 4
		mask |= uint64((*ptr)[5]) << 5
		mask |= uint64((*ptr)[6]) << 6
		mask |= uint64((*ptr)[7]) << 7
		mask |= uint64((*ptr)[8]) << 8
		mask |= uint64((*ptr)[9]) << 9
		mask |= uint64((*ptr)[10]) << 10
		mask |= uint64((*ptr)[11]) << 11
		mask |= uint64((*ptr)[12]) << 12
		mask |= uint64((*ptr)[13]) << 13
		mask |= uint64((*ptr)[14]) << 14
		mask |= uint64((*ptr)[15]) << 15
		mask |= uint64((*ptr)[16]) << 16
		mask |= uint64((*ptr)[17]) << 17
		mask |= uint64((*ptr)[18]) << 18
		mask |= uint64((*ptr)[19]) << 19
		mask |= uint64((*ptr)[20]) << 20
		mask |= uint64((*ptr)[21]) << 21
		mask |= uint64((*ptr)[22]) << 22
		mask |= uint64((*ptr)[23]) << 23
		mask |= uint64((*ptr)[24]) << 24
		mask |= uint64((*ptr)[25]) << 25
		mask |= uint64((*ptr)[26]) << 26
		mask |= uint64((*ptr)[27]) << 27
		mask |= uint64((*ptr)[28]) << 28
		mask |= uint64((*ptr)[29]) << 29
		mask |= uint64((*ptr)[30]) << 30
		mask |= uint64((*ptr)[31]) << 31
		mask |= uint64((*ptr)[32]) << 32
		mask |= uint64((*ptr)[33]) << 33
		mask |= uint64((*ptr)[34]) << 34
		mask |= uint64((*ptr)[35]) << 35
		mask |= uint64((*ptr)[36]) << 36
		mask |= uint64((*ptr)[37]) << 37
		mask |= uint64((*ptr)[38]) << 38
		mask |= uint64((*ptr)[39]) << 39
		mask |= uint64((*ptr)[40]) << 40
		mask |= uint64((*ptr)[41]) << 41
		mask |= uint64((*ptr)[42]) << 42
		mask |= uint64((*ptr)[43]) << 43
		mask |= uint64((*ptr)[44]) << 44
		mask |= uint64((*ptr)[45]) << 45
		mask |= uint64((*ptr)[46]) << 46
		mask |= uint64((*ptr)[47]) << 47
		mask |= uint64((*ptr)[48]) << 48
		mask |= uint64((*ptr)[49]) << 49
		mask |= uint64((*ptr)[50]) << 50
		mask |= uint64((*ptr)[51]) << 51
		mask |= uint64((*ptr)[52]) << 52
		mask |= uint64((*ptr)[53]) << 53
		mask |= uint64((*ptr)[54]) << 54
		mask |= uint64((*ptr)[55]) << 55
		mask |= uint64((*ptr)[56]) << 56
		mask |= uint64((*ptr)[57]) << 57
		mask |= uint64((*ptr)[58]) << 58
		mask |= uint64((*ptr)[59]) << 59
		mask |= uint64((*ptr)[60]) << 60
		mask |= uint64((*ptr)[61]) << 61
		mask |= uint64((*ptr)[62]) << 62
		mask |= uint64((*ptr)[63]) << 63
		(*dst)[blkAt] &^= mask
	}
}
