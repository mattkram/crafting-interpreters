#include <stdio.h>

#include "common.h"
#include "chunk.h"

int main(int argc, const char* argv[]) {
    printf("Welcome to clox!\n");

    Chunk chunk;
    initChunk(&chunk);
    writeChunk(&chunk, OP_RETURN);
    freeChunk(&chunk);

    return 0;
}
