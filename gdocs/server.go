package gdocs

import (
	"errors"
	"fmt"
	"net"
	"net/http"
	"strconv"

	"github.com/spf13/viper"
	"golang.org/x/net/context"
)

var (
	viperConfig *viper.Viper
	port        = 8000
	server      = &http.Server{}
)

const page = `
<html><head> <style>body{position: relative; height: 100%;}.content{position: absolute; top: 50%; left: 50%; transform: translate(-50%, -50%);}.alert{padding: 8px 14px 8px 14px; margin-bottom: 18px; color: #468847; text-shadow: 0 1px 0 rgba(255, 255, 255, 0.5); background-color: #dff0d8; border: 1px solid #d6e9c6; -webkit-border-radius: 4px; -moz-border-radius: 4px; border-radius: 4px; font-family: sans-serif; text-align: center !important;}h2{margin-top: 0; margin-bottom: 0;}hr{border:0.5px solid #5cb85c;}</style></head><body> <div class="content"> <div class="alert"> <pre>
     _____          ___           ___     
    /  /::\        /  /\         /  /\    
   /  /:/\:\      /  /::\       /  /:/    
  /  /:/  \:\    /  /:/\:\     /  /:/     
 /__/:/ \__\:|  /  /:/  \:\   /  /:/  ___ 
 \  \:\ /  /:/ /__/:/ \__\:\ /__/:/  /  /\
  \  \:\  /:/  \  \:\ /  /:/ \  \:\ /  /:/
   \  \:\/:/    \  \:\  /:/   \  \:\  /:/ 
    \  \::/      \  \:\/:/     \  \:\/:/  
     \__\/        \  \::/       \  \::/   
                   \__\/         \__\/    
</pre> <h2>Authentication Complete</h2> <hr> You can now close this tab. </div></div></body></html>`

func handleToken(w http.ResponseWriter, r *http.Request) {
	viperConfig.Set("code", r.FormValue("code"))

	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	fmt.Fprint(w, page)

	server.Shutdown(context.Background())
}

// startServer starts a server that will accept the user token on redirect after the oauth workflow is complete
func startServer(vp *viper.Viper) error {
	viperConfig = vp
	http.HandleFunc("/code", handleToken)

	if err := getFreePort(); err != nil {
		return err
	}

	server.Addr = ":" + strconv.Itoa(port)

	return server.ListenAndServe()
}

func getFreePort() error {
	var (
		listener net.Listener
		err      error
	)

	for i := 0; i < 10; i++ {
		listener, err = net.Listen("tcp", ":"+strconv.Itoa(port))
		if err == nil {
			listener.Close()
			return nil
		}

		port++
	}

	return errors.New("Failed to find available port")
}
