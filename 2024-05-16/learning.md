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




