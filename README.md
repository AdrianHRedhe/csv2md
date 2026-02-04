# csv2md

Convert CSV files to Markdown tables.

## Installation
### Using homebrew
```bash
brew install adrianhredhe/tap/csv2md
```

### Using go
```bash
go install github.com/adrianhredhe/csv2md@latest
```

## Usage
```bash
# From file
csv2md input.csv -o output.md

# From stdin
cat data.csv | csv2md

# Using flags
csv2md -i input.csv -o output.md
```

## Examples
input.csv
```csv
col1, col2, really_long_col3
1,long entry, true
2,shorter, false
```

output.md
```
| col1 |       col2 |  really_long_col3 |
|------|------------|-------------------|
|    1 | long entry |              true |
|    2 |    shorter |             false |
```

## Repo structure
```
.
├── .gignore
├── .github
│   └── workflows
│       └── release.yml  # Automated release on version tags
├── .goreleaser.yml      # Multi-platform build configuration
├── go.mod               # Module definition (no external deps)
├── main.go              # csv2md CLI implementation
├── main_test.go         # unit tests for cli implementation
├── LICENCE
└── README.md
```
