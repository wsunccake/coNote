;Load the Global Descriptor Table

global load_gdt
global load_registers

; 載入 GDT 到 GDTR
load_gdt:
	lgdt [esp + 4]  ; lgdt 是 CPU 指令，載入 GDT 指標到 GDTR（Global Descriptor Table Register）
					; [esp + 4]：呼叫這個函式的 caller（例如 C 程式）會把 GDT pointer 放在 stack 上，
					; 這裡是讀 stack 上的參數。
	ret

; 設定 Segment Registers
load_registers:
	mov ax, 0x10
	mov ds, ax ; 0x10 - an offset into GDT for the third (kernel data segment) record.
	mov ss, ax
	mov es, ax
	mov fs, ax
	mov gs, ax
	jmp 0x08:flush_cs ; 0x08 - an offset into GDT for the second (kernel code segment) record. 

flush_cs:
	ret