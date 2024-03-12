# Median of Two Sorted Arrays

Given two sorted arrays nums1 and nums2 of size m and n respectively, return the median of the two sorted arrays.
The overall run time complexity should be O(log (m+n)).

size がmのnum1, sizeがn のnum2 のソートされた配列がそれぞれわたされる。2つのソートされた配列の中央値を返却しなさい
計算量はO(log(m+n))以下である必要がある

例)
Input: nums1 = [1,3], nums2 = [2]
Output: 2.00000
Explanation: merged array = [1,2,3] and median is 2.

Input: nums1 = [1,2], nums2 = [3,4]
Output: 2.50000
Explanation: merged array = [1,2,3,4] and median is (2 + 3) / 2 = 2.5.


# コードを書く前に考えたこと
- 中央値の計算方法は
  - 配列の長さが奇数の場合、真ん中のindex に紐づく値
  - 偶数の場合は中央に近い2つの値の平均値を計算した値
- 計算量O(log(m+n)) の目安は2分探索2つ分くらいのイメージ？
- 配列結合して、ソートして、偶数の場合と奇数の場合で分岐
  - 奇数の場合の方がルールベースで判断できるので、偶数の場合がキモ？
- 合成した後にソートは遅そう
  - m とn の大きさを比べて、大きい方を基準にループを回す？

## 思考
- 合成した配列の長さが
  - 4 の時、2, 3番目
  - 6 の時、3, 4番目
  - 8 の時、4, 5番目
- (m+n)/2, (m+n)/2 + 1 番目 を足して２で割れば良い
