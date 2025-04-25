# Segmentation

「Segment 分段」是一種記憶體管理技術，主要用於 x86 架構，尤其是早期的 16-bit 和 32-bit 模式中。分段的主要目的是將記憶體劃分為邏輯區塊，每一區塊（Segment）都包含一類特定的資料，例如程式碼、資料或堆疊。

分段（Segmentation）是將記憶體空間分為數個段（Segment），每個段代表一段邏輯意義的區域，例如：

- Code Segment（CS）：儲存程式碼的區域。
- Data Segment（DS）：儲存變數、資料的區域。
- Stack Segment（SS）：堆疊區域，儲存函式呼叫記錄、區域變數等。
- Extra Segment（ES/FS/GS）：額外的資料區段，可由程式控制使用。

## Real Mode / 16-bit

Real mode（實模式）是 x86 CPU 開機後的預設狀態。它有以下特性：

- 只能存取 1MB 的記憶體（實際上是 20-bit 位址線）
- 沒有記憶體保護機制
- 支援 segmentation，但方式非常簡單（且粗糙）

在 real mode 中，每個記憶體位址是由兩部分組成的：

```
Physical Address = Segment × 16 + Offset
```

Segment 是一個 16-bit 的數值，通常是存在 segment register 裡（例如：CS、DS、SS、ES）
Offset 是一個 16-bit 的偏移值
由於 Segment × 16 = Segment << 4（左移四位），所以 Segment 決定了「基底地址」，Offset 則是相對於這個基底的偏移。

Code Segment: CS
Stack Segment: SS
Data Segment: DS, ES, GS, FS

```asm
mov ax, 0x1234  ; 把 0x1234 複製到 AX 暫存器
mov ds, ax      ; 把 AX 中的值 0x1234 複製到段暫存器 DS
                ; 從現在起，所有以 DS 爲 segment base 的資料存取，都從實體位址 0x1234 × 16 = 0x12340 起算
mov bx, 0x0010  ; 把 0x0010 複製到 BX 暫存器
mov al, [ds:bx] ; 從位於 DS:BX 所指定的記憶體位置讀取一個 byte（8-bit），放進 AL
                ; DS:BX = 0x1234:0x0010
                ; → 實體地址 = 0x1234 * 16 + 0x0010 = 0x12340 + 0x0010 = 0x12350
                ; DS:BX 當變數使用 存取在記憶體上
```

ax = ah + al

這段程式碼會存取的實體位址是

```
0x1234 × 16 + 0x0010 = 0x12340 + 0x0010 = 0x12350
```

特點

- 總共可表示的最大位址：0xFFFF × 16 + 0xFFFF = 0x10FFEF（= 1,114,095 bytes，約 1.08MB）
- 實際上 8086 只能使用到前 1MB（0xFFFFF）
- 記憶體空間有「重疊」：不同 segment/offset 組合可能對應同一個實體位址

優點

- 解決 16-bit CPU 的限制：透過這種「segment + offset」方式，實際上可以表達超過 16-bit 的位址範圍（理論上到 20-bit）。
- 相對簡單：不需要表格、不需要額外硬體支援，直接透過左移計算。

缺點

- 沒有記憶體保護：任何程式可以隨意存取任意記憶體，容易錯誤或被惡意利用。
- 記憶體重疊導致混亂：不同的段與偏移可能指向相同位址，對程式設計者來說容易出錯。
- 只能存取到 1MB 記憶體：對現代應用太小了。

段重疊 / Segment Overlap

在 Real Mode 下，segment × 16 + offset 計算出的實體位址（Physical Address）可能會重複（overlap），即使 segment 值不同。也就是說：不同的段組合（Segment:Offset）可以對應到同一個實體位址。

```
0x1234:0x005	0x12340 + 0x005 = 0x12345
0x1200:0x0345	0x12000 + 0x0345 = 0x12345
0x1000:0x2345	0x10000 + 0x2345 = 0x12345
```

---

## Protected Mode / 32-bit

在 Protected Mode 下，記憶體管理的機制比 Real Mode 複雜且強大，並且 Segment（段） 和 GDT（Global Descriptor Table） 是其中非常核心的概念。這些機制使得系統能夠實現虛擬記憶體、記憶體保護、以及多工處理。

### Protected Mode - Segmentation

1. 段的基本概念

在 Protected Mode 中，段不僅僅是將記憶體劃分為不同的區塊，它還可以提供更多的安全和保護功能。每個段都有與之相關聯的 描述符（Descriptor），描述符包含有關段的所有信息，例如段的基址、大小、屬性等。

2. 段的運作方式

每個段都會與一個 段描述符（Segment Descriptor） 對應，這些描述符存儲在 GDT 或 LDT（Local Descriptor Table）中。
段寄存器（如 CS, DS, SS, ES 等）只包含該段描述符在 GDT 或 LDT 中的索引值。這些段描述符提供了段的基址（Base Address）、段大小（Limit）、權限（Access Rights）等信息。

3. 如何存取段

在 Protected Mode 中，段寄存器（如 CS、DS 等）指向 GDT 中的段描述符，通過這些描述符來查詢該段的詳細信息。當 CPU 需要訪問記憶體時，它會查找該段的描述符，然後計算出實際的物理位址。

4. 段描述符的結構

每個段描述符都有以下幾個關鍵字段：

- 基址（Base Address）：指向段的起始地址。
- 段界限（Limit）：定義該段的大小或範圍，範圍從 0 到 Limit。
- 屬性（Access Rights）：指定該段的存取權限，如是否可讀、可寫、可執行等。
- 粒度（Granularity）：定義段大小單位的大小（如 4KB 還是 1 字節等）。

5. 段的類型

- 代碼段（Code Segment）：存放程式代碼，通常設置為可執行且只讀。
- 數據段（Data Segment）：存放變數和數據，通常是可讀可寫的。
- 堆疊段（Stack Segment）：用於程序的堆疊，通常設置為可寫。
- 特權段（System Segment）：為操作系統核心區域提供保護。
