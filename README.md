# Go Channels Demo

a very simple go app shows how to use Go Channels for async tasks and Montior it with Prometheus & Grafana

# Grafana Dashboard Result
<img src="https://i.imgur.com/3aajUZt.gif">

# Installation
- clone the repo
- docker install
- open http://localhost:3000/ 
    - password is `pass`
    - add `Promethues` as `Data Source`
    - then import Dashboard from `./data/grafana-dashboard.json`
- send a requests to http://localhost:8080/email?to=ahmed.ashraf@email.com as a test

# Blog post
- (English) https://dev.to/ahmedash95/understand-golang-channels-and-how-to-monitor-with-grafana-154
- (Arabic) https://ahmedash.com/blog/article/golang-channels-and-monitoring/
