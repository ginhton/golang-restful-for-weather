# Golang Restful for Weather

Golang Restful for Weather is a project that aims provide restful api for weatherapi.com.

<!-- PROJECT SHIELDS -->

[![Contributors][contributors-shield]][contributors-url]
[![Forks][forks-shield]][forks-url]
[![Stargazers][stars-shield]][stars-url]
[![Issues][issues-shield]][issues-url]
[![MIT License][license-shield]][license-url]


<!-- PROJECT LOGO -->
<br />

<p align="center">
  <a href="https://github.com/ginhton/golang-restful-for-weather">
    <img src="images/logo.png" alt="Logo" width="80" height="80">
  </a>

  <h3 align="center">Golang Restful for Weather</h3>
  <p align="center">
    a good attempt to do this!
    <br />
    <a href="#"><strong>explore more about this project in our docs »</strong></a>
    <br />
    <br />
    <a href="#">Demo</a>
    ·
    <a href="#">Bug</a>
    ·
    <a href="#">Code</a>
  </p>

</p>

## toc

- [features](#features)
- [run](#features)


### features

- get temperature for a specific city
- get average temperatures for a specific city in future N days


### run

how to run this *script*

```sh
git clone https://github.com/ginhton/golang-restful-for-weather
cd golang-restful-for-weather
go get -u github.com/gorilla/mux
go run .
```

now open brower and hit "localhost:8081" which will open our homepage with no function actually.

when you hit "localhost:8081/api/CITYNAME", such as "Beijing", you will get a json response including temperature today of that city.

when you hit "localhost:8081/api/CITYNAME/DAYS", such as "localhost:8081/api/Beijing/4", you will get a json response including temperatures of the city that you specified in future some days.



<!-- links -->
[your-project-path]:ginhton/golang-restful-for-weather
[contributors-shield]: https://img.shields.io/github/contributors/ginhton/golang-restful-for-weather.svg?style=flat-square
[contributors-url]: https://github.com/ginhton/golang-restful-for-weather/graphs/contributors
[forks-shield]: https://img.shields.io/github/forks/ginhton/golang-restful-for-weather.svg?style=flat-square
[forks-url]: https://github.com/ginhton/golang-restful-for-weather/network/members
[stars-shield]: https://img.shields.io/github/stars/ginhton/golang-restful-for-weather.svg?style=flat-square
[stars-url]: https://github.com/ginhton/golang-restful-for-weather/stargazers
[issues-shield]: https://img.shields.io/github/issues/ginhton/golang-restful-for-weather.svg?style=flat-square
[issues-url]: https://img.shields.io/github/issues/ginhton/golang-restful-for-weather.svg
[license-shield]: https://img.shields.io/github/license/ginhton/golang-restful-for-weather.svg?style=flat-square
[license-url]: https://github.com/ginhton/golang-restful-for-weather/blob/master/LICENSE.txt

