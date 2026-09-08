# ch1. 練習題與實作指南

## 練習一：暫存器操作與 BIOS 視訊印字 (INT 10h)

- 目標：掌握 x86 暫存器（AX, BX）賦值語法，並透過 BIOS 中斷 INT 10h 於螢幕輸出字元。
- 觀念複習：AH=0x0E 是 BIOS 印字功能碼，AL 存 ASCII 碼，INT 0x10 發動呼叫。

1. 完整程式碼 (lab1_1.asm)

```assembly
[org 0x7c00]        ; 告知編譯器，這段代碼會被 BIOS 載入到記憶體 0x7C00 位址

    mov ah, 0x0E    ; AH = 0x0E (BIOS Teletype 顯示功能)

    mov al, 'O'     ; 將字元 'O' 存入 AL 暫存器
    int 0x10        ; 呼叫 BIOS 視訊中斷

    mov al, 'S'     ; 將字元 'S' 存入 AL 暫存器
    int 0x10        ; 呼叫 BIOS 視訊中斷

    mov al, '!'     ; 將字元 '!' 存入 AL 暫存器
    int 0x10        ; 呼叫 BIOS 視訊中斷

hang:
    jmp hang        ; 無限迴圈，防止 CPU 繼續執行後方的未定義記憶體

; MBR 512 位元組填充與開機簽章
times 510-($-$$) db 0   ; 填充 0 直到第 510 個位元組
dw 0xAA55               ; 第 511-512 位元組寫入 0xAA55 (MBR 魔術數字)
```

1. 操作流程與執行步驟

```bash
# 1. 使用 NASM 編譯成裸機機器碼 (Raw Binary)
nasm -f bin lab1_1.asm -o lab1_1.bin

# 2. 透過 QEMU 模擬 x86 機器啟動該 Binary 檔案
qemu-system-x86_64 -fda lab1_1.bin
```

3. 預期畫面輸出

QEMU 會跳出一個視窗，黑底白字顯示：

```
Booting from Hard Disk...
OS!
```

---

## 練習二：分段機制與記憶體定址 (Segment:Offset)

- 目標：驗證實體位址計算公式 $\text{Physical Address} = (\text{Segment} \times 16) + \text{Offset}$。
- 觀念複習：資料暫存器 DS 無法直接 mov ds, 0x07C0，必須透過 AX 中轉。1.

1. 完整程式碼 (lab1_2.asm)

```assembly
[org 0x0000]        ; 將 Offset 設為 0x0000

    ; 欲存取實體位址 0x7C00：
    ; 設定 Segment = 0x07C0 -> 0x07C0 * 16 (0x10) = 0x7C00
    mov ax, 0x07C0
    mov ds, ax      ; DS = 0x07C0

    mov ah, 0x0E    ; BIOS 印字功能

    ; 從 DS:my_text (即 0x07C0:my_text) 讀取第一個位元組
    mov al, [my_text]
    int 0x10

    ; 讀取第二個位元組 (Offset + 1)
    mov al, [my_text + 1]
    int 0x10

hang:
    jmp hang

my_text:
    db 'HI', 0      ; 定義資料字串 'H', 'I'

times 510-($-$$) db 0
dw 0xAA55
```

2. 操作流程與執行步驟

```bash
# 編譯並執行
nasm -f bin lab1_2.asm -o lab1_2.bin
qemu-system-x86_64 -fda lab1_2.bin
```

3. 預期畫面輸出

視窗左上角輸出：

```
HI
```

---

## 練習三：直接寫入 VGA 顯存 (0xB8000)

- 目標：不透過 BIOS 中斷，直接對獨立顯示記憶體 (Video RAM) 進行指標式讀寫，並設定文字前景/背景顏色。
- 觀念複習：文字模式下 0xB8000 為顯存首位址。每個字元佔 2 位元組（Byte 0 = ASCII 碼, Byte 1 = 顏色屬性）。

### 1. 完整程式碼 (lab1_3.asm)

```assembly
[org 0x7c00]

    ; 1. 將附加段暫存器 ES 指向 VGA 顯存段位址 0xB800
    ; 0xB800 * 16 = 0xB8000 實體位址
    mov ax, 0xB800
    mov es, ax

    ; 2. 寫入第一個字元 'R'，背景黑(0)、前景紅(4) -> 屬性碼 0x04
    mov byte [es:0x0000], 'R'   ; ASCII 字元
    mov byte [es:0x0001], 0x04  ; 顏色屬性

    ; 3. 寫入第二個字元 'U'，背景黑(0)、前景綠(2) -> 屬性碼 0x02
    mov byte [es:0x0002], 'U'
    mov byte [es:0x0003], 0x02

    ; 4. 寫入第三個字元 'S'，背景黑(0)、前景青藍(3) -> 屬性碼 0x03
    mov byte [es:0x0004], 'S'
    mov byte [es:0x0005], 0x03

    ; 5. 寫入第四個字元 'T'，背景黑(0)、前景黃(E) -> 屬性碼 0x0E
    mov byte [es:0x0006], 'T'
    mov byte [es:0x0007], 0x0E

hang:
    jmp hang

times 510-($-$$) db 0
dw 0xAA55
```

### 2. 操作流程與執行步驟

```bash
# 編譯並執行
nasm -f bin lab1_3.asm -o lab1_3.bin
qemu-system-x86_64 -fda lab1_3.bin
```

### 3. 預期畫面輸出

QEMU 畫面左上角會出現帶有彩色的文字：

```
RUST
```

---

## 練習四：Rust 裸機模式 (#![no_std]) 與內嵌彙編

- 目標：理解如何在無作業 systems/C 庫（no_std）的環境下編譯 Rust 程式碼，並利用內嵌彙編 (Inline Assembly) 呼叫底層中斷。

### 1. 專案目錄配置

建立練習四的專案目錄：

```bash
mkdir -p ~/16bit-os-lab/lab1/rust_bare
cd ~/16bit-os-lab/lab1/rust_bare
```

2. 完整程式碼與設定檔

Cargo.toml

```toml
[package]
name = "rust_bare"
version = "0.1.0"
edition = "2021"

[profile.dev]
panic = "abort"     ; 裸機下不支援堆疊展開 (Unwinding)，恐慌時直接中止

[profile.release]
panic = "abort"
```

src/main.rs

```rust
#![no_std]          ; 1. 停用 Rust 標準庫 (std)
#![no_main]         ; 2. 停用 C/Rust 標準 entrypoint (main 函式)

use core::arch::asm;
use core::panic::PanicInfo;

// 3. 自訂 Panic Handler (無 std 時必須手動定義)
#[panic_handler]
fn panic(_info: &PanicInfo) -> ! {
    loop {}
}

// 4. 定義系統進入點 (指定外部符號 _start)
#[no_mangle]
pub extern "C" fn _start() -> ! {
    // 呼叫 BIOS INT 10h 印出字元 'OK'
    unsafe {
        // 印出 'O'
        asm!(
            "int 0x10",
            in("ax") 0x0e4f_u16, // AH=0x0E, AL='O' (0x4F)
        );
        // 印出 'K'
        asm!(
            "int 0x10",
            in("ax") 0x0e4b_u16, // AH=0x0E, AL='K' (0x4B)
        );
    }

    loop {}
}
```

### 3. 操作流程與執行步驟

在 ~/16bit-os-lab/lab1/rust_bare 目錄下執行：

```bash
# 1. 構建並編譯為裸機目標檔 (i686 獨立目標)
cargo build --target i686-unknown-linux-gnu

# 2. 檢視編譯出的 ELF 檔案資訊（驗證 no_std 編譯成功）
file target/i686-unknown-linux-gnu/debug/rust_bare
```

### 4. 預期 Terminal 輸出

file 命令會確認編譯出了一個 32-bit 的獨立可執行檔：

```
target/i686-unknown-linux-gnu/debug/rust_bare: ELF 32-bit LSB executable, Intel 80386, version 1 (SYSV), statically linked...
```

(註：此 Rust 產出的 ELF 檔後續將在階段三與 NASM Bootloader 進行 Link 轉化成真正的 16-bit 磁碟鏡像檔。)
