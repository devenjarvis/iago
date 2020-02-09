# iago
<img src="https://raw.githubusercontent.com/devenjarvis/iago/master/data/parrot.gif" width="50" height="50
"/> + <img src="https://raw.githubusercontent.com/devenjarvis/iago/master/data/batman.png" width="50" height="50
"/> = <img src="https://raw.githubusercontent.com/devenjarvis/iago/master/data/parrot_out.gif" width="50" height="50
"/>

## About
Everyone knows the biggest problem with Slack and other chat apps is the clear lack of focus it allows. And by that I obviously mean lack of focus on partying. iago is a tool that allows you to quickly combine any image with a party parrot gif to create new and exciting ways for you and your team to party! 

## Install

### Using Homebrew
1. Add the tap with: `brew tap devenjarvis/iago`
2. Install the CLI with: `brew install iago`
3. Party Time!

### Using golang
If you have golang setup on your machine you can run `go get -u github.com/devenjarvis/iago` to go get the package. If your GOPATH is setup and visible from your system PATH then iago will be available globally. If you haven't setup gopath because you're on 1.13 or higher, than you'll need to compile iago and move it to your path manually.

## Usage
Once installed, you should be able to run `iago <path_to_image>` from anywhere in your terminal to generate a new parrot gif.

Things to note:
- iago supports .png and .jpeg images
- iago currently outputs the resulting gif to `./parrot_out.gif` so it's relative to wherever you run the command. There's an open issue to make this customizable.
- iago will make no attempts to strip the background of your image before inviting it to the party. Checkout https://remove.bg/ as a possible tool to do this for you - I have no affiliation with the product, just found it when looking for a quick background removal solution.

## Support
iago not meeting your parrot partying needs? Open an issue with your bug or feature request and I'll probably get around to it at some point.