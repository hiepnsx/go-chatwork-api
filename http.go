package gochatwork

type http interface {
    Get() []byte
}

// http interface
type httpImp struct {

}

func (h *httpImp) Get() []byte {
    return make([]byte, 0)
}