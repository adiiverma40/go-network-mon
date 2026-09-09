# Data collection

This docs contains the explantion and my choices for data collection algo, how i made it and what and why i choose that.

The Following project will majorly Be for **Linux**. I dont have a windows machine, neither do i plan to use it, So i wont support it.


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


