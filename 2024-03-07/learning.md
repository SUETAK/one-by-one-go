You are given two non-empty linked lists representing two non-negative integers. 
The digits are stored in reverse order, and each of their nodes contains a single digit. 
Add the two numbers and return the sum as a linked list.

You may assume the two numbers do not contain any leading zero, except the number 0 itself.

あなたは2つの空ではない2つの正のinteger を持つlinked list を与えられます
数字は逆順で保存され、それぞれのノードは2つの数字を持ちます
2つの数字を足して合計をliked list として返却してください

あなたは2つの数字は00123 の00 のようなleading zero は含まれず、0だけが存在します。

例）
2->4->3
5->6->4
----
7->0->8

Input: l1 = [2,4,3], l2 = [5,6,4]
Output: [7,0,8]
Explanation: 342 + 465 = 807.

# コードを書く前の考え
liked list とは、1つのstruct の中に、value と次のnode 情報が入っているもの
output では、liked list を逆順で追って行って、数字を作成し、足し、liked list に入れ直せば良さそう

- そのまま数字を変換して足し算するのは効率が悪そう
- 桁が変わることが考えられる
- for 文で最奥まで掘ってstring で文字列作成して数字に直す？
- linked list になっているのであれば、その特性を利用できるはず

# 人のsolution をみながら考える
- 逆順で与えられているので、筆算をロジックに落とせば良い
- linked list の同じ桁の数字を足して、10で割った時のあまりを繰り上がり（carry） として保存して、次の桁に渡す

# 解いた後に考える
linked list の連鎖はfor 文で次の値がnil になるまで見る必要がある
最終的な値を更新しながら作る場合、都度tmp に保存しつつ次の値に行く必要がある
途中の値の更新とかめっちゃ面倒くさそう
