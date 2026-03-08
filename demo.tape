# nfcfyi-go — Go API Client Demo
Output demo.gif

Set FontSize 20
Set Width 1000
Set Height 580
Set Shell "zsh"
Set Theme "Catppuccin Mocha"
Set TypingSpeed 30ms
Set Padding 24

Type "# NFCFYI Go API Client"
Enter
Sleep 0.5s

Type "curl -s 'https://nfcfyi.com/api/search/?q=mifare' | jq '.results[:3]'"
Enter
Sleep 3s

Type "curl -s 'https://nfcfyi.com/api/chip/ntag-215/' | jq '{name, slug}'"
Enter
Sleep 3s

Sleep 1s
