package proxmox

import (
	"fmt"
	"github.com/gorilla/websocket"
	"net/http"
	"strings"
)

func (c *Client) makeHeaders() http.Header {
	header := make(http.Header)
	header.Add("Accept", "application/json")
	if c.token != "" {
		header.Add("Authorization", fmt.Sprintf("PVEAPIToken=%s", c.token))
	} else if c.session != nil {
		header.Add("Cookie", fmt.Sprintf("PVEAuthCookie=%s", c.session.Ticket))
		header.Add("CSRFPreventionToken", c.session.CSRFPreventionToken)
	}
	return header
}

func (c *Client) Websocket(node string, vnc *VNC) (*websocket.Conn, error) {
	base := strings.Replace(c.baseURL, "https://", "wss://", 1)
	wssUrl := fmt.Sprintf("%s/nodes/%s/vncwebsocket?port=%d&vncticket=%s", base, node, vnc)
	conn, _, err := websocket.DefaultDialer.Dial(wssUrl, c.makeHeaders())
	if err != nil {
		return nil, err
	}
	if err := conn.WriteMessage(websocket.BinaryMessage, []byte(fmt.Sprintf("%s:%s\n", vnc.User, vnc.Ticket))); err != nil {
		return nil, err
	}
	return conn, nil

}
