// [Original] Copyright (c) Valery Carey, Adrian Witas and contributors. All rights reserved.
// see: https://github.com/viant/vec/tree/main/bits
// [Modified] Copyright (c) Roman Atachiants and contributors. All rights reserved.
// Licensed under the APACHE-2.0 license. See LICENSE file in the project root for details.

#include <arm_neon.h>

extern "C" uint64_t arm64count(uint8_t *input, uint64_t size) {
    uint64_t sum = 0;
    uint64_t offset = 0;

    while (size - offset >= 128) {
        uint8x16x4_t input0 = vld1q_u8_x4(input + offset);
        uint8x16x4_t input1 = vld1q_u8_x4(input + offset + 64);
        uint8x16_t t0 = vcntq_u8(input0.val[0]);
        uint8x16_t t1 = vcntq_u8(input0.val[1]);
        uint8x16_t t2 = vcntq_u8(input0.val[2]);
        uint8x16_t t3 = vcntq_u8(input0.val[3]);
        uint8x16_t t4 = vcntq_u8(input1.val[0]);
        uint8x16_t t5 = vcntq_u8(input1.val[1]);
        uint8x16_t t6 = vcntq_u8(input1.val[2]);
        uint8x16_t t7 = vcntq_u8(input1.val[3]);
        t0 = vaddq_u8(t0, t1);
        t2 = vaddq_u8(t2, t3);
        t4 = vaddq_u8(t4, t5);
        t6 = vaddq_u8(t6, t7);
        t0 = vaddq_u8(t0, t2);
        t4 = vaddq_u8(t4, t6);
        t0 = vaddq_u8(t0, t4);
        sum += vaddlvq_u8(t0);
        offset += 128;
    }

    while (size - offset >= 8) {
        sum += vaddv_u8(vcnt_u8(vld1_u8(input + offset)));
        offset += 8;
    }

    if (size - offset > 0) {
        int32x2_t leftover = {0, 0};
        if ((size - offset) >= 4) {
            leftover = vld1_lane_s32(input + offset, leftover, 0);
            offset += 4;
        }
        if (size - offset >= 2) {
            leftover = vld1_lane_u16(input + offset, leftover, 2);
            offset += 2;
        }
        if (size - offset > 0) {
            leftover = vld1_lane_u8(input + offset, leftover, 6);
        }
        sum += vaddv_u8(vcnt_u8(leftover));
    }
    return sum;
}
