// Public Access. Copyright (c) 2021 Kona Arctic. All rights reserved. ABSOLUTELY NO WARRANTY! https://akona.me mailto:arcticjieer@gmail.com
// F̶r̶o̶n̶t̶i̶n̶g̶ ̶v̶i̶a̶ ̶F̶a̶s̶t̶l̶y̶ Turns out Fastly doesnt support the HTTP/1.1 protocol upgrade mechanism required by Public Access
package fronting
import "io"
import "strings"

func Fastly( cookie string )( io.ReadWriteCloser , error ) {
	var outer string = "cdn.reddit.com"
	if len( strings.Split( cookie , "/" ) ) == 2 {
		outer = strings.Split( cookie , "/" )[ 1 ] }
	return Front( outer , strings.Split( cookie , "/" )[ 0 ] )
}

