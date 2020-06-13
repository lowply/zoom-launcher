# Zoom Launcher

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