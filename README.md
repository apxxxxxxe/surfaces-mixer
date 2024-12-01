# surfaces-mixer
yamlからsurfaces.txtを生成するCLIアプリ

## 何をするもの？
yaml形式で定義されたサーフェスをsurfaces.txt形式に変換します。

## どんな特徴があるの？
#### 各パーツの全組み合わせを網羅したsurfaces.txtを作れる
例として、次のパーツで構成されたサーフェスを考えます。
- 目パーツ5つ(目線こっち、目線あっち、半目こっち、半目あっち、閉じ目)
- 眉パーツ4つ(通常、驚き、悲しみ、怒り)
- 口パーツ3つ(笑い、驚き、むっ)
- 腕パーツ3つ(通常、手を振る、ピースサイン)
- その他の頭、胴体など、ベースとなるパーツ1つずつ

このとき、パーツの全組み合わせは 5×4×3×3×1= 180通りです。  
人力では途方もない数ですが、surfaces-mixerはこれを（ほぼ）一瞬でsurfaces.txtに記述することができます。  
また後述しますが、その膨大なサーフェスの各内容（例:笑って手を振るサーフェスは1112番、など）を把握しやすいようなサーフェス番号の割り振りも同時に行います。

## これを使う利点は？
以下の３点です。
#### 各パーツの全組み合わせを網羅的に生成できる
前述の通り、各パーツの数が増えるとその組み合わせは幾何級数的に膨れ上がり、これを書き下すのは至難の業です。  
surfaces-mixerはこの作業を自動化し、不必要な労力と記述ミスの可能性を無くします。

#### サーフェス番号を自動的に、かつ理解しやすい形で定義できる
これも前述したことですが、全組み合わせを定義する上での問題は、冗長かつ分かりづらいことです。  
仮に数十、数百のサーフェス番号を無秩序に割り振ったとして、その中から目的のサーフェスを探し出すのはあまりにも非効率的です。  
surfaces-mixerはこれを解決するため、ゴースト製作者にも分かりやすい、ユニークな方式で番号の割り振りを行います。

#### パーツの追加・削除に柔軟に対応できる
膨大かつ複雑なサーフェス定義において、パーツの追加・削除は容易ではありません。  
1つのパーツ定義を追加するだけでも影響が数十～数百行におよび、その範囲の書き直しを余儀なくされます。  
surfaces-mixerの場合は単純です。単にyamlファイルに定義をちょこっと追加し、surfaces-mixerを実行してsurfaces.txtを生成するだけです。  

## yamlファイルはどんなふうに書けばいいの？
[サンプルファイル](https://github.com/apxxxxxxe/surfaces-mixer/blob/main/sample.yaml)をご確認ください。  
surfaces.txtに変換した結果は[こちら](https://github.com/apxxxxxxe/surfaces-mixer/blob/main/surfaces_sample.txt)です。  
ちなみに、本ファイルは制作中の自作ゴースト[Crave The Grave](https://github.com/apxxxxxxe/Haine)で実際に使われているものです。

## その他
### そんなにいっぱいサーフェスを用意して、動作は重くならないの？
- SSPにおいて、サーフェス定義の数が多くとも、動作パフォーマンスにほぼ影響しないようです（[Reference](https://twitter.com/ponapalt/status/1587364571224092673)）

## ダウンロード
[![surfaces-mixer.exe](https://img.shields.io/github/v/release/apxxxxxxe/surfaces-mixer?color=%2367a5cf&label=surfaces-mixer.exe&logo=github)](https://github.com/apxxxxxxe/surfaces-mixer/releases/latest/download/surfaces-mixer.exe) 

