# python3

## list and dict

```python
# with list ==, !=
list1 = []
list2 =[]
list1 == list2

# without dict ==, !=
dict1 = []
dict2 =[]
dict1 == dict2
```

```python
# slow
dict1 = {}
dict2 = {}

for _ in range(n):
    ...
    if dict1 == dict2:
        return True

# fast
list1 = [0] * 26
list2 = [0] * 26

for _ in range(n):
    ...
    if list1 == list2:
        return True
```

---

## dict

```python
# slow
for char in char_list:
    if char in dict1:
        dict1[char] += 1
    else:
        dict1[char] = 1

# fast
for char in char_list:
    dict1[char] = dict1.get(char, 0) + 1
```

```python
# slow
for _, v in dict1.items():
    if v != 0:
        return False

# fast
for v in dict1.values():
    if v != 0:
        return False
```

---

## string

```python
# slow
sorted(list1) == sorted(list2)

# fast
''.join(sorted(list1)) == ''.join(sorted(list2))
```

```python
m = len(str1)
n = len(str2)
# n > m

# slow
for i in range(n):
    for j in range(m):
        if str1[j] != str2[i+j]:
            break
    if str1[j] != str2[i+j]:
        return True
return False

# normal
for i in range(n-m):
    if str1 == str2[i:i+m]:
        return True
return False

# trick
if str1 in str2:
    return True
return False
```
