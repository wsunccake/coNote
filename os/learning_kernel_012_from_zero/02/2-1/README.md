# real mode / 16-bit boot loader

16-bit 的 bootloader，負責從磁碟載入 setup code 至記憶體特定位置，並顯示訊息後跳轉至 setup

## bootsect.s

```s
SETUPLEN = 4                    # 要讀入的 setup 程式長度 (磁區)
BOOTSEG = 0x07C0                # bootloader 初始 segment，BIOS 載入位置
INITSEG = 0x0900                # bootloader 要搬移到的 segment（更安全）
SETUPSEG = 0x9020               # 第二階段程式載入的 segment
SYSSEG = 0x1000                 # 作業系統主體要載入的 segment
ENDSEG = SYSSEG + SYSSIZE       # 作業系統載入結束位置 (你這邊還沒定義 SYSSIZE)
ROOT_DEV = 0x000                # 根裝置號（保留位）

.code16
.text

.global _start
_start:
    ljmp $BOOTSEG, $start2      # BIOS 開機時載入 bootloader 到 0x0000:0x7C00，這裡強制指定 segment 為 BOOTSEG = 0x07C0

start2:
# DS:SI 是來源記憶體地址。
# ES:DI 是目標記憶體地址。
# 從：DS:SI = 0x07C0:0000 → 實體地址 0x07C00
# 複製到：ES:DI = 0x0900:0000 → 實體地址 0x09000
# 複製 512 bytes
# 因為 bootloader 被 BIOS 載入在 0x7C00 這個位置，
# 這位置可能會被後面載入的 kernel 或 setup 程式蓋掉，
# 因此先搬到 0x9000 是一種 保護自己 的做法，
# 讓 bootloader 可以繼續控制流程而不會被覆蓋。
    movw $BOOTSEG, %ax      # BOOTSEG = 0x07C0
    movw %ax, %ds           # DS 設為目前位置（0x07C0）
    movw $INITSEG, %ax      # INITSEG = 0x0900
    movw %ax, %es           # ES 設為目標位置（0x0900）
    movw $256, %cx          # 256 word（1 word = 2 bytes）= 512 bytes

    xorw %si, %si           # subw %si, %si  # SI = 0（來源偏移）
    xorw %di, %di           # subw %di, %di  # DI = 0（目的偏移）

    rep movsw               # 字串複製指令 複製 512 bytes 從 BOOTSEG:0 → INITSEG:0
                            # movsw：將 DS:SI 指向的 word 複製到 ES:DI，並自動遞增 SI 與 DI。
                            # rep：重複執行 CX 次。
    jmp $INITSEG, $go

go:
    movw %cs, %ax
    movw %ax, %ds
    movw %ax, %es
    movw %ax, %ss           # 設定堆疊段位
    movw $0xFF00, %sp       # 設定堆疊指標（stack pointer）

load_setup:
    movw $0x0000, %dx       # 驅動器 0
    movw $0x0002, %cx       # 從磁區 2 開始（0 為 boot sector，1 多為保留）
    movw $0x0200, %bx       # ES:BX 為資料寫入位置（這邊預設）
    movb $SETUPLEN, %al     # 讀入幾個磁區
    movb $0x02, %ah         # 功能 02h: 讀磁區
    int $0x13
    jnc ok_load_setup       # 無錯則跳

    movw $0x0000, %dx
    movw $0x0002, %ax       # AH = 0: reset disk
    int $0x13
    jnc load_setup          # 成功後重試

ok_load_setup:
    movw $msg, %bp          # msg 字串指標
    movw $0x1301, %ax       # AH=13h (寫字串), AL=01h（使用游標、顏色）
    movw $0x000C, %bx       # 頁面=0，屬性=亮紅色 (text attr = 0x0C)
    movw $21, %cx           # 長度
    movb $0, %dl            # 第 0 行
    int $0x10

    jmp $SETUPSEG, $0

msg:
    .ascii "Setup has been loaded"

.org 508                    # 設定程式碼結尾要在 508 位元組處。
root_dev:
    .word ROOT_DEV          # 被 kernel 認用來設定根檔案系統位置。
boot_flag:
    .word 0xAA55            # BIOS 檢查 MBR 的標準結尾標誌。
```

---

## setup.s

```s
.code16
.text
.global _start_setup        # 讓 linker 或 loader 找到，是 setup 的進入點。

_start_setup:
    movw %cs, %ax
    movw %ax, %ds
    movw %ax, %es
    movw $setup_msg, %bp    # 字串 offset -> BP
    movw $0x1301, %ax       # AH=13h 顯示字串, AL=1 移動游標
    movw $0x0002, %bx       # BH=0（頁），BL=2（顏色：綠）
    movw $16, %cx           # CX=字串長度（手動指定）
    movb $3, %dh            # DH=列（第 4 行）
    movb $0, %dl            # DL=欄（第 1 欄）
    int $0x10               # 呼叫 BIOS 視訊中斷

# | 寄存器 | 用途         |
# |--------|--------------|
# | AH     | 0x13         |
# | AL     | 1 = 移動游標 |
# | BH     | video 頁面   |
# | BL     | 屬性（顏色） |
# | CX     | 長度         |
# | DL     | column       |
# | DH     | row（選擇性）|
# | ES:BP  | 字串指標     |

setup_msg:
    .ascii "setup is running"
```

---

## 16-bit real mode，INT（interrupt）

### 📺 INT 0x10 – 視訊服務（BIOS Video Services）

| AH   | 功能                   | 重要參數                              |
| ---- | ---------------------- | ------------------------------------- |
| 0x00 | 設定視訊模式           | AL = 模式（如 0x03 文字模式）         |
| 0x01 | 設定游標形狀           | CH = 起始掃描線，CL = 結束掃描線      |
| 0x02 | 設定游標位置           | BH = 頁，DH = row，DL = column        |
| 0x03 | 取得游標位置           | BH = 頁 → 回傳 DH/DL 位置             |
| 0x09 | 顯示字元（含屬性）多次 | AL=字元，BH=頁，BL=屬性，CX=次數      |
| 0x0E | 顯示字元（TTY 模式）   | AL=字元，BH=頁，BL=顏色（可略）       |
| 0x13 | 顯示字串               | AL=flag，ES:BP=字串，CX=長度，BL=屬性 |

### 💽 INT 0x13 – 磁碟服務（Disk Services）

| AH   | 功能                | 重要參數                                                         |
| ---- | ------------------- | ---------------------------------------------------------------- |
| 0x00 | 重置磁碟控制器      | DL = 裝置號碼                                                    |
| 0x02 | 讀磁碟區（sectors） | AL=數量，CH=Cylinder，CL=Sector，DH=Head，DL=Drive，ES:BX=buffer |
| 0x03 | 寫磁碟區            | 同上                                                             |
| 0x08 | 取得磁碟參數        | DL=磁碟 → 回傳 CH, CL, DH, DL 等                                 |

### ⌨️ INT 0x16 – 鍵盤服務（Keyboard）

| AH   | 功能                   | 說明                              |
| ---- | ---------------------- | --------------------------------- |
| 0x00 | 取得鍵盤輸入（等按鍵） | 回傳 AL=ASCII，AH=Scan Code       |
| 0x01 | 檢查鍵盤（不等）       | ZF=1 沒有鍵；ZF=0 有 → AL/AH 回傳 |
| 0x02 | 取得 Shift 狀態        | AL = 狀態旗標                     |

### 🖨️ INT 0x17 – 印表機

| AH   | 功能         | 備註                            |
| ---- | ------------ | ------------------------------- |
| 0x00 | 傳送字元     | AH=0x00，AL=字元，DX=印表機編號 |
| 0x01 | 初始化印表機 | DX=印表機                       |
| 0x02 | 取得狀態     | DX=印表機                       |

### ⏱️ INT 0x1A – 時間與 CMOS

| AH   | 功能                | 備註                              |
| ---- | ------------------- | --------------------------------- |
| 0x00 | 取得實時時鐘（RTC） | 回傳 CX:DX = 時間 ticks           |
| 0x02 | 取得 CMOS 時間      | BCD 格式：CH=小時，CL=分鐘，DH=秒 |
