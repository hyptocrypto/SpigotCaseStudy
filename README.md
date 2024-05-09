# Notes
Looks like we hidden some flags or secrets in the DOM that need to manipulated/displayed using the browser window API.

### Steps.
1. Script in the head tag runs before the DOM content is loaded. Wrap in event listener.
2. Along with some text, the second step has a scirpt tag. Sneaky sneaky with the mis-spelling of 'third' lol
3. Okay, this stumped me for a min. The code to generate the rawKey was pretty straight forward. But somehow I missed that extra 'this' in the text of the second step. That caused key to be wrong.
