package main

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
	"os/exec"
	"path/filepath"
	"strconv"
	"strings"
	"text/template"

	"github.com/urfave/cli/v2"
)

type Config struct {
	Path struct {
		Imagemagick string `json:"imagemagick"`
		Exiftool    string `json:"exiftool"`
	} `json:"path"`
	Length int `json:"length"`
	Frame  struct {
		Top    int    `json:"top"`
		Left   int    `json:"left"`
		Right  int    `json:"right"`
		Bottom int    `json:"bottom"`
		Color  string `json:"color"`
	} `json:"frame"`
	Text struct {
		Direction string `json:"direction"`
		Margin    int    `json:"margin"`
		Dpi       int    `json:"dpi"`
		Element   []struct {
			Font         string `json:"font"`
			Color        string `json:"color"`
			Size         int    `json:"size"`
			Margintop    int    `json:"margintop"`
			Marginbottom int    `json:"marginbottom"`
		} `json:"element"`
	} `json:"text"`
}

func (c *Config) loadConfig(configPath string) error {
	file, err := os.Open(configPath)
	if err != nil {
		return fmt.Errorf("failed to open config file '%s': %w", configPath, err)
	}
	defer file.Close()

	decoder := json.NewDecoder(file)
	if err := decoder.Decode(c); err != nil {
		return fmt.Errorf("failed to decode config file '%s': %w", configPath, err)
	}

	if c.Length <= 0 {
		return fmt.Errorf("invalid configuration: Length must be greater than 0")
	}
	if c.Path.Imagemagick == "" {
		return fmt.Errorf("invalid configuration: Imagemagick path must be specified")
	}
	if c.Path.Exiftool == "" {
		return fmt.Errorf("invalid configuration: Exiftool path must be specified")
	}

	return nil
}

var config Config

type ExifData struct {
	Lens         string
	Fnumber      string
	Shutterspeed string
	ISO          string
	Orientation  int
	Author       string
	Copyright    string
	Make         string
	Model        string
	LensID       string
	LensModel    string
	Width        int
	Height       int
}

func (e *ExifData) load(inputPath string) error {
	cmd := exec.Command(config.Path.Exiftool,
		"-Lens", "-FNumber", "-ShutterSpeed", "-ISO", "-Orientation",
		"-Author", "-Copyright", "-Make", "-Model", "-LensID", "-LensModel",
		"-ImageWidth", "-ImageHeight", "-s", "-T", inputPath)
	output, err := cmd.Output()
	if err != nil {
		return err
	}
	outputStr := strings.TrimSpace(string(output))
	lines := strings.Split(outputStr, "\t")

	e.Lens = strings.TrimSpace(lines[0])
	e.Fnumber = strings.TrimSpace(lines[1])
	e.Shutterspeed = strings.TrimSpace(lines[2])
	e.ISO = strings.TrimSpace(lines[3])
	tmpOrientation, err := getOrientation(lines[4])
	if err != nil {
		return err
	}
	e.Orientation, _ = strconv.Atoi(strings.TrimSpace(tmpOrientation))
	e.Author = strings.TrimSpace(lines[5])
	e.Copyright = strings.TrimSpace(lines[6])
	e.Make = strings.TrimSpace(lines[7])
	e.Model = strings.TrimSpace(lines[8])
	e.LensID = strings.TrimSpace(lines[9])
	e.LensModel = strings.TrimSpace(lines[10])
	e.Width, _ = strconv.Atoi(strings.TrimSpace(lines[11]))
	e.Height, _ = strconv.Atoi(strings.TrimSpace(lines[12]))

	if strings.Contains(e.Model, "iPhone") && (e.Orientation == 90 || e.Orientation == 270) {
		e.Width, e.Height = e.Height, e.Width
		e.Orientation = 0
	}

	return nil
}

var exifdata ExifData

func getOrientation(data string) (string, error) {
	orientationStr := strings.TrimSpace(data)
	rotationAngle := "0"

	if orientationStr != "" && orientationStr != "-" {
		switch {
		case strings.Contains(orientationStr, "Rotate 90 CW"):
			rotationAngle = "90"
		case strings.Contains(orientationStr, "Rotate 180 CW"):
			rotationAngle = "180"
		case strings.Contains(orientationStr, "Rotate 270 CW"):
			rotationAngle = "270"
		case strings.Contains(orientationStr, "Rotate 90 CCW"):
			rotationAngle = "270"
		case strings.Contains(orientationStr, "Rotate 270 CCW"):
			rotationAngle = "90"
		default:
			rotationAngle = "0"
		}
	}

	return rotationAngle, nil
}

func createFrame(inputPath string, outputPath string) error {
	var width int
	var height int
	var rotationAngle int
	if exifdata.Width > exifdata.Height {
		width = config.Length
		height = (exifdata.Height * config.Length) / exifdata.Width
		rotationAngle = exifdata.Orientation
	} else {
		width = (exifdata.Width * config.Length) / exifdata.Height
		height = config.Length
		rotationAngle = exifdata.Orientation
	}

	switch rotationAngle {
	case 90, 270:
		width, height = height, width
	}

	w := width + config.Frame.Left + config.Frame.Right
	h := height + config.Frame.Top
	for i := 0 ; i < 3 ; i++ {
		h += config.Text.Element[i].Margintop
		h += config.Text.Element[i].Size * config.Text.Dpi / 72
		h += config.Text.Element[i].Marginbottom
	}

	cmdArgs := []string{
		"-size", fmt.Sprintf("%dx%d", w, h),
		"xc:" + config.Frame.Color,
	}
	lens := ""
	if exifdata.Lens != "" && exifdata.Lens != "-" {
		lens = exifdata.Lens
	} else if exifdata.LensID != "" && exifdata.LensID != "-" {
		lens = exifdata.LensID
	} else if exifdata.LensModel != "" && exifdata.LensModel != "-" {
		lens = exifdata.LensModel
	}

	dpi := strconv.Itoa(config.Text.Dpi)
	x := "+" + strconv.Itoa(config.Text.Margin)
	y := "+" + strconv.Itoa(config.Text.Element[0].Size + config.Text.Element[0].Marginbottom + config.Text.Element[1].Margintop + config.Text.Element[1].Size + config.Text.Element[1].Marginbottom + config.Text.Element[2].Margintop + config.Text.Element[2].Size + config.Text.Element[2].Marginbottom)
	cmdArgs = append(cmdArgs, "-density", dpi, "-pointsize", "24", "-font", config.Text.Element[0].Font, "-fill", config.Text.Element[0].Color, "-gravity", config.Text.Direction, "-annotate", x+y, exifdata.Copyright)

	y = "+" + strconv.Itoa(config.Text.Element[1].Size + config.Text.Element[1].Marginbottom + config.Text.Element[2].Margintop + config.Text.Element[2].Size + config.Text.Element[2].Marginbottom)
	cmdArgs = append(cmdArgs, "-density", dpi, "-pointsize", strconv.Itoa(config.Text.Element[1].Size), "-font", config.Text.Element[1].Font, "-fill", config.Text.Element[1].Color, "-gravity", config.Text.Direction, "-annotate", x+y, fmt.Sprintf("%s / %s", exifdata.Model, exifdata.Make))

	y = "+" + strconv.Itoa(config.Text.Element[2].Size + config.Text.Element[2].Marginbottom)
	cmdArgs = append(cmdArgs, "-density", dpi, "-pointsize", strconv.Itoa(config.Text.Element[2].Size), "-font", config.Text.Element[2].Font, "-fill", config.Text.Element[2].Color, "-gravity", config.Text.Direction, "-annotate", x+y, fmt.Sprintf("%s f%s %ss ISO%s", lens, exifdata.Fnumber, exifdata.Shutterspeed, exifdata.ISO))

	cmdArgs = append(cmdArgs, "-quality", "100", "tmp1.webp")
	cmd := exec.Command(config.Path.Imagemagick, cmdArgs...)
	if err := cmd.Run(); err != nil {
		return err
	}

	if err := rotateImage(inputPath); err != nil {
		return err
	}

	if err := mergeImage(inputPath, outputPath); err != nil {
		return err
	}

	return nil
}

func rotateImage(inputPath string) error {
	var width int
	var height int
	var rotationAngle int
	if exifdata.Width > exifdata.Height {
		width = config.Length
		height = (exifdata.Height * config.Length) / exifdata.Width
		rotationAngle = exifdata.Orientation
	} else {
		width = (exifdata.Width * config.Length) / exifdata.Height
		height = config.Length
		rotationAngle = exifdata.Orientation
	}

	cmd := exec.Command(config.Path.Imagemagick, inputPath, "-resize", fmt.Sprintf("%dx%d", width, height), "-rotate", fmt.Sprintf("+%d", rotationAngle), "-orient", "undefined", "-quality", "100", "tmp2.webp")
	return cmd.Run()
}

func mergeImage(inputPath, outputPath string) error {
	if outputPath == "" {
		outputPath = filepath.Join(".", filepath.Base(inputPath))
		err := copyFile(inputPath, outputPath)
		if err != nil {
			return err
		}
		fmt.Printf("Output file not specified. Copied input file to %s\n", outputPath)
		return nil
	}

	cmd := exec.Command(config.Path.Imagemagick, "tmp1.webp", "tmp2.webp", "-gravity", "north", "-geometry", "+0+32", "-compose", "over", "-composite", outputPath)
	if err := cmd.Run(); err != nil {
		return err
	}

	os.Remove("tmp1.webp")
	os.Remove("tmp2.webp")
	return nil
}

func copyFile(src, dst string) error {
	sourceFile, err := os.Open(src)
	if err != nil {
		return err
	}
	defer sourceFile.Close()

	destFile, err := os.Create(dst)
	if err != nil {
		return err
	}
	defer destFile.Close()

	_, err = io.Copy(destFile, sourceFile)
	return err
}

func main() {
	funcMap := template.FuncMap{
		"join": func(elements []string, sep string) string {
			return strings.Join(elements, sep)
		},
	}

	cli.AppHelpTemplate = `NAME:
   {{.Name}} - {{.Usage}}

VERSION:
   {{.Version}}

USAGE:
   {{.UsageText}}

ARGUMENTS:
   {{range .Flags}}
   {{- with $names := .Names }}{{join $names ", "}}{{end}}
   {{- if .Aliases}} (short: {{join .Aliases ", "}}){{end}}
   :
       {{.Usage}}
   {{end}}
`

	app := &cli.App{
		Name:    "addframe",
		Usage:   "Add a frame to an image",
		UsageText: "addframe [-c configuration_file_path] -i <source_file_path> -o <destination_file_path>",
		Version: "v1.2.0",
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:    "config",
				Aliases: []string{"c"},
				Usage:   "Load configuration from configuration_file_path",
				Value:   "addframe.json in the current directory",
			},
			&cli.StringFlag{
				Name:     "input",
				Aliases:  []string{"i"},
				Usage:    "Input image file path",
				Required: true,
			},
			&cli.StringFlag{
				Name:     "output",
				Aliases:  []string{"o"},
				Usage:    "Output image file path",
				Required: true,
			},
		},
		Action: func(c *cli.Context) error {
			configPath := c.String("config")
			if configPath == "" {
				configPath = "addframe.json"
			}
			if err := config.loadConfig(configPath); err != nil {
				return err
			}

			inputPath := c.String("input")
			if inputPath == "" {
				return fmt.Errorf("input file path must be specified")
			}

			if err := exifdata.load(inputPath); err != nil {
				return err
			}
			outputPath := c.String("output")
			if outputPath == "" {
				return fmt.Errorf("output file path must be specified")
			}

			if err := createFrame(inputPath, outputPath); err != nil {
				return err
			}

			fmt.Println("Done!!")
			return nil
		},
	}

	cli.HelpPrinter = func(w io.Writer, templ string, data interface{}) {
		t := template.Must(template.New("help").Funcs(funcMap).Parse(templ))
		err := t.Execute(w, data)
		if err != nil {
			panic(err)
		}
	}

	err := app.Run(os.Args)
	if err != nil {
		fmt.Println(err)
	}
}
