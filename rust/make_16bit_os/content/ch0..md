# ch0. 規劃學習篇章（共 5 個階段）

## 階段一：底層基礎知識補強

- 瞭解 X86 架構的 實模式 (Real Mode) 與 16-bit 暫存器（AX, BX, CX, DX, SI, DI, SP, BP, CS, DS）。
- 記憶體定址原理：分段機制（Segment:Offset, Segment \* 16 + Offset）。
- 理解 BIOS 中斷機制（例如 INT 10h 用於螢幕輸出、INT 13h 用於讀寫磁碟）。
- 補強基本組合語言 (x86 Assembly) 語法與 C/Rust 的指標（Pointer）概念。

## 階段二：Rust 系統級開發入門 (Bare-metal Rust)

- 學習 Rust 基礎語法（所有權、借用、結構體、Trait）。
- 掌握 #![no_std] 模式：如何在沒有作業系統（無標準庫 std、無 println!、無 dynamic memory allocator）的情況下編譯 Rust。
- 瞭解 Rust unsafe 區塊的使用場景（直接存取裸指標與硬體 IO 埠）。
- 配置 Rust 交叉編譯目標（Target triplet，如 i686-unknown-linux-gnu 或自訂 Target JSON）。

## 階段三：引導程式與系統啟動 (Bootloader & 16-bit Mode)

- 撰寫 512 位元組的 MBR (Master Boot Record) Bootloader（以 NASM 或 Rust 內嵌彙編撰寫）。
- 在螢幕上印出第一行字串 "Hello World"（呼叫 BIOS INT 10h 或直接寫入 VGA 記憶體 0xB8000）。
- 建立 Rust 與 16-bit 組合語言的呼叫介面（C ABI 規範）。

## 階段四：DOS 風格核心功能實作 (Core OS Features)

- 鍵盤輸入處理：透過 BIOS 中斷（INT 16h）擷取使用者按鍵。
- 簡單 Command Line (CLI) 介面：實作字符緩衝區，支援基本的命令解析（如 help, cls, dir）。
- 記憶體管理：實作極簡的靜態記憶體配置器或自訂 Allocator。
- 簡易檔案系統 (FAT12)：解析軟碟 (Floppy Disk) 結構，讀取磁碟區塊與檔案目錄。

## 階段五：除錯、測試與轉移

- 學習使用 QEMU + GDB 觀察 BIOS 引導與記憶體暫存器變化。
- 系統穩定後，將編譯出的 ISO/IMG 映像檔掛載至 VMware 中驗證真實虛擬機器環境的相容性。
