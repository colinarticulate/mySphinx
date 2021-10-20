/* File : example.c */
#include <stdio.h>
#include <stdlib.h>
#include <ctype.h>

#include <stdio.h>
#include <stdlib.h>
#include <ctype.h>

#include <string.h>
#include <assert.h>

//const char *filename="./../data/kl_ay_m__from_wrapper_from_c.jsgf";
const char *filename="./../data/climb1_colin__from_wrapper_from_c.wav";

void create_file(char *buffer, int len, const char *filename) {
    printf("Just called a function\n");
    FILE *file;// = NULL;
    int k = 0;
    printf("About to open a file for writing.\n");
    file =fopen(filename, "wb");
    printf("About to write the file.");
    k = fwrite(buffer, sizeof(char), len, file);
    printf("Just wrote the file.");
    fclose(file);
}

int passing_bytes(char *buffer, int len) {

  char *cmem_buffer = (char*)malloc(len*sizeof(char));

  memcpy(cmem_buffer, buffer, len*sizeof(char));

  printf("About to call a function.\n");

  create_file(buffer, len, filename);
  //  create_file(cmem_buffer, len, filename);
  for(int i = 0; i< len; i++)
    printf("%c",buffer[i]);

  free(cmem_buffer);

  return len;
}


