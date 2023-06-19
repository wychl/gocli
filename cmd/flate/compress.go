/*
Copyright © 2023 Abner Wanyan <abner.wanyan@gmail.com>
*/
package flatcmd

import (
	"bytes"
	"compress/flate"
	"encoding/base64"
	"io"
	"log"
	"os"

	"github.com/spf13/cobra"
)

const compressExp = `# 压缩字符串
echo hello world | gocli flate compress

# 读取文件内容压缩
cat README.md | gocli flate compress`

// compressCmd represents the compress command
var compressCmd = &cobra.Command{
	Use:     "compress",
	Short:   "flate压缩",
	Long:    `flate压缩`,
	Example: compressExp,
	Run: func(cmd *cobra.Command, args []string) {
		data, err := io.ReadAll(os.Stdin)
		if err != nil {
			log.Fatal(err)
		}
		if len(data) > 0 {
			cData, err := comp(data)
			if err != nil {
				log.Fatal(err)
			}
			os.Stdout.Write(cData)
		}
	},
}

func init() {
	FlateCmd.AddCommand(compressCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// compressCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// compressCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

func comp(data []byte) ([]byte, error) {
	var b bytes.Buffer
	zw, err := flate.NewWriter(&b, flate.DefaultCompression)
	if err != nil {
		return nil, err
	}
	_, err = zw.Write(data)
	if err != nil {
		return nil, err
	}
	if err := zw.Close(); err != nil {
		return nil, err
	}
	encodeData := b.Bytes()
	dest := make([]byte, base64.StdEncoding.EncodedLen(len(encodeData)))
	base64.StdEncoding.Encode(dest, encodeData)
	return dest, nil
}
