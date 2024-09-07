# addframe
[English README](README.md)

## これは何？
　デジタルカメラで撮影した写真を縮小し、フレームを付けた写真に変換します。  
　フレーム下部に Exif から取得した著作権者名、メーカー、カメラ、レンズ、撮影データがプリントされます。
<img src="https://raw.githubusercontent.com/mitsugu/addframe/main/images/IMGP6286.JPG" width="512" height="389" alt="example of addframe's result" title="example of addframe's result">

## 必要な他のソフトウェア
  * [ExifTool](https://exiftool.org/)
  * [ImageMagick](https://imagemagick.org/)

## セットアップ
### addframe のインストール
#### go install によるインストール
```
go install github.com/mitsugu/addframe@<tag name>
```
#### リリースされている zip ファイルによるインストール
1. [最新のリリース zip ファイル](https://github.com/mitsugu/addframe/releases) をダウンロードする。
2. zip ファイルを展開する。
3. Linux 版、MacOS 版は addframe にリネームする。
4. 実行ファイルを任意のディレクトリに置く。
5. addframe が置かれたディレクトリへのパスを通します。あるいは addframe 実行可能ファイルをパスが通ったディレクトリに置きます。

### addframe.json を編集する
```
// addframe.json example
{
	"path": {
		"imagemagick": "/path/to//magick",
		"exiftool": "/bin/exiftool"
	},
	"length": 1280,
	"frame": {
		"top": 32,
		"left": 32,
		"right": 32,
		"bottom": 128,
		"color": "#3f3f3f"
	},
	"text": {
		"direction": "South",
		"margin": 0,
		"dpi": 96,
		"element": [{
			"font": "/usr/share/fonts/opentype/noto/NotoSansCJK-Regular.ttc",
			"color": "white",
			"size": 24,
			"margintop": 32,
			"marginbottom": 4
		}, {
			"font": "/usr/share/fonts/opentype/noto/NotoSansCJK-Regular.ttc",
			"color": "white",
			"size": 24,
			"margintop": 4,
			"marginbottom": 4
		}, {
			"font": "/usr/share/fonts/opentype/noto/NotoSansCJK-Regular.ttc",
			"color": "#cfcfcf",
			"size": 16,
			"margintop": 4,
			"marginbottom": 4
		}]
	}
}
```
1. **path.imagemagick**  
ImageMagick の magick コマンドへのパス。絶対パスと相対パスを指定できます。  
Windows ユーザーの場合は、.exe 拡張子を忘れずに含めてください。
2. **path.exiftool**  
exiftool へのパス。絶対パスと相対パスを指定できます。  
Windows ユーザーの場合は、.exe 拡張子を忘れずに含めてください。
3. **length**  
画像の長辺のサイズをピクセル単位で指定します。
4. **frame.top**  
フレーム上端の幅をピクセル単位で指定します。
5. **frame.left**  
フレームの左側の幅をピクセル単位で指定します。
6. **frame.right**  
フレームの右側の幅をピクセル単位で指定します。
7. **frame.bottom**  
フレーム下端の幅をピクセル単位で指定します。  
addframe の最新バージョンでは、プログラムによって自動的に計算され、無視されます。
8. **frame.color**  
[ImageMagick 方式](https://imagemagick.org/script/color.php "ImageMagick - Color Name : 英語")を使用して枠の色を指定します。
9. **text.direction**  
テキストを配置する場所を指定します。[ImageMagick 方式](https://imagemagick.org/script/command-line-options.php#gravity "See gravity type : 英語")を使用して指定します。
10. **text.margin**  
テキスト領域のマージンをピクセル単位で指定します。  
text.direction が South の場合は 0 を推奨します。  
SouthEast の場合、右マージンをピクセル単位で指定します。  
SouthWest の場合、左マージンをピクセル単位で指定します。
11. **text.dpi**  
ディスプレイのDPIを指定します。 ImageMagickのDPIとディスプレイのDPIが異なるとフォントサイズが異常になりますので必ず指定してください。
12. **text.element**  
テキスト 3 行分の書式指定ですので、**必ずすべて** 指定してください。
13. **text.element.font**  
フォントファイルへのパスを指定します。 ttc、ttf、otf フォント ファイルを指定できます。
14. **text.element.color**  
[ImageMagick 形式](https://imagemagick.org/script/color.php "ImageMagick - Color Name : 英語")のフォントカラーを指定します。
15. **text.element.size**  
フォントのサイズを**ポイント数**で指定します。**ピクセル数ではない**ので注意してください。
16. **text.element.margintop**  
Specifies the top margin of a line of text.  
The display position of the line is mainly adjusted using margintop and marginbottom.
17. **text.element.marginbottom**  
Specifies the bottom margin of a line of text.  
The display position of the line is mainly adjusted using margintop and marginbottom.

### usage
```
addframe [--config congiguration_file_path] <--input input_file_path> <--output output_file_path>
or
addframe [-c congiguration_file_path] <-i input_file_path> <-o output_file_path>
```
1. **--config (-c)**  
オプション。省略した場合は、**カレントディレクトリにある設定ファイル addframe.json が使用されます**。  
**省略した場合でも設定ファイル自体が不要になるわけではありません**のでご注意ください。
2. **--input (-i)**  
フレームを追加する画像ファイル。通常は JPEG です。Exif には以下のデータが記録されている必要があります。
* Copyright
* Author
* Camera Maker
* Camera Model
* Lens Model
* ISO Speed
* F Number
* Shutter Speed
3. **--output (-o)**  
フレームを追加した結果のファイルへのパス。

**注)**  
多くのカメラは Exif に著作権者と著作者を記録しません。そのようなカメラをご利用の場合は私が作成した [addcopyright](https://github.com/mitsugu/addcopyright) というプログラムを使うことで画像ファイルの Exif に著作権者と著作者の情報を書き込むことができます。

### 開発リポジトリ
[https://github.com/mitsugu/addframe](https://github.com/mitsugu/addframe)

### ライセンス
[Apache License 2.0](./LICENSE.ja.md)


## 最後に
楽しんで！！ 😀

[![Go Reference](https://pkg.go.dev/badge/github.com/mitsugu/addframe.svg)](https://pkg.go.dev/github.com/mitsugu/addframe)
