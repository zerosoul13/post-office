package publisher

import (
	"reflect"
	"testing"
)

func TestNewPublisher(t *testing.T) {
	type args struct {
		publish func(string) error
	}
	tests := []struct {
		name string
		args args
		want Publisher
	}{
		{
			name: "NewPublisher",
			args: args{
				publish: func(string) error {
					return nil
				},
			},
			want: &EventPublisher{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewPublisher(tt.args.publish); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewPublisher() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestEventPublisher_Suscribe(t *testing.T) {

	evs := EventSuscriber{}

	type fields struct {
		message    chan string
		suscribers []Subscriber
	}
	type args struct {
		s Subscriber
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			name: "Suscribe",
			fields: fields{
				message:    make(chan string),
				suscribers: []Subscriber{&evs},
			},
			args: args{
				s: &EventSuscriber{},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := &EventPublisher{
				message:    tt.fields.message,
				suscribers: tt.fields.suscribers,
			}
			if err := p.Suscribe(tt.args.s); (err != nil) != tt.wantErr {
				t.Errorf("EventPublisher.Suscribe() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestEventPublisher_Publish(t *testing.T) {
	e := EventSuscriber{}
	ev := EventSuscriber{}

	c := make(chan string, 1)

	go func() {
		c <- "Hello World"
	}()

	type fields struct {
		message    chan string
		suscribers []Subscriber
	}
	tests := []struct {
		name    string
		fields  fields
		wantErr bool
	}{
		{
			name: "Publish",
			fields: fields{
				message:    c,
				suscribers: []Subscriber{&e, &ev},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := &EventPublisher{
				message:    tt.fields.message,
				suscribers: tt.fields.suscribers,
			}
			if err := p.Publish(); (err != nil) != tt.wantErr {
				t.Errorf("EventPublisher.Publish() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
