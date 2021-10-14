# mySphinx  
An example of how to change pocketsphinx API. Also, shows an environment for developers to create tailor-made versions of pocketsphinx and any program such as pocketsphinx_continuous or pocketsphinx_batch. 

Non-canonical builds for development of new pocketsphinx libraries using autotools.

Canonical builds for distribution of new pocketsphinx libraries using the same way as the original pocketsphinx library.  

IDE: vs code  
Compilation and fully debuggable code using own made personalised pocketsphinx API of choice.  


*Instructions:*

For rebuilding pocketsphinx locally for development and testing (non-canonical build):

In each directory, pocketsphinx/ and /sphinxbase,

```
./autogen.sh --prefix=$(pwd)/../debug/usr/local/ --exec-prefix=$(pwd)/../debug/usr/local
make
make install
```

Configuration to develop software that uses a pocketsphinx library is already set up in vscode:  
*task.json*, ensuring linking to own pocketsphinx library and including right headers from own pocketsphinx source code.  
*launch.json*, ensures visual debugging works properly and offers an example how to pass command line argumets to your code.  
*c_cpp_properties*, ensures the minimum amount of headers are included.  
Reguires C/C++ extension.
