# CPE

---

## content

- [read data](#read-data)
    - [read data to variable and end of line](#read-data-to-variable-and-end-of-line)
    - [read line and end of line](#read-line-and-end-of-line)
    - [read n line data](#read-n-line-data)
    - [read n line data and specfic data to end](#read-n-line-data-and-specfic-data-to-end)

---

## read data

### read data to variable and end of line

```text
1 10
100 200
201 210
900 1000
```

```cpp
    int a, b
    while (cin >> a >> b)
    {
        ...
    }
```

### read line and end of line

```text
"To be or not to be," quoth the Bard, "that
is the question".
The programming contestant replied: "I must disagree.
To `C' or not to `C', that is The Question!"
```

```cpp
    string line;
    while (getline(cin, line))
    {
        ...
    }
```

### read n line data

```text
3
This is a test.
Count me 1 2 3 4 5.
Wow!!!!  Is this question easy?
```

```cpp
    int n;
    string s;

    cin >> line;

    // method 1.
    for (i = 0; i < n; i++)
    {
        while (cin >> line)
        {
            ...
            if (cin.get() == '\n')
                break;
        }
    }

    // method 2.
    while (--n && cin.get() == '\n')
    {
        while (cin >> s)
        {
            ...
        }
    }
```

### read n line data and specfic data to end

```text
123 456
555 555
123 594
999 1
0 0
```

```cpp
    int a, b;

    // method 1.
    while (cin >> a >> b)
    {
        if (a == 0 && b == 0)
            break;

        ...
    }

    // method 2.
    while (cin >> a >> b && a != 0 && b != 0) {
        ...
    }
```