package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"os/exec"
	"os/signal"
	"syscall"
	"time"
)

func findLineEnd(dat []byte) (out [][]byte) {
	prev := 0;
	for pos,i := range dat {
		if i == []byte("\n")[0] {
			out = append(out, dat[prev:pos]);	
			prev = pos + 1;
		}
	}

	out = append(out, dat[prev:])
	// for pos,i := range out {
	// 	count := 0;
	// 	for _,char := range i {
	// 		if (char == []byte(" ")[0]) {
	// 			count++;
	// 		}
	// 	}
	// 	if count == len(i)  && pos > 0{
	// 		out = append(out[:pos-1],out[pos:]...)
	// 	}
	// }
	return;
}

func copyAndCapture(w io.Writer, r io.Reader) {

	prefix := []byte("Child process: ")
	after  := []byte("\n")
    buf := make([]byte, 1024, 1024)
    for {
        n, err := r.Read(buf[:])
        if n > 0 {
            d := buf[:n]
			lines := findLineEnd(d);
			for _,line := range lines {
				out := append(prefix,line...)
				out = append(out,after...)

				_, err := w.Write(out)
				if err != nil {
					return;
				}
			}
        }
        if err != nil {
            // Read returns io.EOF at the end of file, which is not an error for us
            if err == io.EOF {
                err = nil
            }
                return;
        }
    }
}

func HandleProcess(cmd *exec.Cmd) {
	if cmd == nil {
		return;
	}

	stdoutIn, _ := cmd.StdoutPipe()
	stderrIn, _ := cmd.StderrPipe()
	cmd.Start();
	go func ()  {
		copyAndCapture(os.Stdout, stdoutIn)
	}();
	go func ()  {
		copyAndCapture(os.Stdout, stderrIn)
	}();
	cmd.Wait();

}

func main() {
	engine 	:= "gstreamer"
	envstr  := "dev"
	name 	:= "local"

	args := os.Args[1:]
	for i, arg := range args {
		if arg == "--engine" {
			engine = args[i+1]
		} else if arg == "--env" {
			envstr = args[i+1]
		} else if arg == "--name" {
			name = args[i+1]
		} else if arg == "--help" {
			fmt.Printf("--engine |  encode engine ()\n")
			return
		}
	}

	command := func () string {
		switch engine{
		case "gstreamer":
			return "webrtc-proxy.exe"
		case "screencoder":
			return "screencoder.exe"
		default:
			return "unknown"
		}	
	}()




	
	shutdown := make(chan bool)
	var proxy,devsim *exec.Cmd;
	go func ()  {
		chann := make( chan os.Signal) 
		signal.Notify( chann, syscall.SIGTERM , os.Interrupt);
		<-chann;

		if proxy != nil {
			proxy.Process.Kill()	
		}
		if devsim != nil {
			devsim.Process.Kill()	
		}
		shutdown<-true;
	}()


	go func ()  {
		for {
			time.Sleep(time.Second * 5)
			resp,err := http.Get(fmt.Sprintf("https://auth.thinkmay.net/auth/server/%s",name))
			if err != nil{
				fmt.Printf("%s\n",err.Error());
				continue;
			}

			body := make([]byte,1000);	
			size,err := resp.Body.Read(body);
			if err != nil{
				fmt.Printf("%s\n",err.Error());
				continue;
			}
			str := string(body[:size]);
			if str == "none" {
				fmt.Printf("empty token\n");
				continue;
			}

			fmt.Printf("starting webrtc proxy\n");
			proxy = exec.Command(command,"--token",str,"--env",envstr);
			HandleProcess(proxy);
			time.Sleep(2 * time.Second);
		}
	} ();


	if envstr == "dev" {
	go func ()  {
		for {
			fmt.Printf("starting devsim\n");
			devsim = exec.Command("DevSim.exe");
			HandleProcess(devsim);
			time.Sleep(2 * time.Second);
		}	
	} ()
		
	}

	<-shutdown
}
