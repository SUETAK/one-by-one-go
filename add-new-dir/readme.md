ファイル、ディレクトリの作成

実行当日のディレクトリと、main.go, go.mod を作成する

## 学び
- ファイルの作成は、os.Mkdir で実行できる
- os.MkdirAll を使用すると親ディレクトリから指定できる
- ダブルクォーテーションで囲むときはinterpreted string literal(解釈される文字列リテラル)として解釈される
    - \n が改行と解釈される
- バッククォート（`）で囲まれた時は生文字列リテラルとして解釈される
    - 改行などが解釈されない
