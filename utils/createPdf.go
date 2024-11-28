package utils

import (
	"bytes"
	"fmt"
	"image/jpeg"
	"image/png"
	"io"
	"log"
	"net/http"

	"github.com/johnfercher/maroto/v2"
	"github.com/johnfercher/maroto/v2/pkg/components/code"
	"github.com/johnfercher/maroto/v2/pkg/components/col"
	"github.com/johnfercher/maroto/v2/pkg/components/image"
	"github.com/johnfercher/maroto/v2/pkg/components/page"
	"github.com/johnfercher/maroto/v2/pkg/components/row"
	"github.com/johnfercher/maroto/v2/pkg/components/text"
	"github.com/johnfercher/maroto/v2/pkg/config"
	"github.com/johnfercher/maroto/v2/pkg/consts/align"
	"github.com/johnfercher/maroto/v2/pkg/consts/extension"
	"github.com/johnfercher/maroto/v2/pkg/consts/fontfamily"
	"github.com/johnfercher/maroto/v2/pkg/consts/fontstyle"
	"github.com/johnfercher/maroto/v2/pkg/consts/orientation"
	"github.com/johnfercher/maroto/v2/pkg/core"
	"github.com/johnfercher/maroto/v2/pkg/props"
	"github.com/nfnt/resize"
)

func GetVerticalCard(url, ext, fullname, message, date, time, qrMessage string) core.Maroto {
	response, e := http.Get(url)
	if e != nil {
		log.Fatal(e)
	}
	defer response.Body.Close()

	resp, e := http.Get("http://localhost:8090/api/files/4xn4qtssx56x64m/iybaux6usx8hadz/powered_by_L7QyGFaOsx.png")
	if e != nil {
		log.Fatal(e)
	}
	defer resp.Body.Close()

	// Copy into buffer.
	buf := new(bytes.Buffer)
	if _, err := io.Copy(buf, response.Body); err != nil {
		// TODO: handle error
	}

	// create buffer
	buffImg := new(bytes.Buffer)
	var exten extension.Type

	if ext == "jpg" || ext == "jpeg" {
		// decode jpeg into image.Image
		img, err := jpeg.Decode(buf)
		if err != nil {
			log.Fatal(err)
		}

		newImage := resize.Resize(3000, 2000, img, resize.Lanczos3)

		// encode image to buffer

		err = jpeg.Encode(buffImg, newImage, &jpeg.Options{
			Quality: 100,
		},
		)
		if err != nil {
			fmt.Println("failed to create buffer", err)
		}
		exten = extension.Jpeg
	}
	if ext == "png" {
		// decode jpeg into image.Image
		img, err := png.Decode(buf)
		if err != nil {
			log.Fatal(err)
		}

		newImage := resize.Resize(3000, 2000, img, resize.Lanczos3)

		// encode image to buffer
		err = png.Encode(buffImg, newImage)
		if err != nil {
			fmt.Println("failed to create buffer", err)
		}
		exten = extension.Png
	}

	// Copy into buffer.
	buff := new(bytes.Buffer)
	if _, err := io.Copy(buff, resp.Body); err != nil {
		// TODO: handle error
	}

	b := config.NewBuilder().
		WithDimensions(120, 180).
		WithTopMargin(0).
		WithLeftMargin(0).
		WithRightMargin(0).
		WithBottomMargin(0).
		WithOrientation(orientation.Vertical).
		WithMaxGridSize(20).
		WithBackgroundImage(buffImg.Bytes(), exten)

	mrt := maroto.New(b.Build())
	m := maroto.NewMetricsDecorator(mrt)
	err := m.RegisterFooter(
		image.NewAutoFromBytesRow(buff.Bytes(), extension.Png, props.Rect{
			Center:  true,
			Percent: 90,
		}),
	)
	if err != nil {
		log.Fatal(err)
	}
	m.AddPages(addVerticalPage(fullname, message, date, time, qrMessage))

	return m
}

func GetHorizontalCard(url, ext, fullname, message, date, time, qrMessage string) core.Maroto {
	response, e := http.Get(url)
	if e != nil {
		log.Fatal(e)
	}
	defer response.Body.Close()

	resp, e := http.Get("http://localhost:8090/api/files/4xn4qtssx56x64m/iybaux6usx8hadz/powered_by_L7QyGFaOsx.png")
	if e != nil {
		log.Fatal(e)
	}
	defer resp.Body.Close()

	// Copy into buffer.
	buf := new(bytes.Buffer)
	if _, err := io.Copy(buf, response.Body); err != nil {
		// TODO: handle error
	}

	// create buffer
	buffImg := new(bytes.Buffer)
	var exten extension.Type

	if ext == "jpg" || ext == "jpeg" {
		// decode jpeg into image.Image
		img, err := jpeg.Decode(buf)
		if err != nil {
			log.Fatal(err)
		}

		newImage := resize.Resize(860, 1280, img, resize.Lanczos3)

		// encode image to buffer

		err = jpeg.Encode(buffImg, newImage, &jpeg.Options{
			Quality: 100,
		},
		)
		if err != nil {
			fmt.Println("failed to create buffer", err)
		}
		exten = extension.Jpeg
	}

	if ext == "png" {
		// decode jpeg into image.Image
		img, err := png.Decode(buf)
		if err != nil {
			log.Fatal(err)
		}

		newImage := resize.Resize(860, 1280, img, resize.Lanczos3)

		// encode image to buffer

		err = png.Encode(buffImg, newImage)
		if err != nil {
			fmt.Println("failed to create buffer", err)
		}
		exten = extension.Png
	}
	// Copy into buffer.
	buff := new(bytes.Buffer)
	if _, err := io.Copy(buff, resp.Body); err != nil {
		// TODO: handle error
	}

	b := config.NewBuilder().
		WithDimensions(200, 120).
		WithTopMargin(0).
		WithLeftMargin(0).
		WithRightMargin(0).
		WithBottomMargin(0).
		WithOrientation(orientation.Horizontal).
		WithMaxGridSize(20).
		WithBackgroundImage(buffImg.Bytes(), exten)

	mrt := maroto.New(b.Build())
	m := maroto.NewMetricsDecorator(mrt)

	m.AddPages(addHorizontalPage(buff.Bytes(), fullname, message, date, time, qrMessage))

	return m
}

func addVerticalPage(fullname, message, date, time, qrMessage string) core.Page {
	return page.New().Add(
		row.New(82),
		row.New(10).Add(
			col.New(4),
			text.NewCol(12, fullname, props.Text{
				Size:   15,
				Align:  align.Center,
				Style:  fontstyle.Bold,
				Family: fontfamily.Helvetica,
			}),
			col.New(4),
		),
		row.New(20).Add(
			text.NewCol(20, message, props.Text{
				Style:  fontstyle.Normal,
				Size:   10,
				Left:   5,
				Right:  5,
				Align:  align.Center,
				Family: fontfamily.Helvetica,
				Color:  &props.Color{Red: 95, Green: 95, Blue: 95},
			}),
		),
		row.New(10).Add(
			col.New(4),
			text.NewCol(12, date+" @ "+time, props.Text{
				Size:   10,
				Align:  align.Center,
				Style:  fontstyle.Bold,
				Family: fontfamily.Helvetica,
			}),
			col.New(4),
		),
		row.New(30).Add(
			code.NewQrCol(20,
				fmt.Sprintln(qrMessage),
				props.Rect{
					Center:  true,
					Percent: 100,
				},
			),
		),
		row.New(5),
		row.New(10).Add(
			col.New(4),
			text.NewCol(12, "Please do not distribute or print this file.", props.Text{
				Size:   8,
				Align:  align.Center,
				Style:  fontstyle.Bold,
				Family: fontfamily.Helvetica,
				Color:  &props.Color{Red: 95, Green: 95, Blue: 95},
			}),
			col.New(4),
		),
	)
}

func addHorizontalPage(ft []byte, fullname, message, date, time, qrMessage string) core.Page {
	return page.New().Add(
		row.New(20).Add(
			col.New(8),
			text.NewCol(12, fullname, props.Text{
				Size:   15,
				Align:  align.Center,
				Style:  fontstyle.Bold,
				Family: fontfamily.Helvetica,
				Top:    10,
			}),
		),
		row.New(20).Add(
			col.New(8),
			text.NewCol(12, message, props.Text{
				Style:  fontstyle.Normal,
				Size:   10,
				Left:   5,
				Right:  5,
				Align:  align.Center,
				Family: fontfamily.Helvetica,
				Color:  &props.Color{Red: 95, Green: 95, Blue: 95},
			}),
		),
		row.New(10).Add(
			col.New(8),
			text.NewCol(12, date+" @ "+time, props.Text{
				Size:   10,
				Align:  align.Center,
				Style:  fontstyle.Bold,
				Family: fontfamily.Helvetica,
			}),
		),
		row.New(35).Add(
			col.New(8),
			code.NewQrCol(12,
				fmt.Sprintln(qrMessage),
				props.Rect{
					Center:  true,
					Percent: 100,
				},
			),
		),
		row.New(15).Add(
			col.New(8),
			text.NewCol(12, "Please do not distribute or print this file.", props.Text{
				Size:   8,
				Align:  align.Center,
				Style:  fontstyle.Bold,
				Family: fontfamily.Helvetica,
				Color:  &props.Color{Red: 95, Green: 95, Blue: 95},
				Top:    5,
			}),
		),
		row.New(15).Add(
			col.New(8),
			image.NewFromBytesCol(12, ft, extension.Png, props.Rect{
				Center:  true,
				Percent: 90,
			}),
		),
	)
}
