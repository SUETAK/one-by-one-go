#  Find the Index of the First Occurrence in a String

2つの文字列needleとhaystackが与えられたとき、haystackの中でneedleが最初に出現するインデックスを返す。

Example 1:

Input: haystack = "sadbutsad", needle = "sad"
Output: 0
Explanation: "sad" occurs at index 0 and 6.
The first occurrence is at index 0, so we return 0.

## 考えたこと
neddle でループ回して、見つけたら次の文字に行く
フルでループが回ったら終わり
needle の1文字目を持つ
haystack を回す
見つからなかったら-1を返す
見つかったらindex を保存。
index の次の文字からループ開始
neddle の2文字目が次の文字になかったらindexの次の次のところからループ開始
len(haystack)-1 - len(neddle)-1 を超えたら-1返す 

## 答え
strings.Index を使えば良いや
