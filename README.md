# Repo Analyzer

> go library to find all modules / languages / build-systems / specifications in a project

## Usage (Library)

__Installation__

```bash
go get -u github.com/cidverse/repoanalyzer
```

__Example__

```go
func main() {
    analyzer := repoanalyzer.NewAnalyzer()
    result := analyzer.Scan("/my-project")
    for k, v := range result {
        fmt.Printf("MODULE %d: %+v\n", k, v)
    }
}
```

## Usage (CLI)

...

## License

Released under the [MIT License](./LICENSE).
