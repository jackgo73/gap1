
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
  - [x] [二分搜索](https://labuladong.gitbook.io/algo/bi-du-wen-zhang/er-fen-cha-zhao-xiang-jie)
    - [34. 在排序数组中查找元素的第一个和最后一个位置](problems/p_00034_search_range)
    - [704. 二分查找](problems/p_00704_binary_search)
  - [x] [二叉树第一期](https://labuladong.gitbook.io/algo/bi-du-wen-zhang/er-cha-shu-xi-lie-1)
    - [114.二叉树展开为链表（中等）](problems/p_00114_flatten)
    - [226.翻转二叉树（简单）](problems/p_00226_invert_tree)
    - （重点）[144. 二叉树的前序遍历（中等）](problems/p_00144_preorder_traversal)
    - （重点）[145. 二叉树的后序遍历（中等）](problems/p_00145_postorder_traversal)

- 第一章 数组
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
  - [x] [滑动窗口技巧](https://labuladong.gitbook.io/algo/shu-ju-jie-gou-xi-lie/shou-ba-shou-shua-shu-zu-ti-mu/hua-dong-chuang-kou-ji-qiao-jin-jie)
    -  [3. 无重复字符的最长子串](problems/p_00003_length_of_longest_substring)
  - [ ] O（1）时间查询删除数据任意元素
  

- 第二章 二叉树
  - [x] [二叉树第一期](https://labuladong.gitbook.io/algo/bi-du-wen-zhang/er-cha-shu-xi-lie-1)
    - [114.二叉树展开为链表（中等）](problems/p_00114_flatten)
    - [226.翻转二叉树（简单）](problems/p_00226_invert_tree)
  - [ ] [二叉树第二期](https://labuladong.gitbook.io/algo/shu-ju-jie-gou-xi-lie/shou-ba-shou-shua-er-cha-shu-xun-lian-di-gui-si-wei/er-cha-shu-xi-lie-2)
    - 

- 高频
  - [ ] [如何寻找最长回文子串](https://labuladong.gitbook.io/algo/gao-pin-mian-shi-xi-lie/zui-chang-hui-wen-zi-chuan)
    - [5.最长回文子串（中等）](problems/p_00005_longest_palindrome)
  - [ ] [如何判定括号合法性](https://labuladong.gitbook.io/algo/gao-pin-mian-shi-xi-lie/he-fa-kuo-hao-pan-ding)
    - [20.有效的括号（简单）](problems/p_00020_is_valid)

# 知识点

数据类型？
```
int8    1字节  -128                  到  127
int16   2字节  -32768                到  32767
int32   4字节  -2147483648           到  2147483647
int64   8字节  -9223372036854775808  到  9223372036854775807

默认使用int64，包括a := 1或 var a int其实都是int64
```



golang中map的原理？



golang垃圾回收机制？



快速排序，归并排序
https://labuladong.gitbook.io/algo/bi-du-wen-zhang/er-cha-shu-xi-lie-1
举个例子，比如说我们的经典算法「快速排序」和「归并排序」，对于这两个算法，你有什么理解？如果你告诉我，快速排序就是个二叉树的前序遍历，归并排序就是个二叉树的后序遍历，那么我就知道你是个算法高手了。

深度优先遍历四种主要：：
前序遍历：根结点 ---> 左子树 ---> 右子树
中序遍历：左子树---> 根结点 ---> 右子树
后序遍历：左子树 ---> 右子树 ---> 根结点
广度优先遍历层次遍历：只需按层次遍历即可