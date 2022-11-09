# 收集动态依赖库
//grep 取行，awk 按条件取指定列，cut 按分隔符取指定列。
//sed主要是用来将数据进行选取、替换、删除、新增的命令。可以放在管道符之后处理。

ldd RandomCheckTool | cut -d ">" -f 2 |cut -d "(" -f 1
cp $(ldd RandomCheckTool | cut -d ">" -f 2 |cut -d "(" -f 1 ) lib/
ll $(ldd RandomCheckTool | cut -d ">" -f 2 |cut -d "(" -f 1 )
ll $(ldd RandomCheckTool | cut -d ">" -f 2 |cut -d "(" -f 1 ) | awk '{print $11,$10,$9}' | cut -d '>' -f 2
