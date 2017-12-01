#example 2
#usage: make main OR make
CC = clang          #欲使用的C compiler
CFLAGS = -O3 		#欲使用的參數
INC = -I src/include   	#include headers的位置
LIB = -L src/lib       	#include libraries的位置
       
all: 
	${CC} src/wavereader.c ${CFLAGS} ${INC} ${LIB} -o bin/wavereader
#main.o: main.c target.h                    
#    ${CC} main.c ${CFLAGS} ${INC} ${LIB} -lpthread -c  
#foo1.o: foo1.c target.h                    
#    ${CC} foo1.c ${CFLAGS} ${INC} ${LIB} -c        
clean:                             
	@rm -rf src/*.o  bin/wavereader 