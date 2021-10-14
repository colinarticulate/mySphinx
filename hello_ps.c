#include <pocketsphinx.h>

// //
// #define MODELDIR "/usr/local/share/pocketsphinx/model"
// #define HMM MODELDIR "/en-us/en-us"
// #define LM MODELDIR "/en-us/en-us.lm.bin"
// #define DICT MODELDIR "/en-us/cmudict-en-us.dict"

int
main(int argc, char *argv[])
{
    ps_decoder_t *ps;
    cmd_ln_t *config;
    FILE *fh;
    char const *hyp, *uttid;
    int16 buf[512];
    int rv;
    int32 score;


    printf("---------------------\n");
    printf("MODELDIR directory:\n");
    printf("/usr/local/share/pocketsphinx/model/en-us/en-us");
    printf("\n---------------------\n");

    config = cmd_ln_init(NULL, ps_args(), TRUE,
		         "-hmm", "/usr/local/share/pocketsphinx/model/en-us/en-us",
		         "-lm", "/usr/local/share/pocketsphinx/model/en-us/en-us.lm.bin",
	    		 "-dict", "/usr/local/share/pocketsphinx/model/en-us/cmudict-en-us.dict" ,
		         NULL);

    if (config == NULL) {
	fprintf(stderr, "Failed to create config object, see log for details\n");
	return -1;
    }

    ps = ps_init(config);
    if (ps == NULL) {
	fprintf(stderr, "Failed to create recognizer, see log for details\n");
	return -1;
    }

    fh = fopen("goforward.raw", "rb");
    if (fh == NULL) {
	fprintf(stderr, "Unable to open input file goforward.raw\n");
	return -1;
    }

    rv = ps_start_utt(ps);

    while (!feof(fh)) {
	size_t nsamp;
	nsamp = fread(buf, 2, 512, fh);
	rv = ps_process_raw(ps, buf, nsamp, FALSE, FALSE);
    }

    rv = ps_end_utt(ps);
    hyp = ps_get_hyp(ps, &score);

    printf("\n---------------------\n");
    printf("Recognized:\n%s\n", hyp);
    printf("\n---------------------\n");
    
    fclose(fh);
    ps_free(ps);
    cmd_ln_free_r(config);

    return 0;
}
