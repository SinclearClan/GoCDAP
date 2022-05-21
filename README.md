# GoCDAP
Calendar Discord Availability Provider


## Installation

1. Download a release from the [releases page](https://github.com/SinclearClan/GoCDAP/releases). 
   In most cases it's recommended to use the [latest release](https://github.com/SinclearClan/GoCDAP/releases/latest).
2. Create a file named "config.jsonc" in the same directory you put the release in and where you'll run the program later.
3. Edit the "config.jsonc" file with your favorite text editor and copy-paste the following into it:
   ```jsonc
   {
      /* Calendar */
      "calendar": {
         "type": "ical",                 // only ical is currently supported, other formats might come later
         "url": "",                      // base url of the server, for example 'https://nextcloud.example.com'
                                         // make shure there is no '/' at the end!
         "path": "",                     // path of the calendar (export) on the server, for example '/remote.php/dav/public-calendars/123456789/?export'
                                         // make shure this starts with a '/'!
         "user": "",                     // username, if required by the server to login
         "password": ""                  // password, if required by the server to login
      },
      /* Discord */
      "discord": {
         "app_id": "836637382509068319"  // Disord app id, you can leave the one provided or enter your own app id
      }
   }
   ```
4. Start the program by running the executable.
