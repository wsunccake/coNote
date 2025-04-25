# GDT / Global Descriptor Table

在 x86 架構中，Global Descriptor Table (GDT) 是用來描述不同記憶體段的資料結構，並且其設計會依照不同模式（16-bit, 32-bit, 64-bit）來決定如何處理段選擇器、段描述子，以及如何處理保護模式和虛擬內存。

## Protected Mode / 32-bit

在 Protected Mode 下，記憶體管理的機制比 Real Mode 複雜且強大，並且 Segment（段） 和 GDT（Global Descriptor Table） 是其中非常核心的概念。這些機制使得系統能夠實現虛擬記憶體、記憶體保護、以及多工處理。

### Protected Mode - GDT

1. 什麼是 GDT？

GDT 是一個存放 段描述符 的表格。它是 Protected Mode 記憶體管理的核心之一，允許操作系統定義和管理不同的段及其屬性。

- GDT 中存儲的是 全局段描述符，這些描述符描述了系統中所有段的結構和屬性。
- 每個段都會在 GDT 中有一個條目，這些條目包含基址、界限、屬性等信息。

2. GDT 的結構

GDT 是一個包含 段描述符 的表格，每個段描述符的大小通常為 8 字節。GDT 本身也有一個特殊的 GDT 描述符，它存儲 GDT 的大小和基址。
一個 GDT 描述符的基本結構如下：

```
字段 | 大小 | 描述
基址（Base） | 32位 | 段的基址（存儲段的起始位置）
類型（Type） | 4位 | 指定段的類型（例如代碼段、數據段等）
存取權限 | 8位 | 指定段的存取權限（如是否可讀、可寫等）
顯示性 | 3位 | 控制段是否可見，或段是否可在用戶模式訪問
粒度 | 4位 | 控制段的大小單位（例如每個段的大小是 4KB 還是 1 字節）
邊界（Limit） | 20位 | 段的大小，即該段能包含的最大字節數
```

3. GDT 的初始化

在系統啟動時，操作系統會將 GDT 設置為一個有效的內存區域，並在其中創建一些段描述符。操作系統會設置 GDT 以確保內核段、用戶段以及其他所需段都有相應的描述符。

例如，一個典型的 GDT 可能包含以下描述符：

- 一個 空段（NULL descriptor），用於保護。
- 一個 代碼段 描述符，表示程式代碼段。
- 一個 數據段 描述符，表示程式數據段。
- 一個 堆疊段 描述符，表示堆疊區段。

4. GDT 設置過程

在操作系統啟動時，必須設置 GDT 的基址（基準地址）和大小。這樣 CPU 就能夠正確地查找並使用這些段描述符。
典型的 GDT 初始化過程：

```asm
lgdt [gdt_descriptor]   ; 加載 GDT 描述符
                        ; lgdt（Load Global Descriptor Table Register）是 x86 CPU 的 特權指令
                        ; 用來將一個內存中的 GDT 描述符（limit + base）載入到 CPU 的 GDTR 暫存器中
```

5. GDT 描述符結構

GDT 描述符 會告訴 CPU GDT 的基地址和大小（gdtr 寄存器）。它包含兩個部分：

- 基址（Base Address）：指向 GDT 表格的起始地址。
- 大小（Limit）：GDT 表格的大小，告訴 CPU 表格的範圍。

### Protected Mode 的 Segment 和 GDT 的特性

- 記憶體保護：使用 GDT 和段描述符來設定不同段的存取權限。比如，代碼段可以設置為只讀、數據段可以設置為可讀可寫，這樣防止了越界寫入和錯誤存取。
- 虛擬記憶體支持：GDT 中的段描述符可以與操作系統的虛擬記憶體機制協作，為每個進程提供獨立的虛擬記憶體區域。
- 多任務管理：GDT 支持多任務，操作系統可以為每個進程創建不同的段描述符，並在上下文切換時加載相應的 GDT 描述符。

---

## Long Mode / 64-bit

在 Long Mode（x86-64、64-bit 模式）下，Paging（分頁機制） 和 GDT（Global Descriptor Table） 是進入與運作系統的兩大關鍵基礎。這兩者在 64-bit 模式的地位和功能跟 32-bit 有很大差異。

進入 Long Mode 必須同時：

1. 開啟 PAE（Physical Address Extension）
2. 啟用 分頁（Paging）
3. 並設定 LME（Long Mode Enable） 位元
4. 使用 longjmp 的方式跳轉到 64-bit code 段（CS 有 L-bit）

### Long Mode - GDT

在 64-bit 模式中，Segmentation 幾乎被停用（除了 FS 和 GS 可自訂 base），但 CPU 仍然需要透過 GDT 提供一些基本的段資訊來進行模式轉換、特權管理：

```
用途                  | 是否有效        | 說明
CS code segment       | ✅              | 指示 L-bit=1，進入 64-bit code
DS/SS/其他段          | ❌（忽略 base） | 除 FS/GS 外不再使用 base/limit
FS / GS base          | ✅              | 仍可用，常見於 TLS、per-CPU data
特權等級（DPL）       | ✅              | Ring 0 / Ring 3 控制仍存在
虛擬化 / 保護模式用途 | ✅              | TSS、LDT 還在使用 GDT
```

GDT 最基本要準備的段描述符：

- Null Segment（0x00） ← 必須，防止 0 index 被誤用
- 64-bit Code Segment（0x08） ← 設定 L 位為 1
- Data Segment（0x10） ← 雖然無用，但仍需有效的 selector
- 可選：User-mode Code / Data（Ring 3）

GDT Descriptor

```
欄位  | 說明
base  | 64-bit 中被忽略（除了 FS/GS）
limit | 忽略
type  | 必須設為 code/data
DPL   | 決定是否是 Ring 0 / Ring 3
L-bit | 必須為 1（表示 64-bit code）
D-bit | 必須為 0（在 L=1 下無效）
G-bit | 通常設為 1，粒度 4KB
```

### Long Mode - Paging

Paging 是「必要」的，沒有 Paging 就不能進入 Long Mode！在 Long Mode 下，Segmentation 幾乎被停用，真實的虛擬記憶體管理完全交給 Paging！
Long Mode 使用的 Paging 格式：4-Level Paging

```
虛擬地址結構（48-bit）：

[ 47 - 39 ] → PML4 Index (512 entries)
[ 38 - 30 ] → PDP  Index (512 entries)
[ 29 - 21 ] → PD   Index (512 entries)
[ 20 - 12 ] → PT   Index (512 entries)
[ 11 -  0 ] → Offset within 4KB page
```

各層說明

```
層級 | 名稱                            | 大小 | 說明
L4   | PML4                            | 512  | 每個 entry 指向 PDP 表
L3   | Page Directory Ptr Table (PDPT) | 512  | 每個 entry 指向 PD 表
L2   | Page Directory (PD)             | 512  | 可指向 Page Table，或直接映射 2MB 頁
L1   | Page Table                      | 512  | 每 entry 映射 4KB 實體頁
```

Paging Entry（每層）基本格式

```
位元範圍 | 意義
[0]      | Present (1=存在)
[1]      | R/W (1=可寫)
[2]      | U/S (1=使用者模式可用)
[7]      | Page Size（PS）(2MB/1GB 時有效)
[12-51]  | 實體頁框（Physical Frame Address）
```

啟用 Long Mode 流程

1. 開啟 A20 線
2. 切到 Protected Mode
3. 開啟 PAE：cr4 |= (1 << 5)
4. 設定 4-level Page Tables（PML4 等）
5. 設定 cr3 = pml4_address
6. 設定 efer MSR 中的 LME（Long Mode Enable） 位
7. 設定 cr0 |= (1 << 31) 開啟 Paging
8. 使用 Far Jump 跳進 64-bit code segment（L=1）

怎麼用 NASM 製作 Long Mode bootloader
怎麼寫 page tables（簡化映射 1:1 到實體記憶體）
怎麼做 GDT + paging + jmp far 切入 64-bit code
