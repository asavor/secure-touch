package types

type InteractionPayload struct {
	ApplicationID               string                        `json:"applicationId"`
	DeviceID                    string                        `json:"deviceId"`
	DeviceType                  string                        `json:"deviceType"`
	AppSessionID                string                        `json:"appSessionId"`
	StToken                     string                        `json:"stToken"`
	KeyboardInteractionPayloads []KeyboardInteractionPayloads `json:"keyboardInteractionPayloads"`
	MouseInteractionPayloads    []MouseInteractionPayloads    `json:"mouseInteractionPayloads"`
	IndirectEventsPayload       []IndirectEventsPayload       `json:"indirectEventsPayload"`
	IndirectEventsCounters      IndirectEventsCounters        `json:"indirectEventsCounters"`
	Gestures                    []string                      `json:"gestures"`
	MetricsData                 MetricsData                   `json:"metricsData"`
	AccelerometerData           []string                      `json:"accelerometerData"`
	GyroscopeData               []string                      `json:"gyroscopeData"`
	LinearAccelerometerData     []string                      `json:"linearAccelerometerData"`
	RotationData                []string                      `json:"rotationData"`
	Index                       int                           `json:"index"`
	PayloadID                   string                        `json:"payloadId"`
	Tags                        []Tags                        `json:"tags"`
	Environment                 Environment                   `json:"environment"`
	IsMobile                    bool                          `json:"isMobile"`
	UsernameTs                  int64                         `json:"usernameTs"`
	Username                    string                        `json:"username"`
}
type EventsKeyboard struct {
	Type           string      `json:"type"`
	EventTs        float64     `json:"eventTs"`
	EpochTs        int64       `json:"epochTs"`
	ModifierKeys   []string    `json:"modifierKeys"`
	SelectionStart interface{} `json:"selectionStart"`
	SelectionEnd   interface{} `json:"selectionEnd"`
	Key            interface{} `json:"key"`
	KeyCode        interface{} `json:"keyCode"`
	KeystrokeID    interface{} `json:"keystrokeId"`
	CurrentLength  int         `json:"currentLength"`
	Repeat         bool        `json:"repeat,omitempty"`
	Location       int         `json:"location,omitempty"`
}
type AdditionalDataKeyboard struct {
	WindowID     string `json:"windowId"`
	LocationHref string `json:"locationHref"`
	Checksum     string `json:"checksum"`
	//DevToolsOpen          bool   `json:"DEV_TOOLS_OPEN"`
	//DevToolsOrientation   string `json:"DEV_TOOLS_ORIENTATION"`
	InnerWidth            int  `json:"innerWidth"`
	InnerHeight           int  `json:"innerHeight"`
	OuterWidth            int  `json:"outerWidth"`
	OuterHeight           int  `json:"outerHeight"`
	SnapshotsReduceFactor int  `json:"snapshotsReduceFactor"`
	EventsWereReduced     bool `json:"eventsWereReduced"`
}
type KeyboardInteractionPayloads struct {
	StID           string                 `json:"stId"`
	ElementID      string                 `json:"elementId"`
	Name           string                 `json:"name"`
	Type           string                 `json:"type"`
	Events         []EventsKeyboard       `json:"events"`
	Identified     bool                   `json:"identified"`
	Counter        int                    `json:"counter"`
	AdditionalData AdditionalDataKeyboard `json:"additionalData"`
}
type EventsMouse struct {
	Type         string   `json:"type"`
	EventTs      float64  `json:"eventTs"`
	EpochTs      int64    `json:"epochTs"`
	Button       int      `json:"button"`
	Buttons      int      `json:"buttons"`
	ClientX      int      `json:"clientX"`
	ClientY      int      `json:"clientY"`
	MovementX    int      `json:"movementX"`
	MovementY    int      `json:"movementY"`
	OffsetX      int      `json:"offsetX"`
	OffsetY      int      `json:"offsetY"`
	PageX        int      `json:"pageX"`
	PageY        int      `json:"pageY"`
	ScreenX      int      `json:"screenX"`
	ScreenY      int      `json:"screenY"`
	Which        int      `json:"which"`
	ModifierKeys []string `json:"modifierKeys"`
	TargetBottom int      `json:"targetBottom,omitempty"`
	TargetHeight int      `json:"targetHeight,omitempty"`
	TargetLeft   float64  `json:"targetLeft,omitempty"`
	TargetRight  float64  `json:"targetRight,omitempty"`
	TargetTop    int      `json:"targetTop,omitempty"`
	TargetWidth  int      `json:"targetWidth,omitempty"`
	TargetX      float64  `json:"targetX,omitempty"`
	TargetY      int      `json:"targetY,omitempty"`
}
type AdditionalDataMouse struct {
	WindowID     string `json:"windowId"`
	LocationHref string `json:"locationHref"`
	Checksum     string `json:"checksum"`
	Mouseout     int    `json:"mouseout"`
	Mouseover    int    `json:"mouseover"`
	//DevToolsOpen          bool   `json:"DEV_TOOLS_OPEN"`
	//DevToolsOrientation   string `json:"DEV_TOOLS_ORIENTATION"`
	InnerWidth            int  `json:"innerWidth"`
	InnerHeight           int  `json:"innerHeight"`
	OuterWidth            int  `json:"outerWidth"`
	OuterHeight           int  `json:"outerHeight"`
	SnapshotsReduceFactor int  `json:"snapshotsReduceFactor"`
	EventsWereReduced     bool `json:"eventsWereReduced"`
}
type MouseInteractionPayloads struct {
	Events         []EventsMouse       `json:"events"`
	Identified     bool                `json:"identified"`
	Counter        int                 `json:"counter"`
	AdditionalData AdditionalDataMouse `json:"additionalData"`
}

type AdditionalDataIndirectEventPayload struct {
	StID                   string `json:"stId"`
	ElementID              string `json:"elementId"`
	RelatedTargetType      string `json:"relatedTarget.type"`
	RelatedTargetStID      string `json:"relatedTarget.stId"`
	RelatedTargetElementID string `json:"relatedTarget.elementId"`
	WindowID               string `json:"windowId"`
	LocationHref           string `json:"locationHref"`
	Checksum               string `json:"checksum"`
}
type IndirectEventsPayload struct {
	Category       string                             `json:"category"`
	Type           string                             `json:"type"`
	EventTs        float64                            `json:"eventTs"`
	EpochTs        int64                              `json:"epochTs"`
	AdditionalData AdditionalDataIndirectEventPayload `json:"additionalData,omitempty"`
}
type IndirectEventsCounters struct {
}
type MetricsData struct {
}
type Tags struct {
	Name      string `json:"name"`
	EpochTs   int64  `json:"epochTs"`
	Timestamp int64  `json:"timestamp"`
}
type Environment struct {
	Ops              int    `json:"ops"`
	WebGl            string `json:"webGl"`
	DevicePixelRatio int    `json:"devicePixelRatio"`
	ScreenWidth      int    `json:"screenWidth"`
	ScreenHeight     int    `json:"screenHeight"`
}
