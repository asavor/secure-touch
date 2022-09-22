package generator

import (
	"encoding/json"
	"fmt"
	"github.com/brianvoe/gofakeit"
	"github.com/google/uuid"
	"math"
	mact2 "secure-touch/generator/events/mact"
	"secure-touch/generator/types"
	"time"
)

func (c *Client) GenerateInteractionPayload(mact, kact bool, count int) (string, error) {
	KeyboardEvents := make([]types.KeyboardInteractionPayloads, 0)
	var MactEvents []types.MouseInteractionPayloads

	nowtime := time.Now().UTC().UnixNano()/1e6 - 500
	var focusTime int64
	var epochTime int64

	windowID := uuid.New().String()
	if kact {
		var keyStrokeEvents []types.EventsKeyboard

		keyCount := 0
		eventTs := gofakeit.Float64Range(10000, 11000)

		epocTs := nowtime - toFixed(eventTs, 10)

		for i := 0; i < 20; i++ {

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

				increase := gofakeit.Float64Range(1000, 2000)
				eventTs += increase
				epocTs += toFixed(increase, 10)

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

				increase := gofakeit.Float64Range(50, 100)
				eventTs += increase
				epocTs += toFixed(increase, 10)

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

				increase = gofakeit.Float64Range(100, 250)
				eventTs += increase
				epocTs += toFixed(increase, 10)
				keyStrokeEvents = append(keyStrokeEvents, stroke, up)
			}

		}

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
				InnerWidth:            2560,
				InnerHeight:           617,
				OuterWidth:            2560,
				OuterHeight:           1390,
				SnapshotsReduceFactor: 0,
				EventsWereReduced:     true,
			},
		}
		KeyboardEvents = append(KeyboardEvents, keyboardPayloads)

	}

	if mact {
		var mouseMovementsEvents []types.EventsMouse

		eventTs := gofakeit.Float64Range(10000, 11000)
		epocTs := nowtime - toFixed(eventTs, 10)
		xSlice, ySlice := mact2.TfylGenerateMact(40, 1440, 2560)

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
				ScreenY:      int(xSlice[i]),
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

			mouseMovementsEvents = append(mouseMovementsEvents, mouseMovements)
			increase := gofakeit.Float64Range(100, 250)
			eventTs += increase
			epocTs += toFixed(increase, 10)

		}

		MousePayloads := types.MouseInteractionPayloads{
			Events:     mouseMovementsEvents,
			Identified: false,
			Counter:    0,
			AdditionalData: types.AdditionalDataMouse{
				WindowID:              windowID,
				LocationHref:          `https://my.asos.com/identity/login?signin=` + c.LoginUrl,
				Checksum:              "4426372825",
				Mouseout:              8,
				Mouseover:             8,
				InnerWidth:            2560,
				InnerHeight:           617,
				OuterWidth:            1392,
				OuterHeight:           1392,
				SnapshotsReduceFactor: 0,
				EventsWereReduced:     false,
			},
		}
		MactEvents = append(MactEvents, MousePayloads)

	}

	var IndriectPayload []types.IndirectEventsPayload

	IndriectPayload = append(IndriectPayload, types.IndirectEventsPayload{
		Category:       "FocusEvent",
		Type:           "focusin",
		EventTs:        float64(focusTime),
		EpochTs:        epochTime,
		AdditionalData: types.AdditionalDataIndirectEventPayload{},
	})

	tagPayload := []types.Tags{
		{
			Name:      "SDK started",
			EpochTs:   nowtime - 6000,
			Timestamp: nowtime - 6000,
		},
		{
			Name:      "SDK first load",
			EpochTs:   nowtime - 6000,
			Timestamp: nowtime - 6000,
		},
		{
			Name:      "id-signIn-loginAttempt-click:outlook.com",
			EpochTs:   time.Now().UTC().UnixNano() / 1e6,
			Timestamp: time.Now().UTC().UnixNano() / 1e6,
		},
		{
			Name:      "id-signIn-loginAttempt-hash:87241",
			EpochTs:   time.Now().UTC().UnixNano() / 1e6,
			Timestamp: time.Now().UTC().UnixNano() / 1e6,
		},
	}
	environment := types.Environment{
		Ops:              0,
		WebGl:            "",
		DevicePixelRatio: 1,
		ScreenWidth:      2560,
		ScreenHeight:     1440,
	}

	payload := &types.InteractionPayload{

		ApplicationID:               c.AppID,
		DeviceID:                    c.DeviceID,
		DeviceType:                  c.DeviceType,
		AppSessionID:                c.AppSessionId,
		StToken:                     c.StToken,
		KeyboardInteractionPayloads: KeyboardEvents,
		MouseInteractionPayloads:    MactEvents,
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
		UsernameTs:                  time.Now().UTC().UnixNano()/1e6 - 300,
		Username:                    c.DeviceID,
	}

	jsonPayload, err := json.Marshal(payload)
	if err != nil {
		return "", err
	}
	if !mact && !kact {

		jsonPayload := fmt.Sprintf(`{"applicationId":"asos","deviceId":"%v","deviceType":"Chrome(105.0.0.0)-Windows(10)","appSessionId":"%v","stToken":"%v","keyboardInteractionPayloads":[],"mouseInteractionPayloads":[],"indirectEventsPayload":[],"indirectEventsCounters":{},"gestures":[],"metricsData":{},"accelerometerData":[],"gyroscopeData":[],"linearAccelerometerData":[],"rotationData":[],"index":0,"payloadId":"%v","tags":[{"name":"location:%v","epochTs":%v,"timestamp":%v}],"environment":{"ops":0,"webGl":"","devicePixelRatio":1,"screenWidth":2560,"screenHeight":1440},"isMobile":false,"usernameTs":%v,"username":"%v"}`, c.DeviceID, c.AppSessionId, c.StToken, uuid.New().String(), `https://my.asos.com/identity/login?signin=`+c.LoginUrl, nowtime, nowtime, nowtime-5, c.DeviceID)
		return jsonPayload, nil
	}
	return string(jsonPayload), nil

}
func round(num float64) int64 {
	return int64(num + math.Copysign(0.5, num))
}

func toFixed(num float64, precision int) int64 {
	output := math.Pow(10, float64(precision))
	return round((num * output) / output)

}
