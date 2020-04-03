package gdocs

import (
	"context"
	"errors"
	"net"
	"net/http"
	"strconv"

	"github.com/spf13/viper"
)

var (
	viperConfig *viper.Viper
	port        = 8000
	server      = &http.Server{}
)

func handleToken(w http.ResponseWriter, r *http.Request) {
	viperConfig.Set("code", r.FormValue("code"))
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
