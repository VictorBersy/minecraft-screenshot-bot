package main

import (
	"fmt"
	"io/fs"
	"math/rand"
	"os"
	"os/exec"
	"path/filepath"
	"time"

	"github.com/charmbracelet/log"
	"github.com/go-vgo/robotgo"
)

const (
	SHOTS               = 24
	WAIT_GAME_LAUNCH    = 15
	WAIT_GENERATION     = 15
	WAIT_CHUNKS_LOADING = 45
)

type PlayerRot struct {
	rx string
	rz string
}

func launchGame() {
	log.Info("Starting Minecraft...")
	cmd := exec.Command(
		"/Users/alyx/Minecraft/Install/runtime/java-runtime-gamma/mac-os-arm64/java-runtime-gamma/jre.bundle/Contents/Home/bin/java",
		"-XstartOnFirstThread",
		"-Djava.library.path=/Users/alyx/Minecraft/Install/bin/5ead28993186f55d8844fd30754042b6673e16f3",
		"-Djna.tmpdir=/Users/alyx/Minecraft/Install/bin/5ead28993186f55d8844fd30754042b6673e16f3",
		"-Dorg.lwjgl.system.SharedLibraryExtractPath=/Users/alyx/Minecraft/Install/bin/5ead28993186f55d8844fd30754042b6673e16f3",
		"-Dio.netty.native.workdir=/Users/alyx/Minecraft/Install/bin/5ead28993186f55d8844fd30754042b6673e16f3",
		"-Dminecraft.launcher.brand=minecraft-launcher",
		"-Dminecraft.launcher.version=2.12.17",
		"-cp", "/Users/alyx/Minecraft/Install/libraries/net/fabricmc/tiny-mappings-parser/0.3.0+build.17/tiny-mappings-parser-0.3.0+build.17.jar:/Users/alyx/Minecraft/Install/libraries/net/fabricmc/sponge-mixin/0.12.5+mixin.0.8.5/sponge-mixin-0.12.5+mixin.0.8.5.jar:/Users/alyx/Minecraft/Install/libraries/net/fabricmc/tiny-remapper/0.8.2/tiny-remapper-0.8.2.jar:/Users/alyx/Minecraft/Install/libraries/net/fabricmc/access-widener/2.1.0/access-widener-2.1.0.jar:/Users/alyx/Minecraft/Install/libraries/org/ow2/asm/asm/9.6/asm-9.6.jar:/Users/alyx/Minecraft/Install/libraries/org/ow2/asm/asm-analysis/9.6/asm-analysis-9.6.jar:/Users/alyx/Minecraft/Install/libraries/org/ow2/asm/asm-commons/9.6/asm-commons-9.6.jar:/Users/alyx/Minecraft/Install/libraries/org/ow2/asm/asm-tree/9.6/asm-tree-9.6.jar:/Users/alyx/Minecraft/Install/libraries/org/ow2/asm/asm-util/9.6/asm-util-9.6.jar:/Users/alyx/Minecraft/Install/libraries/net/fabricmc/intermediary/1.20.2/intermediary-1.20.2.jar:/Users/alyx/Minecraft/Install/libraries/net/fabricmc/fabric-loader/0.14.24/fabric-loader-0.14.24.jar:/Users/alyx/Minecraft/Install/libraries/ca/weblite/java-objc-bridge/1.1/java-objc-bridge-1.1.jar:/Users/alyx/Minecraft/Install/libraries/com/github/oshi/oshi-core/6.4.5/oshi-core-6.4.5.jar:/Users/alyx/Minecraft/Install/libraries/com/google/code/gson/gson/2.10.1/gson-2.10.1.jar:/Users/alyx/Minecraft/Install/libraries/com/google/guava/failureaccess/1.0.1/failureaccess-1.0.1.jar:/Users/alyx/Minecraft/Install/libraries/com/google/guava/guava/32.1.2-jre/guava-32.1.2-jre.jar:/Users/alyx/Minecraft/Install/libraries/com/ibm/icu/icu4j/73.2/icu4j-73.2.jar:/Users/alyx/Minecraft/Install/libraries/com/mojang/authlib/5.0.47/authlib-5.0.47.jar:/Users/alyx/Minecraft/Install/libraries/com/mojang/blocklist/1.0.10/blocklist-1.0.10.jar:/Users/alyx/Minecraft/Install/libraries/com/mojang/brigadier/1.1.8/brigadier-1.1.8.jar:/Users/alyx/Minecraft/Install/libraries/com/mojang/datafixerupper/6.0.8/datafixerupper-6.0.8.jar:/Users/alyx/Minecraft/Install/libraries/com/mojang/logging/1.1.1/logging-1.1.1.jar:/Users/alyx/Minecraft/Install/libraries/com/mojang/patchy/2.2.10/patchy-2.2.10.jar:/Users/alyx/Minecraft/Install/libraries/com/mojang/text2speech/1.17.9/text2speech-1.17.9.jar:/Users/alyx/Minecraft/Install/libraries/commons-codec/commons-codec/1.16.0/commons-codec-1.16.0.jar:/Users/alyx/Minecraft/Install/libraries/commons-io/commons-io/2.13.0/commons-io-2.13.0.jar:/Users/alyx/Minecraft/Install/libraries/commons-logging/commons-logging/1.2/commons-logging-1.2.jar:/Users/alyx/Minecraft/Install/libraries/io/netty/netty-buffer/4.1.97.Final/netty-buffer-4.1.97.Final.jar:/Users/alyx/Minecraft/Install/libraries/io/netty/netty-codec/4.1.97.Final/netty-codec-4.1.97.Final.jar:/Users/alyx/Minecraft/Install/libraries/io/netty/netty-common/4.1.97.Final/netty-common-4.1.97.Final.jar:/Users/alyx/Minecraft/Install/libraries/io/netty/netty-handler/4.1.97.Final/netty-handler-4.1.97.Final.jar:/Users/alyx/Minecraft/Install/libraries/io/netty/netty-resolver/4.1.97.Final/netty-resolver-4.1.97.Final.jar:/Users/alyx/Minecraft/Install/libraries/io/netty/netty-transport-classes-epoll/4.1.97.Final/netty-transport-classes-epoll-4.1.97.Final.jar:/Users/alyx/Minecraft/Install/libraries/io/netty/netty-transport-native-unix-common/4.1.97.Final/netty-transport-native-unix-common-4.1.97.Final.jar:/Users/alyx/Minecraft/Install/libraries/io/netty/netty-transport/4.1.97.Final/netty-transport-4.1.97.Final.jar:/Users/alyx/Minecraft/Install/libraries/it/unimi/dsi/fastutil/8.5.12/fastutil-8.5.12.jar:/Users/alyx/Minecraft/Install/libraries/net/java/dev/jna/jna-platform/5.13.0/jna-platform-5.13.0.jar:/Users/alyx/Minecraft/Install/libraries/net/java/dev/jna/jna/5.13.0/jna-5.13.0.jar:/Users/alyx/Minecraft/Install/libraries/net/sf/jopt-simple/jopt-simple/5.0.4/jopt-simple-5.0.4.jar:/Users/alyx/Minecraft/Install/libraries/org/apache/commons/commons-compress/1.22/commons-compress-1.22.jar:/Users/alyx/Minecraft/Install/libraries/org/apache/commons/commons-lang3/3.13.0/commons-lang3-3.13.0.jar:/Users/alyx/Minecraft/Install/libraries/org/apache/httpcomponents/httpclient/4.5.13/httpclient-4.5.13.jar:/Users/alyx/Minecraft/Install/libraries/org/apache/httpcomponents/httpcore/4.4.16/httpcore-4.4.16.jar:/Users/alyx/Minecraft/Install/libraries/org/apache/logging/log4j/log4j-api/2.19.0/log4j-api-2.19.0.jar:/Users/alyx/Minecraft/Install/libraries/org/apache/logging/log4j/log4j-core/2.19.0/log4j-core-2.19.0.jar:/Users/alyx/Minecraft/Install/libraries/org/apache/logging/log4j/log4j-slf4j2-impl/2.19.0/log4j-slf4j2-impl-2.19.0.jar:/Users/alyx/Minecraft/Install/libraries/org/joml/joml/1.10.5/joml-1.10.5.jar:/Users/alyx/Minecraft/Install/libraries/org/lwjgl/lwjgl-glfw/3.3.2/lwjgl-glfw-3.3.2.jar:/Users/alyx/Minecraft/Install/libraries/org/lwjgl/lwjgl-glfw/3.3.2/lwjgl-glfw-3.3.2-natives-macos.jar:/Users/alyx/Minecraft/Install/libraries/org/lwjgl/lwjgl-glfw/3.3.2/lwjgl-glfw-3.3.2-natives-macos-arm64.jar:/Users/alyx/Minecraft/Install/libraries/org/lwjgl/lwjgl-jemalloc/3.3.2/lwjgl-jemalloc-3.3.2.jar:/Users/alyx/Minecraft/Install/libraries/org/lwjgl/lwjgl-jemalloc/3.3.2/lwjgl-jemalloc-3.3.2-natives-macos.jar:/Users/alyx/Minecraft/Install/libraries/org/lwjgl/lwjgl-jemalloc/3.3.2/lwjgl-jemalloc-3.3.2-natives-macos-arm64.jar:/Users/alyx/Minecraft/Install/libraries/org/lwjgl/lwjgl-openal/3.3.2/lwjgl-openal-3.3.2.jar:/Users/alyx/Minecraft/Install/libraries/org/lwjgl/lwjgl-openal/3.3.2/lwjgl-openal-3.3.2-natives-macos.jar:/Users/alyx/Minecraft/Install/libraries/org/lwjgl/lwjgl-openal/3.3.2/lwjgl-openal-3.3.2-natives-macos-arm64.jar:/Users/alyx/Minecraft/Install/libraries/org/lwjgl/lwjgl-opengl/3.3.2/lwjgl-opengl-3.3.2.jar:/Users/alyx/Minecraft/Install/libraries/org/lwjgl/lwjgl-opengl/3.3.2/lwjgl-opengl-3.3.2-natives-macos.jar:/Users/alyx/Minecraft/Install/libraries/org/lwjgl/lwjgl-opengl/3.3.2/lwjgl-opengl-3.3.2-natives-macos-arm64.jar:/Users/alyx/Minecraft/Install/libraries/org/lwjgl/lwjgl-stb/3.3.2/lwjgl-stb-3.3.2.jar:/Users/alyx/Minecraft/Install/libraries/org/lwjgl/lwjgl-stb/3.3.2/lwjgl-stb-3.3.2-natives-macos.jar:/Users/alyx/Minecraft/Install/libraries/org/lwjgl/lwjgl-stb/3.3.2/lwjgl-stb-3.3.2-natives-macos-arm64.jar:/Users/alyx/Minecraft/Install/libraries/org/lwjgl/lwjgl-tinyfd/3.3.2/lwjgl-tinyfd-3.3.2.jar:/Users/alyx/Minecraft/Install/libraries/org/lwjgl/lwjgl-tinyfd/3.3.2/lwjgl-tinyfd-3.3.2-natives-macos.jar:/Users/alyx/Minecraft/Install/libraries/org/lwjgl/lwjgl-tinyfd/3.3.2/lwjgl-tinyfd-3.3.2-natives-macos-arm64.jar:/Users/alyx/Minecraft/Install/libraries/org/lwjgl/lwjgl/3.3.2/lwjgl-3.3.2.jar:/Users/alyx/Minecraft/Install/libraries/org/lwjgl/lwjgl/3.3.2/lwjgl-3.3.2-natives-macos.jar:/Users/alyx/Minecraft/Install/libraries/org/lwjgl/lwjgl/3.3.2/lwjgl-3.3.2-natives-macos-arm64.jar:/Users/alyx/Minecraft/Install/libraries/org/slf4j/slf4j-api/2.0.7/slf4j-api-2.0.7.jar:/Users/alyx/Minecraft/Install/versions/fabric-loader-0.14.24-1.20.2/fabric-loader-0.14.24-1.20.2.jar",
		"-DFabricMcEmu=net.minecraft.client.main.Main",
		"-Xmx4096m",
		"-Xms256m",
		"-Dfml.ignorePatchDiscrepancies=true",
		"-Dfml.ignoreInvalidMinecraftCertificates=true",
		"-Duser.language=en",
		"-Duser.country=US",
		"-Dlog4j.configurationFile=/Users/alyx/Minecraft/Install/assets/log_configs/client-1.12.xml",
		"net.fabricmc.loader.impl.launch.knot.KnotClient",
		"--version", "fabric-loader-0.14.24-1.20.2",
		"--gameDir", "/Users/alyx/Minecraft/Instances/Boosted FPS [FABRIC] Performance Mods",
		"--assetsDir", "/Users/alyx/Minecraft/Install/assets",
		"--assetIndex", "8",
		"--accessToken", os.Getenv("ACCESS_TOKEN"),
		"--clientId", os.Getenv("CLIENT_ID"),
		"--username", os.Getenv("USERNAME"),
		"--uuid", os.Getenv("UUID"),
		"--xuid", os.Getenv("XUID"),
		"--userType", "msa",
		"--versionType", "release",
		"--width", "1920",
		"--height", "1080",
		"--quickPlayPath", "/Users/alyx/Minecraft/Install/quickPlay/java/1698871293341.json",
	)

	if err := cmd.Start(); err != nil {
		log.Fatalf("Failed to start Minecraft: %s", err)
	}
	// Wait for the game to launch
	log.Info("Game launching", "sleep", WAIT_GAME_LAUNCH)
	time.Sleep(WAIT_GAME_LAUNCH * time.Second)
}

func createNewWorld() {
	// Navigate the menu to create a new world
	robotgo.KeySleep = 350
	robotgo.KeyTap("down")  // Singleplayer
	robotgo.KeyTap("enter") // Enter singleplayer
	robotgo.KeyTap("tab")
	robotgo.KeyTap("tab")
	robotgo.KeyTap("tab")   // Create new world
	robotgo.KeyTap("enter") // Select
	robotgo.KeyTap("a", "cmd")
	robotgo.KeyTap("backspace")
	robotgo.TypeStr("BOT_SCREENSHOT")
	robotgo.KeyTap("tab")
	robotgo.KeyTap("space") // Turn hardcore mode
	robotgo.KeyTap("space") // Turn creative mode
	robotgo.KeyTap("tab")
	robotgo.KeyTap("tab")
	robotgo.KeyTap("tab")
	robotgo.KeyTap("enter") // Create new world

	// Wait for world generation to complete
	log.Info("Generating world", "sleep", WAIT_GENERATION)
	time.Sleep(WAIT_GENERATION * time.Second)
}

func setupScreenshot() {
	time.Sleep(1)
	runMinecraftChatCommand("/gamemode spectator")
	time.Sleep(1)
	robotgo.KeyTap("f1") // Hide HUD
}

func setRandomTime() {
	dayTime := getRandomTime()
	log.Info(fmt.Sprintf("Time set to %s", dayTime))
	runMinecraftChatCommand(fmt.Sprintf("/time set %s", dayTime))
}

func setRandomWeather() {
	// Set weather to clear by default
	weather := "clear"

	weatherRand := rand.Float64()
	if weatherRand <= 0.1 { // 10% chance of rain or thunder
		rainy := []string{"rain", "thunder"}
		weather = rainy[rand.Intn(len(rainy))] // 50% chance of rain or thunder
	}

	log.Info(fmt.Sprintf("Weather set to %s", weather))
	runMinecraftChatCommand(fmt.Sprintf("/weather %s", weather))
}

func teleportPlayer() {
	// Teleport the player to random surface location in a 20,000×20,000-block area centered on (0,0)
	runMinecraftChatCommand("/spreadplayers 0 0 0 10000 true @p")
	time.Sleep(2 * time.Second) // Wait for the TP to happen
	// Set player's camera to random rotation
	rot := getRandomAngle()
	log.Info(fmt.Sprintf("Random rotation: 'RX: %s, RZ: %s'", rot.rx, rot.rz))
	runMinecraftChatCommand(fmt.Sprintf("/tp @p ~ ~ ~ %s %s", rot.rx, rot.rz))
	// Teleport player 8 blocks above max
	runMinecraftChatCommand(fmt.Sprintf("/tp @p ~ ~%d ~", rand.Intn(9)))

	log.Info("Loading chunks", "sleep", WAIT_CHUNKS_LOADING)
	time.Sleep((WAIT_CHUNKS_LOADING * time.Second)) // Wait for the chunks generation and rendering
}

func takeRandomScreenshot() {
	// Take a screenshot
	log.Info("Taking screenshot by pressing F2")
	robotgo.KeyTap("f2")        // Take native screenshot
	time.Sleep(2 * time.Second) // Save screenshot
}

func quitGame() {
	// Close the game
	robotgo.KeyTap("q", "cmd")
}

func getLatestScreenshot() string {
	fileInfo, err := getLastCreatedFile()
	if err != nil {
		log.Fatal(err)
	}

	if fileInfo != nil {
		log.Info(fmt.Sprintf("Screenshot found: %s, created at %s\n", fileInfo.Name(), fileInfo.ModTime()))
	} else {
		log.Fatal("Screenshot not found.")
	}

	return filepath.Join(os.Getenv("SCREENSHOTS_DIR_PATH"), fileInfo.Name())
}

// getLastCreatedFile takes a directory path as an argument and returns the FileInfo
// of the most recently created file in that directory. If the directory is empty,
// or if there are no files in the directory, it returns nil.
func getLastCreatedFile() (fs.FileInfo, error) {
	dir := os.Getenv("SCREENSHOTS_DIR_PATH")
	// Initialize variables to store information about the newest file
	var newestFile fs.FileInfo
	var newestTime time.Time

	// Walk through the directory and its subdirectories
	err := filepath.WalkDir(dir, func(path string, d fs.DirEntry, err error) error {
		// If an error occurs, return the error and stop walking the directory
		if err != nil {
			return err
		}
		// If the current entry is a directory, skip it
		if d.IsDir() {
			return nil
		}

		// Retrieve the FileInfo of the current file
		info, err := d.Info()
		if err != nil {
			return err
		}

		// Check if the current file is newer than the newest file found so far
		if info.ModTime().After(newestTime) {
			// If it is, update newestFile and newestTime with the current file's information
			newestFile = info
			newestTime = info.ModTime()
		}
		// Continue walking through the directory
		return nil
	})

	// If an error occurred during the directory walk, return the error
	if err != nil {
		return nil, err
	}

	// Return the FileInfo of the newest file found (or nil if no files were found)
	return newestFile, nil
}

func runMinecraftChatCommand(cmd string) {
	robotgo.KeyTap("t") // Enter chat
	robotgo.TypeStr(cmd)
	robotgo.KeyTap("enter") // Run command
}

func getRandomAngle() PlayerRot {
	return PlayerRot{
		rx: fmt.Sprint(rand.Intn(361)),
		rz: fmt.Sprint(rand.Intn(41) - 20), // Keep between -20 and +20
	}
}

func getRandomTime() string {
	// https://minecraft.fandom.com/fr/wiki/Cycle_jour-nuit#P%C3%A9riodes_d'un_cycle
	return fmt.Sprint(rand.Intn(450) + rand.Intn(13805))
}
