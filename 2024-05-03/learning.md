# Ransom Note

ransomNoteとmagazineの2つの文字列が与えられたとき、
ransomNoteがmagazineの文字を使って構成できる場合はtrueを、
そうでない場合はfalseを返す。

magazineの各文字はransomNoteで一度しか使用できない。

Example 1:

Input: ransomNote = "a", magazine = "b"
Output: false
Example 2:

Input: ransomNote = "aa", magazine = "ab"
Output: false
Example 3:

Input: ransomNote = "aa", magazine = "aab"
Output: true


## 考えたこと
magazine がransomNote を含めば良さそう
↓
違った。magazine でransomNote が構成できるかだから、
ransomNote の各文字が、magagine に含まれているかを数える


