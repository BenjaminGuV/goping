package main

import (
	"fmt"
	"io/ioutil"
	"os/exec"
	"strings"
	"time"
	"bufio"
	"log"
	"runtime"
	"syscall"

	"github.com/getlantern/systray"
	"github.com/gen2brain/beeep"
)

var contador = 0;
var ban_chg  = 1;

func main()  {
	systray.Run( onReady, onExit )
}

func onReady()  {
	

	systray.SetIcon( getIcon("assets/offline.ico") )
	systray.SetTitle( "PING MONITORY" )
	systray.SetTooltip( "PING MONITORY" )

	mQuit := systray.AddMenuItem("Quit", "Quits this app")

	println("START")

	go func() {
        for {
            select {
            case <-mQuit.ClickedCh:
                systray.Quit()
                return
            }
        }
	}()
	
	//get time every 3 seconds
	uptimeTicker := time.NewTicker(3 * time.Second)

	for {
		select {
		case <-uptimeTicker.C:
			sendPing()
		}
	}

}

func onExit()  {
	//clean
}

func getIcon(s string) []byte {
	b, e := ioutil.ReadFile( s )
	if e != nil{
		fmt.Println("Error: ")
		fmt.Print(e)
	}

	return b
}

func sendPing()  {
	cmd := exec.Command("ping", "1.1.1.1")

	if runtime.GOOS == "windows" {
        cmd = exec.Command("ping", "1.1.1.1", "-t")
	}
	
	cmd.SysProcAttr = &syscall.SysProcAttr{HideWindow: true}
	
	stdoutIn, _ := cmd.StdoutPipe()
	stderrIn, _ := cmd.StderrPipe()
	
	err := cmd.Start()
	if err != nil {
		log.Fatalf("cmd.Start() failed with '%s'\n", err)
		log.Fatalf("cmd.Start() failed with '%s'\n", stderrIn)
	}


	//temp
	reader := bufio.NewReader(stdoutIn)
	line, err := reader.ReadString('\n')

	temp_ban := 0

	for err == nil {
		contador++
		fmt.Println( contador )
		fmt.Println("temporal")
		fmt.Println(line)
		

		temp_ban = messageCMD( line )

		fmt.Println("banderas")
		fmt.Println(temp_ban)
		fmt.Println(ban_chg)

		if temp_ban != ban_chg {
			fmt.Println("bandera uno")
			notifyCMD( temp_ban )
			ban_chg = temp_ban
		}else{
			fmt.Println("bandera dos")
		}

		line, err = reader.ReadString('\n')
	}


}


func messageCMD( lines string ) ( x int ) {
	out := ""
	out = string( lines )

	fmt.Println( "hey" )
	fmt.Println( out )

	//only message valid
	if strings.Contains( out, "bytes of data") || strings.Contains( out, "64 bytes from 1.1.1.1" ) || strings.Contains( out, "Haciendo ping a 1.1.1.1" ) || strings.Contains( out, "Respuesta desde 1.1.1.1" ) {
		fmt.Println("IT'S ALIVEEE")
		systray.SetIcon( getIcon("assets/active.ico") )
		x = 1
	} else if strings.Contains( out, "Destination Host Unreachable") || strings.Contains( out, "Host de destino inaccesible") {
		// Host de destino inaccesible
		fmt.Println("TANGO DOWN HOST UNREACHABLE")
		systray.SetIcon( getIcon("assets/fail.ico") )
		x = 2
	} else if strings.Contains( out, "privilegios") {
		fmt.Println("TANGO DOWN PRIVILEGIOS")
		systray.SetIcon( getIcon("assets/fail.ico") )
		x = 2
	} else if strings.Contains( out, "Se debe suministrar un valor") {
		fmt.Println("TANGO DOWN OPCION")
		systray.SetIcon( getIcon("assets/fail.ico") )
		x = 2
	} else if strings.Contains( out, "Error") {
		fmt.Println("TANGO DOWN ERROR")
		systray.SetIcon( getIcon("assets/fail.ico") )
		x = 2
	} else {
		fmt.Println("DESCONOCIDO")
		systray.SetIcon( getIcon("assets/fail.ico") )
		x = 3
	}

	return x
}

func notifyCMD( flags_ban int )  {
	
	if flags_ban == 1 {
		err := beeep.Notify("GO PING", "OK", "assets/ok.png")
		if err != nil {
			panic(err)
		}
	}

	if flags_ban == 2 {
		err := beeep.Notify("GO PING", "FAIL", "assets/fail.png")
		if err != nil {
			panic(err)
		}
	}

	if flags_ban == 3 {
		err := beeep.Notify("GO PING", "FAIL", "assets/fail.png")
		if err != nil {
			panic(err)
		}
	}

}