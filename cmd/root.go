package cmd

import (
	"bytes"
	"fmt"
	"image"
	"image/color/palette"
	"image/draw"
	"image/gif"
	"image/jpeg"
	"image/png"
	"log"
	"math"
	"os"
	"path/filepath"

	"github.com/nfnt/resize"
	"github.com/spf13/cobra"

	homedir "github.com/mitchellh/go-homedir"
	"github.com/spf13/viper"
)

var cfgFile string

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "iago",
	Short: "A brief description of your application",
	Long: `A longer description that spans multiple lines and likely contains
examples and usage of using your application. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		make_parrot()
	},
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	cobra.OnInitialize(initConfig)

	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.iago.yaml)")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

// initConfig reads in config file and ENV variables if set.
func initConfig() {
	if cfgFile != "" {
		// Use config file from the flag.
		viper.SetConfigFile(cfgFile)
	} else {
		// Find home directory.
		home, err := homedir.Dir()
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		// Search config in home directory with name ".iago" (without extension).
		viper.AddConfigPath(home)
		viper.SetConfigName(".iago")
	}

	viper.AutomaticEnv() // read in environment variables that match

	// If a config file is found, read it in.
	if err := viper.ReadInConfig(); err == nil {
		fmt.Println("Using config file:", viper.ConfigFileUsed())
	}
}

func readOverlay(overlayFilepath string, height uint) image.Image {

	//Read in the overlay file
	inOverlay, err := os.Open(overlayFilepath)
	if err != nil {
		log.Fatalln(err)
	}
	defer inOverlay.Close()

	//Determine file type
	extension := filepath.Ext(overlayFilepath)

	//Decode file based on type
	var decodedOverlay image.Image

	if extension == ".jpg" {
		decodedOverlay, err = jpeg.Decode(inOverlay)
	} else if extension == ".png" {
		decodedOverlay, err = png.Decode(inOverlay)
	} else {
		return nil
	}

	//Log out any decoding errors
	if err != nil {
		log.Fatalln(err)
	}

	//Resize overlay image
	resizedOverlay := resize.Resize(0, height, decodedOverlay, resize.Lanczos2)
	return resizedOverlay
}

func buildNewParrot(decodedGif *gif.GIF, overlayImage image.Image, numFrames int) *gif.GIF {
	//For each frame
	for x := 0; x < numFrames; x++ {
		t := float64(x) / float64(numFrames)
		ellipseOffset := math.Pi / 4

		//Get decoded gif frame
		frame := decodedGif.Image[x]

		//Create new frame
		newFrame := image.NewPaletted(frame.Bounds(), append(palette.WebSafe, image.Transparent))

		//Calculate the position for the offset image
		offset := image.Pt(int(40+26*math.Cos(t*2*math.Pi+ellipseOffset)), int(35+-5*math.Sin(t*2*math.Pi+ellipseOffset)))

		//Write frame from decoded gif
		draw.Draw(newFrame, frame.Bounds(), frame, image.ZP, draw.Src)

		//Add overlay image to the frame
		draw.Draw(newFrame, frame.Bounds().Add(offset), overlayImage, image.ZP, draw.Over)

		//Overwrite decoded gif with the newly created frame
		decodedGif.Image[x] = newFrame
	}

	return decodedGif
}

func make_parrot() {
	//Get filepath argument
	var overlayFilepath string
	if len(os.Args) >= 2 {
		overlayFilepath = os.Args[1]
	} else {
		log.Fatalln("ERROR: No overlay filepath provided!")
	}

	//Open the base parrot gif
	baseGif, err := Asset("data/parrot.gif")
	if err != nil {
		log.Fatalln(err)
	}

	//Decode base parrot gif into frames
	decodedGif, err := gif.DecodeAll(bytes.NewReader(baseGif))
	if err != nil {
		log.Fatalln(err)
	}

	//Decode the overlay image
	overlayImage := readOverlay(overlayFilepath, 70)

	//Get number of frames in the decoded gif
	numFrames := len(decodedGif.Image)

	//Build new parrot gif
	parrotGif := buildNewParrot(decodedGif, overlayImage, numFrames)

	//Create output file
	outGif, err := os.Create("./parrot_out.gif")
	if err != nil {
		log.Fatalln(err)
	}
	defer outGif.Close()

	//Encode frames into output gif and write it out
	err = gif.EncodeAll(outGif, parrotGif)
	if err != nil {
		log.Fatalln(err)
	} else {
		fmt.Println("A new party parrot has been born!")
	}
}
