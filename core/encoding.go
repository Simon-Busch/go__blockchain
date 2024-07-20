package core

import "io"

type Encoder[T any] interface {
	Encode(w io.Writer, t T) error
}

type Decoder[T any] interface {
	Decode(r io.Reader, t T) error
}
