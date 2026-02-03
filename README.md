# csv2md

Convert CSV files to Markdown tables.

## Installation
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

output.md (when rendered)
```md
| col1 |       col2 |  really_long_col3 |
|------|------------|-------------------|
|    1 | long entry |              true |
|    2 |    shorter |             false |
```


output.md (as text)
```
| col1 |       col2 |  really_long_col3 |
|------|------------|-------------------|
|    1 | long entry |              true |
|    2 |    shorter |             false |
```
