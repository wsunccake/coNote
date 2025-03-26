CC = gcc
AS = nasm
CFLAGS = -m64
ASFLAGS = -f elf64
OBJS = sum.o main.o
EXE = sum

%.o: %.c
	${CC} ${CFLAGS} -c $< -o $@

%.o: %.asm
	${AS} ${ASFLAGS} $< -o $@

all: $(OBJS)
	${CC} ${OBJS} -o $(EXE)

run:
	./$(EXE)

clean:
	rm *.o $(EXE)

.PHONY:
	all
	run
	clean

