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




//Check options using ./config -help
--program-suffix=SUFFIX


#poscketsphinx non-canonical changing suffix
./autogen.sh --prefix= /home/dbarbera/Repositories/mySphinx/debug/usr/local/ --exec-prefix= /home/dbarbera/Repositories/mySphinx/debug/usr/local/ --program-suffix=xyz
make
make install


#poscketsphinx non-canonical changing suffix
./autogen.sh --prefix= /home/dbarbera/Repositories/mySphinx/xyz/usr/local/ --exec-prefix= /home/dbarbera/Repositories/mySphinx/xyz/usr/local/ --program-suffix="xyz"
make
make install


./autogen.sh --prefix=/home/dbarbera/Repositories/mySphinx/xyz/usr/local/ --program-suffix=_xyz
make
make install



#---------------------- other installations

./autogen.sh --prefix= /home/dbarbera/Installation_Tests/pocketsphinxxyz/pocketsphinx/debug/usr/local/
make
make install



biblio:

Autotools myth buster:
https://autotools.io/index.html
https://gnu.huihoo.org/automake-1.5/html_node/automake_toc.html#TOC1
