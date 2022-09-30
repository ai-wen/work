apt remove cmake

wget https://cmake.org/files/v3.15/cmake-3.15.3-Darwin-x86_64.dmg

wget https://cmake.org/files/v3.15/cmake-3.15.3-Linux-x86_64.tar.gz
wget https://cmake.org/files/v3.15/cmake-3.15.3-Linux-x86_64.sh
tar -xf cmake-3.15.3-Linux-x86_64.tar.gz


https://cmake.org/files/v3.15/cmake-3.15.3-win32-x86.zip
https://cmake.org/files/v3.15/cmake-3.15.3-win64-x64.zip

#mkdir cmake-3.15.3
#tar -xf cmake-3.15.3-Linux-x86_64.tar.gz -C ./cmake-3.15.3

which cmake
sudo ln -sf /home/lm/project/CipherSuite/tools/cmake-3.15.3-Linux-x86_64/bin/cmake /usr/bin/cmake
cmake --version


wget https://cmake.org/files/v3.15/cmake-3.15.3.tar.gz
tar -xf cmake-3.15.3.tar.gz
./configure --prefix=/usr/local/cmake-3.15.3
make
make install


gcc --version
gcc (Ubuntu 5.4.0-6ubuntu1~16.04.11) 5.4.0 20160609
表示此机器已经安装了GCC，此机器的版本号为 5.4.0（注：支持C++11的最低版本为4.8）。
如果没有使用 apt-get install gcc 进行安装。


ubuntu 理论上来讲5.4应该支持C++11，时间20160609，但是实际上__cplusplus的值却是199711L
最后看到一篇文章的一句话，原来是默认不支持。所以只要在编译时加上-std=c++11 或者 -std=c++0x就可以了。

安装gcc-6：
sudo apt-get update && \
sudo apt-get install build-essential software-properties-common -y && \
sudo add-apt-repository ppa:ubuntu-toolchain-r/test -y && \
sudo apt-get update && \
sudo apt-get install gcc-snapshot -y && \
sudo apt-get update && \
sudo apt-get install gcc-6 g++-6 -y && \
sudo update-alternatives --install /usr/bin/gcc gcc /usr/bin/gcc-6 60 --slave /usr/bin/g++ g++ /usr/bin/g++-6 



g++ main.cpp -m32
sudo apt-get install lib32ncurses5 ​lib32z1
sudo apt-get install build-essential module-assistant gcc-multilib g++-multilib