# openssl vs工程
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