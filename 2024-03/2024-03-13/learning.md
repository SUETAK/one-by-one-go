# Longest Palindromic Substring

文字列s を受け取って、最も長いPalindromic substring を返却する

Palindromic substring とは、最長の回文文字列である

例)
Input: s = "babad"
Output: "bab"
Explanation: "aba" is also a valid answer.

Input: s = "cbbd"
Output: "bb"

# コードを書く前に考えたこと
- 回文とは、前から読んでも後ろから読んでも同じ文章になる文字列
- map では回文に出てくる文字の個数しかわからない
- 回文の条件は、中心の文字から数えて同じ文字列が続くこと
- 奇数の場合
  - bc d cb
    - d からみてi-1, i+1 が同じ文字
    - i-2, i-2 が同じ文字。
- 偶数の場合
  - abcd dcba
    - i=3,i=4 が同じ文字 => i=len/2 -1, len/2 が同じ文字
    - i=


# 間違って考えてた
文字列の中で回文になっている箇所を探すのが目的
前からみていって、保存している配列の最後から2つ目の値をみて、次の値が同じ文字ならOK

# ギブアップ。答えをみる
ここで、i, jを含む部分文字列が回文であることがわかっていたとしよう。
もしs[i - 1] == s[j + 1]なら、i - 1, j + 1を含む部分文字列も回文でなければならない。

回文がすでにわかっている場合、i=1, j=5 の時、どちらもa で回文。i-1=0: r, j+1=6:r なので、回文である
r aceca r

長さ1の部分文字列はすべて回文であることがわかっている。
このことから、長さ3の部分文字列が回文であるかどうかは、上記の事実を用いて調べることができる。
j - i = 2となるi, jのペアをすべてチェックすればよい。
長さ3の回文がすべてわかったら、その情報を使って長さ5の回文、さらに7の回文、・・・と見つけることができる。

偶数長の回文は？長さ2の部分文字列は、両方の文字が等しければ回文である。
つまり、s[i] == s[i + 1]であれば、i, i + 1は回文である。
このことから、先ほどの論理を使って、長さ4の回文をすべて見つけ、次に長さ6の回文を見つけることができる。


## めも
回文の判定方法が奇数と偶数に別れることはわかった。
奇数の場合、中心の両端が同じ文字である必要があり、偶数の場合中心となる２文字から両端は常に同じ文字
配列の長さの範囲内で偶数、奇数の最長の回文を探せば良い

## 参考にした回答
https://bb-engineer.com/leetcode_100knock_day4/

