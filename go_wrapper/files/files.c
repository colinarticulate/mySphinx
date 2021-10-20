/* File : example.c */
#include <stdio.h>
#include <stdlib.h>
#include <ctype.h>

#include <string.h>
#include <assert.h>

const char *jsgf_filename="./data/kl_ay_m.jsgf";
const char *jsgf_filename_from_c="./data/kl_ay_m__from_c.jsgf";
const char *wav_filename="./data/climb1_colin.wav";

void* create_buffer(size_t* bsize, const char* filename, const char* mode){
    FILE* file = NULL;
    int k=0;
    file = fopen(filename, mode);
    if (file == NULL) {
        printf("Failed to open %s for parsing", filename);
        return 0;
    }
    fseek(file, 0, SEEK_END);
    long fsize = ftell(file);
    *bsize=fsize;
    fseek(file, 0, SEEK_SET);  /* same as rewind(f); */

    
    void* contents = (void*)malloc(fsize + 1);
    k=fread(contents, 1, fsize, file);
    fclose(file);

    return contents;
}

void delete_buffer(void* buffer){
    free(buffer);
}

void create_file(void *buffer, size_t len, const char *filename) {
    FILE *file = NULL;
    int k = 0;

    file =fopen(filename, "wb");

    k = fwrite(buffer, sizeof(char), len, file);


    fclose(file);
}

int
main(){
    void* jsgf_buffer = NULL;
    void* wav_buffer = NULL;

    size_t jsgf_buffer_size = 0; 
    size_t wav_data_size = 0;
    
    jsgf_buffer = create_buffer( &jsgf_buffer_size, jsgf_filename, "rb");
    wav_buffer = create_buffer( &wav_data_size, wav_filename, "rb");

    create_file(jsgf_buffer, jsgf_buffer_size, jsgf_filename_from_c);

    printf("File size: %zu bytes\n", jsgf_buffer_size);
    printf("File size: %zu bytes\n", wav_data_size);


    delete_buffer(jsgf_buffer);
    delete_buffer(wav_buffer);






    return 0;
}