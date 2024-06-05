#  Isomorphic Strings

2つの文字列sとtが与えられたとき、それらが同型かどうかを判定せよ。

二つの文字列sとtは、sの文字を置き換えてtを得ることができれば同型である。

ある文字が出現する場合はすべて、文字の順序を保ったまま別の文字に置き換えなければならない。
2つの文字が同じ文字に対応することはないが、1つの文字がそれ自身に対応することはある。

# 考え方
s, t の長さは同じなので、文字がそれぞれ同じ感じで配置されてればおk？

例
3Example 1:

Input: s = "egg", t = "add"
Output: true
Example 2:

1文字の英字が1種類、2文字の英字が1種類

Input: s = "foo", t = "bar"
Output: false
Example 3:

s: 1文字の英字が1種類
t: 1文字の英字が3種類

Input: s = "paper", t = "title"
Output: true

英字の種類をhashmap に保存する
e: 1, g: 2
a: 1, d: 2

length 同じか？
f: 1, o: 2
b: 1, a: 1, r: 1

p: 2, a: 1, e: 1, r: 1
t: 2, i: 1, l: 1, e: 1

個数じゃなくて順番も一緒


