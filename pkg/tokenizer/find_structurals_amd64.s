// +build amd64

#include "textflag.h"

DATA ·charLt+0(SB)/1, $'<' // ASCII 60
GLOBL ·charLt(SB), RODATA, $1

DATA ·charGt+0(SB)/1, $'>'
GLOBL ·charGt(SB), RODATA, $1

DATA ·charSlash+0(SB)/1, $'/'
GLOBL ·charSlash(SB), RODATA, $1

DATA ·charEq+0(SB)/1, $'='
GLOBL ·charEq(SB), RODATA, $1

DATA ·charDQuote+0(SB)/1, $"\""
GLOBL ·charDQuote(SB), RODATA, $1

DATA ·charSQuote+0(SB)/1, $'\''
GLOBL ·charSQuote(SB), RODATA, $1

DATA ·charBang+0(SB)/1, $'!'
GLOBL ·charBang(SB), RODATA, $1

DATA ·charDash+0(SB)/1, $'-'
GLOBL ·charDash(SB), RODATA, $1

DATA ·charQMark+0(SB)/1, $'?'
GLOBL ·charQMark(SB), RODATA, $1


// func findStructuralsSIMD(buf *byte, length int, structurals *uint32, maxStructurals int) int
TEXT ·findStructuralsSIMD(SB), NOSPLIT, $0-32
    MOVQ buf+0(FP), DI             // buf pointer
    MOVQ length+8(FP), SI          // length
    MOVQ structurals+16(FP), DX    // structurals pointer
    MOVQ maxStructurals+24(FP), CX // maxStructurals
    XORQ AX, AX                    // struct count
    XORQ R8, R8                    // offset from original start

loop_start:
    CMPQ SI, $32
    JL loop_tail

    VMOVDQU (DI), Y0               // Load 32 bytes from buf
    XORL R9, R9                    // R9: combined match bitmask

    // Compare to '<'
    VPBROADCASTB ·charLt(SB), Y1
    VPCMPEQB Y1, Y0, Y2
    VPMOVMSKB Y2, R10
    ORL R10, R9

    // Compare to '>'
    VPBROADCASTB ·charGt(SB), Y1
    VPCMPEQB Y1, Y0, Y2
    VPMOVMSKB Y2, R10
    ORL R10, R9

    // Compare to '/'
    VPBROADCASTB ·charSlash(SB), Y1
    VPCMPEQB Y1, Y0, Y2
    VPMOVMSKB Y2, R10
    ORL R10, R9

    // Compare to '='
    VPBROADCASTB ·charEq(SB), Y1
    VPCMPEQB Y1, Y0, Y2
    VPMOVMSKB Y2, R10
    ORL R10, R9

    // Compare to '"'
    VPBROADCASTB ·charDQuote(SB), Y1
    VPCMPEQB Y1, Y0, Y2
    VPMOVMSKB Y2, R10
    ORL R10, R9

    // Compare to '\''
    VPBROADCASTB ·charSQuote(SB), Y1
    VPCMPEQB Y1, Y0, Y2
    VPMOVMSKB Y2, R10
    ORL R10, R9

    // Compare to '!'
    VPBROADCASTB ·charBang(SB), Y1
    VPCMPEQB Y1, Y0, Y2
    VPMOVMSKB Y2, R10
    ORL R10, R9

    // Compare to '-'
    VPBROADCASTB ·charDash(SB), Y1
    VPCMPEQB Y1, Y0, Y2
    VPMOVMSKB Y2, R10
    ORL R10, R9

    // Compare to '?'
    VPBROADCASTB ·charQMark(SB), Y1
    VPCMPEQB Y1, Y0, Y2
    VPMOVMSKB Y2, R10
    ORL R10, R9

    // If no matches found, skip storing
    TESTL R9, R9
    JZ skip_store

    // Store matched positions
    XORQ R10, R10
store_loop:
    BSFL R9, R11
    JZ skip_store

    MOVL R8, R12
    ADDL R11, R12
    MOVL R12, (DX)(AX*4)
    INCL AX
    BTRQ R11, R9
    CMPQ AX, CX
    JE done
    JMP store_loop

skip_store:
    ADDQ $32, DI
    ADDQ $32, R8
    SUBQ $32, SI
    JMP loop_start

loop_tail:
    // todo: handle < 32 bytes here
done:
    MOVQ AX, ret+32(FP)
    VZEROUPPER
    RET
