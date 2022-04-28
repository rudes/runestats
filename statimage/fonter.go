package statimage

import (
	"bufio"
	"image"
	"image/color"
	"image/draw"
	"image/png"
	"io/ioutil"
	"os"

	"github.com/golang/freetype"
)

const (
	_dpi           = 72
	_fontName      = "fonts/runescape_chat_font.ttf"
	_skillNum      = 35
	_totalText     = 22
	_baseImageFile = "images/os_rs_base.png"
	_imageWidth    = 315
	_imageHeight   = 434
)

// NewRuneStat generates an image based on the stats argument
// it will name the file by the player's name as a PNG file
func NewRuneStat(player string, stats []string, _staticDir string) error {
	fontColor := color.RGBA{237, 219, 72, 255}
	fontbits, err := ioutil.ReadFile(_staticDir + _fontName)
	if err != nil {
		return err
	}
	font, err := freetype.ParseFont(fontbits)
	if err != nil {
		return err
	}
	baseImage, err := os.Open(_staticDir + _baseImageFile)
	if err != nil {
		return err
	}
	defer baseImage.Close()
	img, err := png.Decode(baseImage)
	if err != nil {
		return err
	}
	dst, ok := img.(draw.Image)
	if !ok {
		return err
	}
	src := image.NewRGBA(image.Rect(0, 0, _imageWidth, _imageHeight))
	draw.Draw(src, src.Bounds(), &image.Uniform{fontColor}, image.Point{}, draw.Src)
	c := freetype.NewContext()
	c.SetDPI(_dpi)
	c.SetFont(font)
	c.SetFontSize(_skillNum)
	c.SetClip(dst.Bounds())
	c.SetDst(dst)
	c.SetSrc(src)
	left := []string{stats[1], stats[3], stats[2], stats[5],
		stats[6], stats[7], stats[21], stats[23]}
	middle := []string{stats[4], stats[17], stats[16], stats[18],
		stats[13], stats[10], stats[19], stats[22]}
	right := []string{stats[15], stats[14], stats[11], stats[8],
		stats[12], stats[9], stats[20]}
	err = drawCol(62, c, left)
	if err != nil {
		return err
	}
	err = drawCol(161, c, middle)
	if err != nil {
		return err
	}
	err = drawCol(261, c, right)
	if err != nil {
		return err
	}
	drawTotal(c, stats[0])
	if err != nil {
		return err
	}

	livestat, err := os.Create(_staticDir + "images/os_rs/" + player + ".png")
	if err != nil {
		return err
	}
	defer livestat.Close()
	b := bufio.NewWriter(livestat)
	err = png.Encode(b, dst)
	if err != nil {
		return err
	}
	err = b.Flush()
	if err != nil {
		return err
	}
	return nil
}

// drawCol takes a column pixel number for X and draws the rows in the column
func drawCol(col int, c *freetype.Context, rows []string) error {
	y := 45
	pt := freetype.Pt(col, y)
	for _, s := range rows {
		_, err := c.DrawString(s, pt)
		if err != nil {
			return err
		}
		y += 50
		pt = freetype.Pt(col, y)
	}
	return nil
}

// drawTotal draws the Total Level data from the stats
func drawTotal(c *freetype.Context, total string) {
	pt := freetype.Pt(213, 384)
	c.SetFontSize(_totalText)
	_, _ = c.DrawString("Total level:", pt)
	pt = freetype.Pt(239, 404)
	_, _ = c.DrawString(total, pt)
}
