# Network

This file contains the docs for the network pkg, and its releavent info.


## Request 204




## Check Portal

This function is used to check if the portal is still connected or not, It checks `204` response for now. 

### Custom client
when Requesting a `204` if the login has timedout, the router will redirect the user request to the `captive portal` login page and that will return status ok instead of 304. To not let this happen we have to tell the go to not follow the redirect link, using `http.ErrUseLastResponse` Thus we need a custom client instead of go default http client 

### Location
When the captive portal redirects to login, It give the url in `location` header, using that, we can define the captive portal for internel pinging, and then auto re-logginin. 


## Internal ping

The internal ping will work like this, After the `check portal` checks for the location header, and returns the http url for redirect, The internal ping will use `time.now`  to measure the amount of time it took to open the webpage, The internal ping will be used to find which is the login page, General if you open the url in my case. 
`172.11.0.1:8090/` It redirects to login page. `/172.11.0.1:8090/login.xml` that will be helpful as the auto reconnect script.



## External ping

The external ping uses cmd.exec for executing ping command in the native terminal. Sometimes the captive portal intercept the data and returns a `200`, the native ping command sends a raw byte. The raw byte is more accurate then `http.get`