# surfaces-mixer
YAMLからsurfaces.txtを生成するCLIアプリケーション

## 機能概要
YAML形式で定義されたサーフェスをsurfaces.txt形式に変換します。

## 特徴
### 各パーツの全組み合わせを網羅したsurfaces.txtを生成
以下のようなパーツで構成されたサーフェスを例に考えます。
- 目パーツ5種類（目線こっち、目線あっち、半目こっち、半目あっち、閉じ目）
- 眉パーツ4種類（通常、驚き、悲しみ、怒り）
- 口パーツ3種類（笑い、驚き、むっ）
- 腕パーツ3種類（通常、手を振る、ピースサイン）
- その他、頭や胴体などのベースパーツ1種類ずつ

この場合、パーツの全組み合わせは 5×4×3×3×1 = 180通りとなります。人力では膨大な作業となりますが、surfaces-mixerはこれを瞬時にsurfaces.txtに記述します。また、後述の通り、その膨大なサーフェスの各内容を把握しやすいように、サーフェス番号の割り振りも自動で行います。

## 利点
以下の3点が挙げられます。
### 各パーツの全組み合わせを網羅的に生成
前述の通り、各パーツの数が増えるとその組み合わせは指数関数的に増加し、手動での記述は非常に困難です。surfaces-mixerはこの作業を自動化し、労力を削減するとともに記述ミスのリスクを排除します。

### サーフェス番号を自動かつ理解しやすい形で定義
全組み合わせを定義する際、冗長で分かりづらい番号の割り振りが問題となります。数十、数百のサーフェス番号を無秩序に割り振るのは非効率です。surfaces-mixerはこれを解決し、ゴースト製作者にとって理解しやすい独自の方式で番号を割り振ります。

### パーツの追加・削除に柔軟に対応
膨大かつ複雑なサーフェス定義において、パーツの追加や削除は容易ではありません。1つのパーツを追加するだけでも影響範囲が広がり、多くの行を書き直す必要があります。surfaces-mixerなら、YAMLファイルに定義を追加し、再実行するだけでsurfaces.txtを生成できます。

## YAMLファイルの記述方法
[サンプルファイル](https://github.com/apxxxxxxe/surfaces-mixer/blob/main/sample.yaml)をご参照ください。変換結果のsurfaces.txtは[こちら](https://github.com/apxxxxxxe/surfaces-mixer/blob/main/surfaces_sample.txt)です。本ファイルは制作中の自作ゴースト[Crave The Grave](https://github.com/apxxxxxxe/Haine)で実際に使用されています。

## FAQ
### 多数のサーフェスを用意すると動作は重くならないのか？
- SSPにおいて、サーフェス定義の数が多くとも動作パフォーマンスにほとんど影響しません（[Reference](https://twitter.com/ponapalt/status/1587364571224092673)）。

## ダウンロード
[![surfaces-mixer.exe](https://img.shields.io/github/v/release/apxxxxxxe/surfaces-mixer?color=%2367a5cf&label=surfaces-mixer.exe&logo=github)](https://github.com/apxxxxxxe/surfaces-mixer/releases/latest/download/surfaces-mixer.exe)
