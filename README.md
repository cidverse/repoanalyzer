# Repo Analyzer

> go library to discover various files of interest in a repository

## Usage (Library)

__Installation__

```bash
go get -u github.com/cidverse/repoanalyzer
```

__Example__

```go
result := analyzer.ScanDirectory("/my-project")
for k, v := range result {
    fmt.Printf("MODULE %d: %+v\n", k, v)
}
```

__Custom Example__

You can add your own scanners and enable result caching.

```go
myAnalyzer := analyzer.NewAnalyzer()
myAnalyzer.EnableCache(true)
// myAnalyzer.Add(yourCustomScanner)
result := myAnalyzer.Scan("/my-project")
for k, v := range result {
    fmt.Printf("MODULE %d: %+v\n", k, v)
}
```

## Usage (CLI)

...

## License

Released under the [MIT License](./LICENSE).
