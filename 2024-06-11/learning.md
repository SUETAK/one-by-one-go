# Word pattern

string S とpattern が渡される。sの中に同じpattern があるか確認する

ここでfollowとは完全一致を意味し、パターン中の文字とs中の空でない単語との間に双射が存在することを意味する。


Example 1:

Input: pattern = "abba", s = "dog cat cat dog"
Output: true
Example 2:

Input: pattern = "abba", s = "dog cat cat fish"
Output: false
Example 3:

Input: pattern = "aaaa", s = "dog cat cat dog"
Output: false

1 <= pattern.length <= 300
パターンが小文字の英字のみを含む
1 <= s.length <= 3000
s は小文字の英字とスペース ' ' のみを含む。
sは先頭にも末尾にもスペースを含まない。
s の単語はすべて空白ひとつで区切られています。

## 考え方
例1 ではa=dog, b=cat, b=cat, a=dog のように対応関係ができている
例2 ではa=dog, b=cat, b=cat, a=fish となっており、対応関係が崩れている。

前の問題のように交差するように対応づけを行なったmap を用意して、同じ対応かどうかを確認する？

