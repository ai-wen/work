# Debug版都是MDd，Release版都是MD，才不会崩溃。


# 《Windows核心编程 第五版》第19章 DLL基础（511页）中给出了一个建议：“当一个MT版本的模块如果提供一个内存分配函数的时候，它必须同时提供另一个用来释放内存的函数。”。说得更加直白一点就是，“对于MT的模块，不要跨模块进行内存释放。”。

- Windows的堆管理器对每个进程都维护了多个“堆”，我们从每个“堆”中分配处理的内存块的地址都不一样。所以我们不能将从“堆A”中分配出来的内存块拿到“堆B”中，让“堆B”来释放，这样就会导致程序异常。

- MT 模式 
    现在有2个模块（A.dll和B.dll）都是使用MT运行时库，即加载的静态库libcmt.lib（可以参考《理解C/C++运行时库》），在A.dll中使用malloc分配100字节的内存，malloc返回的内存地址为0x123456。然后将该地址传给B.dll，在B.dll中调用free函数来释放这个内存。
    从《Windows内存体系（7） – 堆》我们知道，DLL在启动代码_DllMainCRTStartup中会建立一个“堆”（堆句柄存放在_crtheap变量中），所以A.dll和B.dll中都会有一个crt堆。

    为了区分，我们将A.dll中的crt堆称作_crtheap_A，B.dll中的crt堆称作_crtheap_B。

    A.dll中malloc的内存拿到B.dll去中去free，就相当于从堆_crtheap_A中分配的内存拿到另一个堆_crtheap_B中的释放。第一节已经解释了为什么不能这样做了。
- MD 模式
    使用MD运行时库，即加载的动态库msvcr100.dll（可以参考《理解C/C++运行时库》），程序的代码的过程和上面一样，还是在A.dll中使用malloc分配100字节的内存，malloc返回的内存地址为0x123456。然后将该地址传给B.dll，在B.dll中调用free函数来释放这个内存。但是这个时候程序却不会崩溃。
    因为A、B两个dll都是链接的·msvcr100.dll·，同一个dll在一个进程只会被加载一次，所以进程中只会有一个crt堆（_crtheap），malloc和free都是运行时库提供的函数，所以都会调到运行时库里面去，然后从运行时库里面的_crtheap分配和释放内存块。因为分配和释放都是在同一个堆上，所以不会崩溃。


# 测试跨模块传递 STL 对象

```cpp
int main(int argc, char* argv[])
{
    std::string str = "test";

    TestFun(str);

    return 0;
}

//DLL
DLL_API void TestFun( std::string str)
{
    return;
}
```

std::string在进行值传参的过程中会执行一次深拷贝，即：在堆上分配内存块，拷贝“test”到内存块中，然后将临时形参std::string对象传递到dll中，dll中的TestFun函数在作用域结束后对临时形参进行释放时就出现了错误，因为尝试在dll的crt堆中释放由在exe的crt堆中分配的内存块。


一种方案就是让std::string统一在进程的默认堆上分配内存块，而不是在各个模块的crt堆上分配。
下面自定义的内存分配器vm_allocator<char>定义了mystring类，我们只需要将TestFun函数接口中的std::string修改为mystring即可解决崩溃问题。

```cpp
#include <windows.h>
#include <string>
#include <vector>

template <typename T>
class vm_allocator : public std::allocator<T> {
public:
    typedef size_t size_type;
    typedef T* pointer;
    typedef const T* const_pointer;

    template<typename _Tp1>
    struct rebind {
        typedef vm_allocator<_Tp1> other;
    };

    pointer allocate(size_type n, const void *hint = 0) {
        UNREFERENCED_PARAMETER(hint);
        void* pBuffer = HeapAlloc(GetProcessHeap(), HEAP_ZERO_MEMORY, n * sizeof(T));

        return (pointer)pBuffer;
    }

    void deallocate(pointer p, size_type n) {
        UNREFERENCED_PARAMETER(n);
        if (p) {
            HeapFree(GetProcessHeap(), 0, p);
        }
    }

    vm_allocator() throw() : std::allocator<T>() {
    }

    vm_allocator(const vm_allocator &a) throw() : std::allocator<T>(a) {
    }

    template <class U>
    vm_allocator(const vm_allocator<U> &a) throw() : std::allocator<T>(a) {
    }

    ~vm_allocator() throw() {
    }
};

typedef std::basic_string<char, std::char_traits<char>, vm_allocator<char> > mystring;
```