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
