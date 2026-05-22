# arithmetic-vault

A collection of arithmetic operations.

## Packages

- `montgomery` - Montgomery multiplication (three implementations: Bitwise, CIOS, CIOSWords)
- `pollard` - Pollard's rho algorithm for integer factorization using Floyd's cycle detection
- `rabin` - Miller-Rabin probabilistic primality test
- `karatsuba` - Karatsuba multiplication algorithm for fast integer multiplication
- `pseudo-mersenne-reduction` - Fast modular reduction for pseudo-Mersenne primes of the form `2^n - c`

## Development

CLI tools (`golangci-lint`, `lefthook`) are managed by [aqua](https://aquaproj.github.io/) with versions pinned in [aqua.yaml](aqua.yaml).

### Install tools

Install aqua itself first (see the [aqua installation guide](https://aquaproj.github.io/docs/install)), then install the pinned tools:

```
aqua install
```

### Set up git hooks

[lefthook](lefthook.yml) runs `make lint` on staged `*.go` files before each commit. Register the hooks once after cloning:

```
lefthook install
```
