package response

import (
	"io"
	"fmt"
)

type Header struct {
	Key ,Value string
}
type Status struct {
	Code int
	Reason string
}
//通过重写了 printf 中的write 接口，将err 包装为errwrite 中内部的对象，
// 实现了代码的简洁，显式的err 减少
func WriteResponse2 (w io.Writer,st Status,
	headers []Header,body io.Reader) error{
	ew := &errWriter{Writer:w}
	fmt.Fprintf(ew,"HTTP/1.1 %d %s \r\n",st.Code,
			st.Reason,
		)

	for _,h := range headers {
		fmt.Fprintf(w, "%s:%s\r\n",h.Key,h.Value)
	}
	fmt.Fprint(w,"\r\n");

	io.Copy(w,body)
	return ew.err
}
func WriteResponse (w io.Writer,st Status,
	headers []Header,body io.Reader) error{
	_,err := fmt.Fprintf(w,"HTTP/1.1 %d %s \r\n",st.Code,
		st.Reason,
	)
	if err != nil {
		return err
	}
	for _,h := range headers {
		_, err := fmt.Fprintf(w, "%s:%s\r\n",h.Key,h.Value)
		if err != nil {
			return err
		}
	}
	if _,err := fmt.Fprint(w,"\r\n"); err != nil{
		return err
	}
	_, err = io.Copy(w,body)
	return err
}

type errWriter struct{
	io.Writer
	err error
}
//重写了 printf 中的 write 接口
func (e *errWriter) Write(buf []byte)(int, error){
	if e.err != nil {
		return 0,e.err
	}
	var n int
	n,e.err = e.Writer.Write(buf)
	return n,nil
}
