#ifndef INCLUDE_MEMORY_SEGMENTS
#define INCLUDE_MEMORY_SEGMENTS

struct GDT
{
	unsigned short size;
	unsigned int address;
} __attribute__((packed));

struct GDTDescriptor
{
	unsigned short limit_low;
	unsigned short base_low;
	unsigned char base_middle;
	unsigned char access_byte;
	unsigned char limit_and_flags;
	unsigned char base_high;
} __attribute__((packed));

void init_descriptor(int index, unsigned int base_address, unsigned int limit, unsigned char access_byte, unsigned char flags);
void install_gdt();

// Wrappers around ASM.
void load_gdt(struct GDT gdt);
void load_registers();

#endif /* INCLUDE_MEMORY_SEGMENTS */

#define SEGMENT_FLAGS_PART 0x0C
// GDT Descriptor 的 flags 欄位中，0x0C 二進位是 00001100
// Bit 7: G（Granularity）= 1 → 4KB 單位
// Bit 6: D/B（Size）= 1 → 32-bit protected mode
// Bit 5: L（Long mode）= 0 → 不是 64-bit
// Bit 4: AVL（Available for system use）= 0

#define SEGMENT_DESCRIPTOR_COUNT 3
// SEGMENT_DESCRIPTOR_COUNT = 3（定義 null、code、data 三個）

#define SEGMENT_BASE 0
#define SEGMENT_LIMIT 0xFFFFF

#define SEGMENT_CODE_TYPE 0x9A
#define SEGMENT_DATA_TYPE 0x92
