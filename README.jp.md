# addframe
[English](README.md)

## これは何なのか？
  デジタルスチールカメラで撮影した写真のサイズを縮小し、フレームを追加した写真に変換します。  
  フレーム下部にはExifから取得した著作権、メーカー、カメラ、レンズ、撮影データがプリントされます。

## 他に必要なもの
  * [ExifTool](https://exiftool.org/)
  * [ImageMagick](https://imagemagick.org/)

## セットアップ
### addframe のインストール
#### go install を使う方法
```
go install github.com/mitsugu/addframe@<tag name>
```
#### リリース zip ファイルを使う方法
1. [リリース zip ファイル](https://github.com/mitsugu/addframe/releases) をダウンロードしてください。
2. zip ファイルを展開してください。
3. 任意のディレクトリに置いてください。
4. addframe のディレクトリにパスを通してください。またはは addframe の実行可能ファイルをパスが通っているディレクトリに置いてください。

### addframe.json を編集する
```
// addframe.json の例
{
	"length": 1280,
	"top": 32,
	"left": 32,
	"right": 32,
	"bottom": 128,
	"font": "path/to/font.ttf",
	"fontColor": "white",
	"frameColor": "#3f3f3f",
	"imagemagick": "path/to/magick",
	"exiftool": "path/to/exiftool"
}
```
1. font には ttf または otf フォント ファイルへのパスを指定します。ユーザーの環境に合わせて**指定しなければなりません**。
2. imagemagick には、ImageMagick の magick コマンドへのパスを指定します。ユーザーの環境に合わせて**指定しなければなりません**。
3. exiftool は、exiftool コマンドへのパスを指定します。ユーザーの環境に合わせて**指定しなければなりません**。
4. fontColor にユーザーが希望するフォントの色を**指定することができます**。カスケードスタイルシートと同様の方法で指定できます。詳細は [ImageMagick Color Names](https://imagemagick.org/script/color.php) を参照してください。
5. FrameColor は、ユーザーの希望するフレームの色を**指定することができます**。カスケードスタイルシートと同様の方法で指定できます。詳細は[ImageMagick Color Names](https://imagemagick.org/script/color.php) を参照してください。フォントが目立たせ画像を引き立たせる色を選ぶのがコツです。

  注: addframe.json を、addframe を実行するカレントディレクトリに配置してください。

### 使い方
```
addframe --input <input file path> --output <output file path>
or
addframe --i <input file path> --o <output file path>
```
### ライセンス
[Apache License 2.0](./LICENSE.ja.md)


## その他
楽しんでね！！ 😀
