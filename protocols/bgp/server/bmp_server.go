package server

import (
	"fmt"
	"io"
	"net"
	"sync"
	"time"

	"github.com/bio-routing/bio-rd/protocols/bgp/packet"
	"github.com/bio-routing/bio-rd/routingtable/locRIB"
	log "github.com/sirupsen/logrus"
	"github.com/taktv6/tflow2/convert"
)

const (
	defaultBufferLen = 4096
)

// BMPServer represents a BMP server
type BMPServer struct {
	routers       map[string]*router
	routersMu     sync.RWMutex
	reconnectTime uint
}

// NewServer creates a new BMP server
func NewServer() *BMPServer {
	return &BMPServer{
		routers: make(map[string]*router),
	}
}

// AddRouter adds a router to which we connect with BMP
func (b *BMPServer) AddRouter(addr net.IP, port uint16, rib4 *locRIB.LocRIB, rib6 *locRIB.LocRIB) {
	r := &router{
		address:          addr,
		port:             port,
		reconnectTimeMin: 30,  // Suggested by RFC 7854
		reconnectTimeMax: 720, // Suggested by RFC 7854
		reconnectTimer:   time.NewTimer(time.Duration(0)),
		rib4:             rib4,
		rib6:             rib6,
		neighbors:        make(map[[16]byte]*neighbor),
	}

	b.routersMu.Lock()
	b.routers[fmt.Sprintf("%s:%d", r.address.String(), r.port)] = r
	b.routersMu.Unlock()

	go func(r *router) {
		for {
			<-r.reconnectTimer.C
			c, err := net.Dial("tcp", fmt.Sprintf("%s:%d", r.address.String(), r.port))
			if err != nil {
				log.Infof("Unable to connect to BMP router: %v", err)
				if r.reconnectTime == 0 {
					r.reconnectTime = r.reconnectTimeMin
				} else if r.reconnectTime < r.reconnectTimeMax {
					r.reconnectTime *= 2
				}
				r.reconnectTimer = time.NewTimer(time.Second * time.Duration(r.reconnectTime))
				continue
			}

			r.reconnectTime = 0
			r.con = c
			log.Infof("Connected to %s", r.address.String())
			r.serve()
		}
	}(r)
}

func recvBMPMsg(c net.Conn) (msg []byte, err error) {
	buffer := make([]byte, defaultBufferLen)
	_, err = io.ReadFull(c, buffer[0:packet.MinLen])
	if err != nil {
		return nil, fmt.Errorf("Read failed: %v", err)
	}

	l := convert.Uint32b(buffer[1:5])
	if l > defaultBufferLen {
		tmp := buffer
		buffer = make([]byte, l)
		copy(buffer, tmp)
	}

	toRead := l
	_, err = io.ReadFull(c, buffer[packet.MinLen:toRead])
	if err != nil {
		return nil, fmt.Errorf("Read failed: %v", err)
	}

	return buffer[0:toRead], nil
}
