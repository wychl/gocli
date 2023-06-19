/*
Copyright © 2023 Abner Wanyan <abner.wanyan@gmail.com>
*/
package flatcmd

import (
	"bytes"
	"compress/flate"
	"encoding/base64"
	"io"
	"io/ioutil"
	"log"
	"os"

	"github.com/spf13/cobra"
)

const deCompressExp = `# 解压字符串
echo ykjNyclXKM8vyknhAgQAAP// | gocli flate decompress

# 读取文件内容解压
cat demo.txt | gocli flate decompress`

// decompressCmd represents the decompress command
var decompressCmd = &cobra.Command{
	Use:     "decompress",
	Short:   "flate解压缩",
	Long:    `flate解压缩`,
	Example: deCompressExp,
	Run: func(cmd *cobra.Command, args []string) {
		data, err := io.ReadAll(os.Stdin)
		if err != nil {
			log.Fatal(err)
		}
		if len(data) > 0 {
			cData, err := deComp(data)
			if err != nil {
				log.Fatal(err)
			}
			os.Stdout.Write(cData)
		}
	},
}

func init() {
	FlateCmd.AddCommand(decompressCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// decompressCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// decompressCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

func deComp(data []byte) ([]byte, error) {
	dest := make([]byte, base64.StdEncoding.EncodedLen(len(data)))
	base64.StdEncoding.Decode(dest, data)
	zr := flate.NewReader(bytes.NewReader(dest))
	if err := zr.Close(); err != nil {
		return nil, err
	}
	decodeData, err := ioutil.ReadAll(zr)
	if err != nil {
		return nil, err
	}
	return decodeData, nil
}
