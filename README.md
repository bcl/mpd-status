Simple mpd status display

* MPD protocol - https://www.musicpd.org/doc/html/protocol.html
* mpd package docs - https://godoc.org/github.com/fhs/gompd/mpd

Add this to the tmux status line with:

    set -g status-right "#[fg=green]#(~/bin/mpd-status)#[fg=default] | [#H] %H:%M %e-%b-%g"

Fields available:

	STATUS- map[audio:44100:24:2 bitrate:320 consume:0 duration:223.791 elapsed:160.651 mixrampdb:0.000000 nextsong:5 nextsongid:6 playlist:2 playlistlength:19 random:1 repeat:0 single:0 song:0 songid:1 state:pause time:161:224 volume:75]

	SONG- map[Album:Joe Bonamassa Live From the Royal Albert Hall AlbumArtist:Joe Bonamassa Artist:Joe Bonamassa Date:2009 Disc:1 Genre:Blues Id:1 Last-Modified:2017-10-09T18:21:04Z Pos:0 Time:224 Title:Django (Live) Track:1 duration:223.791 file:Joe Bonamassa/Live From The Royal Albert Hall/01 - Django (Live).mp3]
