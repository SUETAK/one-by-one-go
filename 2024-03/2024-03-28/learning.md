# best time to buy and sell stock 2

prices配列にはN日のprice が入っている
売るか買うかを決めます
いつでも保有できる株式は1株まで。ただし、買ってすぐにその日のうちに売ることはできる。

あなたが達成できる最大の利益を見つけて返してください。

例
prices = [7, 1, 5, 3, 6, 4]
output = 7
price=1 の時に買い、price=5の時に売ると利益は4
その後、price=3の時に買い、6の時に売ると利益は3
合計すると最大の利益は7

# コードを書く前に考えたこと
やることは前回と大体同じ
異なるのは、前回は1回だけ売り買いするだけだったが、今回は何回も行える
同じ日に買って売ることもできる

## わからないので答えをみる

## 考え方
2つの性質
- 利益は常にインクリメントされる
- 昨日の価格よりも下がれば利益は最低0になるので、それまでの価格を全て捨てることができる

## アプローチ
以上のことから、本日の価格と昨日の価格の増減計算を維持することができます。
- 価格が上昇し続ければ、販売価格を更新します。
  - 昨日の価格が2, 今日の価格が3であれば、売らない。利益は増えている
- 価格が下がったらすぐに利益を計算し、買いポイントと売りポイントをリセットします。
  - 昨日の価格が2, 今日の価格が1であれば、損をしている可能性がある

たとえ、株価が下がったとしても、その後より上昇する可能性があるが、待つよりも一度利益を確定して、もう一度買い直した方が利益は大きくなる
例えばprices = [1,2,0, 10] という場合、
1->2 で利益は上昇し、2->0 で下降している。その次に0->10 と上昇している。
この時、2->10=8 よりも0->10=10 とした方が利益は多い。
売った当日に買うことができるため、下降したタイミングで利益を取得した方が良い。
この場合、
- 1->2=1, 0->10=10 => 11
- 1->2->0->10=9 => 9

