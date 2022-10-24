# jar
jar 与 zip 唯一的区别就是在 jar 文件的内容中，包含了一个 META-INF/MANIFEST.MF 文件，该文件是在生成 jar 文件的时候自动创建的，作为jar里面的"详情单"，包含了该Jar包的版本、创建人和类搜索路径Class-Path等信息，当然如果是可执行Jar包，会包含Main-Class属性，表明Main方法入口，尤其是较为重要的Class-Path和Main-Class

jar包主要是对class文件进行打包，而java编译生成的class文件是平台无关的，这就意味着jar包是跨平台的，所以不必关心涉及具体平台的问题。


## 查看详情
jar -tf xxx.jar


## class文件打包
jar -cvf lm_cipher_sdk.jar com/imooc/jardemo1/Welcome.class com/imooc/jardemo1/impl/Teacher.class

c表示要创建一个新的jar包，v表示创建的过程中在控制台输出创建过程的一些信息，f表示给生成的jar包命名

这样包中MANIFEST.MF文件中不会创建Main-Class属性。
所有这是一个无法执行的包，只能调用.


## 可执行包
- 1 编辑MANIFEST.MF文件
    jar -cvfm lm_cipher_sdk.jar META-INF/MANIFEST.MF *

    可以编辑后在重新打包：
    编辑MANIFEST.MF文件为：
    Manifest-Version: 1.0 
    Created-By: 11 (Oracle Corporation) 
    Main-Class: com.imooc.jardemo1.Welcome
    注意冒号之后一定要跟英文的空格，整个文件最后有一行空行

    jar -cvfm lm_cipher_sdk.jar META-INF/MANIFEST.MF com/imooc/jardemo1/Welcome.class com/imooc/jardemo1/impl/Teacher.class

    jar -cvfm其中多了一个参数m，表示要定义MANIFEST文件。

- 2 指定main函数所在类进行打包
    java -cp lm_cipher_sdk.jar Test
    
    其中cp表示classpath，后面接上全限的main函数所在的类即可

## 简易打包
javac xxx.java -d target
将class文件都编译到 target 目录中

jar -cvf lm_cipher_sdk.jar  *


##  执行jar包
java -jar lm_cipher_sdk.jar





# 判断操作系统

public enum EPlatform {
	Any("any"),
	Linux("Linux"),
	Mac_OS("Mac OS"),
	Mac_OS_X("Mac OS X"),
	Windows("Windows"),
	OS2("OS/2"),
	Solaris("Solaris"),
	SunOS("SunOS"),
	MPEiX("MPE/iX"),
	HP_UX("HP-UX"),
	AIX("AIX"),
	OS390("OS/390"),
	FreeBSD("FreeBSD"),
	Irix("Irix"),
	Digital_Unix("Digital Unix"),
	NetWare_411("NetWare"),
	OSF1("OSF1"),
	OpenVMS("OpenVMS"),
	Others("Others");
	
	private EPlatform(String desc){
		this.description = desc;
	}
	
	public String toString(){
		return description;
	}
	
	private String description;
}

String OS = System.getProperty("os.name").toLowerCase();
if(OS.indexOf("linux")>=0)
if(OS.indexOf("windows")>=0)
if(OS.indexOf("mac")>=0)


# JNI loadlibrary

```java
public static boolean JNI_LoadLibrary(){
			
		String OS = System.getProperty("os.name").toLowerCase();
	
		String [] libs = {"lm_cipher_sdk.1.0.dll","liblm_cipher_sdk.1.0.so","liblm_cipher_sdk.1.0.dylib"};
	
		String curDir = System.getProperty("user.dir");
	
		String libName = libs[0];
		if(OS.indexOf("windows")>=0)
		{
			libName = libs[0];
		}
		else if(OS.indexOf("mac")>=0)
		{
			libName=libs[2];
		}
		else
		{
			libName=libs[1];
		}
	
		
		try{															
			
			InputStream is = lm_cipher_sdkJNI.class.getResourceAsStream(libName);
			if(is == null){		
				System.exit(1);	
			}

			//File f = new File("." + File.separator + libs[i]);	
			libName = curDir + File.separator + libName;
		
			String path = lm_cipher_sdkJNI.class.getProtectionDomain().getCodeSource().getLocation().getFile();
			if(path.endsWith(".jar"))
			{
				File f = new File(libName);
				if(!f.exists()){
					f.createNewFile();
				}	
			
				FileOutputStream os = new FileOutputStream(f);	
				byte[] cache =  new byte[1024];				
				int realRead = is.read(cache);
				while(realRead != -1){
					os.write(cache, 0, realRead);
					realRead = is.read(cache);
				}
				os.close();
				is.close();
			}
			else
			{
				int count = is.available();
				byte[] cache = new byte[count];
				is.read(cache);
				is.close();

				File f = new File(libName);
				if(!f.exists()){
					f.createNewFile();
				}	
				FileOutputStream os = new FileOutputStream(f);
				os.write(cache, 0, count);
				os.close();
			}

		}catch(Exception e){				
			e.printStackTrace();	
			System.exit(1);	
		}
	
		try {
			//System.loadLibrary("lm_cipher_sdk.1.0");
			System.load(libName);
			return true;
		}catch (UnsatisfiedLinkError e) {
			System.exit(1);	
		}	
		return false;	
	}
```


# java ssl

## 配置ssl证书
```java
System.setProperty("javax.net.ssl.keyStore", "path/to/keystore.jks");
System.setProperty("javax.net.ssl.keyStorePassword", "changeit");
System.setProperty("javax.net.ssl.keyStoreType", "JKS");

System.setProperty("javax.net.ssl.trustStore", "path/to/truststore.jks");
System.setProperty("javax.net.ssl.trustStorePassword", "changeit");
System.setProperty("javax.net.ssl.trustStoreType", "JKS");
```	
## springboot项目配置ssl证书
application.properties 文件中配置
#https对应443端口
server.port=443
#证书路径
server.ssl.key-store=classpath:qcl.30paotui.com.jks
#证书密码
server.ssl.key-store-password=x6qm42y4auvhh
#证书类型
server.ssl.key-store-type=JKS

## keystore
Java也有一个统一管理信任证书的地方， 是一个叫 keystore 的文件 (默认是位于 $JAVA_HOME/lib/security/cacerts ) 
该文件使用 keytool 工具去管理 (该工具默认位于 $JAVA_HOME/bin/keytool )
```c
导入证书到java默认的jdk库中：
keytool -import -alias <证书别名> -keystore $JAVA_HOME/jre/lib/security/cacerts -file your.crt -storepass changeit  -trustcacerts

参数说明：
-keystore  执行后添加证书的位置
-file      证书的位置
-alias     证书的别名

查看证书：
keytool -list -keystore "/usr/java/jdk1.8.0_60/jre/lib/security/cacerts" -storepass changeit |grep "证书的别名"

删除证书,将证书从keystore移除：
keytool -delete -storepass changeit -keystore $JAVA_HOME/jre/lib/security/cacerts -alias "证书的别名"

```
1.keystore文件都受密码保护，访问一个已有的 keystore 文件时，会要求你验证密码，默认密码为 changeit，所以keytool命令要加上-storepass changeit
2.需用root用户执行。

## 自签名pem证书转换jks

首先将 PEM 转换为 PKCS12证书。
openssl pkcs12 -export -out certificate.p12 -inkey key.pem -in cert.pem

keytool -importkeystore -destkeystore keystore.jks -srcstoretype PKCS12 -srckeystore certificate.p12
keytool -list -keystore keystore.jks


openssl pkcs12 -in certificate.p12 -cacerts -nokeys -out chain.pem
keytool -importcert -trustcacerts -file chain.pem -keystore truststore.jks


假设有三个PEM证书文件：
ca.cert.pem —— 根证书文件
cert.cert.pem —— 证书文件
cert.key.pem —— 证书的密钥文件
我们需要将根证书ca.cert.pem转换成JKS格式的根证书truststore.jks；
并将cert.pem和cert.key.pem转换成JKS格式的证书keystore.jks。
- openssl pkcs12 -export -out cert.p12 -in cert.pem -inkey cert.key.pem
- keytool -importkeystore -destkeystore keystore.jks -srcstoretype PKCS12 -srckeystore cert.p12
- keytool -import -file ca.cert.pem -keystore truststore.jks  -trustcacerts

```demo
keytool -importkeystore -destkeystore keystore.jks -srcstoretype PKCS12 -srckeystore client_cert_xxx.pfx -srcstorepass 12345678 -deststorepass 12345678
keytool -list -keystore keystore.jks

keytool -importcert -keystore truststore.jks -file ca_root_cert.pem -storepass 12345678 -trustcacerts -alias ca -noprompt 
keytool -list -keystore truststore.jks

是否信任此证书? [否]:  y
```