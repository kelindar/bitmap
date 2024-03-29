//go:build !noasm && amd64
// AUTO-GENERATED BY GOCC -- DO NOT EDIT

TEXT ·_and(SB), $0-32
	MOVQ a+0(FP), DI
	MOVQ b+8(FP), SI
	MOVQ n+16(FP), DX
	BYTE $0x55               // push	rbp
	WORD $0x8948; BYTE $0xe5 // mov	rbp, rsp
	LONG $0xf8e48348         // and	rsp, -8
	WORD $0x8548; BYTE $0xd2 // test	rdx, rdx
	JE   LBB0_10
	LONG $0x04fa8348         // cmp	rdx, 4
	JAE  LBB0_3
	WORD $0xc031             // xor	eax, eax
	JMP  LBB0_9

LBB0_3:
	LONG $0xd6048d48         // lea	rax, [rsi + 8*rdx]
	WORD $0x3948; BYTE $0xf8 // cmp	rax, rdi
	JBE  LBB0_6
	LONG $0xd7048d48         // lea	rax, [rdi + 8*rdx]
	WORD $0x3948; BYTE $0xf0 // cmp	rax, rsi
	JBE  LBB0_6
	WORD $0xc031             // xor	eax, eax
	JMP  LBB0_9

LBB0_6:
	WORD $0x8948; BYTE $0xd0 // mov	rax, rdx
	LONG $0xfce08348         // and	rax, -4
	WORD $0xc931             // xor	ecx, ecx

LBB0_7:
	LONG $0x0410fcc5; BYTE $0xcf // vmovups	ymm0, ymmword ptr [rdi + 8*rcx]
	LONG $0x0454fcc5; BYTE $0xce // vandps	ymm0, ymm0, ymmword ptr [rsi + 8*rcx]
	LONG $0x0411fcc5; BYTE $0xcf // vmovups	ymmword ptr [rdi + 8*rcx], ymm0
	LONG $0x04c18348             // add	rcx, 4
	WORD $0x3948; BYTE $0xc8     // cmp	rax, rcx
	JNE  LBB0_7
	WORD $0x3948; BYTE $0xd0     // cmp	rax, rdx
	JE   LBB0_10

LBB0_9:
	LONG $0xc60c8b48         // mov	rcx, qword ptr [rsi + 8*rax]
	LONG $0xc70c2148         // and	qword ptr [rdi + 8*rax], rcx
	WORD $0xff48; BYTE $0xc0 // inc	rax
	WORD $0x3948; BYTE $0xc2 // cmp	rdx, rax
	JNE  LBB0_9

LBB0_10:
	WORD $0x8948; BYTE $0xec // mov	rsp, rbp
	BYTE $0x5d               // pop	rbp
	WORD $0xf8c5; BYTE $0x77 // vzeroupper
	BYTE $0xc3               // ret

TEXT ·_andn(SB), $0-32
	MOVQ a+0(FP), DI
	MOVQ b+8(FP), SI
	MOVQ n+16(FP), DX
	BYTE $0x55               // push	rbp
	WORD $0x8948; BYTE $0xe5 // mov	rbp, rsp
	LONG $0xf8e48348         // and	rsp, -8
	WORD $0x8548; BYTE $0xd2 // test	rdx, rdx
	JE   LBB1_10
	LONG $0x04fa8348         // cmp	rdx, 4
	JAE  LBB1_3
	WORD $0xc031             // xor	eax, eax
	JMP  LBB1_9

LBB1_3:
	LONG $0xd6048d48         // lea	rax, [rsi + 8*rdx]
	WORD $0x3948; BYTE $0xf8 // cmp	rax, rdi
	JBE  LBB1_6
	LONG $0xd7048d48         // lea	rax, [rdi + 8*rdx]
	WORD $0x3948; BYTE $0xf0 // cmp	rax, rsi
	JBE  LBB1_6
	WORD $0xc031             // xor	eax, eax
	JMP  LBB1_9

LBB1_6:
	WORD $0x8948; BYTE $0xd0 // mov	rax, rdx
	LONG $0xfce08348         // and	rax, -4
	WORD $0xc931             // xor	ecx, ecx

LBB1_7:
	LONG $0x0410fcc5; BYTE $0xce // vmovups	ymm0, ymmword ptr [rsi + 8*rcx]
	LONG $0x0455fcc5; BYTE $0xcf // vandnps	ymm0, ymm0, ymmword ptr [rdi + 8*rcx]
	LONG $0x0411fcc5; BYTE $0xcf // vmovups	ymmword ptr [rdi + 8*rcx], ymm0
	LONG $0x04c18348             // add	rcx, 4
	WORD $0x3948; BYTE $0xc8     // cmp	rax, rcx
	JNE  LBB1_7
	WORD $0x3948; BYTE $0xd0     // cmp	rax, rdx
	JE   LBB1_10

LBB1_9:
	LONG $0xc60c8b48         // mov	rcx, qword ptr [rsi + 8*rax]
	WORD $0xf748; BYTE $0xd1 // not	rcx
	LONG $0xc70c2148         // and	qword ptr [rdi + 8*rax], rcx
	WORD $0xff48; BYTE $0xc0 // inc	rax
	WORD $0x3948; BYTE $0xc2 // cmp	rdx, rax
	JNE  LBB1_9

LBB1_10:
	WORD $0x8948; BYTE $0xec // mov	rsp, rbp
	BYTE $0x5d               // pop	rbp
	WORD $0xf8c5; BYTE $0x77 // vzeroupper
	BYTE $0xc3               // ret

TEXT ·_or(SB), $0-32
	MOVQ a+0(FP), DI
	MOVQ b+8(FP), SI
	MOVQ n+16(FP), DX
	BYTE $0x55               // push	rbp
	WORD $0x8948; BYTE $0xe5 // mov	rbp, rsp
	LONG $0xf8e48348         // and	rsp, -8
	WORD $0x8548; BYTE $0xd2 // test	rdx, rdx
	JE   LBB2_10
	LONG $0x04fa8348         // cmp	rdx, 4
	JAE  LBB2_3
	WORD $0xc031             // xor	eax, eax
	JMP  LBB2_9

LBB2_3:
	LONG $0xd6048d48         // lea	rax, [rsi + 8*rdx]
	WORD $0x3948; BYTE $0xf8 // cmp	rax, rdi
	JBE  LBB2_6
	LONG $0xd7048d48         // lea	rax, [rdi + 8*rdx]
	WORD $0x3948; BYTE $0xf0 // cmp	rax, rsi
	JBE  LBB2_6
	WORD $0xc031             // xor	eax, eax
	JMP  LBB2_9

LBB2_6:
	WORD $0x8948; BYTE $0xd0 // mov	rax, rdx
	LONG $0xfce08348         // and	rax, -4
	WORD $0xc931             // xor	ecx, ecx

LBB2_7:
	LONG $0x0410fcc5; BYTE $0xcf // vmovups	ymm0, ymmword ptr [rdi + 8*rcx]
	LONG $0x0456fcc5; BYTE $0xce // vorps	ymm0, ymm0, ymmword ptr [rsi + 8*rcx]
	LONG $0x0411fcc5; BYTE $0xcf // vmovups	ymmword ptr [rdi + 8*rcx], ymm0
	LONG $0x04c18348             // add	rcx, 4
	WORD $0x3948; BYTE $0xc8     // cmp	rax, rcx
	JNE  LBB2_7
	WORD $0x3948; BYTE $0xd0     // cmp	rax, rdx
	JE   LBB2_10

LBB2_9:
	LONG $0xc60c8b48         // mov	rcx, qword ptr [rsi + 8*rax]
	LONG $0xc70c0948         // or	qword ptr [rdi + 8*rax], rcx
	WORD $0xff48; BYTE $0xc0 // inc	rax
	WORD $0x3948; BYTE $0xc2 // cmp	rdx, rax
	JNE  LBB2_9

LBB2_10:
	WORD $0x8948; BYTE $0xec // mov	rsp, rbp
	BYTE $0x5d               // pop	rbp
	WORD $0xf8c5; BYTE $0x77 // vzeroupper
	BYTE $0xc3               // ret

TEXT ·_xor(SB), $0-32
	MOVQ a+0(FP), DI
	MOVQ b+8(FP), SI
	MOVQ n+16(FP), DX
	BYTE $0x55               // push	rbp
	WORD $0x8948; BYTE $0xe5 // mov	rbp, rsp
	LONG $0xf8e48348         // and	rsp, -8
	WORD $0x8548; BYTE $0xd2 // test	rdx, rdx
	JE   LBB3_10
	LONG $0x04fa8348         // cmp	rdx, 4
	JAE  LBB3_3
	WORD $0xc031             // xor	eax, eax
	JMP  LBB3_9

LBB3_3:
	LONG $0xd6048d48         // lea	rax, [rsi + 8*rdx]
	WORD $0x3948; BYTE $0xf8 // cmp	rax, rdi
	JBE  LBB3_6
	LONG $0xd7048d48         // lea	rax, [rdi + 8*rdx]
	WORD $0x3948; BYTE $0xf0 // cmp	rax, rsi
	JBE  LBB3_6
	WORD $0xc031             // xor	eax, eax
	JMP  LBB3_9

LBB3_6:
	WORD $0x8948; BYTE $0xd0 // mov	rax, rdx
	LONG $0xfce08348         // and	rax, -4
	WORD $0xc931             // xor	ecx, ecx

LBB3_7:
	LONG $0x0410fcc5; BYTE $0xcf // vmovups	ymm0, ymmword ptr [rdi + 8*rcx]
	LONG $0x0457fcc5; BYTE $0xce // vxorps	ymm0, ymm0, ymmword ptr [rsi + 8*rcx]
	LONG $0x0411fcc5; BYTE $0xcf // vmovups	ymmword ptr [rdi + 8*rcx], ymm0
	LONG $0x04c18348             // add	rcx, 4
	WORD $0x3948; BYTE $0xc8     // cmp	rax, rcx
	JNE  LBB3_7
	WORD $0x3948; BYTE $0xd0     // cmp	rax, rdx
	JE   LBB3_10

LBB3_9:
	LONG $0xc60c8b48         // mov	rcx, qword ptr [rsi + 8*rax]
	LONG $0xc70c3148         // xor	qword ptr [rdi + 8*rax], rcx
	WORD $0xff48; BYTE $0xc0 // inc	rax
	WORD $0x3948; BYTE $0xc2 // cmp	rdx, rax
	JNE  LBB3_9

LBB3_10:
	WORD $0x8948; BYTE $0xec // mov	rsp, rbp
	BYTE $0x5d               // pop	rbp
	WORD $0xf8c5; BYTE $0x77 // vzeroupper
	BYTE $0xc3               // ret

TEXT ·_and_many(SB), $0-32
	MOVQ a+0(FP), DI
	MOVQ b+8(FP), SI
	MOVQ dims+16(FP), DX
	BYTE $0x55                     // push	rbp
	WORD $0x8948; BYTE $0xe5       // mov	rbp, rsp
	WORD $0x5741                   // push	r15
	WORD $0x5641                   // push	r14
	WORD $0x5541                   // push	r13
	WORD $0x5441                   // push	r12
	BYTE $0x53                     // push	rbx
	LONG $0xf8e48348               // and	rsp, -8
	LONG $0x68ec8348               // sub	rsp, 104
	LONG $0xffffffbb; BYTE $0xff   // mov	ebx, 4294967295
	WORD $0x2148; BYTE $0xd3       // and	rbx, rdx
	JE   LBB4_14
	WORD $0x8948; BYTE $0xd0       // mov	rax, rdx
	LONG $0x20e8c148               // shr	rax, 32
	LONG $0x02f88348               // cmp	rax, 2
	LONG $0x0001bc41; WORD $0x0000 // mov	r12d, 1
	LONG $0x24448948; BYTE $0x28   // mov	qword ptr [rsp + 40], rax
	LONG $0xe0430f4c               // cmovae	r12, rax
	LONG $0x000200b9; BYTE $0x00   // mov	ecx, 512
	WORD $0xc031                   // xor	eax, eax
	LONG $0x24448948; BYTE $0x18   // mov	qword ptr [rsp + 24], rax
	WORD $0x8948; BYTE $0xf8       // mov	rax, rdi
	WORD $0xd231                   // xor	edx, edx
	LONG $0x24548948; BYTE $0x10   // mov	qword ptr [rsp + 16], rdx
	WORD $0xd231                   // xor	edx, edx
	LONG $0x24548948; BYTE $0x08   // mov	qword ptr [rsp + 8], rdx
	WORD $0x3145; BYTE $0xc9       // xor	r9d, r9d
	LONG $0x245c8948; BYTE $0x20   // mov	qword ptr [rsp + 32], rbx
	JMP  LBB4_3

LBB4_2:
	LONG $0x2444ff48; BYTE $0x08               // inc	qword ptr [rsp + 8]
	LONG $0x244c8b48; BYTE $0x38               // mov	rcx, qword ptr [rsp + 56]
	LONG $0x00c18148; WORD $0x0002; BYTE $0x00 // add	rcx, 512
	QUAD $0xfffe001024448148; BYTE $0xff       // add	qword ptr [rsp + 16], -512
	QUAD $0x0010001824448148; BYTE $0x00       // add	qword ptr [rsp + 24], 4096
	LONG $0x10000548; WORD $0x0000             // add	rax, 4096
	LONG $0x245c8b48; BYTE $0x20               // mov	rbx, qword ptr [rsp + 32]
	LONG $0x244c8b4c; BYTE $0x30               // mov	r9, qword ptr [rsp + 48]
	WORD $0x3949; BYTE $0xd9                   // cmp	r9, rbx
	JAE  LBB4_14

LBB4_3:
	WORD $0x3948; BYTE $0xd9                   // cmp	rcx, rbx
	WORD $0x8949; BYTE $0xda                   // mov	r10, rbx
	LONG $0x244c8948; BYTE $0x38               // mov	qword ptr [rsp + 56], rcx
	LONG $0xd1420f4c                           // cmovb	r10, rcx
	LONG $0x00898d49; WORD $0x0002; BYTE $0x00 // lea	rcx, [r9 + 512]
	WORD $0x3948; BYTE $0xd9                   // cmp	rcx, rbx
	LONG $0x244c8948; BYTE $0x30               // mov	qword ptr [rsp + 48], rcx
	LONG $0xd9420f48                           // cmovb	rbx, rcx
	LONG $0x247c8348; WORD $0x0028             // cmp	qword ptr [rsp + 40], 0
	JE   LBB4_2
	LONG $0x2454034c; BYTE $0x10               // add	r10, qword ptr [rsp + 16]
	LONG $0xfce28349                           // and	r10, -4
	LONG $0x247c8b4c; BYTE $0x08               // mov	r15, qword ptr [rsp + 8]
	WORD $0x894c; BYTE $0xf9                   // mov	rcx, r15
	LONG $0x09e1c148                           // shl	rcx, 9
	WORD $0x8949; BYTE $0xd8                   // mov	r8, rbx
	WORD $0x2949; BYTE $0xc8                   // sub	r8, rcx
	LONG $0x0ce7c149                           // shl	r15, 12
	LONG $0x3f1c8d4e                           // lea	r11, [rdi + r15]
	LONG $0x247c894c; BYTE $0x60               // mov	qword ptr [rsp + 96], r15
	LONG $0xc73c8d4f                           // lea	r15, [r15 + 8*r8]
	LONG $0x3f0c8d49                           // lea	rcx, [r15 + rdi]
	LONG $0x244c8948; BYTE $0x58               // mov	qword ptr [rsp + 88], rcx
	WORD $0x894c; BYTE $0xc1                   // mov	rcx, r8
	LONG $0xfce18348                           // and	rcx, -4
	LONG $0x244c8948; BYTE $0x48               // mov	qword ptr [rsp + 72], rcx
	WORD $0x014c; BYTE $0xc9                   // add	rcx, r9
	LONG $0x244c8948; BYTE $0x40               // mov	qword ptr [rsp + 64], rcx
	WORD $0x3145; BYTE $0xf6                   // xor	r14d, r14d
	LONG $0x247c894c; BYTE $0x50               // mov	qword ptr [rsp + 80], r15
	JMP  LBB4_5

LBB4_12:
	WORD $0xff49; BYTE $0xc6 // inc	r14
	WORD $0x394d; BYTE $0xe6 // cmp	r14, r12
	JE   LBB4_2

LBB4_5:
	WORD $0x3949; BYTE $0xd9     // cmp	r9, rbx
	JAE  LBB4_12
	LONG $0xf6148b4a             // mov	rdx, qword ptr [rsi + 8*r14]
	WORD $0x894d; BYTE $0xcd     // mov	r13, r9
	LONG $0x04f88349             // cmp	r8, 4
	JB   LBB4_13
	LONG $0x3a0c8d4a             // lea	rcx, [rdx + r15]
	WORD $0x3949; BYTE $0xcb     // cmp	r11, rcx
	JAE  LBB4_9
	LONG $0x244c8b48; BYTE $0x60 // mov	rcx, qword ptr [rsp + 96]
	WORD $0x0148; BYTE $0xd1     // add	rcx, rdx
	WORD $0x894d; BYTE $0xcd     // mov	r13, r9
	LONG $0x244c3b48; BYTE $0x58 // cmp	rcx, qword ptr [rsp + 88]
	JB   LBB4_13

LBB4_9:
	WORD $0x894d; BYTE $0xdf     // mov	r15, r11
	WORD $0x8949; BYTE $0xf3     // mov	r11, rsi
	LONG $0x244c8b48; BYTE $0x18 // mov	rcx, qword ptr [rsp + 24]
	LONG $0x0a348d48             // lea	rsi, [rdx + rcx]
	WORD $0xc931                 // xor	ecx, ecx

LBB4_10:
	LONG $0x0410fcc5; BYTE $0xc8 // vmovups	ymm0, ymmword ptr [rax + 8*rcx]
	LONG $0x0454fcc5; BYTE $0xce // vandps	ymm0, ymm0, ymmword ptr [rsi + 8*rcx]
	LONG $0x0411fcc5; BYTE $0xc8 // vmovups	ymmword ptr [rax + 8*rcx], ymm0
	LONG $0x04c18348             // add	rcx, 4
	WORD $0x3949; BYTE $0xca     // cmp	r10, rcx
	JNE  LBB4_10
	LONG $0x246c8b4c; BYTE $0x40 // mov	r13, qword ptr [rsp + 64]
	LONG $0x24443b4c; BYTE $0x48 // cmp	r8, qword ptr [rsp + 72]
	WORD $0x894c; BYTE $0xde     // mov	rsi, r11
	WORD $0x894d; BYTE $0xfb     // mov	r11, r15
	LONG $0x247c8b4c; BYTE $0x50 // mov	r15, qword ptr [rsp + 80]
	JE   LBB4_12

LBB4_13:
	LONG $0xea0c8b4a         // mov	rcx, qword ptr [rdx + 8*r13]
	LONG $0xef0c214a         // and	qword ptr [rdi + 8*r13], rcx
	WORD $0xff49; BYTE $0xc5 // inc	r13
	WORD $0x3949; BYTE $0xdd // cmp	r13, rbx
	JB   LBB4_13
	JMP  LBB4_12

LBB4_14:
	LONG $0xd8658d48         // lea	rsp, [rbp - 40]
	BYTE $0x5b               // pop	rbx
	WORD $0x5c41             // pop	r12
	WORD $0x5d41             // pop	r13
	WORD $0x5e41             // pop	r14
	WORD $0x5f41             // pop	r15
	BYTE $0x5d               // pop	rbp
	WORD $0xf8c5; BYTE $0x77 // vzeroupper
	BYTE $0xc3               // ret

TEXT ·_andn_many(SB), $0-32
	MOVQ a+0(FP), DI
	MOVQ b+8(FP), SI
	MOVQ dims+16(FP), DX
	BYTE $0x55                     // push	rbp
	WORD $0x8948; BYTE $0xe5       // mov	rbp, rsp
	WORD $0x5741                   // push	r15
	WORD $0x5641                   // push	r14
	WORD $0x5541                   // push	r13
	WORD $0x5441                   // push	r12
	BYTE $0x53                     // push	rbx
	LONG $0xf8e48348               // and	rsp, -8
	LONG $0x68ec8348               // sub	rsp, 104
	LONG $0xffffffbb; BYTE $0xff   // mov	ebx, 4294967295
	WORD $0x2148; BYTE $0xd3       // and	rbx, rdx
	JE   LBB5_14
	WORD $0x8948; BYTE $0xd0       // mov	rax, rdx
	LONG $0x20e8c148               // shr	rax, 32
	LONG $0x02f88348               // cmp	rax, 2
	LONG $0x0001bc41; WORD $0x0000 // mov	r12d, 1
	LONG $0x24448948; BYTE $0x28   // mov	qword ptr [rsp + 40], rax
	LONG $0xe0430f4c               // cmovae	r12, rax
	LONG $0x000200b9; BYTE $0x00   // mov	ecx, 512
	WORD $0xc031                   // xor	eax, eax
	LONG $0x24448948; BYTE $0x18   // mov	qword ptr [rsp + 24], rax
	WORD $0x8948; BYTE $0xf8       // mov	rax, rdi
	WORD $0xd231                   // xor	edx, edx
	LONG $0x24548948; BYTE $0x10   // mov	qword ptr [rsp + 16], rdx
	WORD $0xd231                   // xor	edx, edx
	LONG $0x24548948; BYTE $0x08   // mov	qword ptr [rsp + 8], rdx
	WORD $0x3145; BYTE $0xc9       // xor	r9d, r9d
	LONG $0x245c8948; BYTE $0x20   // mov	qword ptr [rsp + 32], rbx
	JMP  LBB5_3

LBB5_2:
	LONG $0x2444ff48; BYTE $0x08               // inc	qword ptr [rsp + 8]
	LONG $0x244c8b48; BYTE $0x38               // mov	rcx, qword ptr [rsp + 56]
	LONG $0x00c18148; WORD $0x0002; BYTE $0x00 // add	rcx, 512
	QUAD $0xfffe001024448148; BYTE $0xff       // add	qword ptr [rsp + 16], -512
	QUAD $0x0010001824448148; BYTE $0x00       // add	qword ptr [rsp + 24], 4096
	LONG $0x10000548; WORD $0x0000             // add	rax, 4096
	LONG $0x245c8b48; BYTE $0x20               // mov	rbx, qword ptr [rsp + 32]
	LONG $0x244c8b4c; BYTE $0x30               // mov	r9, qword ptr [rsp + 48]
	WORD $0x3949; BYTE $0xd9                   // cmp	r9, rbx
	JAE  LBB5_14

LBB5_3:
	WORD $0x3948; BYTE $0xd9                   // cmp	rcx, rbx
	WORD $0x8949; BYTE $0xdb                   // mov	r11, rbx
	LONG $0x244c8948; BYTE $0x38               // mov	qword ptr [rsp + 56], rcx
	LONG $0xd9420f4c                           // cmovb	r11, rcx
	LONG $0x00898d49; WORD $0x0002; BYTE $0x00 // lea	rcx, [r9 + 512]
	WORD $0x3948; BYTE $0xd9                   // cmp	rcx, rbx
	LONG $0x244c8948; BYTE $0x30               // mov	qword ptr [rsp + 48], rcx
	LONG $0xd9420f48                           // cmovb	rbx, rcx
	LONG $0x247c8348; WORD $0x0028             // cmp	qword ptr [rsp + 40], 0
	JE   LBB5_2
	LONG $0x245c034c; BYTE $0x10               // add	r11, qword ptr [rsp + 16]
	LONG $0xfce38349                           // and	r11, -4
	LONG $0x247c8b4c; BYTE $0x08               // mov	r15, qword ptr [rsp + 8]
	WORD $0x894c; BYTE $0xf9                   // mov	rcx, r15
	LONG $0x09e1c148                           // shl	rcx, 9
	WORD $0x8949; BYTE $0xd8                   // mov	r8, rbx
	WORD $0x2949; BYTE $0xc8                   // sub	r8, rcx
	LONG $0x0ce7c149                           // shl	r15, 12
	LONG $0x3f148d4e                           // lea	r10, [rdi + r15]
	LONG $0x247c894c; BYTE $0x60               // mov	qword ptr [rsp + 96], r15
	LONG $0xc73c8d4f                           // lea	r15, [r15 + 8*r8]
	LONG $0x3f0c8d49                           // lea	rcx, [r15 + rdi]
	LONG $0x244c8948; BYTE $0x58               // mov	qword ptr [rsp + 88], rcx
	WORD $0x894c; BYTE $0xc1                   // mov	rcx, r8
	LONG $0xfce18348                           // and	rcx, -4
	LONG $0x244c8948; BYTE $0x48               // mov	qword ptr [rsp + 72], rcx
	WORD $0x014c; BYTE $0xc9                   // add	rcx, r9
	LONG $0x244c8948; BYTE $0x40               // mov	qword ptr [rsp + 64], rcx
	WORD $0x3145; BYTE $0xf6                   // xor	r14d, r14d
	LONG $0x247c894c; BYTE $0x50               // mov	qword ptr [rsp + 80], r15
	JMP  LBB5_5

LBB5_12:
	WORD $0xff49; BYTE $0xc6 // inc	r14
	WORD $0x394d; BYTE $0xe6 // cmp	r14, r12
	JE   LBB5_2

LBB5_5:
	WORD $0x3949; BYTE $0xd9     // cmp	r9, rbx
	JAE  LBB5_12
	LONG $0xf6148b4a             // mov	rdx, qword ptr [rsi + 8*r14]
	WORD $0x894d; BYTE $0xcd     // mov	r13, r9
	LONG $0x04f88349             // cmp	r8, 4
	JB   LBB5_13
	LONG $0x3a0c8d4a             // lea	rcx, [rdx + r15]
	WORD $0x3949; BYTE $0xca     // cmp	r10, rcx
	JAE  LBB5_9
	LONG $0x244c8b48; BYTE $0x60 // mov	rcx, qword ptr [rsp + 96]
	WORD $0x0148; BYTE $0xd1     // add	rcx, rdx
	WORD $0x894d; BYTE $0xcd     // mov	r13, r9
	LONG $0x244c3b48; BYTE $0x58 // cmp	rcx, qword ptr [rsp + 88]
	JB   LBB5_13

LBB5_9:
	WORD $0x894d; BYTE $0xd7     // mov	r15, r10
	WORD $0x8949; BYTE $0xf2     // mov	r10, rsi
	LONG $0x244c8b48; BYTE $0x18 // mov	rcx, qword ptr [rsp + 24]
	LONG $0x0a348d48             // lea	rsi, [rdx + rcx]
	WORD $0xc931                 // xor	ecx, ecx

LBB5_10:
	LONG $0x0410fcc5; BYTE $0xce // vmovups	ymm0, ymmword ptr [rsi + 8*rcx]
	LONG $0x0455fcc5; BYTE $0xc8 // vandnps	ymm0, ymm0, ymmword ptr [rax + 8*rcx]
	LONG $0x0411fcc5; BYTE $0xc8 // vmovups	ymmword ptr [rax + 8*rcx], ymm0
	LONG $0x04c18348             // add	rcx, 4
	WORD $0x3949; BYTE $0xcb     // cmp	r11, rcx
	JNE  LBB5_10
	LONG $0x246c8b4c; BYTE $0x40 // mov	r13, qword ptr [rsp + 64]
	LONG $0x24443b4c; BYTE $0x48 // cmp	r8, qword ptr [rsp + 72]
	WORD $0x894c; BYTE $0xd6     // mov	rsi, r10
	WORD $0x894d; BYTE $0xfa     // mov	r10, r15
	LONG $0x247c8b4c; BYTE $0x50 // mov	r15, qword ptr [rsp + 80]
	JE   LBB5_12

LBB5_13:
	LONG $0xea0c8b4a         // mov	rcx, qword ptr [rdx + 8*r13]
	WORD $0xf748; BYTE $0xd1 // not	rcx
	LONG $0xef0c214a         // and	qword ptr [rdi + 8*r13], rcx
	WORD $0xff49; BYTE $0xc5 // inc	r13
	WORD $0x3949; BYTE $0xdd // cmp	r13, rbx
	JB   LBB5_13
	JMP  LBB5_12

LBB5_14:
	LONG $0xd8658d48         // lea	rsp, [rbp - 40]
	BYTE $0x5b               // pop	rbx
	WORD $0x5c41             // pop	r12
	WORD $0x5d41             // pop	r13
	WORD $0x5e41             // pop	r14
	WORD $0x5f41             // pop	r15
	BYTE $0x5d               // pop	rbp
	WORD $0xf8c5; BYTE $0x77 // vzeroupper
	BYTE $0xc3               // ret

TEXT ·_or_many(SB), $0-32
	MOVQ a+0(FP), DI
	MOVQ b+8(FP), SI
	MOVQ dims+16(FP), DX
	BYTE $0x55                     // push	rbp
	WORD $0x8948; BYTE $0xe5       // mov	rbp, rsp
	WORD $0x5741                   // push	r15
	WORD $0x5641                   // push	r14
	WORD $0x5541                   // push	r13
	WORD $0x5441                   // push	r12
	BYTE $0x53                     // push	rbx
	LONG $0xf8e48348               // and	rsp, -8
	LONG $0x68ec8348               // sub	rsp, 104
	LONG $0xffffffbb; BYTE $0xff   // mov	ebx, 4294967295
	WORD $0x2148; BYTE $0xd3       // and	rbx, rdx
	JE   LBB6_14
	WORD $0x8948; BYTE $0xd0       // mov	rax, rdx
	LONG $0x20e8c148               // shr	rax, 32
	LONG $0x02f88348               // cmp	rax, 2
	LONG $0x0001bc41; WORD $0x0000 // mov	r12d, 1
	LONG $0x24448948; BYTE $0x28   // mov	qword ptr [rsp + 40], rax
	LONG $0xe0430f4c               // cmovae	r12, rax
	LONG $0x000200b9; BYTE $0x00   // mov	ecx, 512
	WORD $0xc031                   // xor	eax, eax
	LONG $0x24448948; BYTE $0x18   // mov	qword ptr [rsp + 24], rax
	WORD $0x8948; BYTE $0xf8       // mov	rax, rdi
	WORD $0xd231                   // xor	edx, edx
	LONG $0x24548948; BYTE $0x10   // mov	qword ptr [rsp + 16], rdx
	WORD $0xd231                   // xor	edx, edx
	LONG $0x24548948; BYTE $0x08   // mov	qword ptr [rsp + 8], rdx
	WORD $0x3145; BYTE $0xc9       // xor	r9d, r9d
	LONG $0x245c8948; BYTE $0x20   // mov	qword ptr [rsp + 32], rbx
	JMP  LBB6_3

LBB6_2:
	LONG $0x2444ff48; BYTE $0x08               // inc	qword ptr [rsp + 8]
	LONG $0x244c8b48; BYTE $0x38               // mov	rcx, qword ptr [rsp + 56]
	LONG $0x00c18148; WORD $0x0002; BYTE $0x00 // add	rcx, 512
	QUAD $0xfffe001024448148; BYTE $0xff       // add	qword ptr [rsp + 16], -512
	QUAD $0x0010001824448148; BYTE $0x00       // add	qword ptr [rsp + 24], 4096
	LONG $0x10000548; WORD $0x0000             // add	rax, 4096
	LONG $0x245c8b48; BYTE $0x20               // mov	rbx, qword ptr [rsp + 32]
	LONG $0x244c8b4c; BYTE $0x30               // mov	r9, qword ptr [rsp + 48]
	WORD $0x3949; BYTE $0xd9                   // cmp	r9, rbx
	JAE  LBB6_14

LBB6_3:
	WORD $0x3948; BYTE $0xd9                   // cmp	rcx, rbx
	WORD $0x8949; BYTE $0xda                   // mov	r10, rbx
	LONG $0x244c8948; BYTE $0x38               // mov	qword ptr [rsp + 56], rcx
	LONG $0xd1420f4c                           // cmovb	r10, rcx
	LONG $0x00898d49; WORD $0x0002; BYTE $0x00 // lea	rcx, [r9 + 512]
	WORD $0x3948; BYTE $0xd9                   // cmp	rcx, rbx
	LONG $0x244c8948; BYTE $0x30               // mov	qword ptr [rsp + 48], rcx
	LONG $0xd9420f48                           // cmovb	rbx, rcx
	LONG $0x247c8348; WORD $0x0028             // cmp	qword ptr [rsp + 40], 0
	JE   LBB6_2
	LONG $0x2454034c; BYTE $0x10               // add	r10, qword ptr [rsp + 16]
	LONG $0xfce28349                           // and	r10, -4
	LONG $0x247c8b4c; BYTE $0x08               // mov	r15, qword ptr [rsp + 8]
	WORD $0x894c; BYTE $0xf9                   // mov	rcx, r15
	LONG $0x09e1c148                           // shl	rcx, 9
	WORD $0x8949; BYTE $0xd8                   // mov	r8, rbx
	WORD $0x2949; BYTE $0xc8                   // sub	r8, rcx
	LONG $0x0ce7c149                           // shl	r15, 12
	LONG $0x3f1c8d4e                           // lea	r11, [rdi + r15]
	LONG $0x247c894c; BYTE $0x60               // mov	qword ptr [rsp + 96], r15
	LONG $0xc73c8d4f                           // lea	r15, [r15 + 8*r8]
	LONG $0x3f0c8d49                           // lea	rcx, [r15 + rdi]
	LONG $0x244c8948; BYTE $0x58               // mov	qword ptr [rsp + 88], rcx
	WORD $0x894c; BYTE $0xc1                   // mov	rcx, r8
	LONG $0xfce18348                           // and	rcx, -4
	LONG $0x244c8948; BYTE $0x48               // mov	qword ptr [rsp + 72], rcx
	WORD $0x014c; BYTE $0xc9                   // add	rcx, r9
	LONG $0x244c8948; BYTE $0x40               // mov	qword ptr [rsp + 64], rcx
	WORD $0x3145; BYTE $0xf6                   // xor	r14d, r14d
	LONG $0x247c894c; BYTE $0x50               // mov	qword ptr [rsp + 80], r15
	JMP  LBB6_5

LBB6_12:
	WORD $0xff49; BYTE $0xc6 // inc	r14
	WORD $0x394d; BYTE $0xe6 // cmp	r14, r12
	JE   LBB6_2

LBB6_5:
	WORD $0x3949; BYTE $0xd9     // cmp	r9, rbx
	JAE  LBB6_12
	LONG $0xf6148b4a             // mov	rdx, qword ptr [rsi + 8*r14]
	WORD $0x894d; BYTE $0xcd     // mov	r13, r9
	LONG $0x04f88349             // cmp	r8, 4
	JB   LBB6_13
	LONG $0x3a0c8d4a             // lea	rcx, [rdx + r15]
	WORD $0x3949; BYTE $0xcb     // cmp	r11, rcx
	JAE  LBB6_9
	LONG $0x244c8b48; BYTE $0x60 // mov	rcx, qword ptr [rsp + 96]
	WORD $0x0148; BYTE $0xd1     // add	rcx, rdx
	WORD $0x894d; BYTE $0xcd     // mov	r13, r9
	LONG $0x244c3b48; BYTE $0x58 // cmp	rcx, qword ptr [rsp + 88]
	JB   LBB6_13

LBB6_9:
	WORD $0x894d; BYTE $0xdf     // mov	r15, r11
	WORD $0x8949; BYTE $0xf3     // mov	r11, rsi
	LONG $0x244c8b48; BYTE $0x18 // mov	rcx, qword ptr [rsp + 24]
	LONG $0x0a348d48             // lea	rsi, [rdx + rcx]
	WORD $0xc931                 // xor	ecx, ecx

LBB6_10:
	LONG $0x0410fcc5; BYTE $0xc8 // vmovups	ymm0, ymmword ptr [rax + 8*rcx]
	LONG $0x0456fcc5; BYTE $0xce // vorps	ymm0, ymm0, ymmword ptr [rsi + 8*rcx]
	LONG $0x0411fcc5; BYTE $0xc8 // vmovups	ymmword ptr [rax + 8*rcx], ymm0
	LONG $0x04c18348             // add	rcx, 4
	WORD $0x3949; BYTE $0xca     // cmp	r10, rcx
	JNE  LBB6_10
	LONG $0x246c8b4c; BYTE $0x40 // mov	r13, qword ptr [rsp + 64]
	LONG $0x24443b4c; BYTE $0x48 // cmp	r8, qword ptr [rsp + 72]
	WORD $0x894c; BYTE $0xde     // mov	rsi, r11
	WORD $0x894d; BYTE $0xfb     // mov	r11, r15
	LONG $0x247c8b4c; BYTE $0x50 // mov	r15, qword ptr [rsp + 80]
	JE   LBB6_12

LBB6_13:
	LONG $0xea0c8b4a         // mov	rcx, qword ptr [rdx + 8*r13]
	LONG $0xef0c094a         // or	qword ptr [rdi + 8*r13], rcx
	WORD $0xff49; BYTE $0xc5 // inc	r13
	WORD $0x3949; BYTE $0xdd // cmp	r13, rbx
	JB   LBB6_13
	JMP  LBB6_12

LBB6_14:
	LONG $0xd8658d48         // lea	rsp, [rbp - 40]
	BYTE $0x5b               // pop	rbx
	WORD $0x5c41             // pop	r12
	WORD $0x5d41             // pop	r13
	WORD $0x5e41             // pop	r14
	WORD $0x5f41             // pop	r15
	BYTE $0x5d               // pop	rbp
	WORD $0xf8c5; BYTE $0x77 // vzeroupper
	BYTE $0xc3               // ret

TEXT ·_xor_many(SB), $0-32
	MOVQ a+0(FP), DI
	MOVQ b+8(FP), SI
	MOVQ dims+16(FP), DX
	BYTE $0x55                     // push	rbp
	WORD $0x8948; BYTE $0xe5       // mov	rbp, rsp
	WORD $0x5741                   // push	r15
	WORD $0x5641                   // push	r14
	WORD $0x5541                   // push	r13
	WORD $0x5441                   // push	r12
	BYTE $0x53                     // push	rbx
	LONG $0xf8e48348               // and	rsp, -8
	LONG $0x68ec8348               // sub	rsp, 104
	LONG $0xffffffbb; BYTE $0xff   // mov	ebx, 4294967295
	WORD $0x2148; BYTE $0xd3       // and	rbx, rdx
	JE   LBB7_14
	WORD $0x8948; BYTE $0xd0       // mov	rax, rdx
	LONG $0x20e8c148               // shr	rax, 32
	LONG $0x02f88348               // cmp	rax, 2
	LONG $0x0001bc41; WORD $0x0000 // mov	r12d, 1
	LONG $0x24448948; BYTE $0x28   // mov	qword ptr [rsp + 40], rax
	LONG $0xe0430f4c               // cmovae	r12, rax
	LONG $0x000200b9; BYTE $0x00   // mov	ecx, 512
	WORD $0xc031                   // xor	eax, eax
	LONG $0x24448948; BYTE $0x18   // mov	qword ptr [rsp + 24], rax
	WORD $0x8948; BYTE $0xf8       // mov	rax, rdi
	WORD $0xd231                   // xor	edx, edx
	LONG $0x24548948; BYTE $0x10   // mov	qword ptr [rsp + 16], rdx
	WORD $0xd231                   // xor	edx, edx
	LONG $0x24548948; BYTE $0x08   // mov	qword ptr [rsp + 8], rdx
	WORD $0x3145; BYTE $0xc9       // xor	r9d, r9d
	LONG $0x245c8948; BYTE $0x20   // mov	qword ptr [rsp + 32], rbx
	JMP  LBB7_3

LBB7_2:
	LONG $0x2444ff48; BYTE $0x08               // inc	qword ptr [rsp + 8]
	LONG $0x244c8b48; BYTE $0x38               // mov	rcx, qword ptr [rsp + 56]
	LONG $0x00c18148; WORD $0x0002; BYTE $0x00 // add	rcx, 512
	QUAD $0xfffe001024448148; BYTE $0xff       // add	qword ptr [rsp + 16], -512
	QUAD $0x0010001824448148; BYTE $0x00       // add	qword ptr [rsp + 24], 4096
	LONG $0x10000548; WORD $0x0000             // add	rax, 4096
	LONG $0x245c8b48; BYTE $0x20               // mov	rbx, qword ptr [rsp + 32]
	LONG $0x244c8b4c; BYTE $0x30               // mov	r9, qword ptr [rsp + 48]
	WORD $0x3949; BYTE $0xd9                   // cmp	r9, rbx
	JAE  LBB7_14

LBB7_3:
	WORD $0x3948; BYTE $0xd9                   // cmp	rcx, rbx
	WORD $0x8949; BYTE $0xda                   // mov	r10, rbx
	LONG $0x244c8948; BYTE $0x38               // mov	qword ptr [rsp + 56], rcx
	LONG $0xd1420f4c                           // cmovb	r10, rcx
	LONG $0x00898d49; WORD $0x0002; BYTE $0x00 // lea	rcx, [r9 + 512]
	WORD $0x3948; BYTE $0xd9                   // cmp	rcx, rbx
	LONG $0x244c8948; BYTE $0x30               // mov	qword ptr [rsp + 48], rcx
	LONG $0xd9420f48                           // cmovb	rbx, rcx
	LONG $0x247c8348; WORD $0x0028             // cmp	qword ptr [rsp + 40], 0
	JE   LBB7_2
	LONG $0x2454034c; BYTE $0x10               // add	r10, qword ptr [rsp + 16]
	LONG $0xfce28349                           // and	r10, -4
	LONG $0x247c8b4c; BYTE $0x08               // mov	r15, qword ptr [rsp + 8]
	WORD $0x894c; BYTE $0xf9                   // mov	rcx, r15
	LONG $0x09e1c148                           // shl	rcx, 9
	WORD $0x8949; BYTE $0xd8                   // mov	r8, rbx
	WORD $0x2949; BYTE $0xc8                   // sub	r8, rcx
	LONG $0x0ce7c149                           // shl	r15, 12
	LONG $0x3f1c8d4e                           // lea	r11, [rdi + r15]
	LONG $0x247c894c; BYTE $0x60               // mov	qword ptr [rsp + 96], r15
	LONG $0xc73c8d4f                           // lea	r15, [r15 + 8*r8]
	LONG $0x3f0c8d49                           // lea	rcx, [r15 + rdi]
	LONG $0x244c8948; BYTE $0x58               // mov	qword ptr [rsp + 88], rcx
	WORD $0x894c; BYTE $0xc1                   // mov	rcx, r8
	LONG $0xfce18348                           // and	rcx, -4
	LONG $0x244c8948; BYTE $0x48               // mov	qword ptr [rsp + 72], rcx
	WORD $0x014c; BYTE $0xc9                   // add	rcx, r9
	LONG $0x244c8948; BYTE $0x40               // mov	qword ptr [rsp + 64], rcx
	WORD $0x3145; BYTE $0xf6                   // xor	r14d, r14d
	LONG $0x247c894c; BYTE $0x50               // mov	qword ptr [rsp + 80], r15
	JMP  LBB7_5

LBB7_12:
	WORD $0xff49; BYTE $0xc6 // inc	r14
	WORD $0x394d; BYTE $0xe6 // cmp	r14, r12
	JE   LBB7_2

LBB7_5:
	WORD $0x3949; BYTE $0xd9     // cmp	r9, rbx
	JAE  LBB7_12
	LONG $0xf6148b4a             // mov	rdx, qword ptr [rsi + 8*r14]
	WORD $0x894d; BYTE $0xcd     // mov	r13, r9
	LONG $0x04f88349             // cmp	r8, 4
	JB   LBB7_13
	LONG $0x3a0c8d4a             // lea	rcx, [rdx + r15]
	WORD $0x3949; BYTE $0xcb     // cmp	r11, rcx
	JAE  LBB7_9
	LONG $0x244c8b48; BYTE $0x60 // mov	rcx, qword ptr [rsp + 96]
	WORD $0x0148; BYTE $0xd1     // add	rcx, rdx
	WORD $0x894d; BYTE $0xcd     // mov	r13, r9
	LONG $0x244c3b48; BYTE $0x58 // cmp	rcx, qword ptr [rsp + 88]
	JB   LBB7_13

LBB7_9:
	WORD $0x894d; BYTE $0xdf     // mov	r15, r11
	WORD $0x8949; BYTE $0xf3     // mov	r11, rsi
	LONG $0x244c8b48; BYTE $0x18 // mov	rcx, qword ptr [rsp + 24]
	LONG $0x0a348d48             // lea	rsi, [rdx + rcx]
	WORD $0xc931                 // xor	ecx, ecx

LBB7_10:
	LONG $0x0410fcc5; BYTE $0xc8 // vmovups	ymm0, ymmword ptr [rax + 8*rcx]
	LONG $0x0457fcc5; BYTE $0xce // vxorps	ymm0, ymm0, ymmword ptr [rsi + 8*rcx]
	LONG $0x0411fcc5; BYTE $0xc8 // vmovups	ymmword ptr [rax + 8*rcx], ymm0
	LONG $0x04c18348             // add	rcx, 4
	WORD $0x3949; BYTE $0xca     // cmp	r10, rcx
	JNE  LBB7_10
	LONG $0x246c8b4c; BYTE $0x40 // mov	r13, qword ptr [rsp + 64]
	LONG $0x24443b4c; BYTE $0x48 // cmp	r8, qword ptr [rsp + 72]
	WORD $0x894c; BYTE $0xde     // mov	rsi, r11
	WORD $0x894d; BYTE $0xfb     // mov	r11, r15
	LONG $0x247c8b4c; BYTE $0x50 // mov	r15, qword ptr [rsp + 80]
	JE   LBB7_12

LBB7_13:
	LONG $0xea0c8b4a         // mov	rcx, qword ptr [rdx + 8*r13]
	LONG $0xef0c314a         // xor	qword ptr [rdi + 8*r13], rcx
	WORD $0xff49; BYTE $0xc5 // inc	r13
	WORD $0x3949; BYTE $0xdd // cmp	r13, rbx
	JB   LBB7_13
	JMP  LBB7_12

LBB7_14:
	LONG $0xd8658d48         // lea	rsp, [rbp - 40]
	BYTE $0x5b               // pop	rbx
	WORD $0x5c41             // pop	r12
	WORD $0x5d41             // pop	r13
	WORD $0x5e41             // pop	r14
	WORD $0x5f41             // pop	r15
	BYTE $0x5d               // pop	rbp
	WORD $0xf8c5; BYTE $0x77 // vzeroupper
	BYTE $0xc3               // ret

TEXT ·_count(SB), $0-32
	MOVQ a+0(FP), DI
	MOVQ size+8(FP), SI
	MOVQ result+16(FP), DX
	BYTE $0x55               // push	rbp
	WORD $0x8948; BYTE $0xe5 // mov	rbp, rsp
	LONG $0xf8e48348         // and	rsp, -8
	WORD $0x8548; BYTE $0xf6 // test	rsi, rsi
	JE   LBB8_1
	WORD $0xc931             // xor	ecx, ecx
	WORD $0x3145; BYTE $0xc0 // xor	r8d, r8d

LBB8_4:
	LONG $0xb80f48f3; WORD $0xcf04 // popcnt	rax, qword ptr [rdi + 8*rcx]
	WORD $0x0149; BYTE $0xc0       // add	r8, rax
	WORD $0xff48; BYTE $0xc1       // inc	rcx
	WORD $0x3948; BYTE $0xce       // cmp	rsi, rcx
	JNE  LBB8_4
	JMP  LBB8_2

LBB8_1:
	WORD $0x3145; BYTE $0xc0 // xor	r8d, r8d

LBB8_2:
	WORD $0x894c; BYTE $0x02 // mov	qword ptr [rdx], r8
	WORD $0x8948; BYTE $0xec // mov	rsp, rbp
	BYTE $0x5d               // pop	rbp
	BYTE $0xc3               // ret
