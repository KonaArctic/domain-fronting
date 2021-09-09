[Akamai](https://akamai.com/)
------
Supports fronting
```
kona@kona:~$ curl --fail --http1.1 --silent https://pinterest.com/ --header "Host: www.nytimes.com" | grep -m 1 -oEe "New York Times ......"
New York Times - Brea
```
And websocket

[Azure](https://azure.microsoft.com/en-us/services/cdn/)
-----
Probably works
It's the CDN of choice for Tor Meek

[BelugaCDN](https://belugacdn.com/)
---------
Supports fronting
```
kona@kona:~$ curl --fail --http1.1 --silent https://www.comodo.com/ --header "Host: www.surjen.com" | grep --fixed-string --regexp-"<title>" 
<title>Book Doctor Consultation | Ambulance Services | Lab Tests in Nigeria</title>
```
And claims to support websockets
Also cheap ($0.008/Gb)

[BunnyCDN](https://bunny.net/)
--------
Does not support domain fronting
```
kona@kona:~$ curl --fail --http1.1 --silent https://filecoin.io/ --header "Host: tabletmag.com" | grep --fixed-string --regexp-"<title>" 
<html><head> <link href-"http://fonts.googleapis.com/css?family-Rubik:300,400,500,700,900" rel-"stylesheet" type-"text/css"> <link rel-"stylesheet" href-"https://bunnycdn.b-cdn.net/assets/landingpage/css/unconfigured.css"> <title>BunnyCDN Node ES1-703</title></head><body> <div id-"header"> <a href-"https://bunny.net"><img style-"vertical-align:middle; image-rendering: -webkit-optimize-contrast;" src-"https://bunny.net/v2/images/bunnynet-logo.svg"/></a></div><div id-"content"> <h2>Domain suspended or not configured</h2> <p>If you are the administrator and believe this is an error on our side, please check your BunnyCDN account configuration or contact customer support. </p></div></body></html>
```

[CacheFly](https://cachefly.com/)
--------
Min. commitment >$500/mo

[CDN77](https://cdn77.com)
-----
Supports fronting but not websocket
```
kona@kona:~$ curl --fail --http1.1 --silent https://www.datapacket.com/ --header "Host: www.spcdn.org" | grep --fixed-string --regexp-"<title>" 
	<title>Multi-Channel Marketing Automation Platform | SendPulse</title>
```

[CDNetworks](https://cdnetworks.com)
----------
Does not support domain fronting
```
kona@kona:~$ curl --fail --http1.1 https://www.tetrapak.com/ --header "Host: toshiba.com" 
curl: (56) OpenSSL SSL_read: Connection reset by peer, errno 104
```

[CDNVideo](https://cdnvideo.com/)
--------
Does not support domain fronting
```
kona@kona:~$ curl --fail --http1.1 https://rtlab.ru/ --header "Host: gpucloud.ru" 
curl: (22) The requested URL returned error: 503 Service Temporarily Unavailable
```

[ChinaCache](https://en.chinacache.com/)
----------
Supports fronting
```
kona@kona:~$ curl --fail --http1.1 --silent https://hksuning.com/ --header "Host: www.suning.com" | grep --fixed-string --regexp-"<title>" 
<title>苏宁易购(Suning.com)-专注好服务、送货更准时、价格更超值、上新货更快</title>
```

[Cloudflare](https://cloudflare.com/)
----------
[Does not support domain fronting](https://news.ycombinator.com/item?id=9234367)

[CloudFront](https://aws.amazon.com/cloudfront/)
----------
[Does not support domain fronting](https://aws.amazon.com/blogs/networking-and-content-delivery/continually-enhancing-domain-security-on-amazon-cloudfront/)

[Fastly](https://fastly.com/)
------
Supports fronting
```
kona@kona:~$ curl --fail --http1.1 --silent https://cnn.com/ --header "Host: www.bbc.com" | grep --fixed-string --regexp-"<title>" 
        <title>BBC - Homepage</title>
```
But [not websocket](https://fastly.com/blog/server-sent-events-fastly)(still applies last time I checked)

[G-Core Labs](https://gcorelabs.com/cdn/)
-----------
Supports fronting
```
kona@kona:~$ curl --fail --http1.1 --silent https://www.snh48.com/ --header "Host: www.carrotquest.io" | grep --fixed-string --regexp-"<title>" 
	<title>Carrot quest — увеличение продаж на сайте на текущем трафике</title>
```
Also cheap (1Tb/mo free + €0.0025/Gb)
TODO: Check websocket

[Imperva](https://imperva.com/)
-------
Supports fronting
```
kona@kona:~$ curl --fail --http1.1 --silent https://www.cox.com/ --header "Host: www.toyota.com" | grep --fixed-string --regexp-"<title>" 
  <title>New Cars, Trucks, SUVs &amp; Hybrids | Toyota Official Site</title> 
```
[And websocket](https://docs.imperva.com/bundle/cloud-application-security/page/websocket.htm)
Quote needed for pricing

[Kingsoft](https://en.ksyun.com/)
--------
Seems not to support domain fronting
```
kona@kona:~$ curl --fail --http1.1 https://www.wps.cn/ --header "Host: www.bilibili.com" 
curl: (22) The requested URL returned error: 404 Not Found
```

[LeaseWeb](https://leaseweb.com/cdn)
--------
Does not support domain fronting
```
kona@kona:~$ curl --fail --http1.1 https://simpletextadz.com/ --header "Host: thetopad.com" 
curl: (22) The requested URL returned error: 421 Misdirected Request
```

[Limelight](https://limelight.com/)
---------
Supports fronting
```
kona@kona:~$ curl --fail --http1.1 --silent https://www.marvel.com/ --header "Host: www.nintendo.co.uk" | grep --fixed-string --regexp="<title>" 
	<title>Nintendo UK's official site</title>
```
IDK if they support websocket
Quote needed for pricing

[Lumen](https://lumen.com/en-us/networking/cdn.html)
-----
TODO

[Medianova](https://medianova.com/)
---------
Does not support domain fronting
```
kona@kona:~$ curl --fail --http1.1 --silent https://seemeet.live/ --header "Host: kodris.co.uk" | grep --fixed-string --regexp="<title>" 
	<title>SeeMeet</title>
```

[StackPath](https://stackpath.com/products/cdn/)
---------
Supports fronting
```
kona@kona:~$ curl --fail --http1.1 --silent https://mysqltutorial.org/ --header "Host: postgresqltutorial.com" | grep -E --regexp="<title>.........." -m 1 -o 
<title>PostgreSQL
```
[And websocket](https://support.stackpath.com/hc/en-us/articles/360030984211)
$0.03/Gb

[Tata Communications](https://tatacommunications.com/solutions/content-delivery-network/)
------------------
Supports fronting
```
kona@kona:~$ curl --fail --http1.1 --silent https://nasscom.in/ --header "Host: licindia.in" | grep --fixed-string --regexp="<title>" -A 1
<head id="head"><title>
	Life Insurance Corporation of India - Home
```
IDK if they support websockets
Quote needed for pricing

[Tencent Cloud](https://intl.cloud.tencent.com/products/cdn)
-------------
Seems not to support domain fronting
```
kona@kona:~$ curl --fail --http1.1 https://g.58.com/ --header "Host: www.sohu.com"
curl: (22) The requested URL returned error: 508 
```


