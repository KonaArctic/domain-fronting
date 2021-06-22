package fronting
import "bufio"
import "io"
import "net"
import "net/http"
import "time"
import "os"

type any interface{ }

// https://stackoverflow.com/a/50346080/10958912
func mu( one ... any ) [ ]any {
	return one }

// https://stackoverflow.com/questions/26545883
func ternary( test bool, one any , two any ) any {
	if test {
		return one
	} else {
		return two } }

// Convert io.ReadWriteCloser to fake net.Conn
type fakeaddr struct{ }
func ( self fakeaddr ) Network( ) string {
	return "fakeaddr" }
func ( self fakeaddr ) String( ) string {
	return "fakeaddr" }
type fakeconn struct{
	io.ReadWriteCloser }
func ( self fakeconn ) LocalAddr( ) net.Addr {
	var fakeaddr fakeaddr
	return fakeaddr }
func ( self fakeconn ) RemoteAddr( ) net.Addr {
	var fakeaddr fakeaddr
	return fakeaddr }
func ( self fakeconn ) SetDeadline( time time.Time ) error {
	return nil }
func ( self fakeconn ) SetReadDeadline( time time.Time ) error {
	return nil }
func ( self fakeconn ) SetWriteDeadline( time time.Time ) error {
	return nil }
func wrapconn( stream io.ReadWriteCloser ) net.Conn {
	return fakeconn{ stream }
}

// Rewrite HTTP Host header
func httpNewHost( socket io.ReadWriteCloser , inner string ) io.ReadWriteCloser {
	var stream , pipenet net.Conn
	var request http.Request
	stream , pipenet = net.Pipe( )
	go func( ) {
		request = * mu( http.ReadRequest( bufio.NewReader( pipenet ) ) )[ 0 ].( * http.Request )
		request.Host = inner
		request.Write( os.Stderr )
		request.Write( socket )
		go io.Copy( socket , pipenet )
		go io.Copy( pipenet , socket )
	}( )
	return stream
}

