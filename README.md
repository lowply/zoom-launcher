# Zoom Launcher

Open the Zoom URL of the next meeting in your Google Calendar.

## Installation

If you're on macOS, download the latest release from [the releases page](https://github.com/lowply/zoom-launcher/releases) and put the `zoom-launcher` binary to whereever you like, such as `/usr/local/bin`.

If you're not, please [install Go](https://golang.org/) first, then run:

```
go get github.com/lowply/zoom-launcher
```

The `zoom-launcher` command will be installed in your `$GOPATH/bin`.

## Prep

Google Calendar API access

- Run `mkdir ~/.config/google-calendar-api`
- Browse to [Go Quickstart](https://developers.google.com/calendar/quickstart/go),  click "Enable the Google Calendar API" and choose "Desktop App"
- Download the *credentials.json* file to *~/.config/google-calendar-api*

## Configuration

Make sure you export the following env vars.

- `ZL_EMAIL`: Email address of your company's G Suite account
- `ZL_REGEX`: Regular expression to match your company's Zoom URL

Example

```
export ZL_EMAIL="you@example.com"
export ZL_REGEX="^https://company.zoom.us/.*$"
```

Consider adding these to your `~/.bashrc`.

## Usage

Running `zoom-launcher` in your terminal will show you the next 3 meetings that you've accepted on Google Calendar. If the next one has the Zoom meeting URL, it converts the URL from `https://` to `zoommtg://` then opens it.

```
$ zoom-launcher
[2020-06-16 11:00:00 +0900 JST] One on one: https://example.zoom.us/j/id
[2020-06-16 14:30:00 +0900 JST] Team meeting: https://example.zoom.us/j/id
[2020-06-17 16:00:00 +0900 JST] Random chat: No meeting URL
Do you want to join "One on one" now? y/n
```
