# [openssl vs工程](https://github.com/janbar/openssl-cmake)
openssl-cmake-no-asm编译调试版本.zip 
上面的压缩包中构建了cmake脚本用于生成vs可调试的工程


# rsa算法区别
```cpp
# define EVP_PKEY_RSA    NID_rsaEncryption
# define EVP_PKEY_RSA2   NID_rsa
# define EVP_PKEY_RSA_PSS NID_rsassaPss

#define NID_rsaEncryption       6
#define NID_rsa             19
#define NID_rsassaPss           912
```
都是EVP_PKEY_RSA类型，然后覆盖相同的RSA密钥对，但在不同的上下文中使用不同的对象标识符，
即PKCS1或X509证书

从*crypto/objects/obj_dat.h*提取的代码行
对于EVP_PKEY_RSA NID_RSA加密：
{"rsaEncryption","rsaEncryption",NID_rsaEncryption,9,&(lvalues[38]),0},
0x2A,0x86,0x48,0x86,0xF7,0x0D,0x01,0x01,0x01,/* [ 38] OBJ_rsaEncryption */    
这是PKCS1 RSA加密1.2.840.113549.1.1.1


{"RSA", "rsa", NID_rsa, 4, &so[103]},
0x55,0x08,0x01,0x01,                           /* [  103] OBJ_rsa */
这是针对X.500定义算法的rsa加密id ea rsa 2.5.8.1.1


{"RSASSA-PSS", "rsassaPss", NID_rsassaPss, 9, &so[5959]},
0x2A,0x86,0x48,0x86,0xF7,0x0D,0x01,0x01,0x0A,  /* [ 5959] OBJ_rsassaPss */
1.2.840.113549.1.1.10
Rivest, Shamir, Adleman (RSA) Signature Scheme with Appendix - Probabilistic Signature Scheme (RSASSA-PSS)


rsa加解和签名原文数据长度都是模长，签名数据必须<=模长-11
rsa加解密填充方式
case RSA_PKCS1_PADDING:
case RSA_PKCS1_OAEP_PADDING:
case RSA_SSLV23_PADDING:
case RSA_NO_PADDING:


ecc签名原文数据必须是模长
ecc加密原文任意




# [Linux多个库间的符号冲突问题](https://huaweicloud.csdn.net/63566afdd3efff3090b5eceb.html)
[linux c解决多个第三方so动态库包含不同版本openssl造成的符号冲突](https://blog.csdn.net/found/article/details/105263450)

## 解决动态库符号冲突的两个方法，一个是减少导出的符号，一个是优先使用本动态库中的符号

### 减少导出的符号，只导出指定符号:
- 1、加编译器选项fvisibility=hidden，加了这个选项后，默认的符号都不会导出
- 2、在需要导出的函数或者类名前加__attribute__ ((visibility("default")))

### 优先使用本动态库中的符号
很多动态库在制作的时候都是默认的把所有符号导出，你没法保证自己的动态库不错误地引用到其他动态库中的符号。
解决的办法也有两种：
- 1、在编译期解决，就是在编译动态库的是加参数-Wl,-Bsymbolic 这个参数是传给链接器,这个编译参数的作用是：优先使用本动态库中的符号，而不是全局符号。这样即使其他动态库导出的符号和自己动态库中的符号同名，冲突也不会发生，运行自己动态库程序的时候会使用自己本动态库中的函数和类。
- 2、在加载动态的是时候解决，如果你没法重新编译动态，可以在加载动态库的时候自己使用dlopen函数加载动态库，然后在增加RTLD_DEEPBIND这个标志，这个标志的解释是这样的：
RTLD_DEEPBIND (since glibc 2.3.4)将符号的查找范围放在此共享对象的全局范围之前。这意味着自包含对象将优先使用自己的符号，而不是全局符号，这些符号包含在已加载的对象中。
RTLD_LAZY | RTLD_LOCAL | RTLD_DEEPBIND
当希望dlopen载入的库首先从自己和它的依赖库中查找符号，然后再去全局符号中去查找时，就用RTLD_DEEPBIND。这样就允许dlopen载入的库中存在与全局符号重名的符号，而对这个载入的库来说错的或者可能引起问题的全局符号可能是由其它库引入的。


## 特殊情景
当前程序调用了高版本的openssl (静态库)库，同时依赖了低三方库，第三方库动态依赖系统低版本的openssl动态库。
两个openssl动态库之间某些相同的符号产生冲突。

解决符号冲突的一个利器是封装，把代码封装成动态库，只暴露几个必须的符号，对外部看来，表现得就像一个黑盒子。
编译gmssl这个库的时候把它编译成静态库，静态库中的符号也是全部导出的，然后进一步封装sm2加密和解密的方法，封装成gmutil.so动态库，完全屏蔽openssl的头文件，而且这个gmutil.so动态只导出sm2加密和sm2解密的函数符号，这样在gmutil.so的调用者看来，它内部引用的openssl对外部看来就是完全不可见的，而且gmutil.so编译的时候也指明-Wl,-Bsymbolic参数，这样就不会引用到外部的openssl版本的符号。
因为我编译gmssl.a静态库的时候，是用别人制作的makefile编译的，它里面的符号全部导出，即使在编译gmutil.so的时候，使用编译选项fvisibility=hidden，gmssl.a里面的符号也还是全部导出的。因此，加了-Wl,--exclude-libs,ALL后才隐藏了gmssl.a里面所有的符号。


```cpp

#include<stdio.h>
#include<stdlib.h>
#include<string.h>
#include<mysql.h>
#define MAX_COLUMN_LEN    32

#include<openssl/ssl.h>

//静态编译了高板本的openssl库在程序中调用，并且程序使用libmysqlclient进行sql的编程。
// libmysqlclient依赖的系统openssl库，与程序调用的openssl静态库不一致，会产生冲突。
//即一个程序依赖两个不同版本的openssl库。
//解决方式 -Wl,--exclude-libs,ALL  隐藏静态库的符号
//-Wl,--exclude-libs,ALL不自动导出库中的符号，也就是默认将库中的符号隐藏,因为静态依赖高版本的tassl,libmysqlclient本身依赖低版本的openssl,如果不隐藏tassl静态库中的符号，那么就会产生冲突
//apt-get install libmysqlclient-dev
//g++ -I/usr/include/mysql -I./linux/tassl/x86_64/Release/include  main.cpp -o test -L/usr/lib/mysql -lmysqlclient linux/tassl/x86_64/Release/lib/libssl.a linux/tassl/x86_64/Release/lib/libcrypto.a  -ldl -pthread  如此编译会产生冲突
//g++ -I/usr/include/mysql -I./linux/tassl/x86_64/Release/include  main.cpp -o test -L/usr/lib/mysql -lmysqlclient linux/tassl/x86_64/Release/lib/libssl.a linux/tassl/x86_64/Release/lib/libcrypto.a  -ldl -pthread -Wl,--exclude-libs,ALL
int main(int argc , char *argv[])
{
    MYSQL db;
	MYSQL_RES *res;  
    MYSQL_ROW row;  
	
    OPENSSL_init_ssl(OPENSSL_INIT_SSL_DEFAULT, NULL);

	//初始化数据库
    mysql_init(&db);
	//连接数据库
    if(mysql_real_connect(&db,"47.94.149.220","root","longmai","my",3306,NULL,0))
    {
		printf("connect!!!\n");
	}
	
	//查询
	if (mysql_real_query(&db, "select * from S_RESIDENCE", (unsigned long)strlen("select * from S_RESIDENCE")))  
    {  
        printf("mysql_real_query failed\n");
        return 0;  
    }
	
	// 存储结果集  
    res = mysql_store_result(&db);  
    if (NULL == res)  
    {  
        printf("mysql_store_result failed\n");  
        return 0;  
    }  
	
	// 重复读取行，并输出第一个字段的值，直到row为NULL  
    while (row = mysql_fetch_row(res))  
    {  
        printf("%s\n",row[0]);  
    }  
	
	// 释放结果集  
    mysql_free_result(res);  
    // 关闭Mysql连接  
    mysql_close(&db);  
	
	
	return 0;
}

```









## [符号冲突](https://www.jianshu.com/p/fb5a5550f858)
什么是符号冲突，就是库与库之间有相同的符号，使用者不知道用哪个；例如：A SDK有个符号a，B SDK也有个符号a，最终app调用a时，可能用的是A SDK的a，也可能是B SDK的a；这样的话，就会产生歧义，假如app想调用A SDK的a，但可能实际调用的却是B SDK的a，这样就会造成app行为异常，或是崩溃。

### 静态库之间符号冲突
静态库冲突经常会遇到下面几个问题：

- 为什么有些重复符号在链接时会报错，有些不会。
首先静态库包含的是.o文件；.o文件就是对应的每个cpp/c文件编译后的产物。当链接时，链接器会按app使用到的函数逐个扫描静态库里的.o，如果发现要链接的.o里存在着已链接过的符号就会报错。不同编译器的链接算法不同，结果也不同。下面以vs 2015，xcode clang，ndk21来分析。
 xcode clang: app使用到的函数是按字符串排列的，链接器会按这个顺序逐个扫描静态库，看下静态库里的.o是否存在在app使用到的函数，如果有就将.o所有符号放进全局符号表里，如果发现全局符号表里有相同的符号就报错

 vs2015 vc，ndk21 clang：链接器会按静态库链接顺序扫描静态库，看.o是否存在着app使用的函数，如果有就将.o所有符号放进全局符号表里，如果发现全局符号表里有相同的符号就报错

备注：上面的算法并不一定完全准确，因为这些链接器的代码并不开源，只是通过例子推测出来，有问题欢迎指正

- 下面，我们结合例子分析下
情形 1 ：
![WX20200827-111004@2x.png](https://upload-images.jianshu.io/upload_images/15592790-dd035b15b4d944b2.png?imageMogr2/auto-orient/strip|imageView2/2/w/756/format/webp)

上面情况，无论在xcode或是vs2015/ndk，app先链接谁就用谁的d函数，而且不会链接报错。
情形 2：
![WX20200827-112737@2x.png](https://upload-images.jianshu.io/upload_images/15592790-bee1cd1030ceff5a.png?imageMogr2/auto-orient/strip|imageView2/2/w/690/format/webp)

上面的情况：
在xcode下，因为链接器会先链接a函数，他会遍历当前的静态库，发现在a.o里，然后将a.o里的所有符号都放进全局符号里；当链接d函数时，因为d已经在全局符号，因此不需要将b.o放进全局符号，所以无论链接顺序是怎样，app始终用的是liba.a的d；
在vs2015/ndk下，当liba.a先链接时，链接器会发现a.o里存在着app需要的a，d函数，因此将a.o里的所有符号放进全局符号，因为app需要的函数都链接完了，所以不需要将b.o放进全局符号。当libb.a先链接时，链接器会发现b.o存在着app需要的d函数，因此将b.o所有符号放进全局符号。当链接到liba.a时，发现a.o里存在着app需要的a函数，当将a.o所有符号放进全局符号里时，发现已存在了d函数，因此就报符号冲突错误。
情形 3：
![WX20200827-180339@2x.png](https://upload-images.jianshu.io/upload_images/15592790-668851ef391fad83.png?imageMogr2/auto-orient/strip|imageView2/2/w/694/format/webp)
上面情况：无论在xcode或是vs2015/ndk都会报链接出错，因为无论怎么链接，都需要将a.o和b.o里的符号放进全局符号里。
- 链接顺序可以确保app使用的是哪个库的符号吗。
不同编译器结果不同；对于xcode不能保证，对于vs,ndk，只要不报错，app会用先链接的库的符号。
- 怎样查找静态库中的重复符号。
默认情况下，链接器是按需链接静态库，如果app没有用到.o里的函数，.o不会被链接到app，可以添加链接选项，让链接器将所有静态库的.o都链进app。这样重复的符号就会暴露出来，导致链接出错，以便我们分析，修改。
    对于vs2015，在链接选项里加上/WHOLEARCHIVE:a.lib，这样会强制将a里的.o链接到app
    对于xcode clang，在链接选项加上-all_load会强制链接所有静态库库到app，也可以用-force_load liba.a，只将a强制链接。
    对于ndk
        Android.mk: LOCAL_WHOLE_STATIC_LIBRARIES += a；或者通过LOCAL_LDFLAGS += -Wl,--whole-archive /path/liba.a -Wl,--no-whole-archive /path/libb.a
        CMakeLists.txt : target_link_libraries(myapp -Wl,--whole-archive a -Wl,--no-whole-archive b)
- 如何解决静态库之间的符号冲突
    更改名字：最原始有效的方法。
    声明强弱符号：这种方法比较少用，也不太实际，有兴趣的自行查找使用方法

### 动态库与静态库之间符号冲突
- 可以将动态库视为只有一个.o的静态库，链接算法与静态库差不多，但有一点区别：
对于xcode和ndk，当静态库遇到动态库符号时，动态符号会被覆盖掉，而不是报错
对于vs，算法与静态库一样，发现有相同的符号时，一样会报错。
- 在编译链接不报错的情况下，静态库先链接，一定会优先用静态库的符号
- 如何解决动态库与静态库之间的符号冲突
    在xcode和ndk下，是没办法在编译链接时期将冲突暴露出来，所以只能查看动态库的导出符号和静态库的符号，然后更改相同符号的名字。
    对于动态库去除不必要的符号导出，这样能减少与静态库的冲突。
- 同一个静态库里有相同的符号是非常坑的，当编译源文件顺序不同时，最终链接的结果也不同。
### 动态库与动态库之间符号冲突
- 动态库之间相同的符号在链接时不会报错，先链接谁就用谁的符号。所以要解决他们之间的冲突，只能查看动态库的导出符号，更改相同的名字；其次是去除不必要的符号导出，减少冲突的可能性。
- app本质上也是一个动态库。
### 动态库的符号查找问题。
- 动态库是如何查找他依赖的函数呢？
    对于win，动态库会有个导入表，里面存储着他链接时所依赖的库和对应依赖的符号；如下图；可以用 dumpbin a.dll /IMPORTS 来列出所依赖的导入信息。当动态库被加载时，加载器会读取这个表，依次加载所依赖的动态库，从依赖的库中拿到依赖函数的地址填入表中。
![wim.png](https://upload-images.jianshu.io/upload_images/15592790-6773dc8525dca18b.png?imageMogr2/auto-orient/strip|imageView2/2/w/464/format/webp)
- 对于android，动态库存在2个表来存储这些信息。

    一个是链接时所依赖的so的导入库表，这个表的顺序是：直接依赖的编译链接顺序，间接依赖用户库链接顺序，间接依赖系统库顺序，例如，如果a库直接依赖b，c，而b又依赖d和系统库e，那么a的导入库表将是b，c，d，e；
    一个是所依赖的导入符号表。
    当需要使用一个符号时，加载器会去导入库表查找so，会用先找到的so里的符号，找不到则报错
    我们可以用arm-linux-androideabi-readelf -d liba.so来列出所依赖的动态库。然后用arm-linux-androideabi-nm liba.so -D来列出所依赖的导入符号。
![aim.png](https://upload-images.jianshu.io/upload_images/15592790-1dd77fa8d5409509.png?imageMogr2/auto-orient/strip|imageView2/2/w/622/format/webp)
- 对于ios/macos，动态库的导入信息与win类似；如下图；我们可以用/Applications/Xcode.app/Contents/Developer/Toolchains/XcodeDefault.xctoolchain/usr/bin/dyldinfo -lazy_bind liba.dylib 来列出所依赖的导入符号信息。
![mim.png](https://upload-images.jianshu.io/upload_images/15592790-1bcae0f1247f9b73.png?imageMogr2/auto-orient/strip|imageView2/2/w/444/format/webp)
### 动态库链接的一些问题
- vs导入表里存的是他需要的信息，如果将一个无关的库b加入c库编译链接过程里，c库里不会储存b的信息，而且存的是他直接依赖的库信息；
- ndk,xcode，只要参加了链接都会保存对其的依赖信息，而且间接依赖库也会保存其中。
- ndk依赖的符号与库是没有明显的对应关系，这会存在一个问题，如果a.so，b.so同时存在a符号，而c同时依赖a和b库；在开始时c调用的a符号是属于a库的，但如果在某次升级中将a库中的a符号去掉，此时a的调用就会跑到b库去了。vs和xcode在这种情况下，程序直接会出错，反馈说在a库里找不到a符号。
### 总结
- 别以静态库形式提供给客户，静态库的符号冲突比较隐蔽，机率比较大，而且修改成本也大；优先用动态库。
- 通过去掉不必要的导出符号，能降低动态库符号冲突的机率，但是代价比较大，特别是多团队合作的时候。从上面看出在vs和xcode里，依赖符号和依赖库是有明确的对应关系，因此可以将接口和核心功能分成两个动态库，只让接口动态库参加到客户的编译链接。这样客户的代码就不会链接到我们的核心库，冲突的几率会降低很多，万一接口库与客户有相同的符号，要修改的范围也小很多。
### 查看动态库的导出符号
- vs:在vc bin目录下的dumpbin可以查看；例如：dumpbin a.dll /EXPORTS
- ndk:在ndk工具目录下的xxx-nm可以查看；例如：arm-linux-androideabi-nm liba.so -D，其中类型“D”表示的是导出变量，“T”表示的是导出函数，“U”表示依赖别的动态库符号，即导入符号
- xcode:在xcode工具目录下也有个dyldinfo可以查看；例如：/Applications/Xcode.app/Contents/Developer/Toolchains/XcodeDefault.xctoolchain/usr/bin/dyldinfo -export liba.dylib
### 查看静态库的符号
- vs:使用dumpbin工具；例如：dumpbin a.lib /SYMBOLS /ARCHIVEMEMBERS 会列出所有.o的符号
类似02F 00000000 SECTA notype () External | _printf 这种，表示“printf”是.o的符号；
类似029 00000000 UNDEF notype () External | _foo这种，表示“foo”是引用别的模块的符号
- ndk : 使用ndk工具中的readelf；例如 arm-linux-androideabi-readelf liba.a -s
类似 00000001 28 FUNC GLOBAL DEFAULT 11 test这种，表示“test”是.o的符号；
类似 * 00000000 0 NOTYPE GLOBAL DEFAULT UND foo*这种，表示“foo”是引用别的模块的符号
- xcode: 使用xcode工具中的objdump；例如 objdump liba.a --syms
类似：0000000000000020 g F __TEXT,__text _test这种，表示“test”是.o的符号；
类似： 0000000000000000 UND _foo这种，表示“foo”是引用别的模块的符号
### 去除不必要的符号导出
- vs：vs默认是不会将符号导出的，所有要导出的符号都必须声明为__declspec(dllexport)，所以只要将不必要导出的符号去掉这个声明就可以了
- ndk: 可以增加编译选项-fvisibility=hidden，这样默认所有符号就不会导出，对于要导出的符号显示声明__attribute__ ((visibility ("default")))
- xcode：可以通过ndk方式实现；也可以在xcode的设置中将Symbols Hidden by Default设置为Yes，但是这方式有个坑要注意，当Enable Testability为Yes时，前面的设置无效，所以最直接的方式是ndk那种。



-fvisibility=hidden 隐藏本地符号

-fvisibility-inlines-hidden 隐藏本地内联符号

-Wl,-Bsymbolic 优先使用本地符号，消除odr

-Wl,--exclude-libs,ALL 隐藏依赖的静态库符号



### 关于动态库加-fPIC编译选项

是为了生成位置无关的代码，这样多个程序就有可能共享同一个动态库。如果不加这个选项，动态库被加载的时候都要进行地址重定向到自己的进程空间，这样导致每一个使用这个so的进程都会拷贝一份副本。而加了-fPIC这个选项，动态库加载的时候就不需要重定向地址，及位置无关代码，这样多个进程就可以共享同一个so。