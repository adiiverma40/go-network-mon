# Data collection

This docs contains the explantion and my choices for data collection algo, how i made it and what and why i choose that.

The Following project will majorly Be for **Linux**. I dont have a windows machine, neither do i plan to use it, So i wont support it.


## Decleartion
- I am very bad at remebering the spellings. If you see any spelling mistakes, ignore them, or if you want u can correct them, but i wont suggest you to waste your time correcting my spellings
- `Manual speedtest` : manual speed test means the speed test where the program will do the speedtest to some external API,


## Abbreviation
- `no.` : Number
- `btw`: Between

## Network stats

I was considering just running a network test everytime in a specif interval. But running additional speed test is waste when the host machine actually take care of that.

Change of plans : 

- `Read Host machine's network stats`: Like how btop reads all the data of network from the linux machine, make it so it reads from there.

- `Speed test`: The speed test will be optional option that will be only used when the data is not enough. 


**Manual Speed test**

Manual speed test will be used under conditions such as machine is idle or data is not enough for the proper ml. Or in case The docker container is not able to read data from host machine. Eg, permission error or windows. 




## Blind Spots 

There are some situation i will have to be careful about. 

- `Idle machine` : A idle machine is identical to a heavy throtlled network.
- `Localhost` : A machine hosting a local server like navidrome and listening music in local network will look same as downloading data from internet



## Counters 

These might be some solutions that maybe help in blind Spots.

- `Idle machine` : If a machine has low data usages, make a small ping request to external server, and based on the ping result, decide weather to run the speed test or not.

- `Localhost`: instead of monitoring every network interface, monitor the physical interface like wlan0. If for sometime there is no activity in the interface, run a ping test and then speed test. 



## Reading /proc file 

The reading will be done in background, it will read all the data and store them in seperate table in database, making them easy to see the past data. 
The primary goal of this is to reduce the data wasted in doing manual speed test everytime. 

I will have to figure out many things, things like how will the I will be able to accurately choose which data to choose for the baseline, the baseline will be the line that will tell the scrip to run the manual speed test.

I could use Some math functions to decide the baseline but i dont know yet. 


## Ping test 

The script will do primary two ping test. One `external ping test` and `Internal ping test`.

### Internal Ping test

The ping test will be done on the `captive portal` The portal which is used for logging in the wifi and getting internet access. 
This ping test will be used to determine weather the no. of connected device is high or the internet is throtlled. 

### External Ping test

Ping tests to `1.1.1.1` or `8.8.8.8` Will be used to determine if the internet is throtlled or not.


## Why two ping tests?

1. **Better Reasoning**: Without two distinct ping test the program will not be able to determine if the ISP is throttling or the no. of user is high.
    - Ping to Internal: Measures the health of the local Wi-Fi and the hostel router.
    - Ping to External: Measures the health of the college's actual ISP connection.

2. ML likes to have `delta` : The difference btw internal and external ping will be good for machine learnign, as it will be able to associate with the reason better.

    - Bad internal ping: The reason will most likely high student connected.
    - Good external ping: If the internal ping is higher but external ping is lower. The reasons could be that the high priority user is such as office workers are using more wifi bandwith and the server/router is proriting that over normal students 
    - Bad external ping: ISP throtlled.


3. Checks for login : The internal ping will be better way to tell if the wifi is still authenticated rather then requesting a `204`
