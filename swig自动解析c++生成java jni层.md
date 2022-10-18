# 示例

```swig
%module CipherSdk                   /*指定生成的库名*/

%include <typemaps.i>               /*%apply 用一种c++类型替代另一种c++类型*/
%include <arrays_java.i> 
%apply signed char *INOUT { unsigned char * };   /*这种声明方式，对于输入输出参数会在jni层直接取参数的地址使用*/
%apply signed char[] { const unsigned char * };  /*这种声明方式，对于输入参数会在jni层对参数进行拷贝到一个临时变量,临时变量是new的同一类型,调用c库时使用临时变量进行传参*/


%apply int *INOUT { int *piEncDataLength, int *piDataLength};   /*%apply 用一种c++类型替代另一种c++类型，并且指定了参数的具体名称*/

%apply unsigned long *INOUT { void **hkey};
%apply unsigned long [] {const void *hkey}; /*这种声明方式，对于输入参数会在jni层对参数进行临时new且拷贝*/

%{  
#include "../CipherSdk.h"  
%}  
  
%include "../CipherSdk.h" 

%pragma(java) jniclasscode=%{
  static {
    try {
        System.loadLibrary("CipherSdk");
    } catch (UnsatisfiedLinkError e) {
      System.err.println("Native code library failed to load. \n" + e);
      System.exit(1);
    }
  }
%}

```

## 命令
swig -java CipherSdk.i

# 链接
[windows 下载地址](https://sourceforge.net/projects/swig/files/swigwin/)
[2019-04-08 Swig java Jni开发指南](https://cloud.tencent.com/developer/article/2001970)
[OpenSSL for Python](https://github.com/mcepl/M2Crypto/tree/e28c791fd5f51593a197c6ac160aaecf59b25383)
[JNI通过构建工具封装Swig一步生成.so](https://www.jianshu.com/p/745f46b93783)
[](http://web.mit.edu/svn/src/swig-1.3.25/Examples/java/pointer/index.html)