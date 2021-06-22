// Public Access. Copyright (c) 2021 Kona Arctic. All rights reserved. ABSOLUTELY NO WARRANTY! https://akona.me mailto:arcticjieer@gmail.com
// Domain fronting
package fronting
import tls "github.com/refraction-networking/utls"
import "io"
import "net"

var roller * tls.Roller

func front( outer string , inner string )( io.ReadWriteCloser , error ) {
	var err error
	var socket net.Conn

	if roller == nil {
		roller , err = tls.NewRoller( ) }
	if err != nil {
		return socket , err }

	socket , err = roller.Dial( "tcp" , outer + ":443" , outer )
	if err != nil {
		return socket , err }

	return httpNewHost( socket , inner ) , nil
}


