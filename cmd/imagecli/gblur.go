package imagecli

import (
	"fmt"
	"imageCLI/pkg/loader"
	"imageCLI/pkg/service"
	"log"
	"time"

	"github.com/spf13/cobra"
)

var radius float64

func NewGBlurCmd(imageService *service.ImageService) *cobra.Command {
	cmd := cobra.Command{
		Use:   "gblur",
		Short: "Применение гауссова размытия к изображению",
		Run: func(cmd *cobra.Command, args []string) {
			ops := service.Operations{
				Blur:        &service.BlurParams{Radius: radius},
				UseParallel: !noParallel,
			}

			t := time.Now()
			images, err := loader.LoadImages(inputPath, ops.UseParallel)
			if err != nil {
				log.Fatalf("Ошибка загрузки изображений: %v", err)
			}

			processedImages, err := imageService.ProcessImages(images, ops)
			if err != nil {
				log.Fatalf("Ошибка обработки: %v", err)
			}

			err = loader.SaveImagesToDir(outputPath, processedImages)
			if err != nil {
				log.Fatalf("Ошибка сохранения изображений: %v", err)
			}
			fmt.Println(time.Since(t))
			if err != nil {
				log.Fatalf("Ошибка обработки: %v", err)
			}
		},
	}

	// Определяем флаги для команды
	cmd.Flags().Float64VarP(&radius, "radius", "r", 0, "Радиус размытия")
	cmd.MarkFlagRequired("radius")

	return &cmd
}
