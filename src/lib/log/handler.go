package log

import (
	"fmt"
	"io"
	"os"
	"time"
)

type Handler interface {
	Write(p []byte) (n int, err error)
	Close()
}

type base struct {
	msg  chan []byte
	quit chan bool
}

func (h *base) Write(p []byte) (n int, err error) {
	m := make([]byte, len(p))

	copy(m, p)

	h.msg <- m

	return len(p), nil
}

func (h *base) Close() {
	close(h.quit)
}

type StreamHandler struct {
	base
	w io.Writer
}

func NewStreamHandler(w io.Writer, msgNum int) (*StreamHandler, error) {
	h := new(StreamHandler)

	h.w = w

	h.msg = make(chan []byte, msgNum)
	h.quit = make(chan bool)

	go h.run()

	return h, nil
}

func NewDefaultStreamHandler(w io.Writer) (*StreamHandler, error) {
	return NewStreamHandler(w, 1024)
}

func (h *StreamHandler) run() {
	for {
		select {
		case m := <-h.msg:
			h.w.Write(m)
		case <-h.quit:
			return
		}
	}
}

type FileHandler struct {
	base
	fd *os.File
}

func NewFileHandler(fileName string, flag int, msgNum int) (*FileHandler, error) {
	f, err := os.OpenFile(fileName, flag, 0)
	if err != nil {
		return nil, err
	}

	h := new(FileHandler)

	h.fd = f

	h.msg = make(chan []byte, msgNum)
	h.quit = make(chan bool)

	go h.run()

	return h, nil
}

func NewDefaultFileHandler(fileName string, flag int) (*FileHandler, error) {
	return NewFileHandler(fileName, flag, 1024)
}

func (h *FileHandler) run() {
	for {
		select {
		case m := <-h.msg:
			h.fd.Write(m)
		case <-h.quit:
			h.fd.Close()
			return
		}
	}
}

//refer: http://docs.python.org/2/library/logging.handlers.html
//same like python TimedRotatingFileHandler

type TimeRotatingFileHandler struct {
	base

	fd *os.File

	baseName   string
	interval   int64
	suffix     string
	rolloverAt int64
}

const (
	WhenSecond = iota
	WhenMinute
	WhenHour
	WhenDay
)

func NewTimeRotatingFileHandler(baseName string, when int8, interval int, msgNum int) (*TimeRotatingFileHandler, error) {
	h := new(TimeRotatingFileHandler)

	h.baseName = baseName

	switch when {
	case WhenSecond:
		h.interval = 1
		h.suffix = "2006-01-02_15-04-05"
	case WhenMinute:
		h.interval = 60
		h.suffix = "2006-01-02_15-04"
	case WhenHour:
		h.interval = 3600
		h.suffix = "2006-01-02_15"
	case WhenDay:
		h.interval = 3600 * 24
		h.suffix = "2006-01-02"
	default:
		e := fmt.Errorf("invalid when_rotate: %d", when)
		panic(e)
	}

	h.interval = h.interval * int64(interval)

	var err error
	h.fd, err = os.OpenFile(h.baseName, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		panic(err)
	}

	fInfo, _ := h.fd.Stat()
	h.rolloverAt = fInfo.ModTime().Unix() + h.interval
	h.msg = make(chan []byte, msgNum)
	h.quit = make(chan bool)

	go h.run()

	return h, nil
}

func NewDefaultTimeRotatingFileHandler(baseName string, when int8, interval int) (*TimeRotatingFileHandler, error) {
	return NewTimeRotatingFileHandler(baseName, when, interval, 1024)
}

func (h *TimeRotatingFileHandler) doRollover() {
	//refer http://hg.python.org/cpython/file/2.7/Lib/logging/handlers.py
	now := time.Now()

	if h.rolloverAt <= now.Unix() {
		fName := h.baseName + now.Format(h.suffix)
		h.fd.Close()
		e := os.Rename(h.baseName, fName)
		if e != nil {
			panic(e)
		}

		h.fd, _ = os.OpenFile(h.baseName, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)

		h.rolloverAt = time.Now().Unix() + h.interval
	}
}

func (h *TimeRotatingFileHandler) run() {
	for {
		select {
		case m := <-h.msg:
			h.doRollover()
			if h.fd != nil {
				_, e := h.fd.Write(m)
				if e != nil {
					panic(e)
				}
			}
		case <-h.quit:
			if h.fd != nil {
				h.fd.Close()
			}
			return
		}
	}
}
