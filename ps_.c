#include <stdio.h>
#include <string.h>
#include <assert.h>

#include <xyz.h>


char *argv_[] = {
            "/home/dbarbera/Repositories/mySphinx/ps-debug",
            "-alpha", "0.97",
            "-backtrace", "yes",
            "-beam", "1e-10000",
            "-bestpath", "no",
            "-cmn", "live",
            "-cmninit", "52.55,0.14,-3.23,14.29,-7.74,9.03,-7.17,-6.31,-0.13,1.09,5.23,-2.69,1.01",
            "-dict", "/home/dbarbera/Repositories/mySphinx/data/art_db.phone",
            "-dither", "no",
            "-doublebw", "no",
            "-featparams", "/home/dbarbera/Repositories/mySphinx/data/en-us/en-us/feat.params",
            "-fsgusefiller", "no",
            "-frate", "125",
            "-fwdflat", "no",
            "-hmm", "/home/dbarbera/Repositories/mySphinx/data/en-us/en-us",
            "-infile", "/home/dbarbera/Repositories/mySphinx/data/climb1_colin.wav",
            "-jsgf", "/home/dbarbera/Repositories/mySphinx/data/kl_ay_m.jsgf",
            "-logfn", "/home/dbarbera/Repositories/mySphinx/data/output.log",
            "-lpbeam", "1e-10000",
            "-lponlybeam", "1e-10000",
            "-lw", "6",
            "-maxhmmpf", "-1",
            "-maxwpf", "-1",
            "-nfft", "256",
            "-nwpen", "1",
            "-pbeam", "1e-10000",
            "-pip", "1",
            "-pl_window", "0",
            "-remove_dc", "no",
            "-remove_noise", "yes",
            "-remove_silence", "no",
            "-topn", "4",
            "-vad_postspeech", "25",
            "-vad_prespeech", "5",
            "-vad_startspeech", "8",
            "-vad_threshold", "1",
            "-wbeam", "1e-10000",
            "-wip", "0.5",
            "-wlen", "0.016"
        };

int argc_ = 77;

const char* jsgf_filename = "/home/dbarbera/Repositories/mySphinx/data/kl_ay_m.jsgf";
const char* audio_filename = "/home/dbarbera/Repositories/mySphinx/data/climb1_colin.wav";
//const char* audio_filename = "/home/dbarbera/Repositories/sphinx/data/climb1_colin_headerless.raw";

//result_t ps_call(char* jsgf_buffer, int16* audio_buffer, int argc, char *argv[]);
int ps_call(void* jsgf_buffer, int jsgf_buffer_size, void* audio_buffer, int audio_buffer_size, int argc, char *argv[]);

// void print_result(result_t result) {

//     for(int i=0; i<result.n; i++) {
//         printf("%i  %i %i\n", result.phonemes_idx[i], result.starts[i], result.ends[i]);
//     }
// }

void* create_buffer(int* bsize, const char* filename, const char* mode){
    FILE* file = NULL;
    file = fopen(filename, mode);
    if (file == NULL) {
        E_ERROR_SYSTEM("Failed to open %s for parsing", filename);
        return 0;
    }
    fseek(file, 0, SEEK_END);
    long fsize = ftell(file);
    *bsize=fsize;
    fseek(file, 0, SEEK_SET);  /* same as rewind(f); */

    
    void* contents = (void*)malloc(fsize + 1);
    fread(contents, 1, fsize, file);
    fclose(file);

    return contents;
}

void delete_buffer(void* buffer){
    free(buffer);
}


int
main()
{

    void* jsgf_buffer = NULL;
    void* wav_data = NULL;

    int jsgf_buffer_size = 0; 
    int wav_data_size = 0;
    
    jsgf_buffer = create_buffer( &jsgf_buffer_size, jsgf_filename, "rb");
    wav_data = create_buffer(&wav_data_size, audio_filename, "rb");

    //result_t result = ps_call(jsgf_buffer, audio_buffer, argc, argv);
    ps_call(jsgf_buffer, jsgf_buffer_size, wav_data, wav_data_size, argc_, argv_);

    //print_result(result);

    delete_buffer(jsgf_buffer);
    delete_buffer(wav_data);

    return 0;
}