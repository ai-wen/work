

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

