package common

import (
	"bufio"
	"github.com/zeromicro/go-zero/core/logx"
	"os"
	"time"
)

type MultiWriter struct {
	writer        logx.Writer
	consoleWriter logx.Writer
}

func (w *MultiWriter) Debug(v interface{}, fields ...logx.LogField) {
	//TODO implement me
	panic("implement me")
}

func NewMultiWriter(writer logx.Writer) (logx.Writer, error) {
	return &MultiWriter{
		writer:        writer,
		consoleWriter: logx.NewWriter(bufio.NewWriter(os.Stdout)),
	}, nil
}

func (w *MultiWriter) Alert(v interface{}) {
	w.consoleWriter.Alert(v)
	w.writer.Alert(v)
}

func (w *MultiWriter) Close() error {
	w.consoleWriter.Close()
	return w.writer.Close()
}

func (w *MultiWriter) Error(v interface{}, fields ...logx.LogField) {
	w.consoleWriter.Error(v, fields...)
	w.writer.Error(v, fields...)
}

func (w *MultiWriter) Info(v interface{}, fields ...logx.LogField) {
	w.consoleWriter.Info(v, fields...)
	w.writer.Info(v, fields...)
}

func (w *MultiWriter) Severe(v interface{}) {
	w.consoleWriter.Severe(v)
	w.writer.Severe(v)
}

func (w *MultiWriter) Slow(v interface{}, fields ...logx.LogField) {
	w.consoleWriter.Slow(v, fields...)
	w.writer.Slow(v, fields...)
}

func (w *MultiWriter) Stack(v interface{}) {
	w.consoleWriter.Stack(v)
	w.writer.Stack(v)
}

func (w *MultiWriter) Stat(v interface{}, fields ...logx.LogField) {
	w.consoleWriter.Stat(v, fields...)
	w.writer.Stat(v, fields...)
}

func LoggerInit(c logx.LogConf) {

	var d logx.LogConf

	d.Path = "D:\\project\\go\\dcs\\api\\frontend\\logs"
	d.Mode = "file"
	logx.MustSetup(d)

	fileWriter := logx.Reset()
	writer, err := NewMultiWriter(fileWriter)
	logx.Must(err)
	logx.SetWriter(writer)

	logx.Infow("infow foo",
		logx.Field("url", "http://localhost:8080/hello"),
		logx.Field("attempt", 3),
		logx.Field("backoff", time.Second),
	)

}
