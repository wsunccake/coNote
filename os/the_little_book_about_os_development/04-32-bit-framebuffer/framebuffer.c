#include "framebuffer.h"

void fb_write_cell_byte(unsigned int i, char c, unsigned char fg, unsigned char bg)
{
    char *fb = FB_BASE_ADDRESS;
    fb[i] = c;
    fb[i + 1] = ((fg & 0x0F) << 4) | (bg & 0x0F);
    // (bg & 0x0F) 取出背景色的低 4 位（背景色最多 16 種顏色）
    // ((fg & 0x0F) << 4) 把前景色放在屬性 byte 的高 4 位（0xF0）
}

void fb_write_cell_location(unsigned int i, char c, unsigned char fg, unsigned char bg)
{
    char *fb = FB_BASE_ADDRESS;
    fb[i * 2] = c;
    fb[i * 2 + 1] = ((bg & 0x0F) << 4) | (fg & 0x0F);
}
void fb_write_cell_location_rc(unsigned int row, unsigned int col, char c, unsigned char fg, unsigned char bg)
{
    unsigned int pos = row * FB_WIDTH + col;
    fb_write_cell_location(pos, c, fg, bg);
}

void fb_write(unsigned int row, unsigned int col, char *str, unsigned char fg, unsigned char bg)
{
    unsigned int pos = row * FB_WIDTH + col;

    for (unsigned int i = 0; str[i] != '\0'; i++)
    {
        fb_write_cell_location(pos + i, str[i], fg, bg);
    }
}

void fb_clear()
{
    unsigned int i = 0;
    while (i < FB_WIDTH * FB_HEIGHT * 2)
    {
        fb_write_cell_byte(2 * i, ' ', BLACK, BLACK);
        i = i + 1;
    }
    return;
}

void fb_move_cursor(unsigned char pos)
{
    outb(FB_COMMAND_PORT, FB_HIGH_BYTE_COMMAND); // 設定要寫高位元
    outb(FB_DATA_PORT, ((pos >> 8) & 0x00FF));   // 寫高位元

    outb(FB_COMMAND_PORT, FB_LOW_BYTE_COMMAND); // 設定要寫低位元
    outb(FB_DATA_PORT, pos & 0x00FF);           // 寫低位元
}

void fb_enable_cursor(unsigned char cursor_start, unsigned char cursor_end)
{
    outb(FB_COMMAND_PORT, 0x0A);
    outb(FB_DATA_PORT, (inb(FB_DATA_PORT) & 0xC0) | cursor_start);

    outb(FB_COMMAND_PORT, 0x0B);
    outb(FB_DATA_PORT, (inb(FB_DATA_PORT) & 0xE0) | cursor_end);
}

void fb_disable_cursor()
{
    outb(FB_COMMAND_PORT, 0x0A);
    outb(FB_DATA_PORT, 0x20); // Bit 5 設為 1，表示游標關閉
}
