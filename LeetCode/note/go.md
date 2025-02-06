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

```go
m = len(str1)
n = len(str2)
// n > m

strings.Contains(str1, str2)

for i := 0; i < n-m; i++ {
	if str1 == str2[i:i+m] {
		return true
	}
}
return false
```
