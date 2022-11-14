# [现代CMAK](https://ukabuer.me/blog/more-modern-cmake)
对于 C/C++的开发者而言，当涉及到复杂的第三方依赖时，工程的管理往往会变得十分棘手，尤其是还需要支持跨平台开发时。
CMake 做为跨平台的编译流程管理工具，为第三方依赖查找和引入，编译系统创建，程序测试以及安装都提供了成熟的解决方案。

这种编译体验我认为勉强能赶上 Rust, Go 这些现代语言的一半，还有一半则是差在包管理上

CMake 和 C++一样，随着多年的发展，其设计也得到了许多改进，并且和旧版本相比产生了重要的差异，从而有了现代 CMake 的说法。 

## Target 和围绕 Target 的配置
例如: Target 名为 xx

add_executable(xx)
add_library(xx SHARED)
add_library(xx STATIC)

为Target 配置源码和头文件
file(GLOB_RECURSE SRCS ${CMAKE_CURRENT_SOURCE_DIR}/src/*.cpp)
- 参数一 GLOB_RECURSE 表明递归的查找子文件夹，第
- 参数二 SRCS则是存储结果的变量名
- 参数三 为目标文件的匹配模式，找到符合条件的 cpp 文件后，他们的路径会以字符串数组的形式保存在 SRCS 变量中
target_source(xx PRVIATE "main.cpp" ${SRCS})
target_include_directories(xx PRIVATE ${CMAKE_CURRENT_SOURCE_DIR}/include)

为Target 配置语言，编译时宏，编译器传参（比如 gcc, clang, cl）
target_compile_features(xx PRIVATE std_cxx_14)
target_compile_definitions(xx PRIVATE LogLevel=3)
target_compile_options(xx PRIVATE -Werror -Wall -Wextra)

## 传统 CMake 中，配置通常都是以全局变量的形式定义，比如使用include_directories()、set_cxx_flags()等命令，传统方式的问题是灵活度低，当存在多个 target 时无法进行分别配置，导致某个 target 的属性意外遭到污染，因此现代 CMake 基于 target 的配置方式就和引入了 namespace 一样，管理起来更省心。

## PRIVATE/INTERFACE/PUBLIC
Build Specification 和 Usage Requirement
C/C++通过 include 头文件的方式引入依赖，在动态或静态链接后可以调用依赖实现。 一个可执行程序可能会依赖链接库，链接库也同样可能依赖其他的链接库。
使用者如何知道使用这些外部依赖库需要什么条件？ 比方说，其头文件的代码可能需要开启编译器 C++17 的支持、依赖存在许多动态链接库时可能只需要链接其中的一小部分、有哪些间接依赖需要安装、间接依赖的版本要求是什么……
依赖库的作者可以在某个 README、网站、甚至在头文件里说明使用要求

CMake 提供的解决方案是，在对 target 进行配置时，可以规定配置的类型，分为 build specification 和 usage requirement 两类，会影响配置的应用范围。
### Build specification 类型的配置
    仅在编译的时候需要满足，通过PRIVATE关键字声明； 
### Usage requirement 类型的配置
    则是在使用时需要满足，即在其他项目里，使用本项目已编译好的 target 时需要满足，这种类型的配置使用INTERFACE关键词声明。
在实际工程中，有很多配置在编译时以及被使用时都需要被满足的，这种配置通过PUBLIC关键词进行声明。

示例：
一个 library，在编译时静态链接了 Boost，在实现文件中使用了 c++14 的特性，并用到了 Boost 的头文件和函数。
但在对外提供的头文件中只用到 C++03 的语法，也没有引入任何 Boost 的代码，则可以如下配置为Target
target_compile_features(xx PRIVATE cxx_std_14)  
target_link_libraries(xx PRIVATE Boost::Format) 
PRIVATE 说明 c++14 的支持只在编译时需要用到，Boost 库的链接也仅在编译时需要。

但如果我们对外提供的头文件中也使用了 C++14
target_compile_features(xx PUBLIC cxx_std_14)  

当library是 header-only 时，我们的工程是不需要单独编译的，因此也就没有 build specification，通过INTERFACE修饰配置即可
target_compile_features(xx INTERFACE cxx_std_14)

需要注意的是，Usage requirement 类型的配置，即通过INTERFACE或是PUBLIC修饰的配置是会传递的，比如 LibA 依赖 LibB 后，会继承 LibB 的 usage requirement，此后 LibC 依赖 LibB 时，LibA 和 libB 的 usage requirement 都会继承下来，这在存在多级依赖时是非常有用的。


## find_package
在 CMake 中寻找第三方库的命令为find_package，其背后的工作方式有两种，一种基于 Config File 的查找，另一种则是基于 Find File 的查找。 在执行find_package时，实际上 CMake 都是在找这两类文件，找到后从中获取关于库的信息。

### 通过 Config file 找到依赖
Config File 是依赖的开发者提供的 cmake 脚本，通常会随预编译好的二进制一起发布，供下游的使用者使用。 在 Config file 里，会对库里包含的 target 进行描述，说明版本信息以及头文件路径、链接库路径、编译选项等 usage requirement。
CMake 对 Config file 的命名是有规定的，对于find_package(ABC)这样一条命令，CMake 只会去寻找ABCConfig.cmake或是abc-config.cmake。 CMake 默认寻找的路径和平台有关，在 Linux 下寻找路径包括/usr/lib/cmake以及/usr/lib/local/cmake，在这两个路径下可以发现大量的 Config File，一般在安装某个库时，其自带的 Config file 会被放到这里来。
在 Windows 下没有安装库的规范，也因此没有这样的目录，库可能被安装在各种奇奇怪怪的地方。 此外，在 Linux 下，库也可能没有被安装在上述这些默认位置，在这些情况下，CMake 也提供了解决方案，对于find_package(Abc)命令，如果 CMake 没有找到 Config file，使用者可以提供Abc_DIR变量，CMake 会到Abc_DIR指向的路径寻找 Config file。

### 通过 Find file 找到依赖
Config file 看似十分美好，由开发者编写 CMake 脚本，使用者只要能找到 Config file 即可获取到库的 usage requirement。 但现实是，并不是所有的开发者都使用 CMake，很多库并没有提供供 CMake 使用的 Config file，但此时我们还可以使用 Find file。

对于find_package(ABC)命令，如果 CMake 没有找到 Config file，他还会去试着寻找FindABC.cmake。Find file 在功能上和 Config file 相同，区别在于 Find file 是由其他人编写的，而非库的开发者。 如果你使用的某个库没有提供 Config file，你可以去网上搜搜 Find file 或者自己写一个，然后加入到你的 CMake 工程中。

一个好消息是 CMake 官方为我们写好了很多 Find file，在CMake Documentation这一页面可以看到，OpenGL，OpenMP，SDL 这些知名的库官方都为我们写好了 Find 脚本，因此直接调用 find_package 命令即可。 但由于库的安装位置并不是固定的，这些 Find 脚本不一定能找到库，此时根据 CMake 报错的提示设置对应变量即可，通常是需要提供安装路径，这样就可以通过 Find file 获取到库的 usage requirement。 不论是 Config file 还是 Find file，其目的都不只是找到库这么简单，而是告诉 CMake 如何使用这个库。

坏消息是有更大部分库 CMake 官方也没有提供 Find file，这时候就要自己写了或者靠搜索了，写好后放到本项目的目录下，修改CMAKE_MODULE_PATH这个 CMAKE 变量：

list(INSERT CMAKE_MODULE_PATH 0 ${CMAKE_SOURCE_DIR}/cmake)
这样${CMAKE_SOURCE_DIR}/cmake目录下的 Find file 就可以被 CMake 找到了。

## Imported Target
在 C/C++工程里，对于依赖，我们最基本的要求就是知道他们的链接库路径和头文件目录，通过 CMake 的find_library和find_path两个命令就可以完成任务：

find_library(MPI_LIBRARY
  NAMES mpi
  HINTS "${CMAKE_PREFIX_PATH}/lib" ${MPI_LIB_PATH}
  # 如果默认路径没找到libmpi.so，还会去MPI_LIB_PATH找，下游使用者可以设置这个变量值
)
find_path(MPI_INCLUDE_DIR
  NAMES mpi.h
  PATHS "${CMAKE_PREFIX_PATH}/include" ${MPI_INCLUDE_PATH}
  # 如果默认路径没找到mpi.h，还会去MPI_INCLUDE_PATH找，下游使用者可以设置这个变量值
)

## 使用 CMake 来编译
CMake 生成好编译环境后，底层的 make, ninja, MSBuild 编译命令都是不一样的，但 CMake 提供了一个统一的方法进行编译：

cmake --build .
使用--buildflag，CMake 就会调用底层的编译命令，在跨平台时十分方便。

对于 Visual Studio，其 Debug 和 Release 环境是基于 configuration 的，因此CMAKE_BUILD_TYPE变量无效，需要在 build 时指定：
cmake --build . --config Release


## CMake 的缺陷
CMake 的缺陷是很明显的，入门成本很高，其语法的设计也很糟糕，find_package这些函数不会返回结果，而是对全局变量或是 target 产生副作用，函数的行为不查阅文档是很难预测的。 并且在 CMake 中，变量，target，字符串的区分不明确，很容易让人感到迷惑，不知道什么时候应该使用${}去读取值。



# 

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