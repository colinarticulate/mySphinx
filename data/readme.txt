Example parameters for calling pocketsphinx_continuous from the command line:

"-alpha": "0.97",
"-backtrace": "yes",
"-beam": "1e-10000",
"-bestpath": "no",
"-cmn": "live",
"-cmninit": "",
"-dict": "<url for art_db.phone>",
"-dither": "no",
"-doublebw": "no",
"-featparams": "<url for model folder>/en-us/en-us/feat.params",
"-fsgusefiller": "no",
"-frate": "125",
"-fwdflat": "no",
"-hmm": "<url for model folder>/en-us/en-us",
"-infile": "<url for audio file, climb1_colin.wav>",
"-jsgf": "<url for grammar file kl_ay_m.jsgf>",
"-logfn": "<url for log file>",
"-lpbeam": "1e-10000",
"-lponlybeam": "1e-10000",
"-lw": "6",
"-maxhmmpf": "-1",
"-maxwpf": "-1",
"-nfft": "256",
"-nwpen": "1",
"-pbeam": "1e-10000",
"-pip": "1",
"-pl_window": "0",
"-remove_dc": "no",
"-remove_noise": "yes",
"-remove_silence": "no",
"-topn": "4",
"-vad_postspeech": "25",
"-vad_prespeech": "5",
"-vad_startspeech": "8",
"-vad_threshold": "1",
"-wbeam": "1e-10000",
"-wip": "0.5",
"-wlen": "0.016",

<url for model folder> should be the URL for the articulate model files, not the
reference model files that come bundled with the pocketsphinx installation. On my
installation its:

/usr/local/share/pocketsphinx/model

The en-us folder included is the folder that sits directly inside the model folder which
I've included here for completeness.

Note that if you are configuring pocketsphinx in code (so not using the command line
interface) you should remove the -infile parameter. The audiofile that this references
is passed to the pocketsphinx decoder after pocketsphinx has been configured.

Feel free to play around with these settings to see what difference they make to the
output phonemes. Note though that not all settings will necessarily work - it is possible
to provide a combination of settings which causes pocketsphinx_continuous to crash. In
particular the parameters frate, nfft and wlen are all related so changing one of these
values without also changing the others will typically result in trouble. Paul knows more
about the kind of settings that do and don't work as he's experimented with pocketsphinx
far more than I have.

The grammar file, audio file and phoneme dictionary are all included, as is a sample log
file.

The lines of interest in the log file are 265 - 275 inclusive. You need to look for
a line in the file that starts word. The following lines contain the phonemes, one phoneme
per line. Once you read a line that starts INFO (or you reach the end of the file) you've
got all the phonemes. For each phoneme you need to extract the phoneme, start and end
times (columns word, start and end respectively). The (NULL) phonemes aren't worth
extracting although I don't, at the moment, throw them away (see for instance
decodeFromFile in decoder.go on the ps-go branch of the scanScheduler repository). For an
example of go code that extracts phoneme data from the log file see parsePsData in
listen.go on the main branch of the pron repository.

And that's it I think.
