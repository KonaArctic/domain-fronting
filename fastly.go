// Public Access. Copyright (c) 2021 Kona Arctic. All rights reserved. ABSOLUTELY NO WARRANTY! https://akona.me mailto:arcticjieer@gmail.com
// Fronting via Fastly
package fronting
import "io"
import "strings"

func Fastly( cookie string )( io.ReadWriteCloser , error ) {
	var outer string = "cdn.reddit.com"
	if len( strings.Split( cookie , "/" ) ) == 2 {
		outer = strings.Split( cookie , "/" )[ 1 ] }
	return front( outer , strings.Split( cookie , "/" )[ 0 ] )
}

