# Ocelot Media Server [W.I.P.]

### Purpose
A free and open-source alternative to Plex Media Server that should look *beautiful*

### Goal
Reach full feature parity to Plex and possibly even Jellyfin.

- [ ] Users
    - [X] Login with email and password
    - [X] Create admin user 
    - [ ] Create secondary users
        - [ ] Have customized permissions for each user
- [ ] Encoding (ffmpeg)
    - [X] Generate .m3u8 file
    - [ ] Generate .ts files dynamically 
        - [ ] Thumbnails
        - [ ] Subtitles
- [ ] Library
    - [ ] Add directories
    - [ ] Scan and add metadata for directories
        - [ ] Scan on startup
        - [ ] Search metadata from external API's
        - [ ] Detect changes while server is still running
    - [ ] Control permission for which users have access to which libraries
- [ ] Allow custom plugins
    - [ ] Themes
- [ ] Logging
    - [ ] Show user data (such as who's watching etc.)
    - [ ] Analytics

### Questions
Here I'll answer some questions about the project preemptively

#### Why not fork Jellyfin, a project that intends to have the same goal?
During my personal experience moving from Plex to Jellyfin I encountered a few hurdles and friction that eft me less than satisfied.
Not to say Jellyfin isn't an amazing project because it is and will be probably superior to this project for a very long time, but I would like to re-invent the wheel and create something aesthetically pleasing, fast, and written in Golang.


##### Development started on April 29, 2024 by Siddhant Madhur

