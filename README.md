go-atcoder
---

AtCoderの解答を管理するリポジトリ


# ディレクトリ構成

```
archive/YYYY/MM/DD/<contest>/<problem>/   # 日付ごとの解答 (復習で解き直すたびに新しい日付へ)
│   ├── main.go                           # 自己完結した package main (提出物そのもの)
│   └── tests/                            # oj download が取得したサンプル
snippets/<topic>/                         # コピペ元のテスト済みコード (import はしない)
template/main.go                          # task new が複製する雛形
```



# コマンド

```
# 今日の日付のディレクトリを作成し, そのパスをコピー
task today

# コンテストのディレクトリを作成
task new CONTEST=abcXXX

# テストデータで解答をテスト
```



