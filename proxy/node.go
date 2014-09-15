package proxy

import (
	"fmt"
	"github.com/siddontang/go-log/log"
	"github.com/siddontang/mixer/client"
	"github.com/siddontang/mixer/config"
	"sync"
	"time"
)

const (
	Master       = "master"
	MasterBackup = "master_backup"
	Slave        = "slave"
)

type Node struct {
	sync.Mutex

	server *Server

	cfg config.NodeConfig

	//running master db
	db *client.DB

	master       *client.DB
	masterBackup *client.DB
	slave        *client.DB

	downAfterNoAlive time.Duration

	lastMasterPing int64
	lastSlavePing  int64
}

func (n *Node) run() {
	//to do
	//1 check connection alive
	//2 check remove mysql server alive

	t := time.NewTicker(3000 * time.Second)
	defer t.Stop()

	n.lastMasterPing = time.Now().Unix()
	n.lastSlavePing = n.lastMasterPing
	for {
		select {
		case <-t.C:
			n.checkMaster()
			n.checkSlave()
		}
	}
}

func (n *Node) String() string {
	return n.cfg.Name
}

func (n *Node) getMasterConn() (*client.SqlConn, error) {
	n.Lock()
	db := n.db
	n.Unlock()

	if db == nil {
		return nil, fmt.Errorf("master is down")
	}

	return db.GetConn()
}

func (n *Node) getSelectConn() (*client.SqlConn, error) {
	var db *client.DB

	n.Lock()
	if n.cfg.RWSplit && n.slave != nil {
		db = n.slave
	} else {
		db = n.db
	}
	n.Unlock()

	if db == nil {
		return nil, fmt.Errorf("no alive mysql server")
	}

	return db.GetConn()
}

func (n *Node) checkMaster() {
	n.Lock()
	db := n.db
	n.Unlock()

	if db == nil {
		log.Info("no master avaliable")
		return
	}

	if err := db.Ping(); err != nil {
		log.Error("%s ping master %s error %s", n, db.Addr(), err.Error())
	} else {
		n.lastMasterPing = time.Now().Unix()
		return
	}

	if int64(n.downAfterNoAlive) > 0 && time.Now().Unix()-n.lastMasterPing > int64(n.downAfterNoAlive) {
		log.Error("%s down master db %s", n, n.master.Addr())

		n.downMater()
	}
}

func (n *Node) checkSlave() {
	if n.slave == nil {
		return
	}

	db := n.slave
	if err := db.Ping(); err != nil {
		log.Error("%s ping slave %s error %s", n, db.Addr(), err.Error())
	} else {
		n.lastSlavePing = time.Now().Unix()
	}

	if int64(n.downAfterNoAlive) > 0 && time.Now().Unix()-n.lastSlavePing > int64(n.downAfterNoAlive) {
		log.Error("%s slave db %s not alive over %ds, down it",
			n, db.Addr(), int64(n.downAfterNoAlive/time.Second))

		n.downSlave()
	}
}

func (n *Node) openDB(addr string) (*client.DB, error) {
	db, err := client.Open(addr, n.cfg.User, n.cfg.Password, "")
	if err != nil {
		return nil, err
	}

	db.SetMaxIdleConnNum(n.cfg.IdleConns)
	return db, nil
}

func (n *Node) checkUpDB(addr string) (*client.DB, error) {
	db, err := n.openDB(addr)
	if err != nil {
		return nil, err
	}

	if err := db.Ping(); err != nil {
		db.Close()
		return nil, err
	}

	return db, nil
}

func (n *Node) upMaster(addr string) error {
	n.Lock()
	if n.master != nil {
		n.Unlock()
		return fmt.Errorf("%s master must be down first", n)
	}
	n.Unlock()

	db, err := n.checkUpDB(addr)
	if err != nil {
		return err
	}

	n.Lock()
	n.master = db
	n.db = db
	n.Unlock()

	return nil
}

func (n *Node) upMasterBackup(addr string) error {
	n.Lock()
	if n.masterBackup != nil {
		n.Unlock()
		return fmt.Errorf("%s master backup must be down first", n)
	}
	n.Unlock()

	db, err := n.checkUpDB(addr)
	if err != nil {
		return err
	}

	n.Lock()
	n.masterBackup = db
	n.Unlock()

	return nil
}

func (n *Node) upSlave(addr string) error {
	n.Lock()
	if n.slave != nil {
		n.Unlock()
		return fmt.Errorf("%s, slave must be down first", n)
	}
	n.Unlock()

	db, err := n.checkUpDB(addr)
	if err != nil {
		return err
	}

	n.Lock()
	n.slave = db
	n.Unlock()

	return nil
}

func (n *Node) downMater() error {
	n.Lock()
	db := n.master
	if n.master != nil {
		n.master = nil
	}

	//switch db if exists
	n.db = n.masterBackup

	n.Unlock()

	if db != nil {
		db.Close()
	}
	return nil
}

func (n *Node) downMaterBackup() error {
	n.Lock()

	db := n.masterBackup

	n.masterBackup = nil
	n.db = n.master

	n.Unlock()

	if db != nil {
		db.Close()
	}
	return nil
}

func (n *Node) downSlave() error {
	n.Lock()
	db := n.slave
	n.slave = nil
	n.Unlock()

	if db != nil {
		db.Close()
	}

	return nil
}

// Let node use master if using backup before
func (s *Server) UpMaster(node string, addr string) error {
	n := s.getNode(node)
	if n == nil {
		return fmt.Errorf("invalid node %s", node)
	}

	return n.upMaster(addr)
}

func (s *Server) UpSlave(node string, addr string) error {
	n := s.getNode(node)
	if n == nil {
		return fmt.Errorf("invalid node %s", node)
	}

	return n.upSlave(addr)
}

func (s *Server) UpMasterBackup(node string, addr string) error {
	n := s.getNode(node)
	if n == nil {
		return fmt.Errorf("invalid node %s", node)
	}
	return n.upMasterBackup(addr)
}

func (s *Server) DownMaster(node string) error {
	n := s.getNode(node)
	if n == nil {
		return fmt.Errorf("invalid node %s", node)
	}
	return n.downMater()
}

func (s *Server) DownMasterBackup(node string) error {
	n := s.getNode(node)
	if n == nil {
		return fmt.Errorf("invalid node [%s].", node)
	}
	return n.downMaterBackup()
}

func (s *Server) DownSlave(node string) error {
	n := s.getNode(node)
	if n == nil {
		return fmt.Errorf("invalid node [%s].", node)
	}
	return n.downSlave()
}

func (s *Server) getNode(name string) *Node {
	return s.nodes[name]
}

func (s *Server) parseNodes() error {
	cfg := s.cfg
	s.nodes = make(map[string]*Node, len(cfg.Nodes))

	for _, v := range cfg.Nodes {
		if _, ok := s.nodes[v.Name]; ok {
			return fmt.Errorf("duplicate node [%s].", v.Name)
		}

		n, err := s.parseNode(v)
		if err != nil {
			return err
		}

		s.nodes[v.Name] = n
	}

	return nil
}

func (s *Server) parseNode(cfg config.NodeConfig) (*Node, error) {
	n := new(Node)
	n.server = s
	n.cfg = cfg

	n.downAfterNoAlive = time.Duration(cfg.DownAfterNoAlive) * time.Second

	if len(cfg.Master) == 0 {
		return nil, fmt.Errorf("must setting master MySQL node.")
	}

	var err error
	if n.master, err = n.openDB(cfg.Master); err != nil {
		return nil, err
	}

	n.db = n.master

	if len(cfg.MasterBackup) > 0 {
		if n.masterBackup, err = n.openDB(cfg.MasterBackup); err != nil {
			log.Error(err.Error())
			n.masterBackup = nil
		}
	}

	if len(cfg.Slave) > 0 {
		if n.slave, err = n.openDB(cfg.Slave); err != nil {
			log.Error(err.Error())
			n.slave = nil
		}
	}

	go n.run()

	return n, nil
}
