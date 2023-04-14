package main

import (
    "fmt"
    "log"
    "os"
    "time"

    "github.com/joho/godotenv"
    "github.com/nathanwinther/go-hue/pkg/hue"
)

type LightScheme struct {
    Name       string
    LightID    int
    Brightness uint8
    Hue        uint16
    Saturation uint8
}

type Config struct {
    BridgeIP    string
    BridgeUser  string
    LightSchemes []LightScheme
}

func main() {
    // Load configuration from .env file
    err := godotenv.Load()
    if err != nil {
        log.Fatal("Error loading .env file")
    }
    cfg := Config{
        BridgeIP:   os.Getenv("BRIDGE_IP"),
        BridgeUser: os.Getenv("BRIDGE_USER"),
        LightSchemes: []LightScheme{
            {
                Name:       "morning",
                LightID:    1,
                Brightness: 254,
                Hue:        46920,
                Saturation: 254,
            },
            {
                Name:       "evening",
                LightID:    1,
                Brightness: 254,
                Hue:        14910,
                Saturation: 254,
            },
        },
    }

    // Connect to the Hue bridge
    bridge := hue.NewBridge(cfg.BridgeIP, cfg.BridgeUser)

    // Set up the light schemes
    for _, scheme := range cfg.LightSchemes {
        err := bridge.SetLightScheme(scheme.Name, scheme.LightID, scheme.Brightness, scheme.Hue, scheme.Saturation)
        if err != nil {
            log.Fatalf("Error setting up light scheme %s: %s", scheme.Name, err)
        }
    }

    // Start the loop to check the time and adjust the lights accordingly
    for {
        current_time := time.Now().Format("15:04:05")
        if isBetween(current_time, "06:00:00", "08:00:00") {
            err := bridge.SetLightSchemeByName("morning")
            if err != nil {
                log.Fatalf("Error setting light scheme 'morning': %s", err)
            }
        } else if isBetween(current_time, "21:00:00", "23:59:59") || isBetween(current_time, "00:00:00", "06:00:00") {
            err := bridge.SetLightSchemeByName("evening")
            if err != nil {
                log.Fatalf("Error setting light scheme 'evening': %s", err)
            }
        } else {
            err := bridge.SetLightState(1, false)
            if err != nil {
                log.Fatalf("Error turning off light: %s", err)
            }
        }

        // Sleep for one minute
        time.Sleep(time.Minute)
    }
}

// isBetween checks if a given time is between two other times (inclusive)
func isBetween(t, start, end string) bool {
    layout := "15:04:05"
    t1, _ := time.Parse(layout, t)
    s, _ := time.Parse(layout, start)
    e, _ := time.Parse(layout, end)
    return t1.After(s) && t1.Before(e) || t == start || t == end
}
