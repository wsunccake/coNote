#include "segment.h"

/*
 * Flags part of `limit_and_flags`.
 * 1100
 * 0 - Available for system use
 * 0 - Long mode
 * 1 - Size (0 for 16-bit, 1 for 32)
 * 1 - Granularity (0 for 1B - 1MB, 1 for 4KB - 4GB)
 */

static struct GDTDescriptor gdt_descriptors[SEGMENT_DESCRIPTOR_COUNT];
// GDT 的主要資料結構：
// index 0：null descriptor（必要）
// index 1：code segment
// index 2：data segment

// 建立一個 GDT descriptor
void init_descriptor(int index, unsigned int base_address, unsigned int limit, unsigned char access_byte, unsigned char flags)
{
    // GDT 每個 descriptor 組成（從 Intel 規格）如下
    //     name    | bit
    // -----------------------------
    // limit 0–15  | 16 bit
    // base 0–15   | 16 bit
    // base 16–23  | 8 bit
    // access byte | 8 bit
    // limit 16–19 | 4 bit
    // flags       | 4 bit
    // base 24–31  | 8 bit
    gdt_descriptors[index].base_low = base_address & 0xFFFF;
    gdt_descriptors[index].base_middle = (base_address >> 16) & 0xFF;
    gdt_descriptors[index].base_high = (base_address >> 24) & 0xFF;

    gdt_descriptors[index].limit_low = limit & 0xFFFF;
    gdt_descriptors[index].limit_and_flags = (limit >> 16) & 0xF;
    gdt_descriptors[index].limit_and_flags |= (flags << 4) & 0xF0;

    gdt_descriptors[index].access_byte = access_byte;
}

void install_gdt()
{
    gdt_descriptors[0].base_low = 0;
    gdt_descriptors[0].base_middle = 0;
    gdt_descriptors[0].base_high = 0;
    gdt_descriptors[0].limit_low = 0;
    gdt_descriptors[0].access_byte = 0;
    gdt_descriptors[0].limit_and_flags = 0;

    // The null descriptor which is never referenced by the processor.
    // Certain emulators, like Bochs, will complain about limit exceptions if you do not have one present.
    // Some use this descriptor to store a pointer to the GDT itself (to use with the LGDT instruction).
    // The null descriptor is 8 bytes wide and the pointer is 6 bytes wide so it might just be the perfect place for this.
    // From:  http://wiki.osdev.org/GDT_Tutorial
    struct GDT *gdt_ptr = (struct GDT *)gdt_descriptors;
    gdt_ptr->address = (unsigned int)gdt_descriptors;
    gdt_ptr->size = (sizeof(struct GDTDescriptor) * SEGMENT_DESCRIPTOR_COUNT) - 1;
    // sizeof(struct GDTDescriptor) = 8（每個 GDT descriptor 是 8 bytes）
    // SEGMENT_DESCRIPTOR_COUNT = 3（定義 null、code、data 三個）

    // See http://wiki.osdev.org/GDT_Tutorial
    init_descriptor(1, SEGMENT_BASE, SEGMENT_LIMIT, SEGMENT_CODE_TYPE, SEGMENT_FLAGS_PART);
    init_descriptor(2, SEGMENT_BASE, SEGMENT_LIMIT, SEGMENT_DATA_TYPE, SEGMENT_FLAGS_PART);

    load_gdt(*gdt_ptr);
    load_registers();
}