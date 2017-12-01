#example 2
#usage: make main OR make
CC = clang          #欲使用的C compiler
CFLAGS = -O3 		
INC = -I src/include   	
LIB = -L src/lib       
       
all: 
	${CC} src/wavereader.c ${CFLAGS} ${INC} ${LIB} -o bin/wavereader
clean:                             
	@rm -rf src/*.o  bin/wavereader 