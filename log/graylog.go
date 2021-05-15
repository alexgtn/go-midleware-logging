package log

import (
	"fmt"
	"net"
)

type GrayLogLogger struct {
	conn net.Conn
}

func NewGraylogLogger(tcpConn net.Conn) *GrayLogLogger {
	return &GrayLogLogger{
		conn: tcpConn,
	}
}

func (l *GrayLogLogger) Infof(format string, args ...interface{})  {
	l.conn.Write([]byte(fmt.Sprintf(format, args...) + "\n"))
}