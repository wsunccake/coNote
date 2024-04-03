# Big O

## Time Complexity

```c
// time: O(n)
// space: O(n)
int sum(int n) {
    if (n <=) {
        return 0;
    }
    return n + sum(n-1);
}
```

```c
// time: O(n)
// space: O(1)
int pairSumSequence(int n) {
    int sum = 0;
    for (int i =0; i < n; i++) {
        sum += pairSum(i, i+1);
    }
    return sum;
}

int pairSum(int a, int b) {
    return a + b;
}
```

## Space Complexity
