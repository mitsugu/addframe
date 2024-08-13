# addframe
[日本語](README.jp.md)

## what is this?
  Reduces the size of a photo taken with a digital camera and converts it into a photo with a frame added.  
  The copyright, maker, camera, lens, and shooting data obtained from Exif are printed at the bottom of the frame.

## Requirements
  * [ExifTool](https://exiftool.org/)
  * [ImageMagick](https://imagemagick.org/)

## Setup
### install addframe
#### use go install
```
go install github.com/mitsugu/addframe@<tag name>
```
#### download release zip file
1. download [release zip file](https://github.com/mitsugu/addframe/releases)
2. extract the zip file
3. Please put it in any directory.
4. Pass the path to the addframe directory. Or put the addframe executable file in a directory that is in your path.

### edit addframe.json
```
// addframe.json example
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
1. Specify the path to the ttf or otf font file for font.  This **must** also be specified according to the user's environment.
2. For imagemagick, specify the path to ImageMagick's magick command. This **must** also be specified according to the user's environment.
3. exiftool specifies the path to the exiftool command. This **must** also be specified according to the user's environment.
4. fontColor **may** specify the font color desired by the user. It can be specified using the same method as cascading style sheets. See [ImageMagick Color Names](https://imagemagick.org/script/color.php) for details.
5. frameColor **may** be specified as the user's desired frame color. It can be specified using the same method as cascading style sheets. See [ImageMagick Color Names](https://imagemagick.org/script/color.php) for details. The trick is to choose a color that makes the font stand out and brings out the image.

  Note: Place addframe.json in **the current directory** where you run addframe.

### usage
```
addframe --input <input file path> --output <output file path>
or
addframe --i <input file path> --o <output file path>
```
### License
[Apache License 2.0](./LICENSE.en.md)


## comment
enjoy!! 😀
