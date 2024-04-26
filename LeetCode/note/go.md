# go

## slice

```go
dp := make([][]int, m)
for i := range dp {
	dp[i] = make([]int, n)
}
```

## string

```go
num := 1
tmpStr := ""

// slow
tmpStr += fmt.Sprint(count)

// fast
tmpStr += strconv.Itoa(count)
```
