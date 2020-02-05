# Horoscopes Application
Horoscopes is a command line application written in Go that return's a user's (daily, weekly, monthly, or yearly horoscope) based on user zodiac sign. 

This application is secured expanded, to be used on https://localhost:8080, and secured using a variety of features.

### User Stories Basic Command Line Application
- [X] User can save requested horoscope as a text file.
- [X] User can obtain daily reading for desired horoscope. 
- [X] User can obtain weekly reading for desired horoscope. 
- [X] User can obtain monthly reading for desired horoscope. 
- [X] User can obtain yearly reading for desired horoscope. 

### CLI Application Features 
- [X] Documentation
- [X] Unit testing
- [ ] Logging
- [X] API Library
- [X] CLI flags
- [ ] Environment variables
- [X] Concurrency
- [X] Data Persistance
- [X] HTTP/HTTPS API

### Security Features
- [X] Secure Reverse Proxy (HTTPS)
- [ ] Firewall
- [X] Intrusion Detection System
- [ ] Logging Manager
- [ ] Load Balancer
- [X] User Authentication 

### Also Includes
- [ ] Documentation
- [ ] Agile Project Management
- [ ] Unit Testing
- [ ] Logs & Metrics
- [ ] Environment Configuration
- [ ] Security
- [ ] Build & Deploy Scripts
- [ ] Containerization

## Getting Started

### Prerequisites

go get ORM library for Golang
```
go get -u github.com/jinzhu/gorm

```

### Installing 

 To install this project please run the following command:
 
 ```
 go get -u -d github.com/danish287/project-1

```
### Build CLI

From project-1 directory run the following commands:

```
 go build cmd/horoscope/horoscope.go

```
Execute:
```
./horoscope -s="aries"

```

## Flags 

1. ``` -s ``` allows the user to determine the zodiac sign for the requested horscope. This flag is **mandatory** and without it you will be prompted to include an argument for it.

 ```
./horoscope -s="aries" 

```
 Example Output:
 ```

Hi, Aries!
Here is your reading for the Day of February 05, 2020

Ganesha predicts that today you will look up to your family for happiness. Towards the later part of the day, however, there are possibilities of a platonic affair blossoming. Those who are already caught up in a romantic affair, says Ganesha, will experience some thrilling times.

Would you like to save this horoscope? (y/n)


```


2. ``` -date ``` allows the user to determine the date the requested horscope. This flag is optional, and without it the default value is "today". 
>> - today
>> - week
>> - month
>> - year

 
 ```
 ./horoscope -s="Aquarius" -date="year"

```

Example Output:
```
Hi, Aquarius!
Here is your reading for the Year 2020

This year seems to be filled up with adventures and challenges. The ruler of your sign, Saturn, rules Capricorn along with Jupiter and Mercury in the 12th house of your sign. This positioning is predicted to be a tricky one for your sign this year. It can either bring extreme happiness or extreme challenges in your way. You may feel under pressure on the occupational or financial front this year and your progress may be slower than expected. Keep an effective check on your personal as well as miscellaneous expenses for savings that might not go as per your calculations. Also, taking up risky investments is not ideal in this period. From March 23, Saturn enters Aquarius indicating that an uncomfortable situation may finally come to an end.', ' You might take on risky investments this year, says Ganesha. Your increased cash inflow may trigger your instincts to make a well-calculated move in business. The fifth house is directly connected to financial gains and progress on the occupational front. Due to this positioning and planetary movement, the financial gains add more to the power of your already strong occupational as well as financial position. From July 2, Saturn retrogrades into Capricorn. This transition may bring challenges in your way. From June 29, Mars enters Aries through the third house. The third house is directly connected with your profession. Progressive forces shall work effectively which may help your progress in life. Furthermore, Mars retrogrades from September 10 to November 14. Therefore, refrain from taking any major decision regarding occupation and finance in this period. By the end of the year, you shall feel multiple emotions at a time as well!', 'The beginning of the year may not be the best for your marital life. Being adjustable and compromising may prove to be working in your favor. It might as well help you overcome certain existing problems. Nonetheless, be prepared for turbulence during the time frame between May 13 to June 25. Those who are married may find it difficult to cope up with growing issues. Hence, it is advised to be patient, keep calm and resolve issues without creating any disturbances in the relationship. Remember getting angry will only worsen the already delicate situation. Patience is the key to a successful marriage.', 'The influence of Mars due to its movement in the eleventh house can prove to be beneficial which means that you may witness passionate moments arising in your relationship. However, just when your married life looks sorted, your family life screams trouble. Your little ones may annoy you to no end. In fact, you will not see face to face on any topic. But, your children and you need to meet midway to maintain a harmonious balance in your relationship as well. Getting worked up at petty issues may not bring in the desired results so keep your cool with a mature understanding if things go wrong.

Would you like to save this horoscope? (y/n)

```

**Note:** if you type ```y``` + enter this will save the file as a text.
Example Output:

``` Would you like to save this horoscope as a text file? (y/n) ``` 
``` y ``` 
```Your horoscoped has been saved as aquariusYear2020-02-05 under the /savedhoroscopes directory.```  

**Note:** if you type ```n``` + enter this will save the file as a text:
Example Output:
``` Would you like to save this horoscope as a text file? (y/n) ``` 
``` n ``` 

### Local Host

```
 go build cmd/horoscoped/horoscoped.go

```
Execute:
```
./horoscoped

```

### Build Proxy

```
 go build cmd/proxy/proxy.go

```
Execute:
```
sudo ./proxy

```

## Running Tests

These are instructions on how to run all tests from the project-1 directory.

### Package gethoroscope tests

This will test that we are accurately retrieving requested horoscope using the api. Addiitonaly, we tests correct output, promting the user to use valid flag arguments, when the user flag arguments are not valid. 

```
go test internal/gethoroscope/gethoroscope.go  internal/gethoroscope/gethoroscope_test.go -v

```



