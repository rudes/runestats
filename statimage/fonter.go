package statimage

import (
	"bufio"
	"image"
	"image/color"
	"image/draw"
	"image/png"
	"io/ioutil"
	"log"
	"os"

	"github.com/golang/freetype"
)

const (
	_dpi           = 72
	_fontName      = "fonts/runescape_chat_font.ttf"
	_skillNum      = 35
	_totalNum      = 20
	_totalText     = 21
	_baseImageFile = "images/os_rs_base.png"
	_imageWidth    = 315
	_imageHeight   = 434
)

// NewRuneStat generates an image based on the stats argument
// it will name the file by the player's name as a PNG file
func NewRuneStat(player string, stats []string, _staticDir string) {
	fontColor := color.RGBA{237, 219, 72, 255}
	fontbits, err := ioutil.ReadFile(_staticDir + _fontName)
	if err != nil {
		log.Println(err)
		return
	}
	font, err := freetype.ParseFont(fontbits)
	if err != nil {
		log.Println(err)
		return
	}
	baseImage, err := os.Open(_staticDir + _baseImageFile)
	if err != nil {
		log.Println(err)
		return
	}
	defer baseImage.Close()
	img, err := png.Decode(baseImage)
	if err != nil {
		log.Println(err)
		return
	}
	dst, ok := img.(draw.Image)
	if !ok {
		log.Println("failed to make draw image")
		return
	}
	src := image.NewRGBA(image.Rect(0, 0, _imageWidth, _imageHeight))
	draw.Draw(src, src.Bounds(), &image.Uniform{fontColor}, image.ZP, draw.Src)
	c := freetype.NewContext()
	c.SetDPI(_dpi)
	c.SetFont(font)
	c.SetFontSize(_skillNum)
	c.SetClip(dst.Bounds())
	c.SetDst(dst)
	c.SetSrc(src)
	pt := freetype.Pt(62, 10+int(c.PointToFixed(_skillNum)>>6))
	_, err = c.DrawString(stats[1], pt)
	pt.Y += c.PointToFixed(_skillNum * 1.5)
	_, err = c.DrawString(stats[3], pt)
	left := []string{stats[1], stats[3], stats[2], stats[5],
		stats[6], stats[7], stats[21], stats[23]}
	middle := []string{stats[4], stats[17], stats[16], stats[18],
		stats[13], stats[10], stats[19], stats[22]}
	right := []string{stats[15], stats[14], stats[11], stats[8],
		stats[12], stats[9], stats[20]}
	drawCol(62, c, left)
	drawCol(124, c, middle)
	drawCol(186, c, right)
	drawTotal(c, stats[0])
	if err != nil {
		log.Println(err)
		return
	}

	livestat, err := os.Create(_staticDir + "images/" + player + ".png")
	if err != nil {
		log.Println(err)
		return
	}
	defer livestat.Close()
	b := bufio.NewWriter(livestat)
	err = png.Encode(b, dst)
	if err != nil {
		log.Println(err)
		return
	}
	err = b.Flush()
	if err != nil {
		log.Println(err)
		return
	}
}

// drawCol takes a column pixel number for X and draws the rows in the column
func drawCol(col int, c *freetype.Context, rows []string) {
	pt := freetype.Pt(col, 10+int(c.PointToFixed(_skillNum)>>6))
	for _, s := range rows {
		_, err := c.DrawString(s, pt)
		if err != nil {
			log.Println(err)
			return
		}
		pt.Y += c.PointToFixed(_skillNum * 1.5)
	}
}

// drawTotal draws the Total Level data from the stats
func drawTotal(c *freetype.Context, total string) {
	pt := freetype.Pt(214, 10+int(c.PointToFixed(_skillNum)>>6))
	c.SetFontSize(_totalText)
	pt.Y += c.PointToFixed(_totalText * 1.5)
	c.DrawString("Total level:", pt)
	pt.X = 234
	c.SetFontSize(_totalNum)
	c.DrawString(total, pt)
}
