package cmd

import (
	"errors"
	"fmt"
	"image"
	"image/draw"
	"image/jpeg"
	"image/png"
	"net/http"
	"os"
	"path/filepath"
	"regexp"

	"github.com/KhashayarKhm/stamp/internal/config"
	"github.com/spf13/cobra"
)

type Watermark struct{}

func validateImage(path string) (string, error) {
	if file, err := os.Open(path); os.IsNotExist(err) {
		return "", errors.New("The main picture path is not exists")
	} else if err != nil {
		return "", fmt.Errorf("We got an error while opening the main picture:\n%v", err)
	} else {
		defer file.Close()
		buf := make([]byte, 512)
		if _, err = file.Read(buf); err != nil {
			return "", fmt.Errorf("We got an error while reading the main picture:\n%v", err)
		}

		contentType := http.DetectContentType(buf)
		if match, err := regexp.MatchString("^image/(jpg|jpeg|png)", contentType); err != nil {
			return "", fmt.Errorf("We got an error while checking the main picture format:\n%v", err)
		} else if !match {
			return "", fmt.Errorf("The main picture format unsupported. (Supported formats: jpg,jpeg,png - Entered file format: %s)", contentType)
		}

		return contentType, nil
	}
}

func (wmCmd Watermark) Command(trap chan os.Signal) *cobra.Command {
	run := func(cmd *cobra.Command, args []string) error {
		return wmCmd.run(config.Load(true), trap, args, cmd)
	}

	cmd := &cobra.Command{
		Use:   "watermark",
		Short: "Watermark images",
		Args:  cobra.ExactArgs(1),
		RunE:  run,
	}
	cmd.PersistentFlags().StringP("watermark", "w", "", "specify watermark image. default: ~/.stamp/default.png")
	cmd.PersistentFlags().StringP("output", "o", "", "output file name. default: the name of main image with \"stamped\" prefix")

	return cmd
}

func (wmCmd *Watermark) run(config *config.Config, trap chan os.Signal, args []string, cmd *cobra.Command) error {
	mainPicPath := args[0]
	watermarkPath, _ := cmd.Flags().GetString("watermark")
	if watermarkPath == "" {
		watermarkPath = config.WatermarkImg
	}

	outputName, _ := cmd.Flags().GetString("output")
	if outputName == "" {
		outputName = fmt.Sprintf("stamped_%s", filepath.Base(mainPicPath))
	}

	mainPicCT, err := validateImage(mainPicPath)
	if err != nil {
		return err
	}

	watermarkCT, err := validateImage(watermarkPath)
	if err != nil {
		return err
	}

	mainPicBin, err := os.Open(mainPicPath)
	if err != nil {
    return errors.New("Get error on opening the main image")
	}
	defer mainPicBin.Close()

	watermarkBin, err := os.Open(watermarkPath)
	if err != nil {
		return errors.New("Get error on opening watermark image")
	}
	defer watermarkBin.Close()

	var mainPic image.Image
	switch mainPicCT {
	case "image/jpeg":
		mainPic, err = jpeg.Decode(mainPicBin)
		if err != nil {
			return errors.New("Get error on decoding main image")
		}
	case "image/png":
		mainPic, err = png.Decode(mainPicBin)
		if err != nil {
			return errors.New("Get error on decoding main image")
		}
	default:
		return fmt.Errorf("The main image format unsupported. (Supported formats: jpg,jpeg,png - Entered file format: %s)", mainPicCT)
	}

	var watermarkImg image.Image
	switch watermarkCT {
	case "image/jpeg":
		watermarkImg, err = jpeg.Decode(mainPicBin)
		if err != nil {
			return errors.New("Get error on decoding watermark image")
		}
	case "image/png":
		watermarkImg, err = png.Decode(watermarkBin)
		if err != nil {
			return errors.New("Get error on decoding watermark image")
		}
	default:
		return fmt.Errorf("The watermark image format unsupported. (Supported formats: jpg,jpeg,png - Entered file format: %s)", watermarkCT)
	}

	offset := image.Pt(470, 50)
	bounds := mainPic.Bounds()
	dstImage := image.NewRGBA(bounds)

	draw.Draw(dstImage, bounds, mainPic, image.ZP, draw.Src)
	draw.Draw(dstImage, watermarkImg.Bounds().Add(offset), watermarkImg, image.ZP, draw.Over)

	outputImage, err := os.Create(outputName)
	if err != nil {
    return errors.New("Get error on creating output image")
	}
	defer outputImage.Close()

	switch mainPicCT {
	case "image/jpeg":
		err = jpeg.Encode(outputImage, dstImage, &jpeg.Options{Quality: jpeg.DefaultQuality})
		if err != nil {
			return errors.New("Get error on encoding output image")
		}
	case "image/png":
		err = png.Encode(outputImage, dstImage)
		if err != nil {
			return errors.New("Get error on encoding output image")
		}
	default:
		return fmt.Errorf("The output image format unsupported. (Supported formats: jpg,jpeg,png - Entered file format: %s)", mainPicCT)
	}

	return nil
}
