package main

import (
	"database/sql"
	"fmt"
	"image"
	"image/color"
	"image/png"
	"log"
	"os"

	"github.com/dustin/go-heatmap"
	"github.com/dustin/go-heatmap/schemes"
)

func lerp(a uint8, b uint8, c uint8) uint8 {
	return uint8(
		float64(a)*(1.0-float64(c)/255.0) + (float64(b)),
	)
}

func blend_color(color_1 color.RGBA, color_2 color.RGBA) color.RGBA {
	return color.RGBA{
		lerp(color_1.R, color_2.R, color_2.A),
		lerp(color_1.G, color_2.G, color_2.A),
		lerp(color_1.B, color_2.B, color_2.A),
		color_1.A,
	}
}

func create_heatmap_for_level(
	level int,
	db *sql.DB,
	image_width int,
	image_height int,
	level_width int,
	level_height int,
	start_x int,
	start_y int,
) {
	rows, err := db.Query("SELECT * FROM deaths WHERE level = $1", level)
	if err != nil {
		log.Print("can'r read from deaths db: ", err)
	}

	points := []heatmap.DataPoint{}
	defer rows.Close()
	for rows.Next() {
		var death Death
		err := rows.Scan(&death.ID, &death.Level, &death.Time, &death.PositionX, &death.PositionY)
		if err != nil {
			log.Print("can't parse death from db: ", err)
		}
		points = append(
			points,
			heatmap.P(
				(death.PositionX+float64(start_x))/float64(level_width)*float64(image_width),
				(death.PositionY+float64(start_y))/float64(level_height)*float64(image_height),
			),
		)
	}

	if len(points) == 0 {
		log.Print("empty")
		return
	}

	imgFile, err := os.Open(fmt.Sprintf("level_0%d.png", level))
	if err != nil {
		log.Printf("can't open level_0%d.png", level)
	}
	defer imgFile.Close()

	level_img_decoded, err := png.Decode(imgFile)
	if err != nil {
		log.Printf("can't decode level_0%d.png", level)
	}
	level_img := level_img_decoded.(*image.RGBA)

	// scheme, _ := schemes.FromImage("../schemes/fire.png")
	scheme := schemes.AlphaFire

	heatmap_img := heatmap.Heatmap(image.Rect(0, 0, image_width, image_height), points, 150, 128, scheme).(*image.RGBA)
	f, _ := os.Create(fmt.Sprintf("level_0%d_heatmap.png", level))

	log.Printf("color: %d", level_img.RGBAAt(0, 0).A/255)

	for x := 0; x < image_width; x++ {
		for y := 0; y < image_height; y++ {
			level_img.SetRGBA(
				x,
				y,
				blend_color(level_img.RGBAAt(x, y), heatmap_img.RGBAAt(x, y)),
			)
		}
	}

	png.Encode(f, level_img)
}

func create_heatmap(db *sql.DB) {
	create_heatmap_for_level(
		1,
		db,
		982,
		477,
		18923,
		9205,
		2584,
		1688,
	)
	create_heatmap_for_level(
		2,
		db,
		1078,
		395,
		14599,
		5620,
		6285,
		952,
	)
}
