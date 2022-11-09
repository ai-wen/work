# [C++](https://www.zhihu.com/column/c_1360316516295032832)


# C++ 高性能编程实战系列
[概述](https://zhuanlan.zhihu.com/p/533708198)
[并发优化](https://zhuanlan.zhihu.com/p/534004366)
[内存优化](https://zhuanlan.zhihu.com/p/537260855)

# 
[详解三大编译器：gcc、llvm 和 clang](https://zhuanlan.zhihu.com/p/357803433)
[gcc 中遇到的问题，这里有你想要的全部答案](https://zhuanlan.zhihu.com/p/458193070)




# 大小端
```cpp
void tt()
{
    const union {
        int one;
        char little;
    } check_endian = { 1 };

    if (check_endian.little)
    {
        //小端
    }
    else
    {
        //大端
    }
}
```

# 全局只初始化一次的过程

```cpp
#ifdef WIN32
    typedef long CRYPTO_ONCE;
    #define CRYPTO_ONCE_STATIC_INIT 0
#else
    #include <pthread.h>
    typedef pthread_once_t CRYPTO_ONCE;
    #define CRYPTO_ONCE_STATIC_INIT PTHREAD_ONCE_INIT
#endif



#ifdef WIN32

#  define ONCE_UNINITED     0
#  define ONCE_ININIT       1
#  define ONCE_DONE         2

/*
 * We don't use InitOnceExecuteOnce because that isn't available in WinXP which we still have to support.
 */
int CRYPTO_THREAD_run_once(CRYPTO_ONCE *once, void (*init)(void))
{
    long volatile *lock = (long *)once;
    long result;

    if (*lock == ONCE_DONE)
        return 1;

    do {
        result = InterlockedCompareExchange(lock, ONCE_ININIT, ONCE_UNINITED);
        if (result == ONCE_UNINITED) {
            init();
            *lock = ONCE_DONE;
            return 1;
        }
    } while (result == ONCE_ININIT);

    return (*lock == ONCE_DONE);
}

#else

int CRYPTO_THREAD_run_once(CRYPTO_ONCE *once, void (*init)(void))
{
    if (pthread_once(once, init) != 0)
        return 0;

    return 1;
}

#endif

CRYPTO_ONCE engine_lock_init = CRYPTO_ONCE_STATIC_INIT;
CRYPTO_THREAD_run_once(&engine_lock_init, do_engine_lock_init);

```

