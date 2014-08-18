package proxy

import (
	"fmt"
	"github.com/siddontang/mixer/sqlparser"
	"strings"
)

func (c *Conn) handleAdmin(admin *sqlparser.Admin) error {
	name := string(admin.Name)

	var err error
	switch strings.ToLower(name) {
	case "upnode":
		err = c.adminUpNodeServer(admin.Values)
	case "downnode":
		err = c.adminDownNodeServer(admin.Values)
	default:
		return fmt.Errorf("admin %s not supported now", name)
	}

	if err != nil {
		return err
	}

	return c.writeOK(nil)
}

func (c *Conn) adminUpNodeServer(values sqlparser.ValExprs) error {
	if len(values) != 3 {
		return fmt.Errorf("upnode needs 3 args, not %d", len(values))
	}

	nodeName := nstring(values[0])
	sType := strings.ToLower(nstring(values[1]))
	addr := strings.ToLower(nstring(values[2]))

	switch sType {
	case Master:
		return c.server.UpMaster(nodeName, addr)
	case MasterBackup:
		return c.server.UpMasterBackup(nodeName, addr)
	case Slave:
		return c.server.UpSlave(nodeName, addr)
	default:
		return fmt.Errorf("invalid server type %s", sType)
	}
}

func (c *Conn) adminDownNodeServer(values sqlparser.ValExprs) error {
	if len(values) != 2 {
		return fmt.Errorf("upnode needs 2 args, not %d", len(values))
	}

	nodeName := nstring(values[0])
	sType := strings.ToLower(nstring(values[1]))

	switch sType {
	case Master:
		return c.server.DownMaster(nodeName)
	case MasterBackup:
		return c.server.DownMasterBackup(nodeName)
	case Slave:
		return c.server.DownSlave(nodeName)
	default:
		return fmt.Errorf("invalid server type %s", sType)
	}
}
