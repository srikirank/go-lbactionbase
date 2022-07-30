package model

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"os/exec"
)

const DefaultIcon = "icon.png"
const DefaultErrorIcon = "/System/Library/CoreServices/CoreTypes.bundle/Contents/Resources/AlertStopIcon.icns"

const LBCallbackCmd = "open 'x-go-lbactionbase:select?'"

// https://developer.obdev.at/launchbar-developer-documentation/#/script-output
type LBItem struct {
	Title                  string   `json:"title"`
	Subtitle               string   `json:"subtitle,omitempty"`
	AlwaysShowsSubtitle    bool     `json:"alwaysShowsSubtitle"`
	Url                    string   `json:"url,omitempty"`
	Path                   string   `json:"path,omitempty"`
	Label                  string   `json:"label,omitempty"`
	Badge                  string   `json:"badge,omitempty"`
	Icon                   string   `json:"icon,omitempty"`
	IconFont               string   `json:"iconFont,omitempty"`
	Action                 string   `json:"action"`
	ActionArgument         string   `json:"actionArgument,omitempty"`
	ActionReturnsItems     bool     `json:"actionReturnsItems,omitempty"`
	ActionBundleIdentifier string   `json:"actionBundleIdentifier,omitempty"`
	Children               []LBItem `json:"children,omitempty"`
	QuickLookURL           string   `json:"quickLookURL,omitempty"`
}

type LBItems struct {
	Items []LBItem
}

// String To support fuzzy matching:https://github.com/sahilm/fuzzy
func (it *LBItems) String(i int) string {
	return it.Items[i].Title
}

// Len To support fuzzy matching: https://github.com/sahilm/fuzzy
func (it *LBItems) Len() int {
	return len(it.Items)
}

func NewTitleOnlyLBItem(Title string) *LBItem {
	return &LBItem{
		Title: Title,
		Icon:  DefaultIcon,
	}
}

func (i *LBItem) WithSubtitle(SubTitle string) *LBItem {
	i.Subtitle = SubTitle
	i.AlwaysShowsSubtitle = true

	return i
}

func (i *LBItem) WithBadge(Badge string) *LBItem {
	i.Badge = Badge
	return i
}

func (i *LBItem) WithLabel(Label string) *LBItem {
	i.Label = Label
	return i
}

func (i *LBItem) WithIcon(Icon string) *LBItem {
	i.Icon = Icon
	return i
}

func (i *LBItem) WithIconFont(IconFont string) *LBItem {
	i.IconFont = IconFont
	return i
}

func (i *LBItem) WithPath(Path string) *LBItem {
	i.Path = Path
	return i
}

func (i *LBItem) WithUrl(Url string) *LBItem {
	i.Url = Url
	return i
}

func CopyAction(text string) {
	c := exec.Command("/usr/bin/pbcopy")
	stdin, e := c.StdinPipe()
	if e != nil {
		log.Printf("Could not copy : %s", text)
		return
	}
	go func() {
		defer stdin.Close()
		_, _ = io.WriteString(stdin, text)
	}()
	_ = c.Run()
}

func (i *LBItem) WithChildren(Children []LBItem) *LBItem {
	i.Children = append(i.Children, Children...)
	return i
}

func (i *LBItem) WithQuickLookURL(Url string) *LBItem {
	i.QuickLookURL = Url
	return i
}

func (i *LBItem) WithAction(action string) *LBItem {
	i.Action = action
	return i
}

func (i *LBItem) WithActionBundleIdentifier(actionBundleIdentifier string) *LBItem {
	i.ActionBundleIdentifier = actionBundleIdentifier
	return i
}

func (i *LBItem) WithActionArgument(actionArgument string) *LBItem {
	i.ActionArgument = actionArgument
	return i
}

func (i *LBItem) WithActionReturnsItems(actionReturnsItems bool) *LBItem {
	i.ActionReturnsItems = actionReturnsItems
	return i
}

func LBOutput(items LBItems) {
	if len(items.Items) < 1 {
		LBEmptyOutput()
		return
	}
	buf := new(bytes.Buffer)
	enc := json.NewEncoder(buf)
	enc.SetEscapeHTML(false)
	_ = enc.Encode(&items.Items)
	fmt.Printf(buf.String())
}

func LBEmptyOutput() {
	fmt.Printf("[]")
}

func LBErrorItem(err error) *LBItem {
	return NewTitleOnlyLBItem(err.Error()).WithIcon(DefaultErrorIcon)
}

func DefaultLBErrorOutput(e error) {
	title := "An Unknown Error Occurred"
	if e != nil {
		title = e.Error()
		if e.Error() == "exit status 67" {
			title = "Error using curl at /usr/bin/curl. Install: `brew install curl-openssl`"
		}
	}
	var lbItems LBItems
	lbItem := NewTitleOnlyLBItem(title).WithIcon(DefaultErrorIcon)
	lbItems.Items = append(lbItems.Items, *lbItem)

	LBOutput(lbItems)
}
