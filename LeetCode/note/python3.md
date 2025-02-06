# python3

---

## content

- [arithmetic](#arithmetic)
- [list and dict](#list-and-dict)
- [dict](#dict)
- [linked list](#linked-list)
- [string](#string)

---

## arithmetic

```python
# slow
q, r = divmod(s, 10)

# fast
r = s % 10
q = s // 10
```

```python
# slow
if q != 0:
    ...

# fast
if q:
    ...
```

```python
# slow
val, node = (node.val, node.next) if node else (0, None)

# fast
val = 0
if node:
    val = node.val
    node = node.next
```

---

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

## linked list

```python
# slow
while True:
    ...
    node1 = node1.next
    node2 = node2.next
    if node1 == None and node2 == None:
        break

# fast
while node1 != None or node2 != None:
    ...
    node1 = node1.next
    node2 = node2.next

# faster
while node1 or node2:
    ...
    node1 = node1.next
    node2 = node2.next
```

```python
# require head node condition
prev = None
cur = head
while (cur is not None) and (cur != spec):
    prev = cur
    cur = cur.next

if cur == head:
    if cur.next is None:
        head = None
    else:
        head = cur.next
else:
    if cur is None:
        prev.next = None
    else
    prev.next = cur.next
# head

# dummy node, not extra head node
dummy = ListNode(0)
dummy.next = head
prev = dummy
cur = dummy.next
while (cur is not None) and (cur != spec):
    prev = cur
    cur = cur.next

if cur is None:
    prev.next = None
else:
    prev.next = cur.next
head = dummy.next
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
