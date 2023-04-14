package main

import (
    "flag"
    "fmt"
    "time"

    "github.com/studioimaginaire/go.hue"
)

type LightScheme struct {
    LightID     int
    Brightness  uint8
    Hue         uint16
    Saturation  uint8
}

var (
    bridgeIP   = flag.String("bridgeIP", "", "The IP address of the Philips Hue Bridge")
    username   = flag.String("username", "", "The username used to authenticate with the Philips Hue Bridge")
    lightSchemes = map[string]LightScheme{
        "morning": {"light_id": 1, "brightness": 100, "hue": 46920, "saturation": 254},
        "evening": {"light_id": 1, "brightness": 100, "hue": 14910, "saturation": 254},
    }
)

func setLight(client *hue.Client, light LightScheme) {
    state := hue.LightState{
        On:         true,
        Bri:        light.Brightness,
        Hue:        light.Hue,
        Saturation: light.Saturation,
    }
    _, err := client.SetLightState(light.LightID, state)
    if err != nil {
        fmt.Println("Error setting light state:", err)
    }
}

func main() {
    flag.Parse()

    client := hue.NewClient(*bridgeIP, *username)

    for {
        current_time := time.Now().Format("15:04")
        if current_time >= "06:00" && current_time < "08:00" {
            setLight(client, lightSchemes["morning"])
        } else if current_time >= "21:00" || current_time < "06:00" {
            setLight(client, lightSchemes["evening"])
        } else {
            _, err := client.SetLightState(1, hue.LightState{On: false})
            if err != nil {
                fmt.Println("Error setting light state:", err)
            }
        }

        time.Sleep(60 * time.Second)
    }
}
