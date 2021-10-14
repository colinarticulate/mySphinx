#!/bin/sh
# Run this to build pocketsphinx from source. Lib and execs go to /env

rm -r debug

srcdir=`dirname $0`

$srcdir/sphinxbase/autogen.sh --prefix=$(pwd)/../debug/usr/local/ --exec-prefix=$(pwd)/../debug/usr/local
make -C $srcdir/sphinxbase
make install -C $srcdir/sphinxbase

#poscketsphinx
./autogen.sh --prefix=$(pwd)/../debug/usr/local/ --exec-prefix=$(pwd)/../debug/usr/local
make
make install


#poscketsphinx
./autogen.sh --prefix= /home/dbarbera/Repositories/mySphinx/debug/usr/local/ --exec-prefix= /home/dbarbera/Repositories/mySphinx/debug/usr/local/
make
make install
