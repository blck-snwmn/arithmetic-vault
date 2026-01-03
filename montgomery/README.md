# montgomery

Montgomery multiplication implementation for efficient modular arithmetic.

## Implementations

- `MontgomeryBitwise` - Basic bit-by-bit REDC algorithm
- `MontgomeryCIOS` - CIOS algorithm using big.Int internally
- `MontgomeryCIOSWords` - CIOS algorithm using []uint64 for better performance

## Benchmark

```bash
# All benchmarks
go test -bench=. -benchmem

# Single multiplication benchmark
go test -bench=BenchmarkMontgomeryMul -benchmem

# Modular exponentiation benchmark (amortized cost)
go test -bench=BenchmarkModExp -benchmem
```

### Single Multiplication (2048-bit)

Measures the cost of a single modular multiplication including Montgomery form conversion.

| Implementation | ns/op | allocs/op |
|----------------|-------|-----------|
| MontgomeryBitwise | ~38,000 | 0 |
| MontgomeryCIOS | ~5,700 | 396 |
| MontgomeryCIOSWords | ~1,400 | 10 |

### Modular Exponentiation (2048-bit base, 2048-bit exponent)

Demonstrates Montgomery's amortized advantage: conversion cost is paid once at start/end, while many multiplications happen in the Montgomery domain.

| Implementation | ns/op | allocs/op |
|----------------|-------|-----------|
| Montgomery/Bitwise | ~154,000,000 | 6,327 |
| Montgomery/CIOS | ~17,400,000 | 616,535 |
| Montgomery/CIOSWords | ~4,600,000 | 15,808 |
| BigInt/Exp | ~1,900,000 | 22 |

Note: `BigInt/Exp` uses Go's optimized Montgomery multiplication internally, serving as a reference for production-grade performance.
