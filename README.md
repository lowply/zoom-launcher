# Zoom Launcher

Open the Zoom URL of the next meeting in your Google Calendar.

## Prep

- Run `mkdir ~/.config/google-calendar-api`
- Browse to [Go Quickstart](https://developers.google.com/calendar/quickstart/go), Enable the Google Calendar API and choose Desktop App
- Download the *credentials.json* file to *~/.config/google-calendar-api*
- Run the following

```
cat << EOF > ~/.config/zoom-launcher.yaml
# Your company email address
email: you@example.com
# RegExp to match your company's Zoom URL
zoomurl: "^https://company.zoom.us/.*$"
EOF
```

## Usage

Running `zoom-launcher` in your terminal will show you the next 3 meetings that you've accepted on Google Calendar. If the next one has the Zoom meeting URL, it converts the URL from `https://` to `zoommtg://` then opens it.