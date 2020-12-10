package main

import (
	"github.com/hedahong/kcp-go/v6"
	"log"
	"strconv"
	"time"
)

func main() {

	//key := pbkdf2.Key([]byte("demo pass"), []byte("demo salt"), 1024, 32, sha1.New)
	//block, _ := kcp.NewAESBlockCrypt(key)
	block, _ := kcp.NewNoneBlockCrypt([]byte("none encrypt"))
	if listener, err := kcp.ListenWithOptions("0.0.0.0:12345", block, 10, 3); err == nil {
		// spin-up the client
		go client()
		go client2()
		for {
			s, err := listener.AcceptKCP()
			if err != nil {
				log.Fatal(err)
			}
			go handleEcho(s)
		}
	} else {
		log.Fatal(err)
	}
}

// handleEcho send back everything it received
func handleEcho(conn *kcp.UDPSession) {
	buf := make([]byte, 4096)
	for {
		n, err := conn.Read(buf)
		if err != nil {
			log.Println(err)
			return
		}

		n, err = conn.Write(buf[:n])
		if err != nil {
			log.Println(err)
			return
		}
	}
}

func client() {
	//key := pbkdf2.Key([]byte("demo pass"), []byte("demo salt"), 1024, 32, sha1.New)
	//block, _ := kcp.NewAESBlockCrypt(key)
	block, _ := kcp.NewNoneBlockCrypt([]byte("none encrypt"))

	// wait for server to become ready
	time.Sleep(time.Second)

	// dial to the echo server
	if sess, err := kcp.DialWithOptions2("10.168.12.26:12345","10.168.224.223:0", block, 10, 3); err == nil {

		go func() {
			buf := make([]byte, 32*1024)
			for{
				// read back the data

				if rn, err := sess.Read(buf); err == nil {
					log.Println("recv:", string(buf[0:rn]))
				} else {
					log.Fatal(err)
				}
			}

		}()


		for i:=0;i<999;i++{
			data := time.Now().String()+" # "+strconv.Itoa(i)
			//buf := make([]byte, len(data))
			log.Println("sent:", data)
			if _, err := sess.Write([]byte(data)); err == nil {
				// read back the data
				//if _, err := io.ReadFull(sess, buf); err == nil {
				//	log.Println("recv:", string(buf))
				//} else {
				//	log.Fatal(err)
				//}
			} else {
				log.Fatal(err)
			}
			time.Sleep(time.Millisecond*1000)
		}

		<-make(chan int)
	} else {
		log.Fatal(err)
	}
}


func client2() {
	//key := pbkdf2.Key([]byte("demo pass"), []byte("demo salt"), 1024, 32, sha1.New)
	//block, _ := kcp.NewAESBlockCrypt(key)
	block, _ := kcp.NewNoneBlockCrypt([]byte("none encrypt"))

	// wait for server to become ready
	time.Sleep(time.Second)

	// dial to the echo server
	if sess, err := kcp.DialWithOptions2("10.168.224.223:12345","10.168.12.26:0", block, 10, 3); err == nil {

		go func() {
			buf := make([]byte, 32*1024)
			for{
				// read back the data

				if rn, err := sess.Read(buf); err == nil {
					log.Println("recv:", string(buf[0:rn]))
				} else {
					log.Fatal(err)
				}
			}

		}()


		for i:=0;i<999;i++{
			data := time.Now().String()+" $ "+strconv.Itoa(i)
			//buf := make([]byte, len(data))
			log.Println("sent:", data)
			if _, err := sess.Write([]byte(data)); err == nil {
				// read back the data
				//if _, err := io.ReadFull(sess, buf); err == nil {
				//	log.Println("recv:", string(buf))
				//} else {
				//	log.Fatal(err)
				//}
			} else {
				log.Fatal(err)
			}
			time.Sleep(time.Millisecond*1000)
		}

		<-make(chan int)
	} else {
		log.Fatal(err)
	}
}
