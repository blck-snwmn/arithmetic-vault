# montgomery

Montgomery multiplication implementation for efficient modular arithmetic.

## Implementations

- `MontgomeryBitwise` - Basic bit-by-bit REDC algorithm
- `MontgomeryCIOS` - CIOS algorithm using big.Int internally
- `MontgomeryCIOSWords` - CIOS algorithm using []uint64 for better performance

## Benchmark

```bash
go test -bench=. -benchmem
```
