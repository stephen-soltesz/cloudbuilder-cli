package main

import (
	"context"
	"fmt"

	"github.com/m-lab/go/rtx"
	"github.com/stephen-soltesz/pretty"

	"golang.org/x/oauth2/google"
	"google.golang.org/api/cloudbuild/v1"
)

func main() {
	fmt.Println("ok")
	ctx := context.Background()
	client, err := google.DefaultClient(ctx, cloudbuild.CloudPlatformScope)
	rtx.Must(err, "Failed to allocate default client")

	build, err := cloudbuild.New(client)
	rtx.Must(err, "Failed to create new cloudbuilder client")

	tlc := build.Projects.Triggers.List("mlab-sandbox")
	r, err := tlc.Do()
	rtx.Must(err, "Failed to list triggers")
	for i := range r.Triggers {
		pretty.Print(r.Triggers[i])
	}
}
