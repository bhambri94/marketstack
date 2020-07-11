# marketstack

This utility fetches NASDAQ hundred stocks list with its below mentioned EOD figures and saves to Google Sheet. 
Open, High, Low, Close, Volume, AdjHigh, AdjLow, AdjClose, AdjOpen, AdjVolume, Symbol, Exchange, Date

This script can be used with a cron setup and pull out the latest market figures hourly, daily, weekly or monthly.

To run this script you need to install Go.
```
brew install go
```

git clone this project to your GoPATH src

```
cd /Users/{username}/go/src

git clone https://github.com/bhambri94/marketstack.git

go run main.go
```

Results will be updated to https://docs.google.com/spreadsheets/d/1VUAk7tVJU-toOazDW977wjgcaDF1MBTW9j7a-HTFJgI/edit#gid=0
If you want it to be updated at your sheet, please update configs file with your spreadsheetId.
