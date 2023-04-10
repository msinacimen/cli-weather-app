# Command Line Interface Weather App

CLI tool for getting weather information from OpenWeatherMap API with Golang - Cobra.

![](https://github.com/msinacimen/cli-weather-app/blob/main/weather_app.gif)


```
Weather App that shows the weather of a city in the world

Usage:
  wcli [command]

Available Commands:
  city        Get weather of a city
  help        Help about any command
  version     Print the version number of wcli

Flags:
  -h, --help   help for wcli

Use "wcli [command] --help" for more information about a command.
```

### Example Usage
```
>wcli city Istanbul
+----------+-----------------------------+-------------+------------+---------------------+----------+----------+
| City     | Weather                     | Temperature | Feels Like | Min/Max Temperature | Humidity | Pressure |
+----------+-----------------------------+-------------+------------+---------------------+----------+----------+
| Istanbul | Light Intensity Shower Rain | 10°C        | 7°C        | 9 - 10°C            | 87%      | 1013mbar |
+----------+-----------------------------+-------------+------------+---------------------+----------+----------+


>wcli city Poznan
+--------+-----------+-------------+------------+---------------------+----------+----------+
| City   | Weather   | Temperature | Feels Like | Min/Max Temperature | Humidity | Pressure |
+--------+-----------+-------------+------------+---------------------+----------+----------+
| Poznan | Clear Sky | 8°C         | 8°C        | 7 - 10°C            | 94%      | 1025mbar |
+--------+-----------+-------------+------------+---------------------+----------+----------+

```
