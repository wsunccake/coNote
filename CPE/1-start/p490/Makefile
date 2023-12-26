CPP= /usr/bin/g++
CPPFLAGS= -g
my= main

.SUFFIX : .cpp .o

.cpp .o :
	@echo "compile ..."
	$(CPP) $(CPPFLAGS) -c $<

.PHONY: build
build: ${my}.o
	@echo "build ..."
	$(CPP) $(CPPFLAGS) -o ${my}.exe ${my}.o

.PHONY: clean
clean:
	@echo "clean ..."
	-rm *.o
	-rm *.exe
