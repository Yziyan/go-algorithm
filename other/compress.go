// @Author: Ciusyan 5/26/24

package other

import (
	"bytes"
	"encoding/base64"
	"fmt"
	"io"

	"github.com/klauspost/compress/zstd"
)

func Compress(data string) (string, error) {
	// 解码Base64字符串
	decodedData, err := base64.StdEncoding.DecodeString(data)
	if err != nil {
		return "", fmt.Errorf("error decoding base64 data: %v", err)
	}

	// 解压缩Zstd数据
	decompressedData, err := decompressZstd(decodedData)
	if err != nil {

		return "", fmt.Errorf("error decompressing data: %v", err)
	}

	return string(decompressedData), nil
}

func decompressZstd(data []byte) ([]byte, error) {
	decoder, err := zstd.NewReader(bytes.NewReader(data))
	if err != nil {
		return nil, err
	}
	defer decoder.Close()

	var out bytes.Buffer
	_, err = io.Copy(&out, decoder)
	if err != nil {
		return nil, err
	}
	return out.Bytes(), nil
}
