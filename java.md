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



