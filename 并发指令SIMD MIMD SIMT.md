
# [并发优化](https://zhuanlan.zhihu.com/p/534004366)

## 1、单线程中的并发
CPU 指令集的发展经历了 MMX（Multi Media eXtension）、SSE（Streaming SIMD Extensions）、AVX（Advanced Vector Extensions）、IMCI 等。
单核单线程内 SIMD（Single Instruction Multiple Data），
多核多线程中 MIMD（Multiple Instruction Multiple Data）。

## SIMD允许使用单一命令对多个数据值进行操作。这是一种增加CPU的计算能力的便宜的方法。仅需要宽的ALU和较小的控制逻辑。
Intel从MMX开始支持SIMD。ARM已经通过NEON技术将SIMD扩展介绍进了ARM-Cortex架构。NEON SIMD单元128-bit宽，包含16个128-bit的寄存器，能够被用来当做32个64-bit寄存器。这些寄存器能被当做是同等数据类型的的vector。

一般地，SIMD仅包含基础的算数操作（加减乘除）和其他的像求绝对值和求平方根。

对SIMD的高性能有贡献的一个因素是，多个数据元素能够同时从内存读出/写入，利用内存数据总线的所有宽度。图2.11展现了使用SIMD和vector processing的一个简单图示。使用SIMD的代码大约是不使用SIMD的1/4，执行周期大约是不使用的1/4.

## SIMT：single instruction，multiple threads。SIMT类似CPU上的多线程。最简单的理解SIMT的是想象有这样一个多核系统，每一个core有自己的寄存器文件、自己的ALU、自己的data cache，但是没有独立的instruction cache、没有独立的解码器、没有独立的Program Counter register，命令是从单一的instruction cache同时被广播给多个SIMT core的。即所有的core是各有各的执行单元，数据不同，执行的命令确是相同的。多个线程各有各的处理单元，和SIMD公用一个ALU不同。
SIMT在GPU上，GPU是有成百上千的单独的计算单元的。硬件实现上，明显GPU更复杂，成本也更高


```cpp
利用 SSE 指令进行优化 FFT（快速傅立叶变换）：将 Mat1 和 Mat2 矩阵元素乘积之后更新到 Mat2
// 优化前
void MatMulti(Mat m1, Mat m2) {
    for (int i = 0; i < m1.rows; i++) {
        float *pixel_1 = (float *)m1.data + i * m1.step / 4;  // 32f
        float *pixel_2 = (float *)m2.data + i * m2.step / 4;  // 32f
        for (int j = 0; j < m1.cols; j++) {
            *pixel_2 = (*pixel_1) * (*pixel_2);
            pixel_1 += 1;
            pixel_2 += 1;
        }
    }
}

// 优化后
void SSEMatMulti(Mat m1, Mat m2)
{
    for (int i = 0; i < m1.rows; i++)
    {
        float *pixel_1 = (float *)m1.data + i * m1.step / 4;  // 32f
        float *pixel_2 = (float *)m2.data + i * m2.step / 4;  // 32f
        for (int j = 0; j < m1.cols; j++)
        {
            __m128 sse_1 = _mm_load_ps(pixel_1);  // 将 pixel_1 地址指向的值复制给 sse_1
            __m128 sse_2 = _mm_load_ps(pixel_2);  // 将 pixel_2 地址指向的值复制给 sse_2
            __m128 h = _mm_mul_ss(sse_1, sse_2);  
            _mm_storer_ps(pixel_2, h);
            pixel_1 += 1;
            pixel_2 += 1;
        }
    }
}
```

## 2、多线程中的并发
### 临界区保护技术
    Mutual Execlusion(pessimistic locking)：基本的互斥技术，存在某个时间周期，算法没有任何实质进展，典型的悲观锁算法
    Lock Free (optimistic locking)：组成算法的一个线程没有任何实质进展，基于 CAS 同步提交，若遇到冲突，回滚
    Wait Free：任意时间周期，算法的任意一个线程都有实质进展
### 并发队列
### 伪共享