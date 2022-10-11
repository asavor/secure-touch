package generator

import (
	"encoding/json"
	"github.com/brianvoe/gofakeit"
	"github.com/google/uuid"
	"math"
	mact2 "secure-touch/generator/events/mact"
	"secure-touch/generator/types"
	"strconv"
	"time"
)

func (c *Client) GenerateInteractionPayload(mact, kact bool, count int) (string, error) {
	KeyboardEvents := make([]types.KeyboardInteractionPayloads, 0)
	MouseMovementsEvents := make([]types.MouseInteractionPayloads, 0)
	IndriectPayload := make([]types.IndirectEventsPayload, 0)

	nowtime := time.Now().UTC().UnixNano() / 1e6

	environment := types.Environment{
		Ops:              0,
		WebGl:            "",
		DevicePixelRatio: 1,
		ScreenWidth:      c.Device.Screen.Width,
		ScreenHeight:     c.Device.Screen.Height,
	}

	if !mact && !kact {
		ran1 := RanNumber(10, 25)

		tagPayload := []types.Tags{
			{
				Name:      `location:https://my.asos.com/identity/login?signin=` + c.LoginUrl,
				EpochTs:   nowtime,
				Timestamp: nowtime,
			},
		}

		payload := &types.InteractionPayload{

			ApplicationID:               c.AppID,
			DeviceID:                    c.DeviceID,
			DeviceType:                  c.DeviceType,
			AppSessionID:                c.AppSessionId,
			StToken:                     c.StToken,
			KeyboardInteractionPayloads: KeyboardEvents,
			MouseInteractionPayloads:    MouseMovementsEvents,
			IndirectEventsPayload:       IndriectPayload,
			IndirectEventsCounters:      types.IndirectEventsCounters{},
			Gestures:                    []string{},
			MetricsData:                 types.MetricsData{},
			AccelerometerData:           []string{},
			GyroscopeData:               []string{},
			LinearAccelerometerData:     []string{},
			RotationData:                []string{},
			Index:                       int(count),
			PayloadID:                   uuid.New().String(),
			Tags:                        tagPayload,
			Environment:                 environment,
			IsMobile:                    false,
			UsernameTs:                  nowtime - int64(ran1),
			Username:                    c.DeviceID,
		}

		jsonPayload, err := json.Marshal(payload)
		if err != nil {
			return "", err
		}
		time.Sleep(time.Duration(RanNumber(2, 10)) * time.Millisecond)

		return string(jsonPayload), nil

	}

	windowID := uuid.New().String()
	TotalMactTime := 0.0
	var kactStartTime float64
	if mact {
		var mouseDataEvents []types.EventsMouse

		eventTs := float64(nowtime)
		epocTs := nowtime

		kactStartTime = float64(nowtime)

		MaxMactCount := RanNumber(40, 80)
		xSlice, ySlice := mact2.TfylGenerateMact(int16(MaxMactCount), int16(c.Device.Screen.Height), int16(c.Device.Screen.Width))
		ranX := RanNumber(100, 200)

		for i, _ := range xSlice {
			mouseMovements := types.EventsMouse{
				Type:         "mousemove",
				EventTs:      eventTs,
				EpochTs:      epocTs,
				Button:       0,
				Buttons:      0,
				ClientX:      int(ySlice[i]),
				ClientY:      int(xSlice[i]),
				MovementX:    0,
				MovementY:    1,
				OffsetX:      0,
				OffsetY:      0,
				PageX:        int(ySlice[i]),
				PageY:        int(xSlice[i]),
				ScreenX:      int(ySlice[i]),
				ScreenY:      int(xSlice[i]) + ranX,
				Which:        0,
				ModifierKeys: []string{},
				//TargetBottom: 0,
				//TargetHeight: 0,
				//TargetLeft:   0,
				//TargetRight:  0,
				//TargetTop:    0,
				//TargetWidth:  0,
				//TargetX:      0,
				//TargetY:      0,
			}

			mouseDataEvents = append(mouseDataEvents, mouseMovements)
			increase := gofakeit.Float64Range(5, 150)
			eventTs += increase
			TotalMactTime += increase
			epocTs += toFixed(increase, 10)

		}
		//fmt.Println(TotalMactTime)

		MousePayloads := types.MouseInteractionPayloads{
			Events:     mouseDataEvents,
			Identified: false,
			Counter:    0,
			AdditionalData: types.AdditionalDataMouse{
				WindowID:              windowID,
				LocationHref:          `https://my.asos.com/identity/login?signin=` + c.LoginUrl,
				Checksum:              "4426372825",
				Mouseout:              8,
				Mouseover:             8,
				InnerWidth:            c.Device.Window.InnerWidth,
				InnerHeight:           c.Device.Window.InnerWidth,
				OuterWidth:            c.Device.Window.OuterWidth,
				OuterHeight:           c.Device.Window.InnerHeight,
				SnapshotsReduceFactor: 0,
				EventsWereReduced:     false,
			},
		}
		MouseMovementsEvents = append(MouseMovementsEvents, MousePayloads)

	}

	var focusTime int64
	var epochTime int64

	ran1 := RanNumber(40, 100)

	totalKeyEventTime := 0.0
	if kact {
		var keyStrokeEvents []types.EventsKeyboard
		KeyboardInputDelay := float64(RanNumber(300, 600))
		keyCount := 0
		eventTs := kactStartTime + KeyboardInputDelay

		epocTs := nowtime + int64(RanNumber(600, 1200))

		MaxKeyCount := RanNumber(17, 25)

		for i := 0; i < MaxKeyCount; i++ {

			if i == 0 {
				stroke := types.EventsKeyboard{
					Type:         "focus",
					EventTs:      eventTs,
					EpochTs:      epocTs,
					ModifierKeys: []string{},
				}
				focusTime = int64(eventTs)
				epochTime = epocTs

				keyStrokeEvents = append(keyStrokeEvents, stroke)

				increase := gofakeit.Float64Range(100, 300)
				eventTs += increase
				epocTs += toFixed(increase, 10)
				totalKeyEventTime += increase

			} else {
				id := uuid.New()

				stroke := types.EventsKeyboard{
					Type:           "keydown",
					EventTs:        eventTs,
					EpochTs:        epocTs,
					ModifierKeys:   []string{},
					SelectionStart: nil,
					SelectionEnd:   nil,
					Repeat:         false,
					Key:            nil,
					KeyCode:        nil,
					KeystrokeID:    id.String(),
					CurrentLength:  keyCount,
					Location:       0,
				}
				keyCount++

				increase := gofakeit.Float64Range(50, 125)
				eventTs += increase
				epocTs += toFixed(increase, 10)
				totalKeyEventTime += increase

				up := types.EventsKeyboard{
					Type:           "keyup",
					EventTs:        eventTs,
					EpochTs:        epocTs,
					ModifierKeys:   []string{},
					SelectionStart: nil,
					SelectionEnd:   nil,
					Repeat:         false,
					Key:            nil,
					KeyCode:        nil,
					KeystrokeID:    id.String(),
					CurrentLength:  keyCount,
					Location:       0,
				}

				increase = gofakeit.Float64Range(50, 125)
				eventTs += increase
				epocTs += toFixed(increase, 10)
				keyStrokeEvents = append(keyStrokeEvents, stroke, up)
				totalKeyEventTime += increase
			}

		}

		//fmt.Println(totalKeyEventTime)

		keyboardPayloads := types.KeyboardInteractionPayloads{
			StID:       "id-signIn-emailAddress",
			ElementID:  "EmailAddress",
			Name:       "Username",
			Type:       "email",
			Events:     keyStrokeEvents,
			Identified: false,
			Counter:    0,
			AdditionalData: types.AdditionalDataKeyboard{
				WindowID:              windowID,
				LocationHref:          `https://my.asos.com/identity/login?signin=` + c.LoginUrl,
				Checksum:              "4426372825",
				InnerWidth:            c.Device.Window.InnerWidth,
				InnerHeight:           c.Device.Window.InnerWidth,
				OuterWidth:            c.Device.Window.OuterWidth,
				OuterHeight:           c.Device.Window.InnerHeight,
				SnapshotsReduceFactor: 0,
				EventsWereReduced:     true,
			},
		}
		KeyboardEvents = append(KeyboardEvents, keyboardPayloads)

	}

	IndriectPayload = append(IndriectPayload, types.IndirectEventsPayload{
		Category: "FocusEvent",
		Type:     "focusin",
		EventTs:  float64(focusTime),
		EpochTs:  epochTime,
		AdditionalData: types.AdditionalDataIndirectEventPayload{
			StID:                   "id-signIn-emailAddress",
			ElementID:              "EmailAddress",
			RelatedTargetType:      "",
			RelatedTargetStID:      "",
			RelatedTargetElementID: "",
			WindowID:               windowID,
			LocationHref:           "https://my.asos.com/identity/login?" + c.LoginUrl,
			Checksum:               "4426372825",
		},
	}, types.IndirectEventsPayload{
		Category: "FocusEvent",
		Type:     "DOMFocusIn",
		EventTs:  float64(focusTime) + gofakeit.Float64Range(0, 1),
		EpochTs:  epochTime + 1,
		AdditionalData: types.AdditionalDataIndirectEventPayload{
			StID:                   "id-signIn-emailAddress",
			ElementID:              "EmailAddress",
			RelatedTargetType:      "",
			RelatedTargetStID:      "",
			RelatedTargetElementID: "",
			WindowID:               windowID,
			LocationHref:           "https://my.asos.com/identity/login?" + c.LoginUrl,
			Checksum:               "4426372825",
		},
	})

	time.Sleep(time.Duration(TotalMactTime) * time.Millisecond)
	ran2 := RanNumber(100, 300)
	hash := RanNumber(20000, 99999)
	nowTime := time.Now().UTC().UnixNano() / 1e6

	tagPayload := []types.Tags{
		{
			Name:      "SDK started",
			EpochTs:   c.StartTime - int64(ran1),
			Timestamp: c.StartTime - int64(ran1),
		},
		{
			Name:      "SDK first load",
			EpochTs:   c.StartTime - int64(ran1),
			Timestamp: c.StartTime - int64(ran1),
		},
		{
			Name:      "id-signIn-loginAttempt-click:outlook.com",
			EpochTs:   nowTime - int64(ran2),
			Timestamp: nowTime - int64(ran2),
		},
		{
			Name:      "id-signIn-loginAttempt-hash:" + strconv.Itoa(hash),
			EpochTs:   nowTime - int64(ran2),
			Timestamp: nowTime - int64(ran2),
		},
	}
	//time.Sleep(time.Duration(RanNumber(1, 3)) * time.Second)

	payload := &types.InteractionPayload{

		ApplicationID:               c.AppID,
		DeviceID:                    c.DeviceID,
		DeviceType:                  c.DeviceType,
		AppSessionID:                c.AppSessionId,
		StToken:                     c.StToken,
		KeyboardInteractionPayloads: KeyboardEvents,
		MouseInteractionPayloads:    MouseMovementsEvents,
		IndirectEventsPayload:       IndriectPayload,
		IndirectEventsCounters:      types.IndirectEventsCounters{},
		Gestures:                    []string{},
		MetricsData:                 types.MetricsData{},
		AccelerometerData:           []string{},
		GyroscopeData:               []string{},
		LinearAccelerometerData:     []string{},
		RotationData:                []string{},
		Index:                       1,
		PayloadID:                   uuid.New().String(),
		Tags:                        tagPayload,
		Environment:                 environment,
		IsMobile:                    false,
		UsernameTs:                  time.Now().UTC().UnixNano() / 1e6,
		Username:                    c.DeviceID,
	}

	jsonPayload, err := json.Marshal(payload)
	if err != nil {
		return "", err
	}

	time.Sleep(150 * time.Millisecond)
	return string(jsonPayload), nil

}
func round(num float64) int64 {
	return int64(num + math.Copysign(0.5, num))
}

func toFixed(num float64, precision int) int64 {
	output := math.Pow(10, float64(precision))
	return round((num * output) / output)

}
