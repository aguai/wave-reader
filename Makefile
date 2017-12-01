CC = clang         
CFLAGS = -O3 		
INC = -I src/include   	
LIB = -L src/lib       
       
all: 
	${CC} src/wavereader.c ${CFLAGS} ${INC} ${LIB} -o bin/wavereader
clean:                             
	@rm -rf src/*.o  exe/*