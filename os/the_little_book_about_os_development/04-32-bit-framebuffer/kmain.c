#include "framebuffer.h"

char MESSAGE[] = "MakeOS";

void kmain()
{
	fb_clear();

	fb_write_cell_byte(0, 'H', GREEN, DARK_GRAY);
	fb_write_cell_byte(2, 'e', GREEN, DARK_GRAY);
	fb_write_cell_byte(4, 'l', GREEN, DARK_GRAY);
	fb_write_cell_byte(6, 'l', LIGHT_BLUE, YELLOW);
	fb_write_cell_byte(8, 'o', LIGHT_BLUE, YELLOW);

	fb_write_cell_location(7, 'W', WHITE, BLACK);
	fb_write_cell_location(8, 'o', WHITE, BLACK);
	fb_write_cell_location(9, 'r', WHITE, BLACK);
	fb_write_cell_location(10, 'l', LIGHT_GRAY, BLACK);
	fb_write_cell_location(11, 'd', LIGHT_GRAY, BLACK);

	fb_write_cell_location_rc(1, 0, 'O', LIGHT_GRAY, BLACK);
	fb_write_cell_location_rc(1, 1, 'S', LIGHT_GRAY, BLACK);

	fb_write_cell_location(12, ' ', LIGHT_GRAY, BLACK);
	fb_move_cursor(12);

	fb_write(2, 0, MESSAGE, WHITE, BLACK);
}
