# addframe
[日本語 README](README.jp.md)

## what is this?
  Reduces the size of a photo taken with a digital camera and converts it into a photo with a frame added.  
  The copyright, maker, camera, lens, and shooting data obtained from Exif are printed at the bottom of the frame.  
<img src="https://raw.githubusercontent.com/mitsugu/addframe/main/images/IMGP6286.JPG" width="512" height="389" alt="example of addframe's result" title="example of addframe's result">

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
Path to ImageMagick's magick command. Absolute and relative paths can be specified.  
Windows users, don't forget to include the .exe extension.
2. **path.exiftool**  
Path to exiftool command. Absolute and relative paths can be specified.  
Windows users, don't forget to include the .exe extension.
3. **length**  
Specify the size of the long side of the image in pixels.
4. **frame.top**  
Specifies the width of the top edge of the frame in pixels.
5. **frame.left**  
Specifies the width of the left side of the frame in pixels.
6. **frame.right**  
Specifies the width of the right side of the frame in pixels.
7. **frame.bottom**  
Specify the width of the bottom edge of the frame in pixels.  
In the latest version of addframe, it is automatically calculated and ignored by the program.
8. **frame.color**  
Specify the frame color using [ImageMagick method](https://imagemagick.org/script/color.php "ImageMagick - Color Name").
9. **text.direction**  
Specify where to place the text. Specify using [ImageMagick method](https://imagemagick.org/script/command-line-options.php#gravity "See gravity type").
10. **text.margin**  
Specifies the margin of the text area in pixels. 0 is recommended if text.direction is South. For SouthEast, specify the right margin in pixels. For SouthWest, specify the left margin in pixels.
11. **text.dpi**  
Specify the display DPI. If ImageMagick's DPI and the display's DPI are different, the font size will be abnormal, so **be sure to** specify it.
12. **text.element**  
This is a format specification for three lines, so **be sure to** specify all of them.
13. **text.element.font**  
Specify the path to the font file. You can specify ttc, ttf, and otf font files.
14. **text.element.color**  
Specify the font color using [ImageMagick method](https://imagemagick.org/script/color.php "ImageMagick - Color Name").
15. **text.element.size**  
Specify the font size.
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
1. --config (-c)  
Optional. If omitted, the configuration file addframe.json in the current directory will be used.  
Even if you omit it, **the configuration file is not unnecessary**.
2. --input (-i)  
The image file to which the frame will be added. Usually a JPEG file. The following data must be written to Exif.
* Copyright
* Author
* Camera Maker
* Camera Model
* Lens Model
* ISO Speed
* F Number
* Shutter Speed
3. --output (-o)  
Path to the resulting file after adding frames.

**NOTE)** Many cameras do not record Copyright and Author in Exif. If you have such a camera, you can use a program I wrote called [addcopyright](https://github.com/mitsugu/addcopyright).

### development repository
[https://github.com/mitsugu/addframe](https://github.com/mitsugu/addframe)

### License
[Apache License 2.0](./LICENSE.en.md)


## comment
enjoy!! 😀

[![Go Reference](https://pkg.go.dev/badge/github.com/mitsugu/addframe.svg)](https://pkg.go.dev/github.com/mitsugu/addframe)
