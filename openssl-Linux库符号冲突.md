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