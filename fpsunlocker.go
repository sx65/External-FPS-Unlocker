package main

import (
    "encoding/json"
    "fmt"
    "io/ioutil"
    "net/http"
    "os"
    "path/filepath"
)

func main() {
    // The path to the local Roblox directory
    localAppData := os.Getenv("LOCALAPPDATA")
    if localAppData == "" {
        fmt.Println("appdata env variable not found")
        return
    }

    // Get the ver variable value from the github raw site
    resp, err := http.Get("https://raw.githubusercontent.com/lxyobaba/version/main/robloxversion")
    if err != nil {
        fmt.Println("error getting ver from our website")
        return
    }
    defer resp.Body.Close()

    body, err := ioutil.ReadAll(resp.Body)
    if err != nil {
        fmt.Println("error reading version")
        return
    }

    var data map[string]interface{}
    err = json.Unmarshal(body, &data)
    if err != nil {
        fmt.Println("error parsing version")
        return
    }

    ver, ok := data["ver"].(string)
    if !ok {
        fmt.Println("invalid version")
        return
    }

    robloxPath := filepath.Join(localAppData, "Roblox")
    versionsPath := filepath.Join(robloxPath, "Versions")
    versionPath := filepath.Join(versionsPath, ver)
    fmt.Println("Versions folder path:", versionPath)


    // Create the ClientSettings folder if it does not exist
    clientSettingsPath := filepath.Join(versionPath, "ClientSettings")
    if _, err := os.Stat(clientSettingsPath); os.IsNotExist(err) {
        err := os.Mkdir(clientSettingsPath, 0755)
        if err != nil {
            fmt.Println("error configuration 0x01")
            return
        }
    }

    // Read the maximum FPS value from user input
    fmt.Print("Enter the maximum FPS value: ")
    var maxFPS int
    fmt.Scanln(&maxFPS)

    // Create the ClientAppSettings.json file with the maximum FPS value
    settings := map[string]interface{}{
        "DFIntTaskSchedulerTargetFps": maxFPS,
    }
    settingsJSON, err := json.Marshal(settings)
    if err != nil {
        fmt.Println("error configuring 0x02")
        return
    }

    settingsFilePath := filepath.Join(clientSettingsPath, "ClientAppSettings.json")
    err = os.WriteFile(settingsFilePath, settingsJSON, 0644)
    if err != nil {
        fmt.Println("error configuring 0x03")
        return
    }

    fmt.Println("discord.gg/streamable | max fps has been set up: ", maxFPS)
}