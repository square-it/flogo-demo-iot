# flogo-demo-iot

The application is the screen part of the Square IT Flogo demo

![Flogo Demo](https://github.com/square-it/flogo-demo/blob/master/Flogo%20Demo.png)

## Description

The application is running on a raspberry pi.
His goal is to display a smiley image on the display of the raspberry.
The smiley image is chosen by a API Rest HTTP.

This API is described by the file [swagger.yaml](swagger.yaml)

## Implementation

The application is built with [Flogo](http://www.flogo.io/) and [Square IT custom activities](https://github.com/square-it/flogo-contrib-activities).

The flow is composed of :

- A REST trigger
- A log activity to print the input request
- A command activity to hide the previous smiley image displayed
- A command activity to display the new smiley image
- A return activity to response to the client

The program [fbi](https://linux.die.net/man/1/fbi) is used by the application to display the images.

The smiley images used come from the [openmoji site](http://openmoji.org/).
