package link

import (
	"io"
	"reflect"
	"strings"
	"testing"
)

func TestGetLinks(t *testing.T) {
	var htmlExamples = []string{
		`<html>
		<body>
		<h1>Hello!</h1>
		<a href="/other-page">A link to another page</a>
		<a href="/2ndPage">a second age</a>
		</body>
		</html>`,
		`<html>
		<head>
		<link rel="stylesheet" href="https://maxcdn.bootstrapcdn.com/font-awesome/4.7.0/css/font-awesome.min.css">
		</head>
		<body>
		<h1>Social stuffs</h1>
		<div>
			<a href="https://www.twitter.com/joncalhoun">
			Check me out on twitter
			<i class="fa fa-twitter" aria-hidden="true"></i>
			</a>
			<a href="https://github.com/gophercises">
			Gophercises is on 
			<strong>Github</strong>!
			</a>
		</div>
		</body>
		</html>`,
	}
	type args struct {
		r io.Reader
	}
	tests := []struct {
		name    string
		args    args
		want    []Link
		wantErr bool
	}{
		{
			"first test",
			args{strings.NewReader(htmlExamples[0])},
			[]Link{{"/other-page", "A link to another page"}, {"/2ndPage", "a second age"}},
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := GetLinks(tt.args.r)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetLinks() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetLinks() = %v, want %v", got, tt.want)
			}
		})
	}
}
