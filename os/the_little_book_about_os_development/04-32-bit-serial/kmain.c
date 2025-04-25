#include "serial.h"

char MESSAGE[] = "MakeOS";

void kmain()
{
	serial_write(MESSAGE, sizeof(MESSAGE));
}
