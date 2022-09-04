package main

import (
	"fmt"
	"net/http"
	"os/exec"
	"time"
)

func main() {
	id := 0

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
			process := exec.Command("webrtc-proxy.exe","--token",str);
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
}
