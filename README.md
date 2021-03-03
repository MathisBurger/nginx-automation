# nginx-automation
This REST service is written for the automatisation of the nginx reverse proxy.
I wrote this appication especially for an server I am working on. But it can also be
used on every other server.

## further developed
I will develop this, if I need more features or I get some new ideas to add for other users of this software.

## Important
This application is only creating your configurations. It does not restart your nginx server.
If you want to use this application in production, make sure you set your allowed-origins correctly in the config.json
Otherwise everyone can create configurations on your server.
Furthermore this application is not working in your network, because it does not match you needs.

## Installation
This is a step by step tutorial
- Download the executable file from the releases
- Create a folder in the same directory as your executable and create a file config.json
- Use this sample config to configure your application: <a href="https://github.com/MathisBurger/nginx-automation/blob/main/config/config.json">sample config</a>
- Start the application as service or in a screen
