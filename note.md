
# cmd

```
go test -run=^$ -bench=twoSumS1
go test -run=^$ -bench=twoSumS2
go test -run=^$ -bench=.
go test -run=^$ -bench=. -benchtime=2s

go test -run=^$ -bench=. -benchmem -cpuprofile profile.out
go test -run=^$ -bench=. -benchmem -memprofile memprofile.out -cpuprofile profile.out

go tool pprof profile.out
```

# 常用

```
go test -run=^$ -bench=. -benchmem -memprofile memprofile.out -cpuprofile profile.out
go tool pprof -http=":8089" profile.out
go tool pprof -http=":8089" memprofile.out
```

# 含义

```
Benchmark_twoSumS1-12           26250096                46.0 ns/op            16 B/op          1 allocs/op
Benchmark_twoSumS2-12           49687029                22.8 ns/op            16 B/op          1 allocs/op
                   |                |                       |                    ｜                 |
               GOMAXPROCS   调用被测试代码的次数          每次执行时间         每次执行内存分配   每次执行内存分配次数


ok      github.com/mjoker73/gap/problems/p_00001_two_sum        3.814s
                                                                  |
                                                              累计消耗时间
```

# 日常算法训练大纲

- 必读
  - [x] 二分搜索

- 第一章 数组题目
  - [ ] 如何运用二分查找算法
  - [x] [双指针技巧总结](https://labuladong.gitbook.io/algo/shu-ju-jie-gou-xi-lie/shou-ba-shou-shua-shu-zu-ti-mu/shuang-zhi-zhen-ji-qiao)
    - [19.删除链表倒数第 N 个元素（中等）](problems/p_00019_remove_nth_from_end)
      - 双指针，前面的指针慢N步即可，注意前面的指针停止条件（`right! = nil or right.Next != nil`）
    - [141.环形链表（简单）](problems/p_00141_has_cycle)
      - 快慢指针，遇不到即可。注意前面的指针停止条件（`right! = nil or right.Next != nil`）
    - [167.两数之和 II - 输入有序数组（中等）](problems/p_00167_two_sum)
      - 从第二个数开始遍历，每次拿出一个来里面做Logn的二分查找即可，复杂度nlogn
    - [344.反转字符串（简单）](problems/p_00344_reverse_string)
      - 指针反转即可
    - [876. 链表的中间结点](problems/p_00876_middle_node)
      - 快慢指针
  - [ ] 滑动窗口技巧
  - [ ] O（1）时间查询删除数据任意元素
  



# 日常算法训练明细




