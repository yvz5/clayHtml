# Plan9 Syntax
Here is a list of plan9 syntax equivalent to c intrinsics.

| Concept                      | Plan 9 Assembly                      | C Intrinsic                    |
|-----------------------------|--------------------------------------|--------------------------------|
| Load 32 bytes               | `VMOVDQU (DI), Y0`                   | `_mm256_loadu_si256`          |
| Broadcast 1 byte            | `VPBROADCASTB ·charLt(SB), Y1`       | `_mm256_set1_epi8('<')`       |
| SIMD compare (byte == byte) | `VPCMPEQB Y1, Y0, Y2`                | `_mm256_cmpeq_epi8`           |
| Extract match bitmask       | `VPMOVMSKB Y2, R10`                  | `_mm256_movemask_epi8`        |
| Combine masks               | `ORL R10, R9`                        | `mask |= mask_for_char`       |

## _mm256_loadu_si256
```
VMOVDQU (DI), Y0
```
Loads 32 bytes from memory (pointed to by **DI**) into YMM0 (**Y0**).

## _mm256_set1_epi8
```
VPBROADCASTB ·charLt(SB), Y1
```
Broadcasts the byte stored at **charLt(SB)** (i.e., **'<'**) to all 32 bytes of YMM1 (**Y1**).

## _mm256_cmpeq_epi8
```
VPCMPEQB Y1, Y0, Y2
```
Compares each byte in **Y0** (input) to corresponding byte in **Y1** (target = '<'), and stores result in **Y2**.

## _mm256_movemask_epi8
```
VPMOVMSKB Y2, R10
```
Converts the most significant bit of each byte in **Y2** into a 32-bit integer in **R10**.