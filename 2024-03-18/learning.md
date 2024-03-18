# MergeSortedArray

2つの整数配列 nums1 と nums2 が与えられ、それぞれ nums1 と nums2 の要素数を表す。
nums1 と nums2 を、減らない順にソートされた1つの配列にマージする。

最終的にソートされた配列は関数から返されるのではなく、 配列 nums1 の中に格納される。

これに対応するため、nums1はm + nの長さを持ち、
最初のm個の要素はマージされるべき要素を表し、
最後のn個の要素は0に設定され無視される。

# コードを書く前に考えたこと
- nums1.length == m + n なので、常にnが入る文の場所は空いている
  - [1,0], [1,2] ということはない
- nums1 をm 文回せば0ではない部分だけ取得できる
- nums1を0~m まで取得して、nums2と合体してsort すればいいのでは

# 答えを見て考えたこと
- 1ms で、ベストな回答ではなさそうだった
- memory も2.40MB で多めに使ってるぽい
- [参考になりそうな回答](https://leetcode.com/problems/merge-sorted-array/solutions/4467339/iterate-backwards-with-two-pointers)

# 回答時間
- 00:7:38
  - 問題読解：3分
  - 方針考える：3分
  - コード作成：1分
