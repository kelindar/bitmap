// Copyright (c) Roman Atachiants and contributors. All rights reserved.
// Licensed under the MIT license. See LICENSE file in the project root for details.

#include <stdint.h>

void _and_avx512(uint64_t* a, uint64_t* b, uint64_t n) {
    #pragma clang loop vectorize(enable)
    for (uint64_t i = 0; i < n; ++i) {
        a[i] &= b[i];
    }
}

void _andn_avx512(uint64_t* a, uint64_t* b, uint64_t n) {
    #pragma clang loop vectorize(enable) interleave(enable)
    for (uint64_t i = 0; i < n; ++i) {
        a[i] &= ~b[i];
    }
}

void _or_avx512(uint64_t* a, uint64_t* b, uint64_t n) {
    #pragma clang loop vectorize(enable) interleave(enable)
    for (uint64_t i = 0; i < n; ++i) {
        a[i] |= b[i];
    }
}

void _xor_avx512(uint64_t* a, uint64_t* b, uint64_t n) {
    #pragma clang loop vectorize(enable) interleave(enable)
    for (uint64_t i = 0; i < n; ++i) {
        a[i] ^= b[i];
    }
}

void _and_many_avx512(uint64_t* a, uint64_t** b, uint64_t dims) {
    int64_t n = (dims & 0xffffffff);
    int64_t m = (dims >> 32);
    const int64_t chunk_size = 512;

    // Loop over chunks of b
    for (int64_t chunk = 0; chunk < n; chunk += chunk_size) {
        int64_t chunk_end = chunk + chunk_size;
        if (chunk_end > n) {
            chunk_end = n;
        }

        for (int64_t j = 0; j < m; ++j) {
            #pragma clang loop vectorize(enable) interleave(enable)
            for (int64_t i = chunk; i < chunk_end; ++i) {
                a[i] &= b[j][i];
            }
        }
    }
}

void _andn_many_avx512(uint64_t* a, uint64_t** b, uint64_t dims) {
    int64_t n = (dims & 0xffffffff);
    int64_t m = (dims >> 32);
    const int64_t chunk_size = 512;

    // Loop over chunks of b
    for (int64_t chunk = 0; chunk < n; chunk += chunk_size) {
        int64_t chunk_end = chunk + chunk_size;
        if (chunk_end > n) {
            chunk_end = n;
        }

        for (int64_t j = 0; j < m; ++j) {
            #pragma clang loop vectorize(enable) interleave(enable)
            for (int64_t i = chunk; i < chunk_end; ++i) {
                a[i] &= ~b[j][i];
            }
        }
    }
}

void _or_many_avx512(uint64_t* a, uint64_t** b, uint64_t dims) {
    int64_t n = (dims & 0xffffffff);
    int64_t m = (dims >> 32);
    const int64_t chunk_size = 512;

    // Loop over chunks of b
    for (int64_t chunk = 0; chunk < n; chunk += chunk_size) {
        int64_t chunk_end = chunk + chunk_size;
        if (chunk_end > n) {
            chunk_end = n;
        }

        for (int64_t j = 0; j < m; ++j) {
            #pragma clang loop vectorize(enable) interleave(enable)
            for (int64_t i = chunk; i < chunk_end; ++i) {
                a[i] |= b[j][i];
            }
        }
    }
}

void _xor_many_avx512(uint64_t* a, uint64_t** b, uint64_t dims) {
    int64_t n = (dims & 0xffffffff);
    int64_t m = (dims >> 32);
    const int64_t chunk_size = 512;

    // Loop over chunks of b
    for (int64_t chunk = 0; chunk < n; chunk += chunk_size) {
        int64_t chunk_end = chunk + chunk_size;
        if (chunk_end > n) {
            chunk_end = n;
        }

        for (int64_t j = 0; j < m; ++j) {
            #pragma clang loop vectorize(enable) interleave(enable)
            for (int64_t i = chunk; i < chunk_end; ++i) {
                a[i] ^= b[j][i];
            }
        }
    }
}
