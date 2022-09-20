package utils

import (
	"bytes"
	"compress/gzip"
	"encoding/base64"
	"io"
)

func Compress(raw []byte) ([]byte, error) {
	var compressed bytes.Buffer

	w := gzip.NewWriter(&compressed)
	w.Write(raw)
	err := w.Close()
	if err != nil {
		return nil, err
	}

	return compressed.Bytes(), nil
}

func Decompress(raw []byte) ([]byte, error) {
	compressed := bytes.NewBuffer(raw)

	r, err := gzip.NewReader(compressed)
	if err != nil {
		return nil, err
	}

	decompressed, err := io.ReadAll(r)
	if err != nil {
		// fmt.Println("hello")
		return nil, err
	}

	err = r.Close()
	if err != nil {
		return nil, err
	}
	return decompressed, nil
}

func Encode(raw []byte) []byte {
	encoded := make([]byte, base64.StdEncoding.EncodedLen(len(raw)))
	base64.StdEncoding.Encode(encoded, raw)
	return encoded
}

func Decode(raw []byte) ([]byte, error) {
	decoded := make([]byte, base64.StdEncoding.DecodedLen(len(raw)))
	n, err := base64.StdEncoding.Decode(decoded, raw)
	if err != nil {
		return nil, err
	}
	return decoded[:n], nil
}

func CompressEncode(raw []byte) ([]byte, error) {
	compressed, err := Compress(raw)
	if err != nil {
		return nil, err
	}
	return Encode(compressed), nil
}

func DecodeDecompress(raw []byte) ([]byte, error) {
	decoded, err := Decode(raw)
	if err != nil {
		return nil, err
	}
	decompressed, err := Decompress(decoded)
	if err != nil {
		return nil, err
	}
	return decompressed, nil
}
