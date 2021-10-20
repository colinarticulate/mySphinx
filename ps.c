#include <stdio.h>
#include <string.h>
#include <assert.h>

//#include "ps.h"

const char *jsgf_filename="./../data/kl_ay_m__from_wrapper_from_c.jsgf";
const char *wav_filename="./../data/climb1_colin__from_wrapper_from_c.wav";
const char *params_filename="./../data/params__from_wrapper_from_c.txt";

void create_file(char *buffer, int len, const char *filename) {
    //printf("Just called a function\n");
    FILE *file;// = NULL;
    int k = 0;
    //printf("About to open a file for writing.\n");
    file =fopen(filename, "wb");
    if (file == NULL) {
        printf("Failed to open %s for writing", filename);

    }
    //printf("About to write the file.");
    k = fwrite(buffer, sizeof(char), len, file);
    //printf("Just wrote the file.");
    fclose(file);
}

void create_file_params(char *argv[], int argc, const char *filename){
        //printf("Just called a function\n");
    FILE *file;// = NULL;
    int k = 0;
    //printf("About to open a file for writing.\n");
    file =fopen(filename, "wb");
    if (file == NULL) {
        printf("Failed to open %s for writing", filename);

    }

    for(int i=0; i<argc; i++) {
        fprintf(file, "%d\t%s\n", i, argv[i]);
    }

    fclose(file);
}

//result_t ps_call(char* jsgf_buffer, int16* audio_buffer, int argc, char *argv[]);
int ps_call(void* jsgf_buffer, int jsgf_buffer_size, void* audio_buffer, int audio_buffer_size, char *argv[], int argc){

    create_file(jsgf_buffer, jsgf_buffer_size, jsgf_filename);
    create_file(audio_buffer, audio_buffer_size, wav_filename);
    create_file_params(argv, argc, params_filename);


    return 0;
} 

