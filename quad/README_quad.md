# Quad

Five small Go programs that draw rectangles ("quads") in the terminal. Each quad prints the same shape with a different set of characters, given a width `x` and a height `y`.

Built with Go and the [`z01`](https://github.com/01-edu/z01) package for printing runes.

## Project structure

```
quad/
├── quad/                 # Go module "quad" (one program per quad)
│   ├── go.mod
│   ├── go.sum
│   ├── quadA/main.go
│   ├── quadB/main.go
│   ├── quadC/main.go
│   ├── quadD/main.go
│   ├── quadE/main.go
│   └── test/main.go      # Runner that calls all five quads
└── test.go/              # Separate module "test": all five quads in one file
    ├── go.mod
    ├── go.sum
    └── test.go
```

## Requirements

- Go 1.22.2 or newer
- Internet access on first run, to download `github.com/01-edu/z01`

## Usage

### Run a single quad

From the inner `quad/` folder:

```bash
cd quad
go run ./quadA
go run ./quadB
go run ./quadC
go run ./quadD
go run ./quadE
```

Each program calls its function with four sample sizes: `(5,3)`, `(5,1)`, `(1,1)` and `(1,5)`.

### Try any size

Open `test.go/test.go`, uncomment a call in `main()` and change the numbers:

```go
func main() {
	QuadA(8, 4)
}
```

Then run it:

```bash
cd test.go
go run .
```

## The quads

Every quad has the signature `QuadX(x, y int)`: `x` is the width, `y` is the height. If either is negative, nothing is printed.

| Quad | Corners (top-left, top-right, bottom-left, bottom-right) | Top/bottom edge | Side edge |
|------|----------------------------------------------------------|-----------------|-----------|
| A    | `o` `o` `o` `o`                                          | `-`             | `\|`      |
| B    | `/` `\` `\` `/`                                          | `*`             | `*`       |
| C    | `A` `A` `C` `C`                                          | `B`             | `B`       |
| D    | `A` `C` `A` `C`                                          | `B`             | `B`       |
| E    | `A` `C` `C` `A`                                          | `B`             | `B`       |

### Example output for `(5, 3)`

**Quad A**
```
o---o
|   |
o---o
```

**Quad B**
```
/***\
*   *
\***/
```

**Quad C**
```
ABBBA
B   B
CBBBC
```

**Quad D**
```
ABBBC
B   B
ABBBC
```

**Quad E**
```
ABBBC
B   B
CBBBA
```

### Edge cases

| Size  | Result                                                 |
|-------|--------------------------------------------------------|
| `1×1` | A single corner character (the top-left one)           |
| `5×1` | One row: left corner, edge characters, right corner    |
| `1×5` | One column: top corner, side characters, bottom corner |

## Known issues

- `quad/test/main.go` imports `quad` as a package, but the quad functions are in `package main` inside each `quadX/` folder, so it won't compile as is. Use `test.go/` instead, or move the functions into a shared non-`main` package.
