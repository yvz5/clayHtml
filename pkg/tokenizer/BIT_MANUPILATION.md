# Bit Manupilation

| 6 bits     | 32 bits    | 26 bits   |
| ---------- | ---------- | --------- |
| Type       | Start      | Length    |
| bits 58–63 | bits 26–57 | bits 0–25 |

we want to pack all three fields into a single `uint64`.
To write or read them, we use:

- **Shift left/right** to move bits into position
- **Masks** to isolate or clear regions

# Packing

### Type Field

```go
tapeMaskType  = uint64(0b111111) << tapeShiftType
```

- `0b111111` = 6 bits set to 1 = 63 in decimal
- `tapeShiftType = 58` => we're saying `put these 6 ones at the top of 64-bit space`
  Result:

```
1111110000000000000000000000000000000000000000000000000000000000
```

We’ll use this to extract or clear the "Type" field at the top.

### Start Field

```go
tapeMaskStart = uint64(0xFFFFFFFF) << tapeShiftStart
```

- `0xFFFFFFFF` = 32 bits of 1s (for the Start offset)
- `tapeShiftStart = 26`
  Result:

```
0000001111111111111111111111111111111111000000000000000000000000
```

Mask for the middle `32` bits.

### Length Field

```go
tapeMaskLen = uint64((1 << 26) - 1)
```

- `1 << 26` = 2²⁶ = 67,108,864
- `-1` = fills the bottom 26 bits with 1s:
  Result:

```
0000000000000000000000000000000000000000111111111111111111111111
```

This is the length field mask for the lowest `26` bits.

### Summary:

```go
return HTMLTapeEntry(
    (uint64(tokenType & 0b111111) << tapeShiftType) |
    (uint64(start) << tapeShiftStart) |
    (uint64(length) << tapeShiftLen),
)
```

`(uint64(tokenType & 0b111111) << tapeShiftType)`

- Keeps only the lowest 6 bits of tokenType (just in case)
- Shifts it left to bit position 58
- Sets the top 6 bits of the tape

`(uint64(start) << tapeShiftStart)`

- Shifts start left by 26
- Places it into the middle 32 bits

`(uint64(length) << tapeShiftLen)`

- tapeShiftLen is 0, so this is just length
- Fills the bottom 26 bits

`|` **= bitwise OR**

- Combines all three parts into one 64-bit value, each in its proper position.

# Unpacking

```go
func (t HTMLTapeEntry) Type() TokenType {
    return TokenType((uint64(t) >> tapeShiftType) & 0b111111)
}
```

- Shift right by 58 → bring the top 6 bits down to the bottom
- Mask them with 0b111111 to extract just the bits we want

# Summary

| Action       | Description                    |
| ------------ | ------------------------------ |
| Encode       | Shift left then OR together    |
| Decode       | Shift right then AND with mask |
| Type field   | 6 bits at top (bit 58–63)      |
| Start field  | 32 bits in middle (bit 26–57)  |
| Length field | 26 bits at bottom (bit 0–25)   |
