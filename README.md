# addframe
## what is this?
  Reduces the size of a photo taken with a digital camera and converts it into a photo with a frame added.  
  The copyright, maker, camera, lens, and shooting data obtained from Exif are printed at the bottom of the frame.

## Requirements
  * [ExifTool](https://exiftool.org/)
  * [ImageMagick](https://imagemagick.org/)

## Setup
### install addcopyright
```
go install github.com/mitsugu/addframe@<tag name>
```
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
  Note: Place addframe.json in **the current directory** where you run addcopyright.

### usage
```
addframe --input <input file path> --output <output file path>
```
### License
Apache License 2.0


## comment
enjoy!! 😀
