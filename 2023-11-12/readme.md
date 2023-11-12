# 最小公倍数と最大公約数を使った問題

## between 2 sets

[問題URL](https://www.hackerrank.com/challenges/between-two-sets/problem?isFullScreen=true)

### 内容
2つの配列が与えられる

a = [2, 4]
b = [16, 32, 96]

a の最小公倍数Xとb の最大公約数Yを求め、X〜Yの間でXとY両方の倍数である数字の個数を求める。


### ポイント

- a に含まれる数字全てに共通する最小公倍数を求める
- b に含まれる数字全てに共通する最大公約数を求める
- 最小公倍数X〜最大公約数Yまでの間にある数字を、Xの倍数ごとに条件を満たすか確認する


