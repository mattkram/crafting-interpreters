CC = gcc

EXE = lox
SRC = src

C_FILES = $(wildcard $(SRC)/*.c)
H_FILES = $(wildcard $(SRC)/*.h)
O_FILES = $(C_FILES:.c=.o)

default: $(EXE)

$(EXE): $(O_FILES)
	$(CC) -o $(EXE) $(O_FILES)

%.o: %.c $(H_FILES)
	$(CC) -c -o $@ $< $(CFLAGS)

run: default 
	./$(EXE)

clean:
	$(RM) $(O_FILES)
	$(RM) $(EXE)
