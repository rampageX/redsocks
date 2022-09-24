#!/bin/sh

[ -n "$1" ] && ssVersion=$1 || ssVersion="git"

mkdir -p done/${ssVersion}/OpenSSL/packed
mkdir -p done/${ssVersion}/PolarSSL/packed

echo "Compiling OpenSSL Version..."
make distclean
make -j6 CC="distcc arm-linux-gcc" CXX="distcc arm-linux-g++" V=1 ENABLE_STATIC=true DISABLE_SHADOWSOCKS=true ENABLE_HTTPS_PROXY=true
mv -f ./redsocks2 done/${ssVersion}/OpenSSL/

#exit

echo "Compiling PolarSSL Version..."
make distclean
make -j6 CC="distcc arm-linux-gcc" CXX="distcc arm-linux-g++" V=1 ENABLE_STATIC=true DISABLE_SHADOWSOCKS=true USE_CRYPTO_POLARSSL=true ENABLE_HTTPS_PROXY=true
#make USE_CRYPTO_POLARSSL=true ENABLE_HTTPS_PROXY=true
mv -f ./redsocks2 done/${ssVersion}/PolarSSL/

#exit

#UPX all
cp -f done/${ssVersion}/OpenSSL/redsocks2 done/${ssVersion}/OpenSSL/packed/
upx -9 done/${ssVersion}/OpenSSL/packed/*

cp -f done/${ssVersion}/PolarSSL/redsocks2 done/${ssVersion}/PolarSSL/packed/
upx -9 done/${ssVersion}/PolarSSL/packed/*

file done/${ssVersion}/PolarSSL/packed/redsocks2
done/${ssVersion}/PolarSSL/packed/redsocks2 -v

file done/${ssVersion}/OpenSSL/packed/redsocks2
done/${ssVersion}/OpenSSL/packed/redsocks2 -v

