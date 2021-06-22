// Public Access. Copyright (c) 2021 Kona Arctic. All rights reserved. ABSOLUTELY NO WARRANTY! https://akona.me mailto:arcticjieer@gmail.com
// Domain fronting
package fronting
import tls "github.com/refraction-networking/utls"
import "io"
import "net"

func front( outer string , inner string )( io.ReadWriteCloser , error ) {
	var err error
	var socket net.Conn
	var config tls.Config

	socket , err = net.Dial( "tcp" , outer + ":443" )
	if err != nil {
		return socket , err }

	config.ServerName = outer
	socket = tls.UClient( socket , & config , tls.HelloRandomizedNoALPN )

	return httpNewHost( socket , inner ) , nil
}


