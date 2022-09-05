package main

import (
	"fmt"
	"net/http"
	"os/exec"
	"os"
	"time"
)

func main() {
	engine := "gstreamer"
	envstr    := "dev"
	id := 0

	args := os.Args[1:]
	for i, arg := range args {
		if arg == "--engine" {
			engine = args[i+1]
		} else if arg == "--env" {
			envstr = args[i+1]
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
	



	go func ()  {
		for {
			time.Sleep(time.Second * 5)
			resp,err := http.Get(fmt.Sprintf("https://auth.thinkmay.net/auth/server/%d",id))
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
			process := exec.Command(command,"--token",str,"--env",envstr);
			process.Start();
			process.Wait();
		}
	} ();


	go func ()  {
		for {
			fmt.Printf("starting devsim\n");
			process := exec.Command("DevSim.exe");
			process.Start();
			process.Wait();
		}	
	} ()

	shutdown := make(chan bool)
	<-shutdown
}
