# ch4 數位邏輯設計

## 4-1 邏輯電路 / Logic Circuit

```
X   Y   sum / X^Y   carry / X*Y
0   0   0           0
0   1   1           0
1   0   1           0
1   1   0           1
```

```math
sum(x, y) = x + y - 2 * x * y
carry(x, y) = x * y
```

### AND gate

```
X   Y   X*Y
0   0   0
0   1   0
1   0   0
1   1   1
```

```math
f(x, y) = x * y
```

- [及閘](https://zh.wikipedia.org/zh-tw/%E4%B8%8E%E9%97%A8)
- [AND gate](https://en.wikipedia.org/wiki/AND_gate)

### OR gate

```
X   Y   X+Y
0   0   0
0   1   1
1   0   1
1   1   1
```

```math
f(x, y) = x + y - x * y
```

- [或閘](https://zh.wikipedia.org/zh-tw/%E6%88%96%E9%97%A8)
- [OR gate](https://en.wikipedia.org/wiki/OR_gate)

### NOT gate

```
X   X'
0   1
1   0
```

```math
f(x) = 1 - x
```

- [反相器](https://zh.wikipedia.org/zh-tw/%E5%8F%8D%E7%9B%B8%E5%99%A8)
- [Inverter](<https://en.wikipedia.org/wiki/Inverter_(logic_gate)>)

### NAND gate

```
X   Y   (X*Y)'
0   0   1
0   1   1
1   0   1
1   1   0
```

```math
f(x, y) = 1 - (x * y)
= 1 - x * y
```

- [反及閘](https://zh.wikipedia.org/zh-tw/%E4%B8%8E%E9%9D%9E%E9%97%A8)
- [NAND gate](https://en.wikipedia.org/wiki/NAND_gate)

### NOR gate

```
X   Y   (X+Y)'
0   0   1
0   1   0
1   0   0
1   1   0
```

```math
f(x, y) = 1 - (x + y - x * y)
= 1 - x - y + x * y
```

- [反或閘](https://zh.wikipedia.org/zh-tw/%E6%88%96%E9%9D%9E%E9%97%A8)
- [NOR gate](https://en.wikipedia.org/wiki/NOR_gate)

### XOR gate

```
X   Y   X^Y = X * Y' + X' * Y
0   0   0
0   1   1
1   0   1
1   1   0
```

```math
f(x, y) = x * (1 - y) + (1 - x) * y
= x + y - 2 * x * y
```

- [互斥或閘](https://zh.wikipedia.org/zh-tw/%E5%BC%82%E6%88%96%E9%97%A8)
- [XOR gate](https://en.wikipedia.org/wiki/XOR_gate)

### XNOR gate

```
X   Y   (X^Y)'
0   0   1
0   1   0
1   0   0
1   1   1
```

```math
f(x, y) = 1 - (x + y - 2 * x * y)
= 1 - x - y + 2 * x * y
```

- [反互斥或閘](https://zh.wikipedia.org/zh-tw/%E5%90%8C%E6%88%96%E9%97%A8)
- [XNOR gate](https://en.wikipedia.org/wiki/XNOR_gate)

---

## 4-2 布林代數 / Boolean Algebra

---

## 4-3 邏輯閘

---

## 4-4 邏輯簡化

---

## 4-5 組合電路

---

## 4-6 常見的組合電路
