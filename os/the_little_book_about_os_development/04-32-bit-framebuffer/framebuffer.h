#include "io.h"

enum VGAColor
{
    BLACK = 0,
    BLUE = 1,
    GREEN = 2,
    CYAN = 3,
    RED = 4,
    MAGENTA = 5,
    BROWN = 6,
    LIGHT_GRAY = 7,
    DARK_GRAY = 8,
    LIGHT_BLUE = 9,
    LIGHT_GREEN = 10,
    LIGHT_CYAN = 11,
    LIGHT_RED = 12,
    LIGHT_MAGENTA = 13,
    YELLOW = 14,
    WHITE = 15
};
typedef enum VGAColor FB_COLOR;

/* I/O ports */
#define FB_COMMAND_PORT 0x3D4
#define FB_DATA_PORT 0x3D5

/* I/O port commands */
#define FB_HIGH_BYTE_COMMAND 14
#define FB_LOW_BYTE_COMMAND 15
#define FB_WIDTH 80
#define FB_HEIGHT 25

#define FRAMEBUFFER_ADDRESS 0x000B8000
#define FB_BASE_ADDRESS ((char *)FRAMEBUFFER_ADDRESS)

void fb_write_cell_byte(unsigned int i, char c, unsigned char fg, unsigned char bg);
void fb_write_cell_location(unsigned int i, char c, unsigned char fg, unsigned char bg);
void fb_write_cell_location_rc(unsigned int row, unsigned int col, char c, unsigned char fg, unsigned char bg);
void fb_write(unsigned int row, unsigned int col, char *str, unsigned char fg, unsigned char bg);
void fb_clear();

void fb_move_cursor(unsigned char pos);
void fb_enable_cursor(unsigned char cursor_start, unsigned char cursor_end);
void fb_disable_cursor();
