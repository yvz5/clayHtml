# Tape Format

We use a 64-bit packed tape format similarly defined in simdjson repository. We change it slightly to make it work for html parsing.

represented by a single `uint64`

## Tape Layout Design

| Bits    | Field    | Description                            |
| ------- | -------- | -------------------------------------- |
| 6 Bits  | `Type`   | Token type (`<tag>`, attr, text, etc.) |
| 32 Bits | `Start`  | Offset into original buffer            |
| 26 Bits | `Length` | Length in bytes                        |

```
6 + 32 + 26 = 64 bits
```

## Bit positions

| Field  | Shift |
| ------ | ----- |
| Type   | 58    |
| Start  | 26    |
| Length | 0     |

## Advantages of this layout

- 4 GB+ input support
- 64 MB max token length (enough for most <script> blocks)
