package imagecli

import (
	"fmt"
	"imageCLI/pkg/service"
	"os"

	"github.com/spf13/cobra"
)

var (
	inputPath  string
	outputPath string
	noParallel bool
)

func Execute(service *service.ImageService) {
	rootCmd := &cobra.Command{
		Use:   "imagecli",
		Short: "Инструмент для обработки изображений",
		Long: `Инструмент для обработки изображений с поддержкой различных операций:
- Изменение размера
- Гауссово размытие
- Настройка цветовых параметров

Более подробная документация доступна на GitHub: https://github.com/yourusername/imageCLI 
Для каждой команды доступен флаг --help, который выводит вспомогающую информацию.`,
		CompletionOptions: cobra.CompletionOptions{
			DisableDefaultCmd: true,
		},
		PersistentPreRunE: func(cmd *cobra.Command, args []string) error {
			if cmd.Name() == "help" {
				return nil
			}
			if inputPath == "" || outputPath == "" {
				return fmt.Errorf("необходимо указать параметры --input и --output (-i и -o)")
			}
			return nil
		},
	}

	rootCmd.PersistentFlags().StringVarP(&inputPath, "input", "i", "", "Путь к входному файлу или директории")
	rootCmd.PersistentFlags().StringVarP(&outputPath, "output", "o", "", "Путь к выходному файлу или директории")
	rootCmd.PersistentFlags().BoolVar(&noParallel, "no-parallel", false, "Отключить параллельную обработку")

	rootCmd.AddCommand(NewResizeCmd(service))
	rootCmd.AddCommand(NewGBlurCmd(service))
	rootCmd.AddCommand(NewAdjustCmd(service))

	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
