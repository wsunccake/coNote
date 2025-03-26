#include <stdint.h>

void kernel_main() {
    const char *str = "Hello, Kernel!\n";
    volatile uint16_t *vga = (uint16_t*)0xB8000;

    for (int i = 0; str[i] != '\0'; i++) {
        vga[i] = (uint16_t) str[i] | (0x0F << 8);
    }

    while (1) {
        __asm__ volatile ("hlt");
    }
}

