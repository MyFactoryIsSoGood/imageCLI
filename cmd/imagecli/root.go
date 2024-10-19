// cmd/imagecli/root.go
package imagecli

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var (
	inputPath  string
	outputPath string
	noParallel bool
)

var rootCmd = &cobra.Command{
	Use:   "imagecli",
	Short: "Инструмент для обработки изображений",
	PersistentPreRunE: func(cmd *cobra.Command, args []string) error {
		if inputPath == "" || outputPath == "" {
			return fmt.Errorf("необходимо указать параметры --input и --output (-i и -o)")
		}
		return nil
	},
}

func init() {
	rootCmd.PersistentFlags().StringVarP(&inputPath, "input", "i", "", "Путь к входному файлу или директории")
	rootCmd.PersistentFlags().StringVarP(&outputPath, "output", "o", "", "Путь к выходному файлу или директории")
	rootCmd.PersistentFlags().BoolVar(&noParallel, "no-parallel", false, "Отключить параллельную обработку")

	rootCmd.AddCommand(resizeCmd)
	rootCmd.AddCommand(gblurCmd)
	rootCmd.AddCommand(adjustCmd)
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
