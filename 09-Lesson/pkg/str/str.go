package str

import (
	"encoding/json"
	"io"
	"log"
)

type Data struct {
	strings []string
}

func (s *Data) Write(p []byte) (n int, err error) {
	s.strings = append(s.strings, string(p))
	return len(s.strings), nil

}

func Add(w io.Writer, any ...any) {
	for _, a := range any {
		if s, ok := a.(string); ok {
			b, err := json.Marshal(s)
			if err != nil {
				log.Println(err)
			}

			_, err = w.Write(b)
			if err != nil {
				log.Println(err)
			}
		}

	}

}
