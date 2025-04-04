# clayHtml

An **experimental** SIMD-accelerated HTML parser written in Go.  
Built for fun, speed, and possibly breaking everything.

⚠️ **Warning: This is experimental software. I don’t really know what I’m doing. Please do not use this in production.** ⚠️

---

## 🚧 What is this?

This is a fast, hand-crafted HTML parser that attempts to leverage SIMD (Single Instruction, Multiple Data) operations for blazing-fast parsing performance.

> **Heavily inspired by [`simdjson`](https://github.com/simdjson/simdjson)** — a ridiculously fast JSON parser that showed what's possible with SIMD. This project tries to apply similar ideas to HTML parsing, with varying degrees of success (and failure).

## 🤖 Why?

- I wanted to mess around with SIMD.
- I wanted to see how far I could push Go for this kind of task.
- I have a weird love for performance tuning.

## 🔍 Features

- Parses some HTML using SIMD
- Written in Go, no C bindings
- Uses Plan9 assembly
- Uses bitwise operations

---

# 🛠️ Contributing
Pull requests are welcome, especially ones that fix my obvious mistakes.
Open issues, file bugs, or just fork it and build your own better version.

# 🙈 Disclaimer
This project is purely experimental and built for educational purposes.
It is not production-ready and should not be relied upon in any critical system.

I take no responsibility for broken sites or corrupted HTML.

# 🧠 License
MIT – use it, break it, learn from it.