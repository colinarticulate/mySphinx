
Examples:

Command line:

$ pocketsphinx_continuous -infile pocketsphinx/test/data/goforward.raw -hmm pocketsphinx/model/en-us/en-us -lm pocketsphinx/model/en-us/en-us.lm.bin -dict pocketsphinx/model/en-us/cmudict-en-us.dict


Compilation:

gcc -o hello_ps hello_ps.c \
    -DMODELDIR=\"`pkg-config --variable=modeldir pocketsphinx`\" \
    `pkg-config --cflags --libs pocketsphinx sphinxbase`
	
https://en.wikipedia.org/wiki/Pkg-config
pkg-config is a computer program that defines and supports a unified interface for querying installed libraries for the purpose of compiling software that depends on them. It allows programmers and installation scripts to work without explicit knowledge of detailed library path information. 
This would refer to the files pocketsphinx/pocketsphinx.pc and sphinxbase/sphinxbase.pc
	pkg-config --cflags --libs pocketsphinx sphinxbase
	
on the command line and make sure the output looks like this:
	-I/usr/local/include -I/usr/local/include/sphinxbase -I/usr/local/include/pocketsphinx  
	-L/usr/local/lib -lpocketsphinx -lsphinxbase -lsphinxad
	
I got:

-I/usr/local/include -I/usr/local/include/sphinxbase -I/usr/local/include/pocketsphinx 
-I/usr/local/include -I/usr/local/include/sphinxbase 
-L/usr/local/lib -lpocketsphinx -lsphinxbase -lsphinxad 
-lpthread -lm
