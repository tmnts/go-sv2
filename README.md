# go-sv2 (WIP) ⚒️

**Golang implementation of the Stratum v2 Mining Protocol.**

An efficient, type-safe, and tested library for building next-generation Bitcoin mining infrastructure. Born from a mix of low-level audio engineering (VST/C++) vibes and high-precision financial math.

## 🚀 Status: Early Alpha
Current progress: **Handshake Lifecycle Completed**
- [x] `SetupConnection` (Serialization/Deserialization)
- [x] `OpenStandardMiningChannel` (Request/Success flow)
- [x] 24-bit Binary Framing (uint24)
- [x] 100% Test Coverage for implemented messages

## 🛠 Tech Stack
- **Language:** Go (Golang)
- **Math:** Fixed-precision (Decimal) & IEEE 754 (float32)
- **Testing:** Standard `testing` package with binary validation

## 🧪 Quick Start
```bash
go mod tidy
go test -v ./...
```

![Project Banner](./assets/go-sv2-test.png)

## 🚧 Roadmap

- [ ] **Noise Protocol Integration** (Encrypted Handshake)
- [ ] **Job Negotiation Protocol**
- [ ] **Template Provider** implementation