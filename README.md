# iago
![alt text](https://raw.githubusercontent.com/devenjarvis/iago/master/data/parrot.gif)

## About
Everyone knows the biggest problem with Slack and other chat apps is the clear lack of focus it allows. And by that I obviously mean lack of focus on partying. iago is a tool that allows you to quickly combine any image with a party parrot gif to create new and exciting ways for you and your team to party! 

## Install
To get the lastest version of iago, make sure you have golang setup on your machine and run:
`go get -u github.com/devenjarvis/iago`

## Usage
Once installed, you should be able to run `iago <path_to_image>` from anywhere in your terminal to generate a new parrot gif.

Things to note:
- iago supports .png and .jpeg images
- iago currently outputs the resulting gif to `./parrot_out.gif` so it's relative to wherever you run the command. There's an open issue to make this customizable.
- iago will make no attempts to strip the background of your image before inviting it to the party. Checkout https://remove.bg/ as a possible tool to do this for you - I have no affiliation with the product, just found it when looking for a quick background removal solution.

## Support
iago not meeting your parrot partying needs? Open an issue with your bug or feature request and I'll probably get around to it at some point.