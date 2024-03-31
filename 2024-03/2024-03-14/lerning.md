# ZigZag Conversion

文字列 "PAYPALISHIRING "は、次のように指定された行数にジグザグのパターンで書かれている
文字列を受け取り、行数を指定してこの変換を行うコードを書きなさい：
P   A   H   N
A P L S I I G
Y   I   R


Input: s = "PAYPALISHIRING", numRows = 3
Output: "PAHNAPLSIIGYIR"


Input: s = "PAYPALISHIRING", numRows = 4
Output: "PINALSIGYAHRPI"
Explanation:
     0 1 2 3 4 5 6 
row1 P     I     N
row2 A   L S   I G
row3 Y A   H R
row4 P     I

row0 ->P~I: 6 I~N: 6               -> 4*2-2=6 6を基準のstep とする
row1 ->A~L: 4 L~S: 2 S~I: 4 I~G: 2 -> 2*
row2 ->Y~A: 2 A~H: 4 H~R: 2 R~?: 0
row3 ->P~I: 6 I~?: 0

1,6-1

Input: s = "PAYPALISHIRING", numRows = 5
Output: "PINALSIGYAHRPI"
Explanation:
     0 1 2 3 4
row0 P       H
row1 A     S I
row2 Y   I   R
row3 P L     I G
row4 A       N

row0 ->P~H: 8
row1 ->A~S: 6
row2 ->Y~I: 4
row3 ->P~L: 2 
row4 ->A~N: 8

# コードを書く前に考えたこと 5分
- ジグザグパターンってどういうものなんだろう
- row が4の時はrow に４文字、次の列の3行目に1文字、さらに次の列の2行目に1文字、1行目に4文字。が連続する
- output は横に並んだ文字を1つの文字列にして返す
- ２次元配列で表現して、該当する文字を抜き出す？
- ルールを記載して、頑張る？
  - row文を下に降ってからrow-1 上るので
    - 0,row+row-1,(row+row-1)*2,3,,,
    - 1,row+row-2,(row+row-2)*2,3,,,
    - 
# わからんので答えみる
1つ1つたどっていって、ルールをコードに落とし込んでいくことで実装できる。
最初の行と最後の行のstep数は、間の行のstep 数を計算する時に使用できることに気付きたかった。

index のイメージが具体的になればより安定して解くことができそう。
他の回答では、登ったり降ったりをコードに落とし込んでいた。

頭の中だけの理解では混乱するので、紙に書くなどの工夫が必要になりそう
今回は5分問題理解、5分解法を考える、わからなかったら答えをみる。という方法でやった。
答えを見てもパッとわからなかったけど、1つ1つなぞったらわかったので、今後に期待。
