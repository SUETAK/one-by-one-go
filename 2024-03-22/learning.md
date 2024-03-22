# Majority Element

size がNのnums が与えられる。majority element を返却してください
majority element とは、nums の中で過半数を閉めるelement のことです。
このmajority element は常に存在すると仮定します

## コードを書く前に考えたこと:8min
- 過半数とはどんなことか？
  - size の1/2 以上登場する値
  - 1/2 以上登場した段階でカウントは止めていい
  - 奇数と偶数のsize で何か条件が変わる？=> 変わらない
    - 奇数
      - 1,3,5,7,9
        - golang で3/2 = 1.5 = 1, 5/2 = 2.5 = 2
        - x/2 + 1 が過半数基準
    - 偶数 
      - 2, 4, 6, 8
        - 4/2 = 2 => 3 が過半数
        - x/2 +1 が過半数基準
- 処理(3min)
  - nums のループ
    - map 作成
    - map に入れるO(N)
    - map に入れた値のうち、size/2 +1 になったらその値を返却

## コードの結果 31min
- 21ms, 6.95MB
