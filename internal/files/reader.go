package files

import (
	"bufio"
	"bytes"
	"io"

	"github.com/simon-co/fltr-cli/internal/apperr"
)

type FileReader struct {
	filepath string
	buffer   []byte
	file     *bufio.Reader
	Data     *bytes.Buffer
}

func (_ FileReader) New(filepath string) *FileReader{
    return &FileReader{
    	filepath: filepath,
    	buffer:   make([]byte, 4096),
    	file:     &bufio.Reader{},
    	Data:     &bytes.Buffer{},
    }
}  

func (self *FileReader) Read() error{
    if err := self.openFile(); err != nil {
        return apperr.Parse(err)
    }
    for {
         n, err := self.file.Read(self.buffer)
         if err != nil {
             if err != io.EOF {
                 return apperr.Parse(err)
             }
             break
         }else {
             self.Data.Write(self.buffer[:n])
         }
    }
    return nil
}

func (self *FileReader) openFile() error {
    f, err := efs.Open(self.filepath)
    if err != nil {
        return apperr.Parse(err)
    }
    s, _ := f.Stat()
    self.file = bufio.NewReader(f)
    self.Data = bytes.NewBuffer(make([]byte, 0, s.Size()))
    return nil
}
