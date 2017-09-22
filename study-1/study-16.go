package main

import (
	"errors"
	"fmt"
	"io"
	"log"
	"math/rand"
	"strconv"
	"sync"
	"sync/atomic"
	"time"
)

func main() {
	//randFunc()
	//intFunc()
	poolFunc()
}

const (
	//模拟的最大goroutine
	maxGoroutine = 5
	//资源池的大小
	poolRes = 2
)

func poolFunc() {
	var wg sync.WaitGroup
	wg.Add(maxGoroutine)
	p, err := New(createConnection, poolRes)
	if err != nil {
		log.Println(err)
		return
	}
	for query := 0; query < maxGoroutine; query++ {
		go func(q int) {
			dbquery(q, p)
			wg.Done()
		}(query)
	}
	wg.Wait()
	log.Println("开始关闭资源池")
	p.Closed()
}

func dbquery(query int, pool *Pool) {
	conn, err := pool.Acquire()
	if err != nil {
		log.Println(err)
		return
	}
	defer pool.Release(conn)
	//模拟查询
	time.Sleep(time.Duration(rand.Intn(1000)) * time.Millisecond)
	log.Printf("第%d个查询，使用的是ID为%d的数据库连接", query, conn.(*dbConnection).Id)
}

var idCounter int32

type dbConnection struct {
	Id int32
}

//实现io.Closer接口
func (db *dbConnection) Close() error {
	log.Println("关闭连接", db.ID)
	return nil
}

//生成数据库连接的方法，以供资源池使用
func createConnection() (io.Reader, error) {
	id := atomic.AddInt32(&idCounter, 1)
	return &dbConnection{Id: id}, nil
}

type Pool struct {
	m       sync.Mutex
	res     chan io.Closer
	closed  bool
	factory func() (io.Reader, error)
}

var ErrPoolClosed = errors.New("资源池已经被关闭。。。。")

//创建一个资源池
func New(f func() (io.Reader, error), size uint) (*Pool, error) {
	if size <= 0 {
		return nil, errors.New("size 太小了。。。")
	}
	return &Pool{
		factory: f,
		res:     make(io.Closer, size),
	}, nil
}

//从资源池里获取一个资源
func (p *Pool) Acquire() (io.Reader, error) {
	select {
	case r, ok := <-p.res:
		log.Println("Acquire:共享资源")
		if !ok {
			return nil, ErrPoolClosed
		}
	default:
		log.Println("Acquire:新生资源。。。")
		return p.factory()
	}
}

//关闭资源池，释放资源
func (p *Pool) Closed() {
	p.m.Lock()
	defer p.m.Unlock()
	if p.closed {
		return
	}
	p.closed = true
	close(p.res)
	for r := range p.res {
		r.Close()
	}
}

func (p *Pool) Release(r io.Closer) {
	p.m.Lock()
	defer p.m.Unlock()
	if p.closed {
		r.Close()
		return
	}
	select {
	case p.res <- r:
		log.Println("资源释放到池子里了....")
	default:
		log.Println("资源池满了，释放这个资源吧....")
		r.Close()
	}
}

func intFunc() {
	f, _ := strconv.ParseFloat("1.23", 64)
	fmt.Println(f)
	i, _ := strconv.ParseInt("1234", 0, 64)
	fmt.Println(i)
	d, _ := strconv.ParseInt("0x1c8", 0, 64)
	fmt.Println(d)
	u, _ := strconv.ParseUint("789", 0, 64)
	fmt.Println(u)
	k, _ := strconv.Atoi("444")
	fmt.Println(k)
	_, e := strconv.Atoi("heeh")
	fmt.Println(e)
}

func randFunc() {
	var p = fmt.Println
	p(rand.Intn(100), ",")
	p(rand.Intn(100))
	p(rand.Float64())
	p((rand.Float64()*5)+5, ",")
	p((rand.Float64() * 5) + 5)
	s1 := rand.NewSource(time.Now().UnixNano())
	r1 := rand.New(s1)
	p(r1.Intn(100), ",")
	p(r1.Intn(100))
	s2 := rand.NewSource(42)
	r2 := rand.New(s2)
	p(r2.Intn(100), ",")
	p(r2.Intn(100))
	s3 := rand.NewSource(43)
	r3 := rand.New(s3)
	p(r3.Intn(100), ",")
	p(r3.Intn(100))
}
