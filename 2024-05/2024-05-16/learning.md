# minimum size subarray sum

nums に正の値が入っている
正の値target がある
target の値以上の最小の長さのsubarray の長さを取得して返却する
subarray がないなら0を返す

## 考え方
subarray だから、順番を変えずに、地続きの状態である必要がある
最小のsubarray の長さが必要なので、1が最小。
最小を更新していくイメージ
左から進めて、target を超えたらその一からスタート


[2,3,1,2,4,3], target=7 なら
i=0の2 からスタート
2+3=5, 2+3+1=6, 2+3+1+2=8 なのでi=3までleftを動かす
leftをright に引き寄せつつ、引き算していく。left=right になったら同じようにright を進める

## 考え方2。復習
sliding window の問題。
窓の右側を最大幅までみる
ウィンドウのサイズの判定条件
窓の右側を増やす→部分配列の和がtarget 以上になる
窓の左側を増やす→条件を満たす

sum がtarget を上回ったら、条件に応じてminLength の長さを更新する。
今の長さはtarget を上回っているので、窓枠を大きくしても絶対にsum がtarget を上回るのは変わらない
そのため、このループの中で左を進めてsumの合計を少なくする。
仮にleft で合計が少なくなりすぎた（次の右側でsumがtarget を上回らない）場合でも、leftを増やすのはsumがtarget を上回っている間だけ



