# Notes
Looks like we hidden some flags or secrets in the DOM that need to manipulated/displayed using the browser window API.

## Steps.

### Javascript hackery
1. Script in the head tag runs before the DOM content is loaded. Easy hack, wrap it in event listener.
2. Along with some text, the second step has a scirpt tag. Sneaky sneaky with the mis-spelling of 'third' lol
3. Okay, this stumped me for a sec. The code to generate the rawKey was pretty straight forward. But somehow I missed that extra 'this' in the text of the second step. That caused key to be wrong.

### Go App
I use a mac, and macOS does not have any simple native API to get the color of some window or dock. So we check the system theme and return the hex color of that. On windows it looks like you can set the color of the task bar to the color of your choosing. I would imagine there is some windows API that can return this info. However im not going to take the time to install VirtalBox and test this. Maybe what I hacked together will work!
Other than that, C# and Swift would be a better choice this for this task. Or some horrible automation library that can grab a screen shot / frame of the task bar and parse the color from that.  
