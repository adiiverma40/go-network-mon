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


