# container with most water

長さnの整数配列heightが与えられ、
i番目の線の2つの端点が(i, 0)と(i, height[i])となるようにn本の垂直線が引かれる。

x軸と一緒になって容器を形成し，その容器が最も多くの水を含むような2本の直線を求めよ．
容器が貯蔵できる水の最大量を返せ．
容器を斜めにしてはいけないことに注意しなさい。


例
Input: height = [1,8,6,2,5,4,8,3,7]
Output: 49
Explanation: 
The above vertical lines are represented by array [1,8,6,2,5,4,8,3,7]. 
In this case, the max area of water (blue section) 
the container can contain is 49.

# 考え
最大の面積を返す。高さ*横
1番目の線の2つの端点が(1,0)と(1,height[1])=(1,8)となるようにn本の線がひかれる
箱を作らないといけないので、(1,0) が確実に存在している。
あとは引かれた線のうち2つを選んで、面積が最大になるようにする

2つのポインタを使って、最大面積を更新していけば良い
全探索以外の方法。。。

高さは2つのpointer のうち、低い方に合わせる。
(index2-index1)*height[i] = 面積

## 双方向ポインター(two-pointer)というアプローチ
最大化問題に使える




