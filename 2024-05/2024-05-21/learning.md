# Longest Substring Without Repeating Characters

文字列s が与えられる。最大の部分文字列を出せ。
ただし、繰り返しがあるものは除く

## 考え方
sliding window を使う。
最大の部分文字列が窓枠になる。
繰り返しの定義は、前の文字と同じ文字であること

右枠はsの長さまで
求めるの最大の部分文字列maxLength
マイループ、繰り返しかどうか判定する。
問題意図としては、同じ文字が入っていてはいけないらしいので、判断はこの文字列に同じ文字があるかどうかを判断する



|a|bcabcbb
|ab|cabcbb
|abc|abcbb → abc が後ろの配列の中に存在する
|abca|bcbb → a が中で重複するのでleft を進める
a|bca|bcbb 
a|bcab|cbb → b が中で重複するのでleft を進める
ab|cab|cbb → b が中で重複するのでleft を進める
ab|cabc|bb → right を進める


