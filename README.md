# shovel
wget powered go program for multi-threaded web scraping

###Running Shovel:

Install golang (osx):

`brew install hg`

`brew install go`

Pull repo:

`git clone https://github.com/phact/shovel.git`

Run shovel:

`go run shovel.go`

Config:

Your list of URL's goes in data/urls.txt

By default we run 100 threads of wget at a time. Change the maxFutures int to alter the number of concurrent jobs.
